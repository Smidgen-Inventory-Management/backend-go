/*
 * Smidgen
 *
 * API for interacting with Smidgen.
 *
 *   Smidgen aims to simplify and automate common tasks that logisticians
 *   conduct on a daily basis so they can focus on the effective distribution
 *   of materiel, as well as maintain an accurate record keeping book of
 *   receiving, issuance, audits, surpluses, amongst other logistical tasks.
 *   Copyright (C) 2024  Jose Hernandez
 *
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package smidgen

import (
	"context"
	api "smidgen-backend/go/api"
	utils "smidgen-backend/go/utils"
	"time"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
}

type healthCheck struct {
	Service string        `json:"service"`
	Status  string        `json:"status"`
	Latency time.Duration `json:"latency"`
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService() api.DefaultAPIServicer {
	return &DefaultAPIService{}
}

// HealthCheck - Check
func (s *DefaultAPIService) HealthCheck(ctx context.Context) (utils.ImplResponse, error) {
	log.Debug("checking status of core Smidgen services")
	healthcheckStart := time.Now()
	var services []healthCheck
	db, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "read")
	if err != nil {
		log.Errorf("failed to create database connection: %v", err)
		healthcheckEnd := time.Since(healthcheckStart)
		services = append(services, healthCheck{"API Server", "OK", healthcheckEnd})
		services = append(services, healthCheck{"Database", "DOWN", -1})
		return utils.Response(500, services), nil
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	if db == nil {
		log.Errorf("database connection is nil")
		healthcheckEnd := time.Since(healthcheckStart)
		services = append(services, healthCheck{"API Server", "OK", healthcheckEnd})
		services = append(services, healthCheck{"Database", "DEGREDADED", -1})
		return utils.Response(500, services), nil
	}

	start := time.Now()
	err = db.Ping()
	latency := time.Since(start)

	if err != nil {
		log.Errorf("failed to ping database: %v", err)
		healthcheckEnd := time.Since(healthcheckStart)
		services = append(services, healthCheck{"API Server", "OK", healthcheckEnd})
		services = append(services, healthCheck{"Database", "DEGREDADED", -1})
		return utils.Response(500, services), nil
	}

	log.Infof("connection to database successfully established in %s", latency)

	healthcheckEnd := time.Since(healthcheckStart)
	services = append(services, healthCheck{"API Server", "OK", healthcheckEnd})
	services = append(services, healthCheck{"Database", "OK", latency})
	return utils.Response(200, services), nil
}

// RootGet - Root
func (s *DefaultAPIService) RootGet(ctx context.Context) (utils.ImplResponse, error) {
	return utils.Response(403, nil), nil
}
