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

// PackagedProductDefinition is documented here http://hl7.org/fhir/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinition struct {
	Id                    *string                                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                                        `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name                  *string                                        `bson:"name,omitempty" json:"name,omitempty"`
	Type                  *CodeableConcept                               `bson:"type,omitempty" json:"type,omitempty"`
	PackageFor            []Reference                                    `bson:"packageFor,omitempty" json:"packageFor,omitempty"`
	Status                *CodeableConcept                               `bson:"status,omitempty" json:"status,omitempty"`
	StatusDate            *string                                        `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	ContainedItemQuantity []Quantity                                     `bson:"containedItemQuantity,omitempty" json:"containedItemQuantity,omitempty"`
	Description           *string                                        `bson:"description,omitempty" json:"description,omitempty"`
	LegalStatusOfSupply   []PackagedProductDefinitionLegalStatusOfSupply `bson:"legalStatusOfSupply,omitempty" json:"legalStatusOfSupply,omitempty"`
	MarketingStatus       []MarketingStatus                              `bson:"marketingStatus,omitempty" json:"marketingStatus,omitempty"`
	CopackagedIndicator   *bool                                          `bson:"copackagedIndicator,omitempty" json:"copackagedIndicator,omitempty"`
	Manufacturer          []Reference                                    `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	AttachedDocument      []Reference                                    `bson:"attachedDocument,omitempty" json:"attachedDocument,omitempty"`
	Packaging             *PackagedProductDefinitionPackaging            `bson:"packaging,omitempty" json:"packaging,omitempty"`
	Characteristic        []PackagedProductDefinitionPackagingProperty   `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
}
type PackagedProductDefinitionLegalStatusOfSupply struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Jurisdiction      *CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
}
type PackagedProductDefinitionPackaging struct {
	Id                *string                                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                                      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              *CodeableConcept                                  `bson:"type,omitempty" json:"type,omitempty"`
	ComponentPart     *bool                                             `bson:"componentPart,omitempty" json:"componentPart,omitempty"`
	Quantity          *int                                              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Material          []CodeableConcept                                 `bson:"material,omitempty" json:"material,omitempty"`
	AlternateMaterial []CodeableConcept                                 `bson:"alternateMaterial,omitempty" json:"alternateMaterial,omitempty"`
	ShelfLifeStorage  []ProductShelfLife                                `bson:"shelfLifeStorage,omitempty" json:"shelfLifeStorage,omitempty"`
	Manufacturer      []Reference                                       `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	Property          []PackagedProductDefinitionPackagingProperty      `bson:"property,omitempty" json:"property,omitempty"`
	ContainedItem     []PackagedProductDefinitionPackagingContainedItem `bson:"containedItem,omitempty" json:"containedItem,omitempty"`
	Packaging         []PackagedProductDefinitionPackaging              `bson:"packaging,omitempty" json:"packaging,omitempty"`
}
type PackagedProductDefinitionPackagingProperty struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 CodeableConcept  `bson:"type" json:"type"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueDate            *string          `bson:"valueDate,omitempty" json:"valueDate,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueAttachment      *Attachment      `bson:"valueAttachment,omitempty" json:"valueAttachment,omitempty"`
}
type PackagedProductDefinitionPackagingContainedItem struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Item              CodeableReference `bson:"item" json:"item"`
	Amount            *Quantity         `bson:"amount,omitempty" json:"amount,omitempty"`
}
type OtherPackagedProductDefinition PackagedProductDefinition

// MarshalJSON marshals the given PackagedProductDefinition as JSON into a byte slice
func (r PackagedProductDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPackagedProductDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherPackagedProductDefinition: OtherPackagedProductDefinition(r),
		ResourceType:                   "PackagedProductDefinition",
	})
}
func (r PackagedProductDefinition) ResourceType() string {
	return "PackagedProductDefinition"
}

// UnmarshalPackagedProductDefinition unmarshals a PackagedProductDefinition.
func UnmarshalPackagedProductDefinition(b []byte) (PackagedProductDefinition, error) {
	var packagedProductDefinition PackagedProductDefinition
	if err := json.Unmarshal(b, &packagedProductDefinition); err != nil {
		return packagedProductDefinition, err
	}
	return packagedProductDefinition, nil
}
