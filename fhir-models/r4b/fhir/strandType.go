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

// StrandType is documented here http://hl7.org/fhir/ValueSet/strand-type
type StrandType int

const (
	StrandTypeWatson StrandType = iota
	StrandTypeCrick
)

func (code StrandType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StrandType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "watson":
		*code = StrandTypeWatson
	case "crick":
		*code = StrandTypeCrick
	default:
		return fmt.Errorf("unknown StrandType code `%s`", s)
	}
	return nil
}
func (code StrandType) String() string {
	return code.Code()
}
func (code StrandType) Code() string {
	switch code {
	case StrandTypeWatson:
		return "watson"
	case StrandTypeCrick:
		return "crick"
	}
	return "<unknown>"
}
func (code StrandType) Display() string {
	switch code {
	case StrandTypeWatson:
		return "Watson strand of starting sequence"
	case StrandTypeCrick:
		return "Crick strand of starting sequence"
	}
	return "<unknown>"
}
func (code StrandType) Definition() string {
	switch code {
	case StrandTypeWatson:
		return "Watson strand of starting sequence."
	case StrandTypeCrick:
		return "Crick strand of starting sequence."
	}
	return "<unknown>"
}