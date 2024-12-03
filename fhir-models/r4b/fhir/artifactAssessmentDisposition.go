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

// ArtifactAssessmentDisposition is documented here http://hl7.org/fhir/ValueSet/artifactassessment-disposition
type ArtifactAssessmentDisposition int

const (
	ArtifactAssessmentDispositionUnresolved ArtifactAssessmentDisposition = iota
	ArtifactAssessmentDispositionNotPersuasive
	ArtifactAssessmentDispositionPersuasive
	ArtifactAssessmentDispositionPersuasiveWithModification
	ArtifactAssessmentDispositionNotPersuasiveWithModification
)

func (code ArtifactAssessmentDisposition) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ArtifactAssessmentDisposition) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "unresolved":
		*code = ArtifactAssessmentDispositionUnresolved
	case "not-persuasive":
		*code = ArtifactAssessmentDispositionNotPersuasive
	case "persuasive":
		*code = ArtifactAssessmentDispositionPersuasive
	case "persuasive-with-modification":
		*code = ArtifactAssessmentDispositionPersuasiveWithModification
	case "not-persuasive-with-modification":
		*code = ArtifactAssessmentDispositionNotPersuasiveWithModification
	default:
		return fmt.Errorf("unknown ArtifactAssessmentDisposition code `%s`", s)
	}
	return nil
}
func (code ArtifactAssessmentDisposition) String() string {
	return code.Code()
}
func (code ArtifactAssessmentDisposition) Code() string {
	switch code {
	case ArtifactAssessmentDispositionUnresolved:
		return "unresolved"
	case ArtifactAssessmentDispositionNotPersuasive:
		return "not-persuasive"
	case ArtifactAssessmentDispositionPersuasive:
		return "persuasive"
	case ArtifactAssessmentDispositionPersuasiveWithModification:
		return "persuasive-with-modification"
	case ArtifactAssessmentDispositionNotPersuasiveWithModification:
		return "not-persuasive-with-modification"
	}
	return "<unknown>"
}
func (code ArtifactAssessmentDisposition) Display() string {
	switch code {
	case ArtifactAssessmentDispositionUnresolved:
		return "Unresolved"
	case ArtifactAssessmentDispositionNotPersuasive:
		return "Not Persuasive"
	case ArtifactAssessmentDispositionPersuasive:
		return "Persuasive"
	case ArtifactAssessmentDispositionPersuasiveWithModification:
		return "Persuasive with Modification"
	case ArtifactAssessmentDispositionNotPersuasiveWithModification:
		return "Not Persuasive with Modification"
	}
	return "<unknown>"
}
func (code ArtifactAssessmentDisposition) Definition() string {
	switch code {
	case ArtifactAssessmentDispositionUnresolved:
		return "The comment is unresolved"
	case ArtifactAssessmentDispositionNotPersuasive:
		return "The comment is not persuasive (rejected in full)"
	case ArtifactAssessmentDispositionPersuasive:
		return "The comment is persuasive (accepted in full)"
	case ArtifactAssessmentDispositionPersuasiveWithModification:
		return "The comment is persuasive with modification (partially accepted)"
	case ArtifactAssessmentDispositionNotPersuasiveWithModification:
		return "The comment is not persuasive with modification (partially rejected)"
	}
	return "<unknown>"
}
