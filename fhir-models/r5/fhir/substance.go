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

// Substance is documented here http://hl7.org/fhir/StructureDefinition/Substance
type Substance struct {
	Id                *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Instance          bool                  `bson:"instance" json:"instance"`
	Status            *FHIRSubstanceStatus  `bson:"status,omitempty" json:"status,omitempty"`
	Category          []CodeableConcept     `bson:"category,omitempty" json:"category,omitempty"`
	Code              CodeableReference     `bson:"code" json:"code"`
	Description       *string               `bson:"description,omitempty" json:"description,omitempty"`
	Expiry            *string               `bson:"expiry,omitempty" json:"expiry,omitempty"`
	Quantity          *Quantity             `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Ingredient        []SubstanceIngredient `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
}
type SubstanceIngredient struct {
	Id                       *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity                 *Ratio          `bson:"quantity,omitempty" json:"quantity,omitempty"`
	SubstanceCodeableConcept CodeableConcept `bson:"substanceCodeableConcept" json:"substanceCodeableConcept"`
	SubstanceReference       Reference       `bson:"substanceReference" json:"substanceReference"`
}
type OtherSubstance Substance

// MarshalJSON marshals the given Substance as JSON into a byte slice
func (r Substance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstance
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstance: OtherSubstance(r),
		ResourceType:   "Substance",
	})
}

// UnmarshalSubstance unmarshals a Substance.
func UnmarshalSubstance(b []byte) (Substance, error) {
	var substance Substance
	if err := json.Unmarshal(b, &substance); err != nil {
		return substance, err
	}
	return substance, nil
}
