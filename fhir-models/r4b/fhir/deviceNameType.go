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

// DeviceNameType is documented here http://hl7.org/fhir/ValueSet/device-nametype
type DeviceNameType int

const (
	DeviceNameTypeRegisteredName DeviceNameType = iota
	DeviceNameTypeUserFriendlyName
	DeviceNameTypePatientReportedName
)

func (code DeviceNameType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceNameType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "registered-name":
		*code = DeviceNameTypeRegisteredName
	case "user-friendly-name":
		*code = DeviceNameTypeUserFriendlyName
	case "patient-reported-name":
		*code = DeviceNameTypePatientReportedName
	default:
		return fmt.Errorf("unknown DeviceNameType code `%s`", s)
	}
	return nil
}
func (code DeviceNameType) String() string {
	return code.Code()
}
func (code DeviceNameType) Code() string {
	switch code {
	case DeviceNameTypeRegisteredName:
		return "registered-name"
	case DeviceNameTypeUserFriendlyName:
		return "user-friendly-name"
	case DeviceNameTypePatientReportedName:
		return "patient-reported-name"
	}
	return "<unknown>"
}
func (code DeviceNameType) Display() string {
	switch code {
	case DeviceNameTypeRegisteredName:
		return "Registered name"
	case DeviceNameTypeUserFriendlyName:
		return "User Friendly name"
	case DeviceNameTypePatientReportedName:
		return "Patient Reported name"
	}
	return "<unknown>"
}
func (code DeviceNameType) Definition() string {
	switch code {
	case DeviceNameTypeRegisteredName:
		return "The term assigned to a medical device by the entity who registers or submits information about it to a jurisdiction or its databases. This may be considered the manufacturer assigned name (e.g., brand name assigned by the labeler or manufacturer in US, or device name assigned by the manufacturer in EU) and may also be synonymous with proprietary name or trade name of the device."
	case DeviceNameTypeUserFriendlyName:
		return "The term that generically describes the device by a name as assigned by the manufacturer that is recognized by lay person.  This common or generic name may be printed on the package it came in or some combination of that name with the model number, serial number, or other attribute that makes the name easy to understand for the user of that device. It is often exposed in communicating devices transport protocols. It is provided to help users identify the device when reported in discovery operations."
	case DeviceNameTypePatientReportedName:
		return "the term used by the patient associated with the device when describing the device, for example 'knee implant', when documented as a self-reported device."
	}
	return "<unknown>"
}