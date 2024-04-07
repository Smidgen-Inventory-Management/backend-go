package smidgen

import (
	"net/http"
	utils "smidgen-backend/go/utils"
	"strings"
)

// DefaultAPIController binds http requests to an api service and writes the service results to the http response
type DefaultAPIController struct {
	service      DefaultAPIServicer
	errorHandler utils.ErrorHandler
}

// DefaultAPIOption for how the controller is set up.
type DefaultAPIOption func(*DefaultAPIController)

// WithDefaultAPIErrorHandler inject ErrorHandler into controller
func WithDefaultAPIErrorHandler(h utils.ErrorHandler) DefaultAPIOption {
	return func(c *DefaultAPIController) {
		c.errorHandler = h
	}
}

// NewDefaultAPIController creates a default api controller
func NewDefaultAPIController(s DefaultAPIServicer, opts ...DefaultAPIOption) utils.Router {
	controller := &DefaultAPIController{
		service:      s,
		errorHandler: utils.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultAPIController
func (c *DefaultAPIController) Routes() utils.Routes {
	return utils.Routes{
		"CheckHealthcheckGet": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "healthcheck",
			HandlerFunc: c.CheckHealthcheckGet,
		},
		"RootGet": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "",
			HandlerFunc: c.RootGet,
		},
	}
}

// CheckHealthcheckGet - Check
func (c *DefaultAPIController) CheckHealthcheckGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.CheckHealthcheckGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// RootGet - Root
func (c *DefaultAPIController) RootGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.RootGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}
