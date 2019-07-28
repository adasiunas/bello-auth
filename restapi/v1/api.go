package v1

import (
	"github.com/adasiunas/bello-auth/apimodel"
	"github.com/adasiunas/bello-auth/business"
	belloerr "github.com/adasiunas/bello-auth/error"
	"github.com/adasiunas/bello-auth/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
)

// /v1/user [POST]
func RegisterUser(f *business.Factory) func(params user.RegisterUserV1Params) middleware.Responder {
	return func(params user.RegisterUserV1Params) middleware.Responder {
		token, err := business.RegisterUser(f.DB(), params.Register.Email, params.Register.Password)
		if err != nil {
			if err == belloerr.ErrUserAlreadyExist {
				return user.NewRegisterUserV1BadRequest().WithPayload(&apimodel.ErrorResponse{Message: err.Error(), Type: apimodel.ErrorResponseMessageUserAlreadyExists})
			}

			return user.NewRegisterUserV1InternalServerError().WithPayload(&apimodel.ErrorResponse{Message: err.Error(), Type: apimodel.ErrorResponseMessageInternalServerError})
		}

		return user.NewRegisterUserV1OK().WithPayload(&apimodel.TokenResponse{AccessToken: token.AccessToken, RefreshToken: token.RefreshToken})
	}
}

// /v1/user/login [POST]
func LoginUser(f *business.Factory) func(params user.LoginUserParams) middleware.Responder {
	return func(params user.LoginUserParams) middleware.Responder {
		token, err := business.LoginUser(f.DB(), params.Login.Email, params.Login.Password)
		if err != nil {
			if err == belloerr.ErrUserNotFound {
				return user.NewLoginUserBadRequest().WithPayload(&apimodel.ErrorResponse{Message:"Incorrect credentials", Type:apimodel.ErrorResponseMessageIncorrectCredentials})
			}

			return user.NewLoginUserInternalServerError().WithPayload(&apimodel.ErrorResponse{Message: err.Error(), Type: apimodel.ErrorResponseMessageInternalServerError})
		}

		return user.NewRegisterUserV1OK().WithPayload(&apimodel.TokenResponse{AccessToken: token.AccessToken, RefreshToken: token.RefreshToken})
	}
}
