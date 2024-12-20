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

// Organization is documented here http://hl7.org/fhir/StructureDefinition/Organization
type Organization struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active            *bool                       `bson:"active,omitempty" json:"active,omitempty"`
	Type              []CodeableConcept           `bson:"type,omitempty" json:"type,omitempty"`
	Name              *string                     `bson:"name,omitempty" json:"name,omitempty"`
	Alias             []string                    `bson:"alias,omitempty" json:"alias,omitempty"`
	Description       *string                     `bson:"description,omitempty" json:"description,omitempty"`
	Contact           []ExtendedContactDetail     `bson:"contact,omitempty" json:"contact,omitempty"`
	PartOf            *Reference                  `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Endpoint          []Reference                 `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	Qualification     []OrganizationQualification `bson:"qualification,omitempty" json:"qualification,omitempty"`
}

func (r Organization) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OrganizationQualification struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Period            *Period         `bson:"period,omitempty" json:"period,omitempty"`
	Issuer            *Reference      `bson:"issuer,omitempty" json:"issuer,omitempty"`
}
type OtherOrganization Organization

// MarshalJSON marshals the given Organization as JSON into a byte slice
func (r Organization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOrganization
		ResourceType string `json:"resourceType"`
	}{
		OtherOrganization: OtherOrganization(r),
		ResourceType:      "Organization",
	})
}
func (r Organization) ResourceType() string {
	return "Organization"
}

// UnmarshalOrganization unmarshals a Organization.
func UnmarshalOrganization(b []byte) (Organization, error) {
	var organization Organization
	if err := json.Unmarshal(b, &organization); err != nil {
		return organization, err
	}
	return organization, nil
}
