package business

import (
	"github.com/adasiunas/bello-auth/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"os"
	"sync"
	"time"
)

const (
	ACCESS_TOKEN_EXPIRE_TIME  = 5 * time.Minute
	REFRESH_TOKEN_EXPIRE_TIME  = 20 * time.Second
	ACCESS_TOKEN TokenType = "access"
	REFRESH_TOKEN TokenType = "refresh"
)

type TokenType string
var jwtSecret string
var once sync.Once

func IssueToken(userID uuid.UUID, tokenTp TokenType, expireAt time.Time) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{}
	claims["sub"] = userID
	claims["exp"] = expireAt.Unix()
	claims["iat"] = now.Unix()
	claims["tp"] = tokenTp

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(JWTSecret())
}

func JWTSecret() []byte {
	once.Do(func() {
		jwtSecret = os.Getenv("JWT_SECRET")
	})

	return []byte(jwtSecret)
}

func ValidateRefreshToken(token string) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret(), nil
	})

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, errors.New("refresh token is expired")
	}

	tokenType := claims["tp"]
	if tokenType != string(REFRESH_TOKEN) {
		return nil, errors.New("refresh token is invalid")
	}

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("refresh token is expired")
		}

		return nil, errors.Wrap(err, "messed up refresh token")
	}

	if !tkn.Valid {
		return nil, errors.New("refresh token is invalid")
	}

	return claims, nil
}

func GeneratePasswordToken(pass string) string {
	salt := util.GenerateSalt()
	pwHash := util.GetHash(pass, salt)

	return pwHash.String() + "." + salt.String()
}