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

// ConsentState is documented here http://hl7.org/fhir/ValueSet/consent-state-codes
type ConsentState int

const (
	ConsentStateDraft ConsentState = iota
	ConsentStateActive
	ConsentStateInactive
	ConsentStateNotDone
	ConsentStateEnteredInError
	ConsentStateUnknown
)

func (code ConsentState) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConsentState) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "draft":
		*code = ConsentStateDraft
	case "active":
		*code = ConsentStateActive
	case "inactive":
		*code = ConsentStateInactive
	case "not-done":
		*code = ConsentStateNotDone
	case "entered-in-error":
		*code = ConsentStateEnteredInError
	case "unknown":
		*code = ConsentStateUnknown
	default:
		return fmt.Errorf("unknown ConsentState code `%s`", s)
	}
	return nil
}
func (code ConsentState) String() string {
	return code.Code()
}
func (code ConsentState) Code() string {
	switch code {
	case ConsentStateDraft:
		return "draft"
	case ConsentStateActive:
		return "active"
	case ConsentStateInactive:
		return "inactive"
	case ConsentStateNotDone:
		return "not-done"
	case ConsentStateEnteredInError:
		return "entered-in-error"
	case ConsentStateUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code ConsentState) Display() string {
	switch code {
	case ConsentStateDraft:
		return "Pending"
	case ConsentStateActive:
		return "Active"
	case ConsentStateInactive:
		return "Inactive"
	case ConsentStateNotDone:
		return "Abandoned"
	case ConsentStateEnteredInError:
		return "Entered in Error"
	case ConsentStateUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code ConsentState) Definition() string {
	switch code {
	case ConsentStateDraft:
		return "The consent is in development or awaiting use but is not yet intended to be acted upon."
	case ConsentStateActive:
		return "The consent is to be followed and enforced."
	case ConsentStateInactive:
		return "The consent is terminated or replaced."
	case ConsentStateNotDone:
		return "The consent development has been terminated prior to completion."
	case ConsentStateEnteredInError:
		return "The consent was created wrongly (e.g. wrong patient) and should be ignored."
	case ConsentStateUnknown:
		return "The resource is in an indeterminate state."
	}
	return "<unknown>"
}
