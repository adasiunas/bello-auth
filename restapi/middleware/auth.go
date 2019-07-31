package middlew

import (
	"github.com/adasiunas/bello-auth/apimodel"
	"github.com/adasiunas/bello-auth/business"
	"github.com/dgrijalva/jwt-go"
	apierrors "github.com/go-openapi/errors"
	"net/http"
	"strings"
	"time"
)

func BearerAuthentication(token string) (interface{}, error) {
	const prefix = "Bearer "

	if !strings.HasPrefix(token, prefix) {
		return nil, apierrors.New(http.StatusBadRequest, "incorrect jwt provided")
	}

	token = strings.TrimPrefix(token, prefix)
	claims := jwt.MapClaims{}
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return business.JWTSecret(), nil
	})

	if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
		return nil, apierrors.New(http.StatusUnauthorized, "access token is expired")
	}

	tokenType := claims["tp"]
	if tokenType != string(business.ACCESS_TOKEN) {
		return nil, apierrors.New(http.StatusUnauthorized, apimodel.ErrorResponseMessageUnauthorized)
	}

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, apierrors.New(http.StatusUnauthorized, apimodel.ErrorResponseMessageUnauthorized)
		}

		return nil, apierrors.New(http.StatusInternalServerError, apimodel.ErrorResponseMessageInternalServerError)
	}

	if !tkn.Valid {
		return nil, apierrors.New(http.StatusUnauthorized, apimodel.ErrorResponseMessageUnauthorized)
	}

	return claims, nil
}
