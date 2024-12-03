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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// MarketingStatus is documented here http://hl7.org/fhir/StructureDefinition/MarketingStatus
type MarketingStatus struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Country           *CodeableConcept `bson:"country,omitempty" json:"country,omitempty"`
	Jurisdiction      *CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Status            CodeableConcept  `bson:"status" json:"status"`
	DateRange         *Period          `bson:"dateRange,omitempty" json:"dateRange,omitempty"`
	RestoreDate       *string          `bson:"restoreDate,omitempty" json:"restoreDate,omitempty"`
}
