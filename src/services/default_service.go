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
	"fmt"
	api "smidgen-backend/src/api"
	utils "smidgen-backend/src/utils"
	"time"
)

type DefaultAPIService struct {
}

type healthCheck struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Latency string `json:"latency"`
}


func NewDefaultAPIService() api.DefaultAPIServicer {
	return &DefaultAPIService{}
}

func (s *DefaultAPIService) HealthCheck(ctx context.Context) (utils.ImplResponse, error) {
	log.Debug("checking status of core Smidgen services")
	healthcheckStart := time.Now()
	var services []healthCheck
	db, err := utils.NewDatabaseConnection(utils.DatabaseConfigPath, "read")
	if err != nil {
		log.Errorf("failed to create database connection: %v", err)
		healthcheckEnd := time.Since(healthcheckStart).Milliseconds()
		services = append(services, healthCheck{"Overall", "DEGRADED", "DEGRADED"})
		services = append(services, healthCheck{"API Server", "OK", fmt.Sprintf("%dms", healthcheckEnd)})
		services = append(services, healthCheck{"Database", "DOWN", "DOWN"})
		return utils.Response(500, services), nil
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	if db == nil {
		log.Errorf("database connection is nil")
		healthcheckEnd := time.Since(healthcheckStart).Milliseconds()
		services = append(services, healthCheck{"Overall", "DEGRADED", "DEGRADED"})
		services = append(services, healthCheck{"API Server", "OK", fmt.Sprintf("%dms", healthcheckEnd)})
		services = append(services, healthCheck{"Database", "DEGRADED", "DEGRADED"})
		return utils.Response(500, services), nil
	}

	start := time.Now()
	err = db.Ping()
	databaseLatency := time.Since(start).Milliseconds()

	if err != nil {
		log.Errorf("failed to ping database: %v", err)
		healthcheckEnd := time.Since(healthcheckStart).Milliseconds()
		services = append(services, healthCheck{"Overall", "DEGRADED", "DEGRADED"})
		services = append(services, healthCheck{"API Server", "OK", fmt.Sprintf("%dms", healthcheckEnd)})
		services = append(services, healthCheck{"Database", "DEGRADED", "DEGRADED"})
		return utils.Response(500, services), nil
	}

	log.Infof("connection to database successfully established in %d", databaseLatency)

	healthcheckEnd := time.Since(healthcheckStart).Milliseconds()
	totalLatency := healthcheckEnd + databaseLatency
	services = append(services, healthCheck{"Overall", "OK", fmt.Sprintf("%dms", totalLatency)})
	services = append(services, healthCheck{"API Server", "OK", fmt.Sprintf("%dms", healthcheckEnd)})
	services = append(services, healthCheck{"Database", "OK", fmt.Sprintf("%dms", databaseLatency)})
	return utils.Response(200, services), nil
}

func (s *DefaultAPIService) RootGet(ctx context.Context) (utils.ImplResponse, error) {
	return utils.Response(403, nil), nil
}
