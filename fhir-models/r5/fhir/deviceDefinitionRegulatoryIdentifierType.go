// Copyright 2024 Medical Center University Freiburg
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

// THIS FILE IS GENERATED BY https://github.com/mcuf-idim/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DeviceDefinitionRegulatoryIdentifierType is documented here http://hl7.org/fhir/ValueSet/devicedefinition-regulatory-identifier-type
type DeviceDefinitionRegulatoryIdentifierType int

const (
	DeviceDefinitionRegulatoryIdentifierTypeBasic DeviceDefinitionRegulatoryIdentifierType = iota
	DeviceDefinitionRegulatoryIdentifierTypeMaster
	DeviceDefinitionRegulatoryIdentifierTypeLicense
)

func (code DeviceDefinitionRegulatoryIdentifierType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceDefinitionRegulatoryIdentifierType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "basic":
		*code = DeviceDefinitionRegulatoryIdentifierTypeBasic
	case "master":
		*code = DeviceDefinitionRegulatoryIdentifierTypeMaster
	case "license":
		*code = DeviceDefinitionRegulatoryIdentifierTypeLicense
	default:
		return fmt.Errorf("unknown DeviceDefinitionRegulatoryIdentifierType code `%s`", s)
	}
	return nil
}
func (code DeviceDefinitionRegulatoryIdentifierType) String() string {
	return code.Code()
}
func (code DeviceDefinitionRegulatoryIdentifierType) Code() string {
	switch code {
	case DeviceDefinitionRegulatoryIdentifierTypeBasic:
		return "basic"
	case DeviceDefinitionRegulatoryIdentifierTypeMaster:
		return "master"
	case DeviceDefinitionRegulatoryIdentifierTypeLicense:
		return "license"
	}
	return "<unknown>"
}
func (code DeviceDefinitionRegulatoryIdentifierType) Display() string {
	switch code {
	case DeviceDefinitionRegulatoryIdentifierTypeBasic:
		return "Basic"
	case DeviceDefinitionRegulatoryIdentifierTypeMaster:
		return "Master"
	case DeviceDefinitionRegulatoryIdentifierTypeLicense:
		return "License"
	}
	return "<unknown>"
}
func (code DeviceDefinitionRegulatoryIdentifierType) Definition() string {
	switch code {
	case DeviceDefinitionRegulatoryIdentifierTypeBasic:
		return "EUDAMED's basic UDI-DI identifier."
	case DeviceDefinitionRegulatoryIdentifierTypeMaster:
		return "EUDAMED's master UDI-DI identifier."
	case DeviceDefinitionRegulatoryIdentifierTypeLicense:
		return "The identifier is a license number."
	}
	return "<unknown>"
}
