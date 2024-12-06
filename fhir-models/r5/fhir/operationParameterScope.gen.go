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

// OperationParameterScope is documented here http://hl7.org/fhir/ValueSet/operation-parameter-scope
type OperationParameterScope int

const (
	OperationParameterScopeInstance OperationParameterScope = iota
	OperationParameterScopeType
	OperationParameterScopeSystem
)

func (code OperationParameterScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *OperationParameterScope) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "instance":
		*code = OperationParameterScopeInstance
	case "type":
		*code = OperationParameterScopeType
	case "system":
		*code = OperationParameterScopeSystem
	default:
		return fmt.Errorf("unknown OperationParameterScope code `%s`", s)
	}
	return nil
}
func (code OperationParameterScope) String() string {
	return code.Code()
}
func (code OperationParameterScope) Code() string {
	switch code {
	case OperationParameterScopeInstance:
		return "instance"
	case OperationParameterScopeType:
		return "type"
	case OperationParameterScopeSystem:
		return "system"
	}
	return "<unknown>"
}
func (code OperationParameterScope) Display() string {
	switch code {
	case OperationParameterScopeInstance:
		return "Instance"
	case OperationParameterScopeType:
		return "Type"
	case OperationParameterScopeSystem:
		return "System"
	}
	return "<unknown>"
}
func (code OperationParameterScope) Definition() string {
	switch code {
	case OperationParameterScopeInstance:
		return "This is a parameter that can be used at the instance level."
	case OperationParameterScopeType:
		return "This is a parameter that can be used at the type level."
	case OperationParameterScopeSystem:
		return "This is a parameter that can be used at the system level."
	}
	return "<unknown>"
}