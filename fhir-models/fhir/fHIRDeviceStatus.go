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

// FHIRDeviceStatus is documented here http://hl7.org/fhir/ValueSet/device-status
type FHIRDeviceStatus int

const (
	FHIRDeviceStatusActive FHIRDeviceStatus = iota
	FHIRDeviceStatusInactive
	FHIRDeviceStatusEnteredInError
)

func (code FHIRDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FHIRDeviceStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = FHIRDeviceStatusActive
	case "inactive":
		*code = FHIRDeviceStatusInactive
	case "entered-in-error":
		*code = FHIRDeviceStatusEnteredInError
	default:
		return fmt.Errorf("unknown FHIRDeviceStatus code `%s`", s)
	}
	return nil
}
func (code FHIRDeviceStatus) String() string {
	return code.Code()
}
func (code FHIRDeviceStatus) Code() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "active"
	case FHIRDeviceStatusInactive:
		return "inactive"
	case FHIRDeviceStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code FHIRDeviceStatus) Display() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "Active"
	case FHIRDeviceStatusInactive:
		return "Inactive"
	case FHIRDeviceStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code FHIRDeviceStatus) Definition() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "The device record is current and is appropriate for reference in new instances."
	case FHIRDeviceStatusInactive:
		return "The device record is not current and is not appropriate for reference in new instances."
	case FHIRDeviceStatusEnteredInError:
		return "The device record is not current and is not appropriate for reference in new instances."
	}
	return "<unknown>"
}
