// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	apimodel "github.com/adasiunas/bello-auth/apimodel"
)

// NewRegisterUserV1Params creates a new RegisterUserV1Params object
// no default values defined in spec.
func NewRegisterUserV1Params() RegisterUserV1Params {

	return RegisterUserV1Params{}
}

// RegisterUserV1Params contains all the bound params for the register user v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters registerUserV1
type RegisterUserV1Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Register *apimodel.RegistrationRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRegisterUserV1Params() beforehand.
func (o *RegisterUserV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body apimodel.RegistrationRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("register", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Register = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
