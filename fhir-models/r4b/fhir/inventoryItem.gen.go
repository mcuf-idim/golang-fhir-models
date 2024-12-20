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

// InventoryItem is documented here http://hl7.org/fhir/StructureDefinition/InventoryItem
type InventoryItem struct {
	Id                      *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Meta                    *Meta                                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules           *string                                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                *string                                `bson:"language,omitempty" json:"language,omitempty"`
	Text                    *Narrative                             `bson:"text,omitempty" json:"text,omitempty"`
	Contained               []json.RawMessage                      `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension               []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier              []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                  InventoryItemStatusCodes               `bson:"status" json:"status"`
	Category                []CodeableConcept                      `bson:"category,omitempty" json:"category,omitempty"`
	Code                    []CodeableConcept                      `bson:"code,omitempty" json:"code,omitempty"`
	Name                    []InventoryItemName                    `bson:"name,omitempty" json:"name,omitempty"`
	ResponsibleOrganization []InventoryItemResponsibleOrganization `bson:"responsibleOrganization,omitempty" json:"responsibleOrganization,omitempty"`
	Description             *InventoryItemDescription              `bson:"description,omitempty" json:"description,omitempty"`
	InventoryStatus         []CodeableConcept                      `bson:"inventoryStatus,omitempty" json:"inventoryStatus,omitempty"`
	BaseUnit                *CodeableConcept                       `bson:"baseUnit,omitempty" json:"baseUnit,omitempty"`
	NetContent              *Quantity                              `bson:"netContent,omitempty" json:"netContent,omitempty"`
	Association             []InventoryItemAssociation             `bson:"association,omitempty" json:"association,omitempty"`
	Characteristic          []InventoryItemCharacteristic          `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Instance                *InventoryItemInstance                 `bson:"instance,omitempty" json:"instance,omitempty"`
	ProductReference        *Reference                             `bson:"productReference,omitempty" json:"productReference,omitempty"`
}

func (r InventoryItem) ContainedResources() []json.RawMessage {
	return r.Contained
}

type InventoryItemName struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	NameType          Coding      `bson:"nameType" json:"nameType"`
	Language          string      `bson:"language" json:"language"`
	Name              string      `bson:"name" json:"name"`
}
type InventoryItemResponsibleOrganization struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              CodeableConcept `bson:"role" json:"role"`
	Organization      Reference       `bson:"organization" json:"organization"`
}
type InventoryItemDescription struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Language          *string     `bson:"language,omitempty" json:"language,omitempty"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
}
type InventoryItemAssociation struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	AssociationType   CodeableConcept `bson:"associationType" json:"associationType"`
	RelatedItem       Reference       `bson:"relatedItem" json:"relatedItem"`
	Quantity          Ratio           `bson:"quantity" json:"quantity"`
}
type InventoryItemCharacteristic struct {
	Id                   *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	CharacteristicType   CodeableConcept `bson:"characteristicType" json:"characteristicType"`
	ValueString          string          `bson:"valueString" json:"valueString"`
	ValueInteger         int             `bson:"valueInteger" json:"valueInteger"`
	ValueDecimal         json.Number     `bson:"valueDecimal" json:"valueDecimal"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean"`
	ValueUrl             string          `bson:"valueUrl" json:"valueUrl"`
	ValueDateTime        string          `bson:"valueDateTime" json:"valueDateTime"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange"`
	ValueRatio           Ratio           `bson:"valueRatio" json:"valueRatio"`
	ValueAnnotation      Annotation      `bson:"valueAnnotation" json:"valueAnnotation"`
	ValueAddress         Address         `bson:"valueAddress" json:"valueAddress"`
	ValueDuration        Duration        `bson:"valueDuration" json:"valueDuration"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept"`
}
type InventoryItemInstance struct {
	Id                *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	LotNumber         *string      `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	Expiry            *string      `bson:"expiry,omitempty" json:"expiry,omitempty"`
	Subject           *Reference   `bson:"subject,omitempty" json:"subject,omitempty"`
	Location          *Reference   `bson:"location,omitempty" json:"location,omitempty"`
}
type OtherInventoryItem InventoryItem

// MarshalJSON marshals the given InventoryItem as JSON into a byte slice
func (r InventoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInventoryItem
		ResourceType string `json:"resourceType"`
	}{
		OtherInventoryItem: OtherInventoryItem(r),
		ResourceType:       "InventoryItem",
	})
}
func (r InventoryItem) ResourceType() string {
	return "InventoryItem"
}

// UnmarshalInventoryItem unmarshals a InventoryItem.
func UnmarshalInventoryItem(b []byte) (InventoryItem, error) {
	var inventoryItem InventoryItem
	if err := json.Unmarshal(b, &inventoryItem); err != nil {
		return inventoryItem, err
	}
	return inventoryItem, nil
}
