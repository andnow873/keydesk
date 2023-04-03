// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vpngen/keydesk/gen/models"
)

// PostTokenCreatedCode is the HTTP code returned for type PostTokenCreated
const PostTokenCreatedCode int = 201

/*
PostTokenCreated Token created.

swagger:response postTokenCreated
*/
type PostTokenCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Token `json:"body,omitempty"`
}

// NewPostTokenCreated creates PostTokenCreated with default headers values
func NewPostTokenCreated() *PostTokenCreated {

	return &PostTokenCreated{}
}

// WithPayload adds the payload to the post token created response
func (o *PostTokenCreated) WithPayload(payload *models.Token) *PostTokenCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post token created response
func (o *PostTokenCreated) SetPayload(payload *models.Token) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTokenCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTokenInternalServerErrorCode is the HTTP code returned for type PostTokenInternalServerError
const PostTokenInternalServerErrorCode int = 500

/*
PostTokenInternalServerError Internal server error

swagger:response postTokenInternalServerError
*/
type PostTokenInternalServerError struct {
}

// NewPostTokenInternalServerError creates PostTokenInternalServerError with default headers values
func NewPostTokenInternalServerError() *PostTokenInternalServerError {

	return &PostTokenInternalServerError{}
}

// WriteResponse to the client
func (o *PostTokenInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

/*
PostTokenDefault error

swagger:response postTokenDefault
*/
type PostTokenDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostTokenDefault creates PostTokenDefault with default headers values
func NewPostTokenDefault(code int) *PostTokenDefault {
	if code <= 0 {
		code = 500
	}

	return &PostTokenDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post token default response
func (o *PostTokenDefault) WithStatusCode(code int) *PostTokenDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post token default response
func (o *PostTokenDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post token default response
func (o *PostTokenDefault) WithPayload(payload *models.Error) *PostTokenDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post token default response
func (o *PostTokenDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTokenDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
