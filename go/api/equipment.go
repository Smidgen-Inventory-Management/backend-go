package smidgen

import (
	"encoding/json"
	"net/http"
	models "smidgen-backend/go/models"
	utils "smidgen-backend/go/utils"
	"strings"

	"github.com/gorilla/mux"
)

// EquipmentAPIController binds http requests to an api service and writes the service results to the http response
type EquipmentAPIController struct {
	service      EquipmentAPIServicer
	errorHandler utils.ErrorHandler
}

// EquipmentAPIOption for how the controller is set up.
type EquipmentAPIOption func(*EquipmentAPIController)

// WithEquipmentAPIErrorHandler inject ErrorHandler into controller
func WithEquipmentAPIErrorHandler(h utils.ErrorHandler) EquipmentAPIOption {
	return func(c *EquipmentAPIController) {
		c.errorHandler = h
	}
}

// NewEquipmentAPIController creates a default api controller
func NewEquipmentAPIController(s EquipmentAPIServicer, opts ...EquipmentAPIOption) utils.Router {
	controller := &EquipmentAPIController{
		service:      s,
		errorHandler: utils.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the EquipmentAPIController
func (c *EquipmentAPIController) Routes() utils.Routes {
	return utils.Routes{
		"AddEquipmentEquipmentPost": utils.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "equipment",
			HandlerFunc: c.AddEquipmentEquipmentPost,
		},
		"DeleteEquipmentEquipmentEquipmentIdDelete": utils.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "equipment/{equipment_id}",
			HandlerFunc: c.DeleteEquipmentEquipmentEquipmentIdDelete,
		},
		"GetEquipmentEquipmentGet": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "equipment/",
			HandlerFunc: c.GetEquipmentEquipmentGet,
		},
		"GetEquipmentsEquipmentEquipmentIdGet": utils.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "equipment/{equipment_id}",
			HandlerFunc: c.GetEquipmentsEquipmentEquipmentIdGet,
		},
		"UpdateEquipmentEquipmentEquipmentIdPut": utils.Route{
			Method:      strings.ToUpper("Put"),
			Pattern:     "equipment/{equipment_id}",
			HandlerFunc: c.UpdateEquipmentEquipmentEquipmentIdPut,
		},
	}
}

// AddEquipmentEquipmentPost - Create equipment
func (c *EquipmentAPIController) AddEquipmentEquipmentPost(w http.ResponseWriter, r *http.Request) {
	equipmentParam := models.Equipment{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&equipmentParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertEquipmentRequired(equipmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertEquipmentConstraints(equipmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddEquipmentEquipmentPost(r.Context(), equipmentParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteEquipmentEquipmentEquipmentIdDelete - Delete equipment
func (c *EquipmentAPIController) DeleteEquipmentEquipmentEquipmentIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	equipmentIdParam, err := utils.ParseNumericParameter[int32](
		params["equipment_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteEquipmentEquipmentEquipmentIdDelete(r.Context(), equipmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetEquipmentEquipmentGet - Get equipments
func (c *EquipmentAPIController) GetEquipmentEquipmentGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetEquipmentEquipmentGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetEquipmentsEquipmentEquipmentIdGet - Get equipment
func (c *EquipmentAPIController) GetEquipmentsEquipmentEquipmentIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	equipmentIdParam, err := utils.ParseNumericParameter[int32](
		params["equipment_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetEquipmentsEquipmentEquipmentIdGet(r.Context(), equipmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateEquipmentEquipmentEquipmentIdPut - Update equipment
func (c *EquipmentAPIController) UpdateEquipmentEquipmentEquipmentIdPut(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	equipmentIdParam, err := utils.ParseNumericParameter[int32](
		params["equipment_id"],
		utils.WithRequire[int32](utils.ParseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	equipmentParam := models.Equipment{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&equipmentParam); err != nil {
		c.errorHandler(w, r, &utils.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertEquipmentRequired(equipmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertEquipmentConstraints(equipmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateEquipmentEquipmentEquipmentIdPut(r.Context(), equipmentIdParam, equipmentParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	utils.EncodeJSONResponse(result.Body, &result.Code, w)
}
