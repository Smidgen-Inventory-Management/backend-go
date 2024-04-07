package smidgen

import (
	"context"
	"errors"
	"fmt"
	"log"
	api "smidgen-backend/go/api"
	models "smidgen-backend/go/models"
	utils "smidgen-backend/go/utils"
)

// EquipmentAssignmentAPIService is a service that implements the logic for the EquipmentAssignmentAPIServicer
// This service should implement the business logic for every endpoint for the EquipmentAssignmentAPI API.
// Include any external packages or services that will be required by this service.
type EquipmentAssignmentAPIService struct {
}

// NewEquipmentAssignmentAPIService creates a default api service
func NewEquipmentAssignmentAPIService() api.EquipmentAssignmentAPIServicer {
	return &EquipmentAssignmentAPIService{}
}

// AddAssignmentEquipmentAssignmentPost - Create assignment
func (s *EquipmentAssignmentAPIService) AddAssignmentEquipmentAssignmentPost(ctx context.Context, equipmentAssignment models.EquipmentAssignment) (utils.ImplResponse, error) {
	privilege := "write"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.InsertRow("equipmentAssignment", equipmentAssignment)
	if err != nil {
		fmt.Println(err)
		return utils.Response(500, nil), errors.New("an error has occured while adding new data")
	}
	return utils.Response(202, nil), nil
}

// DeleteAssignmentEquipmentAssignmentAssignmentIdDelete - Delete assignment
func (s *EquipmentAssignmentAPIService) DeleteAssignmentEquipmentAssignmentAssignmentIdDelete(ctx context.Context, assignmentId int32) (utils.ImplResponse, error) {
	privilege := "delete"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.DeleteRow("equipment", "equipmentid", assignmentId)
	if err != nil {
		fmt.Println("Error:", err)
		return utils.Response(500, nil), err
	}

	return utils.Response(200, nil), nil
}

// GetAssignmentEquipmentAssignmentGet - Get assignments
func (s *EquipmentAssignmentAPIService) GetAssignmentEquipmentAssignmentGet(ctx context.Context) (utils.ImplResponse, error) {
	// Add api_user_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	privilege := "read"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.EquipmentAssignment
	rows, err := dbConnection.GetRows("equipmentAssignment", &dest)
	if err != nil {
		fmt.Println("Error:", err)

	}

	var Assignments []models.EquipmentAssignment
	for _, row := range rows {
		Assignment, ok := row.(models.EquipmentAssignment)
		if !ok {
			fmt.Println("Error: Unexpected type in row")
			continue
		}
		Assignments = append(Assignments, Assignment)
	}

	return utils.Response(200, Assignments), nil
}

// GetAssignmentsEquipmentAssignmentAssignmentIdGet - Get assignment
func (s *EquipmentAssignmentAPIService) GetAssignmentsEquipmentAssignmentAssignmentIdGet(ctx context.Context, assignmentId int32) (utils.ImplResponse, error) {
	privilege := "read"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	var dest models.Equipment
	row, err := dbConnection.GetByID("equipmetnAssignment", "assignmentId", assignmentId, &dest)
	if err != nil {
		fmt.Println("Data Not Found:", err)
		return utils.Response(404, nil), fmt.Errorf("the requested ID was not found")
	}

	assignment, ok := row.(models.EquipmentAssignment)
	if !ok {
		fmt.Println("Error: Unexpected type in row")
		return utils.Response(500, nil), errors.New("unexpected type in row")
	}
	return utils.Response(200, assignment), nil
}

// UpdateAssignmentEquipmentAssignmentAssignmentIdPut - Update assignment
func (s *EquipmentAssignmentAPIService) UpdateAssignmentEquipmentAssignmentAssignmentIdPut(ctx context.Context, assignmentId int32, equipmentAssignment models.EquipmentAssignment) (utils.ImplResponse, error) {
	privilege := "write"
	dbConnection, err := utils.NewDatabaseConnection(privilege)
	if err != nil {
		log.Fatalf("Failed to establish database connection as %s: %v", privilege, err)
	}

	err = dbConnection.UpdateRow("equipmentAssignment", "assignmentid", assignmentId, equipmentAssignment)
	if err != nil {
		fmt.Println(err)
		return utils.Response(500, nil), errors.New("an error has occured while updating the data")
	}
	return utils.Response(202, nil), nil
}
