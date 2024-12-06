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

// BiologicallyDerivedProductDispense is documented here http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProductDispense
type BiologicallyDerivedProductDispense struct {
	Id                     *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                                         `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                                       `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                                       `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                                    `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             []Identifier                                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn                []Reference                                   `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf                 []Reference                                   `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                 BiologicallyDerivedProductDispenseCodes       `bson:"status" json:"status"`
	OriginRelationshipType *CodeableConcept                              `bson:"originRelationshipType,omitempty" json:"originRelationshipType,omitempty"`
	Product                Reference                                     `bson:"product" json:"product"`
	Patient                Reference                                     `bson:"patient" json:"patient"`
	MatchStatus            *CodeableConcept                              `bson:"matchStatus,omitempty" json:"matchStatus,omitempty"`
	Performer              []BiologicallyDerivedProductDispensePerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	Location               *Reference                                    `bson:"location,omitempty" json:"location,omitempty"`
	Quantity               *Quantity                                     `bson:"quantity,omitempty" json:"quantity,omitempty"`
	PreparedDate           *string                                       `bson:"preparedDate,omitempty" json:"preparedDate,omitempty"`
	WhenHandedOver         *string                                       `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	Destination            *Reference                                    `bson:"destination,omitempty" json:"destination,omitempty"`
	Note                   []Annotation                                  `bson:"note,omitempty" json:"note,omitempty"`
	UsageInstruction       *string                                       `bson:"usageInstruction,omitempty" json:"usageInstruction,omitempty"`
}
type BiologicallyDerivedProductDispensePerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}
type OtherBiologicallyDerivedProductDispense BiologicallyDerivedProductDispense

// MarshalJSON marshals the given BiologicallyDerivedProductDispense as JSON into a byte slice
func (r BiologicallyDerivedProductDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBiologicallyDerivedProductDispense
		ResourceType string `json:"resourceType"`
	}{
		OtherBiologicallyDerivedProductDispense: OtherBiologicallyDerivedProductDispense(r),
		ResourceType:                            "BiologicallyDerivedProductDispense",
	})
}
func (r BiologicallyDerivedProductDispense) ResourceType() string {
	return "BiologicallyDerivedProductDispense"
}

// UnmarshalBiologicallyDerivedProductDispense unmarshals a BiologicallyDerivedProductDispense.
func UnmarshalBiologicallyDerivedProductDispense(b []byte) (BiologicallyDerivedProductDispense, error) {
	var biologicallyDerivedProductDispense BiologicallyDerivedProductDispense
	if err := json.Unmarshal(b, &biologicallyDerivedProductDispense); err != nil {
		return biologicallyDerivedProductDispense, err
	}
	return biologicallyDerivedProductDispense, nil
}
