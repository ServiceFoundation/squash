// Code generated by go-swagger; DO NOT EDIT.

package debugrequest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/solo-io/squash/pkg/models"
)

// GetDebugRequestOKCode is the HTTP code returned for type GetDebugRequestOK
const GetDebugRequestOKCode int = 200

/*GetDebugRequestOK OK

swagger:response getDebugRequestOK
*/
type GetDebugRequestOK struct {

	/*
	  In: Body
	*/
	Payload *models.DebugRequest `json:"body,omitempty"`
}

// NewGetDebugRequestOK creates GetDebugRequestOK with default headers values
func NewGetDebugRequestOK() *GetDebugRequestOK {
	return &GetDebugRequestOK{}
}

// WithPayload adds the payload to the get debug request o k response
func (o *GetDebugRequestOK) WithPayload(payload *models.DebugRequest) *GetDebugRequestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get debug request o k response
func (o *GetDebugRequestOK) SetPayload(payload *models.DebugRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDebugRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDebugRequestNotFoundCode is the HTTP code returned for type GetDebugRequestNotFound
const GetDebugRequestNotFoundCode int = 404

/*GetDebugRequestNotFound Not found

swagger:response getDebugRequestNotFound
*/
type GetDebugRequestNotFound struct {
}

// NewGetDebugRequestNotFound creates GetDebugRequestNotFound with default headers values
func NewGetDebugRequestNotFound() *GetDebugRequestNotFound {
	return &GetDebugRequestNotFound{}
}

// WriteResponse to the client
func (o *GetDebugRequestNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
