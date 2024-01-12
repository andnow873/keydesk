// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/vpngen/keydesk/gen/models"
)

// NewPutMessageParams creates a new PutMessageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutMessageParams() *PutMessageParams {
	return &PutMessageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutMessageParamsWithTimeout creates a new PutMessageParams object
// with the ability to set a timeout on a request.
func NewPutMessageParamsWithTimeout(timeout time.Duration) *PutMessageParams {
	return &PutMessageParams{
		timeout: timeout,
	}
}

// NewPutMessageParamsWithContext creates a new PutMessageParams object
// with the ability to set a context for a request.
func NewPutMessageParamsWithContext(ctx context.Context) *PutMessageParams {
	return &PutMessageParams{
		Context: ctx,
	}
}

// NewPutMessageParamsWithHTTPClient creates a new PutMessageParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutMessageParamsWithHTTPClient(client *http.Client) *PutMessageParams {
	return &PutMessageParams{
		HTTPClient: client,
	}
}

/*
PutMessageParams contains all the parameters to send to the API endpoint

	for the put message operation.

	Typically these are written to a http.Request.
*/
type PutMessageParams struct {

	/* Message.

	   The user to create.
	*/
	Message *models.Message

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutMessageParams) WithDefaults() *PutMessageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put message params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutMessageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put message params
func (o *PutMessageParams) WithTimeout(timeout time.Duration) *PutMessageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put message params
func (o *PutMessageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put message params
func (o *PutMessageParams) WithContext(ctx context.Context) *PutMessageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put message params
func (o *PutMessageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put message params
func (o *PutMessageParams) WithHTTPClient(client *http.Client) *PutMessageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put message params
func (o *PutMessageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessage adds the message to the put message params
func (o *PutMessageParams) WithMessage(message *models.Message) *PutMessageParams {
	o.SetMessage(message)
	return o
}

// SetMessage adds the message to the put message params
func (o *PutMessageParams) SetMessage(message *models.Message) {
	o.Message = message
}

// WriteToRequest writes these params to a swagger request
func (o *PutMessageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Message != nil {
		if err := r.SetBodyParam(o.Message); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
