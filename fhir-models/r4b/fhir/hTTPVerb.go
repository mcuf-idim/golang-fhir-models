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

// HTTPVerb is documented here http://hl7.org/fhir/ValueSet/http-verb
type HTTPVerb int

const (
	HTTPVerbGET HTTPVerb = iota
	HTTPVerbHEAD
	HTTPVerbPOST
	HTTPVerbPUT
	HTTPVerbDELETE
	HTTPVerbPATCH
)

func (code HTTPVerb) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *HTTPVerb) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "GET":
		*code = HTTPVerbGET
	case "HEAD":
		*code = HTTPVerbHEAD
	case "POST":
		*code = HTTPVerbPOST
	case "PUT":
		*code = HTTPVerbPUT
	case "DELETE":
		*code = HTTPVerbDELETE
	case "PATCH":
		*code = HTTPVerbPATCH
	default:
		return fmt.Errorf("unknown HTTPVerb code `%s`", s)
	}
	return nil
}
func (code HTTPVerb) String() string {
	return code.Code()
}
func (code HTTPVerb) Code() string {
	switch code {
	case HTTPVerbGET:
		return "GET"
	case HTTPVerbHEAD:
		return "HEAD"
	case HTTPVerbPOST:
		return "POST"
	case HTTPVerbPUT:
		return "PUT"
	case HTTPVerbDELETE:
		return "DELETE"
	case HTTPVerbPATCH:
		return "PATCH"
	}
	return "<unknown>"
}
func (code HTTPVerb) Display() string {
	switch code {
	case HTTPVerbGET:
		return "GET"
	case HTTPVerbHEAD:
		return "HEAD"
	case HTTPVerbPOST:
		return "POST"
	case HTTPVerbPUT:
		return "PUT"
	case HTTPVerbDELETE:
		return "DELETE"
	case HTTPVerbPATCH:
		return "PATCH"
	}
	return "<unknown>"
}
func (code HTTPVerb) Definition() string {
	switch code {
	case HTTPVerbGET:
		return "HTTP GET Command."
	case HTTPVerbHEAD:
		return "HTTP HEAD Command."
	case HTTPVerbPOST:
		return "HTTP POST Command."
	case HTTPVerbPUT:
		return "HTTP PUT Command."
	case HTTPVerbDELETE:
		return "HTTP DELETE Command."
	case HTTPVerbPATCH:
		return "HTTP PATCH Command."
	}
	return "<unknown>"
}
