// Copyright 2019 - 2022 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// IngredientManufacturerRole is documented here http://hl7.org/fhir/ValueSet/ingredient-manufacturer-role
type IngredientManufacturerRole int

const (
	IngredientManufacturerRoleAllowed IngredientManufacturerRole = iota
	IngredientManufacturerRolePossible
	IngredientManufacturerRoleActual
)

func (code IngredientManufacturerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *IngredientManufacturerRole) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "allowed":
		*code = IngredientManufacturerRoleAllowed
	case "possible":
		*code = IngredientManufacturerRolePossible
	case "actual":
		*code = IngredientManufacturerRoleActual
	default:
		return fmt.Errorf("unknown IngredientManufacturerRole code `%s`", s)
	}
	return nil
}
func (code IngredientManufacturerRole) String() string {
	return code.Code()
}
func (code IngredientManufacturerRole) Code() string {
	switch code {
	case IngredientManufacturerRoleAllowed:
		return "allowed"
	case IngredientManufacturerRolePossible:
		return "possible"
	case IngredientManufacturerRoleActual:
		return "actual"
	}
	return "<unknown>"
}
func (code IngredientManufacturerRole) Display() string {
	switch code {
	case IngredientManufacturerRoleAllowed:
		return "Manufacturer is specifically allowed for this ingredient"
	case IngredientManufacturerRolePossible:
		return "Manufacturer is known to make this ingredient in general"
	case IngredientManufacturerRoleActual:
		return "Manufacturer actually makes this particular ingredient"
	}
	return "<unknown>"
}
func (code IngredientManufacturerRole) Definition() string {
	switch code {
	}
	return "<unknown>"
}
