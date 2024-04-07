package smidgen

import (
	"encoding/json"
	"net/http"
	models "smidgen-backend/go/models"
	utils "smidgen-backend/go/utils"
	"strings"

	"github.com/gorilla/mux"
)

// BusinessUnitAPIController binds http requests to an api service and writes the service results to the http response
type BusinessUnitAPIController struct {
	service      BusinessUnitAPIServicer
	errorHandler utils.ErrorHandler
}

// BusinessUnitAPIOption for how the controller is set up.
type BusinessUnitAPIOption func(*BusinessUnitAPIController)

// WithBusinessUnitAPIErrorHandler inject ErrorHandler into controller
func WithBusinessUnitAPIErrorHandler(h utils.ErrorHandler) BusinessUnitAPIOption {
	return func(c *BusinessUnitAPIController) {
		c.errorHandler = h
	}
}

// NewBusinessUnitAPIController creates a default api controller
func NewBusinessUnitAPIController(s BusinessUnitAPIServicer, opts ...BusinessUnitAPIOption) utils.Router {
	controller := &BusinessUnitAPIController{
		service:      s,
		errorHandler: utils.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the BusinessUnitAPIController
func (c *BusinessUnitAPIController) Routes() utils.Routes {
	return utils.Routes{
		"AddUnitBusinessUnitPost": utils.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "business_unit",
			HandlerFunc: c.AddUnitBusinessUnitPost,
		},
		"DeleteUnitBusinessUnitUnitIdDelete": utils.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "business_unit/{unit_id}",
			HandlerFunc: c.DeleteUnitBusinessUnitUnitIdDelete,
		},
		"GetUnitBusinessUnitGet": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "business_unit/",
			HandlerFunc: c.GetUnitBusinessUnitGet,
		},
		"GetUnitsBusinessUnitUnitIdGet": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "business_unit/{unit_id}",
			HandlerFunc: c.GetUnitsBusinessUnitUnitIdGet,
		},
		"UpdateUnitBusinessUnitUnitIdPut": utils.Route{
			Method:      strings.ToUpper("Put"),
			Pattern:     "business_unit/{unit_id}",
			HandlerFunc: c.UpdateUnitBusinessUnitUnitIdPut,
		},
	}
}

// AddUnitBusinessUnitPost - Create Business Unit
func (c *BusinessUnitAPIController) AddUnitBusinessUnitPost(w http.ResponseWriter, r *http.Request) {
	businessUnitParam := models.BusinessUnit{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&businessUnitParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertBusinessUnitRequired(businessUnitParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertBusinessUnitConstraints(businessUnitParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddUnitBusinessUnitPost(r.Context(), businessUnitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteUnitBusinessUnitUnitIdDelete - Delete Business Unit
func (c *BusinessUnitAPIController) DeleteUnitBusinessUnitUnitIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	unitIdParam, err := utils.ParseNumericParameter[int32](
		params["unit_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteUnitBusinessUnitUnitIdDelete(r.Context(), unitIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetUnitBusinessUnitGet - Get Business Units
func (c *BusinessUnitAPIController) GetUnitBusinessUnitGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetUnitBusinessUnitGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetUnitsBusinessUnitUnitIdGet - Get Business Unit
func (c *BusinessUnitAPIController) GetUnitsBusinessUnitUnitIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	unitIdParam, err := utils.ParseNumericParameter[int32](
		params["unit_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetUnitsBusinessUnitUnitIdGet(r.Context(), unitIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateUnitBusinessUnitUnitIdPut - Update Business Unit
func (c *BusinessUnitAPIController) UpdateUnitBusinessUnitUnitIdPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	unitIdParam, err := utils.ParseNumericParameter[int32](
		params["unit_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	businessUnitParam := models.BusinessUnit{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&businessUnitParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertBusinessUnitRequired(businessUnitParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertBusinessUnitConstraints(businessUnitParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateUnitBusinessUnitUnitIdPut(r.Context(), unitIdParam, businessUnitParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}
