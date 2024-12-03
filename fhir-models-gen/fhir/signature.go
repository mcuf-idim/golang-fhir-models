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

// THIS FILE IS GENERATED BY https://github.com/mcuf-idim/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Signature is documented here http://hl7.org/fhir/StructureDefinition/Signature
type Signature struct {
	Id           *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension    []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Type         []Coding    `bson:"type,omitempty" json:"type,omitempty"`
	When         *string     `bson:"when,omitempty" json:"when,omitempty"`
	Who          *Reference  `bson:"who,omitempty" json:"who,omitempty"`
	OnBehalfOf   *Reference  `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	TargetFormat *string     `bson:"targetFormat,omitempty" json:"targetFormat,omitempty"`
	SigFormat    *string     `bson:"sigFormat,omitempty" json:"sigFormat,omitempty"`
	Data         *string     `bson:"data,omitempty" json:"data,omitempty"`
}
