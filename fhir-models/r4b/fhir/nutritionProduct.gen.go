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

// NutritionProduct is documented here http://hl7.org/fhir/StructureDefinition/NutritionProduct
type NutritionProduct struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept                 `bson:"code,omitempty" json:"code,omitempty"`
	Status            NutritionProductStatus           `bson:"status" json:"status"`
	Category          []CodeableConcept                `bson:"category,omitempty" json:"category,omitempty"`
	Manufacturer      []Reference                      `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Nutrient          []NutritionProductNutrient       `bson:"nutrient,omitempty" json:"nutrient,omitempty"`
	Ingredient        []NutritionProductIngredient     `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	KnownAllergen     []CodeableReference              `bson:"knownAllergen,omitempty" json:"knownAllergen,omitempty"`
	Characteristic    []NutritionProductCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Instance          []NutritionProductInstance       `bson:"instance,omitempty" json:"instance,omitempty"`
	Note              []Annotation                     `bson:"note,omitempty" json:"note,omitempty"`
}
type NutritionProductNutrient struct {
	Id                *string            `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Item              *CodeableReference `bson:"item,omitempty" json:"item,omitempty"`
	Amount            []Ratio            `bson:"amount,omitempty" json:"amount,omitempty"`
}
type NutritionProductIngredient struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Item              CodeableReference `bson:"item" json:"item"`
	Amount            []Ratio           `bson:"amount,omitempty" json:"amount,omitempty"`
}
type NutritionProductCharacteristic struct {
	Id                   *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 CodeableConcept `bson:"type" json:"type"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept"`
	ValueString          string          `bson:"valueString" json:"valueString"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity"`
	ValueBase64Binary    string          `bson:"valueBase64Binary" json:"valueBase64Binary"`
	ValueAttachment      Attachment      `bson:"valueAttachment" json:"valueAttachment"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean"`
}
type NutritionProductInstance struct {
	Id                    *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Quantity              *Quantity    `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Identifier            []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name                  *string      `bson:"name,omitempty" json:"name,omitempty"`
	LotNumber             *string      `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	Expiry                *string      `bson:"expiry,omitempty" json:"expiry,omitempty"`
	UseBy                 *string      `bson:"useBy,omitempty" json:"useBy,omitempty"`
	BiologicalSourceEvent *Identifier  `bson:"biologicalSourceEvent,omitempty" json:"biologicalSourceEvent,omitempty"`
}
type OtherNutritionProduct NutritionProduct

// MarshalJSON marshals the given NutritionProduct as JSON into a byte slice
func (r NutritionProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherNutritionProduct: OtherNutritionProduct(r),
		ResourceType:          "NutritionProduct",
	})
}
func (r NutritionProduct) ResourceType() string {
	return "NutritionProduct"
}

// UnmarshalNutritionProduct unmarshals a NutritionProduct.
func UnmarshalNutritionProduct(b []byte) (NutritionProduct, error) {
	var nutritionProduct NutritionProduct
	if err := json.Unmarshal(b, &nutritionProduct); err != nil {
		return nutritionProduct, err
	}
	return nutritionProduct, nil
}
