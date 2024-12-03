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

// QuestionnaireItemDisabledDisplay is documented here http://hl7.org/fhir/ValueSet/questionnaire-disabled-display
type QuestionnaireItemDisabledDisplay int

const (
	QuestionnaireItemDisabledDisplayHidden QuestionnaireItemDisabledDisplay = iota
	QuestionnaireItemDisabledDisplayProtected
)

func (code QuestionnaireItemDisabledDisplay) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *QuestionnaireItemDisabledDisplay) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "hidden":
		*code = QuestionnaireItemDisabledDisplayHidden
	case "protected":
		*code = QuestionnaireItemDisabledDisplayProtected
	default:
		return fmt.Errorf("unknown QuestionnaireItemDisabledDisplay code `%s`", s)
	}
	return nil
}
func (code QuestionnaireItemDisabledDisplay) String() string {
	return code.Code()
}
func (code QuestionnaireItemDisabledDisplay) Code() string {
	switch code {
	case QuestionnaireItemDisabledDisplayHidden:
		return "hidden"
	case QuestionnaireItemDisabledDisplayProtected:
		return "protected"
	}
	return "<unknown>"
}
func (code QuestionnaireItemDisabledDisplay) Display() string {
	switch code {
	case QuestionnaireItemDisabledDisplayHidden:
		return "Hidden"
	case QuestionnaireItemDisabledDisplayProtected:
		return "Protected"
	}
	return "<unknown>"
}
func (code QuestionnaireItemDisabledDisplay) Definition() string {
	switch code {
	case QuestionnaireItemDisabledDisplayHidden:
		return "The item (and its children) should not be visible to the user at all."
	case QuestionnaireItemDisabledDisplayProtected:
		return "The item (and possibly its children) should not be selectable or editable but should still be visible - to allow the user to see what questions *could* have been completed had other answers caused the item to be enabled."
	}
	return "<unknown>"
}
