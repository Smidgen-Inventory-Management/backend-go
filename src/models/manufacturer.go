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
	utils "smidgen-backend/src/utils"
	"time"
)

type Manufacturer struct {
	ManufacturerId int32     `json:"manufacturer_id"`
	Name           string    `json:"name"`
	PrimaryService string    `json:"primary_service"`
	PointOfContact string    `json:"point_of_contact"`
	Location       string    `json:"location"`
	DateAdded      time.Time `json:"date_added"`
}

func AssertManufacturerRequired(obj Manufacturer) error {
	elements := map[string]interface{}{
		"manufacturer_id":  obj.ManufacturerId,
		"name":             obj.Name,
		"primary_service":  obj.PrimaryService,
		"point_of_contact": obj.PointOfContact,
		"location":         obj.Location,
		"date_added":       obj.DateAdded,
	}
	for name, el := range elements {
		if isZero := utils.IsZeroValue(el); isZero {
			return &utils.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertEquipmentConstraints checks if the values respects the defined constraints
func AssertManufacturerConstraints(obj Manufacturer) error {
	return nil
}
