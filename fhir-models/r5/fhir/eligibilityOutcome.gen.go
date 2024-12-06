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

// EligibilityOutcome is documented here http://hl7.org/fhir/ValueSet/eligibility-outcome
type EligibilityOutcome int

const (
	EligibilityOutcomeQueued EligibilityOutcome = iota
	EligibilityOutcomeComplete
	EligibilityOutcomeError
	EligibilityOutcomePartial
)

func (code EligibilityOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EligibilityOutcome) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "queued":
		*code = EligibilityOutcomeQueued
	case "complete":
		*code = EligibilityOutcomeComplete
	case "error":
		*code = EligibilityOutcomeError
	case "partial":
		*code = EligibilityOutcomePartial
	default:
		return fmt.Errorf("unknown EligibilityOutcome code `%s`", s)
	}
	return nil
}
func (code EligibilityOutcome) String() string {
	return code.Code()
}
func (code EligibilityOutcome) Code() string {
	switch code {
	case EligibilityOutcomeQueued:
		return "queued"
	case EligibilityOutcomeComplete:
		return "complete"
	case EligibilityOutcomeError:
		return "error"
	case EligibilityOutcomePartial:
		return "partial"
	}
	return "<unknown>"
}
func (code EligibilityOutcome) Display() string {
	switch code {
	case EligibilityOutcomeQueued:
		return "Queued"
	case EligibilityOutcomeComplete:
		return "Processing Complete"
	case EligibilityOutcomeError:
		return "Error"
	case EligibilityOutcomePartial:
		return "Partial Processing"
	}
	return "<unknown>"
}
func (code EligibilityOutcome) Definition() string {
	switch code {
	case EligibilityOutcomeQueued:
		return "The Claim/Pre-authorization/Pre-determination has been received but processing has not begun."
	case EligibilityOutcomeComplete:
		return "The processing has completed without errors"
	case EligibilityOutcomeError:
		return "One or more errors have been detected in the Claim"
	case EligibilityOutcomePartial:
		return "No errors have been detected in the Claim and some of the adjudication has been performed."
	}
	return "<unknown>"
}