// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetResourceOKCode is the HTTP code returned for type GetResourceOK
const GetResourceOKCode int = 200

/*GetResourceOK get resource o k

swagger:response getResourceOK
*/
type GetResourceOK struct {

	/*
	  In: Body
	*/
	Payload *GetResourceOKBody `json:"body,omitempty"`
}

// NewGetResourceOK creates GetResourceOK with default headers values
func NewGetResourceOK() *GetResourceOK {

	return &GetResourceOK{}
}

// WithPayload adds the payload to the get resource o k response
func (o *GetResourceOK) WithPayload(payload *GetResourceOKBody) *GetResourceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get resource o k response
func (o *GetResourceOK) SetPayload(payload *GetResourceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResourceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
