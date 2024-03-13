// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetMessagesHandlerFunc turns a function with the right signature into a get messages handler
type GetMessagesHandlerFunc func(GetMessagesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMessagesHandlerFunc) Handle(params GetMessagesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetMessagesHandler interface for that can handle valid get messages params
type GetMessagesHandler interface {
	Handle(GetMessagesParams, interface{}) middleware.Responder
}

// NewGetMessages creates a new http.Handler for the get messages operation
func NewGetMessages(ctx *middleware.Context, handler GetMessagesHandler) *GetMessages {
	return &GetMessages{Context: ctx, Handler: handler}
}

/*
	GetMessages swagger:route GET /messages getMessages

# Get messages

Get messages, triggered by frontend
*/
type GetMessages struct {
	Context *middleware.Context
	Handler GetMessagesHandler
}

func (o *GetMessages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMessagesParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
