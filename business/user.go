package business

import (
	"github.com/adasiunas/bello-auth/db"
	belloerr "github.com/adasiunas/bello-auth/error"
	"github.com/adasiunas/bello-auth/model"
	"github.com/jinzhu/gorm"
	"time"
)

func RegisterUser(d *gorm.DB, email, pass string) (token model.Token, err error) {
	if _, err = db.GetUserByEmail(d, email); err != nil && err != belloerr.ErrUserNotFound {
		return token, err
	}

	if err == belloerr.ErrUserNotFound {
		now := time.Now().UTC()
		user, err := db.CreateUser(d, email, pass)
		if err != nil {
			return token, err
		}

		token.AccessToken, err = IssueToken(user.ID, now.Add(ACCESS_TOKEN_EXPIRE_TIME))
		return token, err
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

	now := time.Now().UTC()
	token.AccessToken, err = IssueToken(user.ID, now.Add(ACCESS_TOKEN_EXPIRE_TIME))
	return token, err
}