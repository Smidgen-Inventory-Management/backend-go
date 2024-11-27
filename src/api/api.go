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
	models "smidgen-backend/src/models"
	utils "smidgen-backend/src/utils"
)


type BusinessUnitAPIServicer interface {
	AddBusinessUnit(context.Context, models.BusinessUnit) (utils.ImplResponse, error)
	DeleteBusinessUnit(context.Context, int32) (utils.ImplResponse, error)
	GetBusinessUnits(context.Context) (utils.ImplResponse, error)
	GetBusinessUnitById(context.Context, int32) (utils.ImplResponse, error)
	UpdateBusinessUnit(context.Context, int32, models.BusinessUnit) (utils.ImplResponse, error)
}

type DefaultAPIServicer interface {
	HealthCheck(context.Context) (utils.ImplResponse, error)
	RootGet(context.Context) (utils.ImplResponse, error)
}

type EquipmentAPIServicer interface {
	AddEquipment(context.Context, models.Equipment) (utils.ImplResponse, error)
	DeleteEquipment(context.Context, int32) (utils.ImplResponse, error)
	GetEquipments(context.Context) (utils.ImplResponse, error)
	GetEquipmentById(context.Context, int32) (utils.ImplResponse, error)
	UpdateEquipment(context.Context, int32, models.Equipment) (utils.ImplResponse, error)
}

type ManufacturerAPIServicer interface {
	AddManufacturer(context.Context, models.Manufacturer) (utils.ImplResponse, error)
	DeleteManufacturer(context.Context, int32) (utils.ImplResponse, error)
	GetManufacturers(context.Context) (utils.ImplResponse, error)
	GetManufacturerById(context.Context, int32) (utils.ImplResponse, error)
	UpdateManufacturer(context.Context, int32, models.Manufacturer) (utils.ImplResponse, error)
}

type EquipmentAssignmentAPIServicer interface {
	AddEquipmentAssignment(context.Context, models.EquipmentAssignment) (utils.ImplResponse, error)
	DeleteEquipmentAssignment(context.Context, int32) (utils.ImplResponse, error)
	GetEquipmentAssignments(context.Context) (utils.ImplResponse, error)
	GetEquipmentAssignmentById(context.Context, int32) (utils.ImplResponse, error)
	UpdateEquipmentAssignment(context.Context, int32, models.EquipmentAssignment) (utils.ImplResponse, error)
}

type UserAPIServicer interface {
	AddUser(context.Context, models.User) (utils.ImplResponse, error)
	DeleteUser(context.Context, int32) (utils.ImplResponse, error)
	GetUsers(context.Context) (utils.ImplResponse, error)
	GetUserById(context.Context, int32) (utils.ImplResponse, error)
	UpdateUser(context.Context, int32, models.User) (utils.ImplResponse, error)
}

type AuditLogAPIServicer interface {
	GetAuditLogs(context.Context) (utils.ImplResponse, error)
	GetAuditLogById(context.Context, int32) (utils.ImplResponse, error)
}
