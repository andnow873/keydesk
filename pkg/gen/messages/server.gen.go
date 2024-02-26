// Package messages provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package messages

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get messages
	// (GET /messages)
	GetMessages(ctx echo.Context, params GetMessagesParams) error
	// Create message
	// (POST /messages)
	PostMessages(ctx echo.Context) error
	// Mark message as read
	// (POST /messages/{id}/read)
	PostMessagesIdRead(ctx echo.Context, id MessageID) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetMessages converts echo context to params.
func (w *ServerInterfaceWrapper) GetMessages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetMessagesParams
	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "read" -------------

	err = runtime.BindQueryParameter("form", true, false, "read", ctx.QueryParams(), &params.Read)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter read: %s", err))
	}

	// ------------- Optional query parameter "priority" -------------

	err = runtime.BindQueryParameter("deepObject", true, false, "priority", ctx.QueryParams(), &params.Priority)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter priority: %s", err))
	}

	// ------------- Optional query parameter "sort" -------------

	err = runtime.BindQueryParameter("form", true, false, "sort", ctx.QueryParams(), &params.Sort)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter sort: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMessages(ctx, params)
	return err
}

// PostMessages converts echo context to params.
func (w *ServerInterfaceWrapper) PostMessages(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostMessages(ctx)
	return err
}

// PostMessagesIdRead converts echo context to params.
func (w *ServerInterfaceWrapper) PostMessagesIdRead(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id MessageID

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: false})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostMessagesIdRead(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/messages", wrapper.GetMessages)
	router.POST(baseURL+"/messages", wrapper.PostMessages)
	router.POST(baseURL+"/messages/:id/read", wrapper.PostMessagesIdRead)

}

type GetMessagesRequestObject struct {
	Params GetMessagesParams
}

type GetMessagesResponseObject interface {
	VisitGetMessagesResponse(w http.ResponseWriter) error
}

type GetMessages200JSONResponse PaginatedMessages

func (response GetMessages200JSONResponse) VisitGetMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetMessagesdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response GetMessagesdefaultJSONResponse) VisitGetMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostMessagesRequestObject struct {
	Body *PostMessagesJSONRequestBody
}

type PostMessagesResponseObject interface {
	VisitPostMessagesResponse(w http.ResponseWriter) error
}

type PostMessages200JSONResponse Message

func (response PostMessages200JSONResponse) VisitPostMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostMessagesdefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response PostMessagesdefaultJSONResponse) VisitPostMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostMessagesIdReadRequestObject struct {
	Id MessageID `json:"id,omitempty"`
}

type PostMessagesIdReadResponseObject interface {
	VisitPostMessagesIdReadResponse(w http.ResponseWriter) error
}

type PostMessagesIdRead200Response struct {
}

func (response PostMessagesIdRead200Response) VisitPostMessagesIdReadResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type PostMessagesIdReaddefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response PostMessagesIdReaddefaultJSONResponse) VisitPostMessagesIdReadResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get messages
	// (GET /messages)
	GetMessages(ctx context.Context, request GetMessagesRequestObject) (GetMessagesResponseObject, error)
	// Create message
	// (POST /messages)
	PostMessages(ctx context.Context, request PostMessagesRequestObject) (PostMessagesResponseObject, error)
	// Mark message as read
	// (POST /messages/{id}/read)
	PostMessagesIdRead(ctx context.Context, request PostMessagesIdReadRequestObject) (PostMessagesIdReadResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetMessages operation middleware
func (sh *strictHandler) GetMessages(ctx echo.Context, params GetMessagesParams) error {
	var request GetMessagesRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetMessages(ctx.Request().Context(), request.(GetMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetMessagesResponseObject); ok {
		return validResponse.VisitGetMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostMessages operation middleware
func (sh *strictHandler) PostMessages(ctx echo.Context) error {
	var request PostMessagesRequestObject

	var body PostMessagesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostMessages(ctx.Request().Context(), request.(PostMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostMessagesResponseObject); ok {
		return validResponse.VisitPostMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostMessagesIdRead operation middleware
func (sh *strictHandler) PostMessagesIdRead(ctx echo.Context, id MessageID) error {
	var request PostMessagesIdReadRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostMessagesIdRead(ctx.Request().Context(), request.(PostMessagesIdReadRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostMessagesIdRead")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostMessagesIdReadResponseObject); ok {
		return validResponse.VisitPostMessagesIdReadResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
