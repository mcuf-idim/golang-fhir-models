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

// ExampleScenarioActorType is documented here http://hl7.org/fhir/ValueSet/examplescenario-actor-type
type ExampleScenarioActorType int

const (
	ExampleScenarioActorTypePerson ExampleScenarioActorType = iota
	ExampleScenarioActorTypeSystem
)

func (code ExampleScenarioActorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ExampleScenarioActorType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "person":
		*code = ExampleScenarioActorTypePerson
	case "system":
		*code = ExampleScenarioActorTypeSystem
	default:
		return fmt.Errorf("unknown ExampleScenarioActorType code `%s`", s)
	}
	return nil
}
func (code ExampleScenarioActorType) String() string {
	return code.Code()
}
func (code ExampleScenarioActorType) Code() string {
	switch code {
	case ExampleScenarioActorTypePerson:
		return "person"
	case ExampleScenarioActorTypeSystem:
		return "system"
	}
	return "<unknown>"
}
func (code ExampleScenarioActorType) Display() string {
	switch code {
	case ExampleScenarioActorTypePerson:
		return "Person"
	case ExampleScenarioActorTypeSystem:
		return "System"
	}
	return "<unknown>"
}
func (code ExampleScenarioActorType) Definition() string {
	switch code {
	case ExampleScenarioActorTypePerson:
		return "A human actor"
	case ExampleScenarioActorTypeSystem:
		return "A software application or other system"
	}
	return "<unknown>"
}
