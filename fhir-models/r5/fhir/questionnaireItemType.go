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

// QuestionnaireItemType is documented here http://hl7.org/fhir/ValueSet/item-type
type QuestionnaireItemType int

const (
	QuestionnaireItemTypeGroup QuestionnaireItemType = iota
	QuestionnaireItemTypeDisplay
	QuestionnaireItemTypeQuestion
	QuestionnaireItemTypeBoolean
	QuestionnaireItemTypeDecimal
	QuestionnaireItemTypeInteger
	QuestionnaireItemTypeDate
	QuestionnaireItemTypeDateTime
	QuestionnaireItemTypeTime
	QuestionnaireItemTypeString
	QuestionnaireItemTypeText
	QuestionnaireItemTypeUrl
	QuestionnaireItemTypeCoding
	QuestionnaireItemTypeAttachment
	QuestionnaireItemTypeReference
	QuestionnaireItemTypeQuantity
)

func (code QuestionnaireItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *QuestionnaireItemType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "group":
		*code = QuestionnaireItemTypeGroup
	case "display":
		*code = QuestionnaireItemTypeDisplay
	case "question":
		*code = QuestionnaireItemTypeQuestion
	case "boolean":
		*code = QuestionnaireItemTypeBoolean
	case "decimal":
		*code = QuestionnaireItemTypeDecimal
	case "integer":
		*code = QuestionnaireItemTypeInteger
	case "date":
		*code = QuestionnaireItemTypeDate
	case "dateTime":
		*code = QuestionnaireItemTypeDateTime
	case "time":
		*code = QuestionnaireItemTypeTime
	case "string":
		*code = QuestionnaireItemTypeString
	case "text":
		*code = QuestionnaireItemTypeText
	case "url":
		*code = QuestionnaireItemTypeUrl
	case "coding":
		*code = QuestionnaireItemTypeCoding
	case "attachment":
		*code = QuestionnaireItemTypeAttachment
	case "reference":
		*code = QuestionnaireItemTypeReference
	case "quantity":
		*code = QuestionnaireItemTypeQuantity
	default:
		return fmt.Errorf("unknown QuestionnaireItemType code `%s`", s)
	}
	return nil
}
func (code QuestionnaireItemType) String() string {
	return code.Code()
}
func (code QuestionnaireItemType) Code() string {
	switch code {
	case QuestionnaireItemTypeGroup:
		return "group"
	case QuestionnaireItemTypeDisplay:
		return "display"
	case QuestionnaireItemTypeQuestion:
		return "question"
	case QuestionnaireItemTypeBoolean:
		return "boolean"
	case QuestionnaireItemTypeDecimal:
		return "decimal"
	case QuestionnaireItemTypeInteger:
		return "integer"
	case QuestionnaireItemTypeDate:
		return "date"
	case QuestionnaireItemTypeDateTime:
		return "dateTime"
	case QuestionnaireItemTypeTime:
		return "time"
	case QuestionnaireItemTypeString:
		return "string"
	case QuestionnaireItemTypeText:
		return "text"
	case QuestionnaireItemTypeUrl:
		return "url"
	case QuestionnaireItemTypeCoding:
		return "coding"
	case QuestionnaireItemTypeAttachment:
		return "attachment"
	case QuestionnaireItemTypeReference:
		return "reference"
	case QuestionnaireItemTypeQuantity:
		return "quantity"
	}
	return "<unknown>"
}
func (code QuestionnaireItemType) Display() string {
	switch code {
	case QuestionnaireItemTypeGroup:
		return "Group"
	case QuestionnaireItemTypeDisplay:
		return "Display"
	case QuestionnaireItemTypeQuestion:
		return "Question"
	case QuestionnaireItemTypeBoolean:
		return "Boolean"
	case QuestionnaireItemTypeDecimal:
		return "Decimal"
	case QuestionnaireItemTypeInteger:
		return "Integer"
	case QuestionnaireItemTypeDate:
		return "Date"
	case QuestionnaireItemTypeDateTime:
		return "Date Time"
	case QuestionnaireItemTypeTime:
		return "Time"
	case QuestionnaireItemTypeString:
		return "String"
	case QuestionnaireItemTypeText:
		return "Text"
	case QuestionnaireItemTypeUrl:
		return "Url"
	case QuestionnaireItemTypeCoding:
		return "Coding"
	case QuestionnaireItemTypeAttachment:
		return "Attachment"
	case QuestionnaireItemTypeReference:
		return "Reference"
	case QuestionnaireItemTypeQuantity:
		return "Quantity"
	}
	return "<unknown>"
}
func (code QuestionnaireItemType) Definition() string {
	switch code {
	case QuestionnaireItemTypeGroup:
		return "An item with no direct answer but should have at least one child item."
	case QuestionnaireItemTypeDisplay:
		return "Text for display that will not capture an answer or have child items."
	case QuestionnaireItemTypeQuestion:
		return "An item that defines a specific answer to be captured, and which may have child items. (the answer provided in the QuestionnaireResponse should be of the defined datatype)."
	case QuestionnaireItemTypeBoolean:
		return "Question with a yes/no answer (valueBoolean)."
	case QuestionnaireItemTypeDecimal:
		return "Question with is a real number answer (valueDecimal).  There is an extension 'http://hl7.org/fhir/StructureDefinition/questionnaire-unit' that can be used to computably convey the unit of measure associated with the answer for use when performing data extraction to an element of type Quantity."
	case QuestionnaireItemTypeInteger:
		return "Question with an integer answer (valueInteger).  There is an extension 'http://hl7.org/fhir/StructureDefinition/questionnaire-unit' that can be used to computably convey the unit of measure associated with the answer for use when performing data extraction to an element of type Quantity."
	case QuestionnaireItemTypeDate:
		return "Question with a date answer (valueDate)."
	case QuestionnaireItemTypeDateTime:
		return "Question with a date and time answer (valueDateTime)."
	case QuestionnaireItemTypeTime:
		return "Question with a time (hour:minute:second) answer independent of date. (valueTime)."
	case QuestionnaireItemTypeString:
		return "Question with a short (few words to short sentence) free-text entry answer (valueString).  Strings SHOULD NOT contain carriage return or newline characters.  If multi-line answers are needed, use the 'text' type."
	case QuestionnaireItemTypeText:
		return "Question with a long (potentially multi-paragraph) free-text entry answer (valueString)."
	case QuestionnaireItemTypeUrl:
		return "Question with a URL (website, FTP site, etc.) answer (valueUri)."
	case QuestionnaireItemTypeCoding:
		return "Question with a Coding - generally drawn from a list of possible answers (valueCoding)"
	case QuestionnaireItemTypeAttachment:
		return "Question with binary content such as an image, PDF, etc. as an answer (valueAttachment)."
	case QuestionnaireItemTypeReference:
		return "Question with a reference to another resource (practitioner, organization, etc.) as an answer (valueReference)."
	case QuestionnaireItemTypeQuantity:
		return "Question with a combination of a numeric value and unit as an answer. (valueSimpleQuantity)  There are two extensions ('http://hl7.org/fhir/StructureDefinition/questionnaire-unitOption' and 'http://hl7.org/fhir/StructureDefinition/questionnaire-unitValueSet')  that can be used to define what unit should be selected for the Quantity.code and Quantity.system."
	}
	return "<unknown>"
}
