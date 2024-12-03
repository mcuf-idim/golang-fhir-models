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

// SubstanceProtein is documented here http://hl7.org/fhir/StructureDefinition/SubstanceProtein
type SubstanceProtein struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	SequenceType      *CodeableConcept          `bson:"sequenceType,omitempty" json:"sequenceType,omitempty"`
	NumberOfSubunits  *int                      `bson:"numberOfSubunits,omitempty" json:"numberOfSubunits,omitempty"`
	DisulfideLinkage  []string                  `bson:"disulfideLinkage,omitempty" json:"disulfideLinkage,omitempty"`
	Subunit           []SubstanceProteinSubunit `bson:"subunit,omitempty" json:"subunit,omitempty"`
}
type SubstanceProteinSubunit struct {
	Id                      *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Subunit                 *int        `bson:"subunit,omitempty" json:"subunit,omitempty"`
	Sequence                *string     `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Length                  *int        `bson:"length,omitempty" json:"length,omitempty"`
	SequenceAttachment      *Attachment `bson:"sequenceAttachment,omitempty" json:"sequenceAttachment,omitempty"`
	NTerminalModificationId *Identifier `bson:"nTerminalModificationId,omitempty" json:"nTerminalModificationId,omitempty"`
	NTerminalModification   *string     `bson:"nTerminalModification,omitempty" json:"nTerminalModification,omitempty"`
	CTerminalModificationId *Identifier `bson:"cTerminalModificationId,omitempty" json:"cTerminalModificationId,omitempty"`
	CTerminalModification   *string     `bson:"cTerminalModification,omitempty" json:"cTerminalModification,omitempty"`
}
type OtherSubstanceProtein SubstanceProtein

// MarshalJSON marshals the given SubstanceProtein as JSON into a byte slice
func (r SubstanceProtein) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceProtein
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceProtein: OtherSubstanceProtein(r),
		ResourceType:          "SubstanceProtein",
	})
}

// UnmarshalSubstanceProtein unmarshals a SubstanceProtein.
func UnmarshalSubstanceProtein(b []byte) (SubstanceProtein, error) {
	var substanceProtein SubstanceProtein
	if err := json.Unmarshal(b, &substanceProtein); err != nil {
		return substanceProtein, err
	}
	return substanceProtein, nil
}
