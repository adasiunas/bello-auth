package business

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"os"
	"sync"
	"time"
)

const ACCESS_TOKEN_EXPIRE_TIME  = 5 * time.Minute
const REFRESH_TOKEN_EXPIRE_TIME  = 20 * time.Minute
var jwtSecret string
var once sync.Once

func IssueToken(userID uuid.UUID, expireAt time.Time) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{}
	claims["sub"] = userID
	claims["exp"] = expireAt.Unix()
	claims["iat"] = now.Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(JWTSecret())
}

func JWTSecret() []byte {
	once.Do(func() {
		jwtSecret = os.Getenv("JWT_SECRET")
	})

	return []byte(jwtSecret)
}