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

// BiologicallyDerivedProduct is documented here http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProduct
type BiologicallyDerivedProduct struct {
	Id                      *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                    *Meta                                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules           *string                               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                *string                               `bson:"language,omitempty" json:"language,omitempty"`
	Text                    *Narrative                            `bson:"text,omitempty" json:"text,omitempty"`
	Extension               []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ProductCategory         *Coding                               `bson:"productCategory,omitempty" json:"productCategory,omitempty"`
	ProductCode             *CodeableConcept                      `bson:"productCode,omitempty" json:"productCode,omitempty"`
	Parent                  []Reference                           `bson:"parent,omitempty" json:"parent,omitempty"`
	Request                 []Reference                           `bson:"request,omitempty" json:"request,omitempty"`
	Identifier              []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BiologicalSourceEvent   *Identifier                           `bson:"biologicalSourceEvent,omitempty" json:"biologicalSourceEvent,omitempty"`
	ProcessingFacility      []Reference                           `bson:"processingFacility,omitempty" json:"processingFacility,omitempty"`
	Division                *string                               `bson:"division,omitempty" json:"division,omitempty"`
	ProductStatus           *Coding                               `bson:"productStatus,omitempty" json:"productStatus,omitempty"`
	ExpirationDate          *string                               `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	Collection              *BiologicallyDerivedProductCollection `bson:"collection,omitempty" json:"collection,omitempty"`
	StorageTempRequirements *Range                                `bson:"storageTempRequirements,omitempty" json:"storageTempRequirements,omitempty"`
	Property                []BiologicallyDerivedProductProperty  `bson:"property,omitempty" json:"property,omitempty"`
}
type BiologicallyDerivedProductCollection struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Collector         *Reference  `bson:"collector,omitempty" json:"collector,omitempty"`
	Source            *Reference  `bson:"source,omitempty" json:"source,omitempty"`
	CollectedDateTime *string     `bson:"collectedDateTime,omitempty" json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period     `bson:"collectedPeriod,omitempty" json:"collectedPeriod,omitempty"`
}
type BiologicallyDerivedProductProperty struct {
	Id                   *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 CodeableConcept `bson:"type" json:"type"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean"`
	ValueInteger         int             `bson:"valueInteger" json:"valueInteger"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept"`
	ValuePeriod          Period          `bson:"valuePeriod" json:"valuePeriod"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange"`
	ValueRatio           Ratio           `bson:"valueRatio" json:"valueRatio"`
	ValueString          string          `bson:"valueString" json:"valueString"`
	ValueAttachment      Attachment      `bson:"valueAttachment" json:"valueAttachment"`
}
type OtherBiologicallyDerivedProduct BiologicallyDerivedProduct

// MarshalJSON marshals the given BiologicallyDerivedProduct as JSON into a byte slice
func (r BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBiologicallyDerivedProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherBiologicallyDerivedProduct: OtherBiologicallyDerivedProduct(r),
		ResourceType:                    "BiologicallyDerivedProduct",
	})
}
func (r BiologicallyDerivedProduct) ResourceType() string {
	return "BiologicallyDerivedProduct"
}

// UnmarshalBiologicallyDerivedProduct unmarshals a BiologicallyDerivedProduct.
func UnmarshalBiologicallyDerivedProduct(b []byte) (BiologicallyDerivedProduct, error) {
	var biologicallyDerivedProduct BiologicallyDerivedProduct
	if err := json.Unmarshal(b, &biologicallyDerivedProduct); err != nil {
		return biologicallyDerivedProduct, err
	}
	return biologicallyDerivedProduct, nil
}
