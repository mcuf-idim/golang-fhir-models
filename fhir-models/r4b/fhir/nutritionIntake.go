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

// NutritionIntake is documented here http://hl7.org/fhir/StructureDefinition/NutritionIntake
type NutritionIntake struct {
	Id                    *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string                         `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                         `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference                      `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf                []Reference                      `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                EventStatus                      `bson:"status" json:"status"`
	StatusReason          []CodeableConcept                `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Code                  *CodeableConcept                 `bson:"code,omitempty" json:"code,omitempty"`
	Subject               Reference                        `bson:"subject" json:"subject"`
	Encounter             *Reference                       `bson:"encounter,omitempty" json:"encounter,omitempty"`
	OccurrenceDateTime    *string                          `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                          `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	Recorded              *string                          `bson:"recorded,omitempty" json:"recorded,omitempty"`
	ReportedBoolean       *bool                            `bson:"reportedBoolean,omitempty" json:"reportedBoolean,omitempty"`
	ReportedReference     *Reference                       `bson:"reportedReference,omitempty" json:"reportedReference,omitempty"`
	ConsumedItem          []NutritionIntakeConsumedItem    `bson:"consumedItem" json:"consumedItem"`
	IngredientLabel       []NutritionIntakeIngredientLabel `bson:"ingredientLabel,omitempty" json:"ingredientLabel,omitempty"`
	Performer             []NutritionIntakePerformer       `bson:"performer,omitempty" json:"performer,omitempty"`
	Location              *Reference                       `bson:"location,omitempty" json:"location,omitempty"`
	DerivedFrom           []Reference                      `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Reason                []CodeableReference              `bson:"reason,omitempty" json:"reason,omitempty"`
	Note                  []Annotation                     `bson:"note,omitempty" json:"note,omitempty"`
}
type NutritionIntakeConsumedItem struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `bson:"type" json:"type"`
	NutritionProduct  CodeableReference `bson:"nutritionProduct" json:"nutritionProduct"`
	Schedule          *Timing           `bson:"schedule,omitempty" json:"schedule,omitempty"`
	Amount            *Quantity         `bson:"amount,omitempty" json:"amount,omitempty"`
	Rate              *Quantity         `bson:"rate,omitempty" json:"rate,omitempty"`
	NotConsumed       *bool             `bson:"notConsumed,omitempty" json:"notConsumed,omitempty"`
	NotConsumedReason *CodeableConcept  `bson:"notConsumedReason,omitempty" json:"notConsumedReason,omitempty"`
}
type NutritionIntakeIngredientLabel struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Nutrient          CodeableReference `bson:"nutrient" json:"nutrient"`
	Amount            Quantity          `bson:"amount" json:"amount"`
}
type NutritionIntakePerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}
type OtherNutritionIntake NutritionIntake

// MarshalJSON marshals the given NutritionIntake as JSON into a byte slice
func (r NutritionIntake) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionIntake
		ResourceType string `json:"resourceType"`
	}{
		OtherNutritionIntake: OtherNutritionIntake(r),
		ResourceType:         "NutritionIntake",
	})
}

// UnmarshalNutritionIntake unmarshals a NutritionIntake.
func UnmarshalNutritionIntake(b []byte) (NutritionIntake, error) {
	var nutritionIntake NutritionIntake
	if err := json.Unmarshal(b, &nutritionIntake); err != nil {
		return nutritionIntake, err
	}
	return nutritionIntake, nil
}
