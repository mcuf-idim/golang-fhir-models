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

// AssertionOperatorType is documented here http://hl7.org/fhir/ValueSet/assert-operator-codes
type AssertionOperatorType int

const (
	AssertionOperatorTypeEquals AssertionOperatorType = iota
	AssertionOperatorTypeNotEquals
	AssertionOperatorTypeIn
	AssertionOperatorTypeNotIn
	AssertionOperatorTypeGreaterThan
	AssertionOperatorTypeLessThan
	AssertionOperatorTypeEmpty
	AssertionOperatorTypeNotEmpty
	AssertionOperatorTypeContains
	AssertionOperatorTypeNotContains
	AssertionOperatorTypeEval
	AssertionOperatorTypeManualEval
)

func (code AssertionOperatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AssertionOperatorType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "equals":
		*code = AssertionOperatorTypeEquals
	case "notEquals":
		*code = AssertionOperatorTypeNotEquals
	case "in":
		*code = AssertionOperatorTypeIn
	case "notIn":
		*code = AssertionOperatorTypeNotIn
	case "greaterThan":
		*code = AssertionOperatorTypeGreaterThan
	case "lessThan":
		*code = AssertionOperatorTypeLessThan
	case "empty":
		*code = AssertionOperatorTypeEmpty
	case "notEmpty":
		*code = AssertionOperatorTypeNotEmpty
	case "contains":
		*code = AssertionOperatorTypeContains
	case "notContains":
		*code = AssertionOperatorTypeNotContains
	case "eval":
		*code = AssertionOperatorTypeEval
	case "manualEval":
		*code = AssertionOperatorTypeManualEval
	default:
		return fmt.Errorf("unknown AssertionOperatorType code `%s`", s)
	}
	return nil
}
func (code AssertionOperatorType) String() string {
	return code.Code()
}
func (code AssertionOperatorType) Code() string {
	switch code {
	case AssertionOperatorTypeEquals:
		return "equals"
	case AssertionOperatorTypeNotEquals:
		return "notEquals"
	case AssertionOperatorTypeIn:
		return "in"
	case AssertionOperatorTypeNotIn:
		return "notIn"
	case AssertionOperatorTypeGreaterThan:
		return "greaterThan"
	case AssertionOperatorTypeLessThan:
		return "lessThan"
	case AssertionOperatorTypeEmpty:
		return "empty"
	case AssertionOperatorTypeNotEmpty:
		return "notEmpty"
	case AssertionOperatorTypeContains:
		return "contains"
	case AssertionOperatorTypeNotContains:
		return "notContains"
	case AssertionOperatorTypeEval:
		return "eval"
	case AssertionOperatorTypeManualEval:
		return "manualEval"
	}
	return "<unknown>"
}
func (code AssertionOperatorType) Display() string {
	switch code {
	case AssertionOperatorTypeEquals:
		return "equals"
	case AssertionOperatorTypeNotEquals:
		return "notEquals"
	case AssertionOperatorTypeIn:
		return "in"
	case AssertionOperatorTypeNotIn:
		return "notIn"
	case AssertionOperatorTypeGreaterThan:
		return "greaterThan"
	case AssertionOperatorTypeLessThan:
		return "lessThan"
	case AssertionOperatorTypeEmpty:
		return "empty"
	case AssertionOperatorTypeNotEmpty:
		return "notEmpty"
	case AssertionOperatorTypeContains:
		return "contains"
	case AssertionOperatorTypeNotContains:
		return "notContains"
	case AssertionOperatorTypeEval:
		return "evaluate"
	case AssertionOperatorTypeManualEval:
		return "manualEvaluate"
	}
	return "<unknown>"
}
func (code AssertionOperatorType) Definition() string {
	switch code {
	case AssertionOperatorTypeEquals:
		return "Default value. Equals comparison."
	case AssertionOperatorTypeNotEquals:
		return "Not equals comparison."
	case AssertionOperatorTypeIn:
		return "Compare value within a known set of values."
	case AssertionOperatorTypeNotIn:
		return "Compare value not within a known set of values."
	case AssertionOperatorTypeGreaterThan:
		return "Compare value to be greater than a known value."
	case AssertionOperatorTypeLessThan:
		return "Compare value to be less than a known value."
	case AssertionOperatorTypeEmpty:
		return "Compare value is empty."
	case AssertionOperatorTypeNotEmpty:
		return "Compare value is not empty."
	case AssertionOperatorTypeContains:
		return "Compare value string contains a known value."
	case AssertionOperatorTypeNotContains:
		return "Compare value string does not contain a known value."
	case AssertionOperatorTypeEval:
		return "Evaluate the FHIRPath expression as a boolean condition."
	case AssertionOperatorTypeManualEval:
		return "Manually evaluate the condition described by this assert. The test engine SHALL pause and provide an input mechanism to set the outcome of this assert to 'pass', 'fail', 'skip' or 'stop'."
	}
	return "<unknown>"
}
