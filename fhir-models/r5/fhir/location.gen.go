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

// Location is documented here http://hl7.org/fhir/StructureDefinition/Location
type Location struct {
	Id                   *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status               *LocationStatus         `bson:"status,omitempty" json:"status,omitempty"`
	OperationalStatus    *Coding                 `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	Name                 *string                 `bson:"name,omitempty" json:"name,omitempty"`
	Alias                []string                `bson:"alias,omitempty" json:"alias,omitempty"`
	Description          *string                 `bson:"description,omitempty" json:"description,omitempty"`
	Mode                 *LocationMode           `bson:"mode,omitempty" json:"mode,omitempty"`
	Type                 []CodeableConcept       `bson:"type,omitempty" json:"type,omitempty"`
	Contact              []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	Address              *Address                `bson:"address,omitempty" json:"address,omitempty"`
	Form                 *CodeableConcept        `bson:"form,omitempty" json:"form,omitempty"`
	Position             *LocationPosition       `bson:"position,omitempty" json:"position,omitempty"`
	ManagingOrganization *Reference              `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	PartOf               *Reference              `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Characteristic       []CodeableConcept       `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	HoursOfOperation     []Availability          `bson:"hoursOfOperation,omitempty" json:"hoursOfOperation,omitempty"`
	VirtualService       []VirtualServiceDetail  `bson:"virtualService,omitempty" json:"virtualService,omitempty"`
	Endpoint             []Reference             `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type LocationPosition struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Longitude         json.Number  `bson:"longitude" json:"longitude"`
	Latitude          json.Number  `bson:"latitude" json:"latitude"`
	Altitude          *json.Number `bson:"altitude,omitempty" json:"altitude,omitempty"`
}
type OtherLocation Location

// MarshalJSON marshals the given Location as JSON into a byte slice
func (r Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLocation
		ResourceType string `json:"resourceType"`
	}{
		OtherLocation: OtherLocation(r),
		ResourceType:  "Location",
	})
}
func (r Location) ResourceType() string {
	return "Location"
}

// UnmarshalLocation unmarshals a Location.
func UnmarshalLocation(b []byte) (Location, error) {
	var location Location
	if err := json.Unmarshal(b, &location); err != nil {
		return location, err
	}
	return location, nil
}
