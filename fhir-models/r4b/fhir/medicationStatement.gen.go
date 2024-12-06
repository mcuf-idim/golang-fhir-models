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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/mcuf-idim/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// MedicationStatement is documented here http://hl7.org/fhir/StructureDefinition/MedicationStatement
type MedicationStatement struct {
	Id                         *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta                       *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules              *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                   *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text                       *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension                  []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension          []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                 []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	PartOf                     []Reference                    `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                     MedicationStatementStatusCodes `bson:"status" json:"status"`
	Category                   []CodeableConcept              `bson:"category,omitempty" json:"category,omitempty"`
	Medication                 CodeableReference              `bson:"medication" json:"medication"`
	Subject                    Reference                      `bson:"subject" json:"subject"`
	Encounter                  *Reference                     `bson:"encounter,omitempty" json:"encounter,omitempty"`
	EffectiveDateTime          *string                        `bson:"effectiveDateTime,omitempty" json:"effectiveDateTime,omitempty"`
	EffectivePeriod            *Period                        `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	EffectiveTiming            *Timing                        `bson:"effectiveTiming,omitempty" json:"effectiveTiming,omitempty"`
	DateAsserted               *string                        `bson:"dateAsserted,omitempty" json:"dateAsserted,omitempty"`
	InformationSource          []Reference                    `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	DerivedFrom                []Reference                    `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Reason                     []CodeableReference            `bson:"reason,omitempty" json:"reason,omitempty"`
	Note                       []Annotation                   `bson:"note,omitempty" json:"note,omitempty"`
	RelatedClinicalInformation []Reference                    `bson:"relatedClinicalInformation,omitempty" json:"relatedClinicalInformation,omitempty"`
	RenderedDosageInstruction  *string                        `bson:"renderedDosageInstruction,omitempty" json:"renderedDosageInstruction,omitempty"`
	Dosage                     []Dosage                       `bson:"dosage,omitempty" json:"dosage,omitempty"`
	Adherence                  *MedicationStatementAdherence  `bson:"adherence,omitempty" json:"adherence,omitempty"`
}
type MedicationStatementAdherence struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept  `bson:"code" json:"code"`
	Reason            *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}
type OtherMedicationStatement MedicationStatement

// MarshalJSON marshals the given MedicationStatement as JSON into a byte slice
func (r MedicationStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationStatement: OtherMedicationStatement(r),
		ResourceType:             "MedicationStatement",
	})
}
func (r MedicationStatement) ResourceType() string {
	return "MedicationStatement"
}

// UnmarshalMedicationStatement unmarshals a MedicationStatement.
func UnmarshalMedicationStatement(b []byte) (MedicationStatement, error) {
	var medicationStatement MedicationStatement
	if err := json.Unmarshal(b, &medicationStatement); err != nil {
		return medicationStatement, err
	}
	return medicationStatement, nil
}
