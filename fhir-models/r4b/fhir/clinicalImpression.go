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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ClinicalImpression is documented here http://hl7.org/fhir/StructureDefinition/ClinicalImpression
type ClinicalImpression struct {
	Id                       *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta                     *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules            *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                 *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text                     *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension                []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier               []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                   EventStatus                 `bson:"status" json:"status"`
	StatusReason             *CodeableConcept            `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Description              *string                     `bson:"description,omitempty" json:"description,omitempty"`
	Subject                  Reference                   `bson:"subject" json:"subject"`
	Encounter                *Reference                  `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime        *string                     `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod          *Period                     `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Date                     *string                     `bson:"date,omitempty" json:"date,omitempty"`
	Performer                *Reference                  `bson:"performer,omitempty" json:"performer,omitempty"`
	Previous                 *Reference                  `bson:"previous,omitempty" json:"previous,omitempty"`
	Problem                  []Reference                 `bson:"problem,omitempty" json:"problem,omitempty"`
	ChangePattern            *CodeableConcept            `bson:"changePattern,omitempty" json:"changePattern,omitempty"`
	Protocol                 []string                    `bson:"protocol,omitempty" json:"protocol,omitempty"`
	Summary                  *string                     `bson:"summary,omitempty" json:"summary,omitempty"`
	Finding                  []ClinicalImpressionFinding `bson:"finding,omitempty" json:"finding,omitempty"`
	PrognosisCodeableConcept []CodeableConcept           `bson:"prognosisCodeableConcept,omitempty" json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference       []Reference                 `bson:"prognosisReference,omitempty" json:"prognosisReference,omitempty"`
	SupportingInfo           []Reference                 `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Note                     []Annotation                `bson:"note,omitempty" json:"note,omitempty"`
}
type ClinicalImpressionFinding struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Item              *CodeableReference `bson:"item,omitempty" json:"item,omitempty"`
	Basis             *string            `bson:"basis,omitempty" json:"basis,omitempty"`
}
type OtherClinicalImpression ClinicalImpression

// MarshalJSON marshals the given ClinicalImpression as JSON into a byte slice
func (r ClinicalImpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClinicalImpression
		ResourceType string `json:"resourceType"`
	}{
		OtherClinicalImpression: OtherClinicalImpression(r),
		ResourceType:            "ClinicalImpression",
	})
}

// UnmarshalClinicalImpression unmarshals a ClinicalImpression.
func UnmarshalClinicalImpression(b []byte) (ClinicalImpression, error) {
	var clinicalImpression ClinicalImpression
	if err := json.Unmarshal(b, &clinicalImpression); err != nil {
		return clinicalImpression, err
	}
	return clinicalImpression, nil
}