package business

import (
	"github.com/adasiunas/bello-auth/db"
	belloerr "github.com/adasiunas/bello-auth/error"
	"github.com/adasiunas/bello-auth/model"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

func RegisterUser(d *gorm.DB, email, pass string) (token model.Token, err error) {
	if _, err = db.GetUserByEmail(d, email); err != nil && err != belloerr.ErrUserNotFound {
		return token, err
	}

	if err == belloerr.ErrUserNotFound {
		user, err := db.CreateUser(d, email, pass)
		if err != nil {
			return token, err
		}

		return RefreshTokens(d, user.ID)
	}

	return token, belloerr.ErrUserAlreadyExist
}

func LoginUser(d *gorm.DB, email, pass string) (token model.Token, err error) {
	user, err := db.GetUserByEmail(d, email)
	if err != nil {
		return token, err
	}

	if user.Password != pass {
		return token, belloerr.ErrUserNotFound
	}

	return RefreshTokens(d, user.ID)
}

func RefreshTokens(d *gorm.DB, userID uuid.UUID) (token model.Token, err error) {
	user, err := db.GetUserByID(d, userID)
	if err != nil {
		return
	}

	token.AccessToken, err = IssueToken(user.ID, ACCESS_TOKEN, time.Now().Add(ACCESS_TOKEN_EXPIRE_TIME))
	if err != nil {
		return token, err
	}

	token.RefreshToken, err = IssueToken(user.ID, REFRESH_TOKEN, time.Now().Add(REFRESH_TOKEN_EXPIRE_TIME))
	if err != nil {
		return token, err
	}

	return token, err
}