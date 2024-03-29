// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	apimodel "github.com/adasiunas/bello-auth/apimodel"
)

// RegisterUserV1OKCode is the HTTP code returned for type RegisterUserV1OK
const RegisterUserV1OKCode int = 200

/*RegisterUserV1OK User successfully registered

swagger:response registerUserV1OK
*/
type RegisterUserV1OK struct {

	/*
	  In: Body
	*/
	Payload *apimodel.TokenResponse `json:"body,omitempty"`
}

// NewRegisterUserV1OK creates RegisterUserV1OK with default headers values
func NewRegisterUserV1OK() *RegisterUserV1OK {

	return &RegisterUserV1OK{}
}

// WithPayload adds the payload to the register user v1 o k response
func (o *RegisterUserV1OK) WithPayload(payload *apimodel.TokenResponse) *RegisterUserV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user v1 o k response
func (o *RegisterUserV1OK) SetPayload(payload *apimodel.TokenResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterUserV1BadRequestCode is the HTTP code returned for type RegisterUserV1BadRequest
const RegisterUserV1BadRequestCode int = 400

/*RegisterUserV1BadRequest Bad request payload

swagger:response registerUserV1BadRequest
*/
type RegisterUserV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *apimodel.ErrorResponse `json:"body,omitempty"`
}

// NewRegisterUserV1BadRequest creates RegisterUserV1BadRequest with default headers values
func NewRegisterUserV1BadRequest() *RegisterUserV1BadRequest {

	return &RegisterUserV1BadRequest{}
}

// WithPayload adds the payload to the register user v1 bad request response
func (o *RegisterUserV1BadRequest) WithPayload(payload *apimodel.ErrorResponse) *RegisterUserV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user v1 bad request response
func (o *RegisterUserV1BadRequest) SetPayload(payload *apimodel.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterUserV1InternalServerErrorCode is the HTTP code returned for type RegisterUserV1InternalServerError
const RegisterUserV1InternalServerErrorCode int = 500

/*RegisterUserV1InternalServerError Service down

swagger:response registerUserV1InternalServerError
*/
type RegisterUserV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *apimodel.ErrorResponse `json:"body,omitempty"`
}

// NewRegisterUserV1InternalServerError creates RegisterUserV1InternalServerError with default headers values
func NewRegisterUserV1InternalServerError() *RegisterUserV1InternalServerError {

	return &RegisterUserV1InternalServerError{}
}

// WithPayload adds the payload to the register user v1 internal server error response
func (o *RegisterUserV1InternalServerError) WithPayload(payload *apimodel.ErrorResponse) *RegisterUserV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user v1 internal server error response
func (o *RegisterUserV1InternalServerError) SetPayload(payload *apimodel.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
