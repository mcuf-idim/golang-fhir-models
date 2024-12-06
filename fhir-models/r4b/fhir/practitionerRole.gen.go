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

// PractitionerRole is documented here http://hl7.org/fhir/StructureDefinition/PractitionerRole
type PractitionerRole struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                   `bson:"active,omitempty" json:"active,omitempty"`
	Period            *Period                 `bson:"period,omitempty" json:"period,omitempty"`
	Practitioner      *Reference              `bson:"practitioner,omitempty" json:"practitioner,omitempty"`
	Organization      *Reference              `bson:"organization,omitempty" json:"organization,omitempty"`
	Code              []CodeableConcept       `bson:"code,omitempty" json:"code,omitempty"`
	Specialty         []CodeableConcept       `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Location          []Reference             `bson:"location,omitempty" json:"location,omitempty"`
	HealthcareService []Reference             `bson:"healthcareService,omitempty" json:"healthcareService,omitempty"`
	Contact           []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	Characteristic    []CodeableConcept       `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Communication     []CodeableConcept       `bson:"communication,omitempty" json:"communication,omitempty"`
	Availability      []Availability          `bson:"availability,omitempty" json:"availability,omitempty"`
	Endpoint          []Reference             `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type OtherPractitionerRole PractitionerRole

// MarshalJSON marshals the given PractitionerRole as JSON into a byte slice
func (r PractitionerRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitionerRole
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitionerRole: OtherPractitionerRole(r),
		ResourceType:          "PractitionerRole",
	})
}
func (r PractitionerRole) ResourceType() string {
	return "PractitionerRole"
}

// UnmarshalPractitionerRole unmarshals a PractitionerRole.
func UnmarshalPractitionerRole(b []byte) (PractitionerRole, error) {
	var practitionerRole PractitionerRole
	if err := json.Unmarshal(b, &practitionerRole); err != nil {
		return practitionerRole, err
	}
	return practitionerRole, nil
}
