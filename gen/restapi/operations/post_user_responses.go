// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"test/gen/models"
)

// PostUserCreatedCode is the HTTP code returned for type PostUserCreated
const PostUserCreatedCode int = 201

/*PostUserCreated New user created.

swagger:response postUserCreated
*/
type PostUserCreated struct {
	/*the value is `attachment; filename="wg.conf"`

	 */
	ContentDisposition string `json:"Content-Disposition"`

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewPostUserCreated creates PostUserCreated with default headers values
func NewPostUserCreated() *PostUserCreated {

	return &PostUserCreated{}
}

// WithContentDisposition adds the contentDisposition to the post user created response
func (o *PostUserCreated) WithContentDisposition(contentDisposition string) *PostUserCreated {
	o.ContentDisposition = contentDisposition
	return o
}

// SetContentDisposition sets the contentDisposition to the post user created response
func (o *PostUserCreated) SetContentDisposition(contentDisposition string) {
	o.ContentDisposition = contentDisposition
}

// WithPayload adds the payload to the post user created response
func (o *PostUserCreated) WithPayload(payload io.ReadCloser) *PostUserCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user created response
func (o *PostUserCreated) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Disposition

	contentDisposition := o.ContentDisposition
	if contentDisposition != "" {
		rw.Header().Set("Content-Disposition", contentDisposition)
	}

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostUserForbiddenCode is the HTTP code returned for type PostUserForbidden
const PostUserForbiddenCode int = 403

/*PostUserForbidden You do not have necessary permissions for the resource

swagger:response postUserForbidden
*/
type PostUserForbidden struct {
}

// NewPostUserForbidden creates PostUserForbidden with default headers values
func NewPostUserForbidden() *PostUserForbidden {

	return &PostUserForbidden{}
}

// WriteResponse to the client
func (o *PostUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

/*PostUserDefault error

swagger:response postUserDefault
*/
type PostUserDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostUserDefault creates PostUserDefault with default headers values
func NewPostUserDefault(code int) *PostUserDefault {
	if code <= 0 {
		code = 500
	}

	return &PostUserDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post user default response
func (o *PostUserDefault) WithStatusCode(code int) *PostUserDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post user default response
func (o *PostUserDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post user default response
func (o *PostUserDefault) WithPayload(payload *models.Error) *PostUserDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user default response
func (o *PostUserDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
