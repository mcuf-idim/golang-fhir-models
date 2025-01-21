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

// SubstanceReferenceInformation is documented here http://hl7.org/fhir/StructureDefinition/SubstanceReferenceInformation
type SubstanceReferenceInformation struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                    `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                                 `bson:"text,omitempty" json:"text,omitempty"`
	Contained         []json.RawMessage                          `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Comment           *string                                    `bson:"comment,omitempty" json:"comment,omitempty"`
	Gene              []SubstanceReferenceInformationGene        `bson:"gene,omitempty" json:"gene,omitempty"`
	GeneElement       []SubstanceReferenceInformationGeneElement `bson:"geneElement,omitempty" json:"geneElement,omitempty"`
	Target            []SubstanceReferenceInformationTarget      `bson:"target,omitempty" json:"target,omitempty"`
}

func (r SubstanceReferenceInformation) ContainedResources() []json.RawMessage {
	return r.Contained
}

type SubstanceReferenceInformationGene struct {
	Id                 *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	GeneSequenceOrigin *CodeableConcept `bson:"geneSequenceOrigin,omitempty" json:"geneSequenceOrigin,omitempty"`
	Gene               *CodeableConcept `bson:"gene,omitempty" json:"gene,omitempty"`
	Source             []Reference      `bson:"source,omitempty" json:"source,omitempty"`
}
type SubstanceReferenceInformationGeneElement struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Element           *Identifier      `bson:"element,omitempty" json:"element,omitempty"`
	Source            []Reference      `bson:"source,omitempty" json:"source,omitempty"`
}
type SubstanceReferenceInformationTarget struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Target            *Identifier      `bson:"target,omitempty" json:"target,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Interaction       *CodeableConcept `bson:"interaction,omitempty" json:"interaction,omitempty"`
	Organism          *CodeableConcept `bson:"organism,omitempty" json:"organism,omitempty"`
	OrganismType      *CodeableConcept `bson:"organismType,omitempty" json:"organismType,omitempty"`
	AmountQuantity    *Quantity        `bson:"amountQuantity,omitempty" json:"amountQuantity,omitempty"`
	AmountRange       *Range           `bson:"amountRange,omitempty" json:"amountRange,omitempty"`
	AmountString      *string          `bson:"amountString,omitempty" json:"amountString,omitempty"`
	AmountType        *CodeableConcept `bson:"amountType,omitempty" json:"amountType,omitempty"`
	Source            []Reference      `bson:"source,omitempty" json:"source,omitempty"`
}
type OtherSubstanceReferenceInformation SubstanceReferenceInformation

// MarshalJSON marshals the given SubstanceReferenceInformation as JSON into a byte slice
func (r SubstanceReferenceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceReferenceInformation
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceReferenceInformation: OtherSubstanceReferenceInformation(r),
		ResourceType:                       "SubstanceReferenceInformation",
	})
}
func (r SubstanceReferenceInformation) ResourceType() string {
	return "SubstanceReferenceInformation"
}
func (r SubstanceReferenceInformation) ResourceIdentifier() string {
	if r.Id != nil {
		return *r.Id
	}
	return ""
}

// UnmarshalSubstanceReferenceInformation unmarshals a SubstanceReferenceInformation.
func UnmarshalSubstanceReferenceInformation(b []byte) (SubstanceReferenceInformation, error) {
	var substanceReferenceInformation SubstanceReferenceInformation
	if err := json.Unmarshal(b, &substanceReferenceInformation); err != nil {
		return substanceReferenceInformation, err
	}
	return substanceReferenceInformation, nil
}
