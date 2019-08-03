package db

import (
	belloerr "github.com/adasiunas/bello-auth/error"
	"github.com/adasiunas/bello-auth/model"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"time"
)

func CreateUser(db *gorm.DB, email string, passwordToken string) (user model.User, err error) {
	now := time.Now().UTC()
	user = model.User{
		ID:            uuid.NewV1(),
		Email:         email,
		PasswordToken: passwordToken,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	err = db.Create(&user).Error
	return user, err
}

func GetUserByEmail(db *gorm.DB, email string) (user model.User, err error) {
	err = db.First(&user, model.User{Email: email}).Error
	if err == gorm.ErrRecordNotFound {
		return user, belloerr.ErrUserNotFound
	}

	return user, err
}

func GetUserByID(db *gorm.DB, ID uuid.UUID) (user model.User, err error) {
	err = db.First(&user, model.User{ID: ID}).Error
	if err == gorm.ErrRecordNotFound {
		return user, belloerr.ErrUserNotFound
	}

	return user, err
}
