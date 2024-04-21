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
	utils "smidgen-backend/go/utils"
)

type BusinessUnit struct {
	UnitId         int32  `json:"unit_id"`
	Name           string `json:"name"`
	PointOfContact string `json:"point_of_contact"`
	AddressLineOne string `json:"address_line_one"`
	AddressLineTwo string `json:"address_line_two"`
	State          string `json:"state"`
	City           string `json:"city"`
	Country        string `json:"country"`
}

func AssertBusinessUnitRequired(obj BusinessUnit) error {
	elements := map[string]interface{}{
		"name":             obj.Name,
		"point_of_contact": obj.PointOfContact,
		"address_line_one": obj.AddressLineOne,
		"address_line_two": obj.AddressLineTwo,
		"state":            obj.State,
		"city":             obj.City,
		"country":          obj.Country,
	}
	for name, el := range elements {
		if isZero := utils.IsZeroValue(el); isZero {
			return &utils.RequiredError{Field: name}
		}
	}

	return nil
}

func AssertBusinessUnitConstraints(obj BusinessUnit) error {
	return nil
}
