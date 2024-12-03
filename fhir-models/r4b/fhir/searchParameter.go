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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// SearchParameter is documented here http://hl7.org/fhir/StructureDefinition/SearchParameter
type SearchParameter struct {
	Id                     *string                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                    `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    string                     `bson:"url" json:"url"`
	Identifier             []Identifier               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string                    `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString *string                    `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                    `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                   string                     `bson:"name" json:"name"`
	Title                  *string                    `bson:"title,omitempty" json:"title,omitempty"`
	DerivedFrom            *string                    `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Status                 PublicationStatus          `bson:"status" json:"status"`
	Experimental           *bool                      `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string                    `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string                    `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail            `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            string                     `bson:"description" json:"description"`
	UseContext             []UsageContext             `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept          `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose                *string                    `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright              *string                    `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel         *string                    `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	Code                   string                     `bson:"code" json:"code"`
	Base                   []string                   `bson:"base" json:"base"`
	Type                   SearchParamType            `bson:"type" json:"type"`
	Expression             *string                    `bson:"expression,omitempty" json:"expression,omitempty"`
	ProcessingMode         *SearchProcessingModeType  `bson:"processingMode,omitempty" json:"processingMode,omitempty"`
	Constraint             *string                    `bson:"constraint,omitempty" json:"constraint,omitempty"`
	Target                 []string                   `bson:"target,omitempty" json:"target,omitempty"`
	MultipleOr             *bool                      `bson:"multipleOr,omitempty" json:"multipleOr,omitempty"`
	MultipleAnd            *bool                      `bson:"multipleAnd,omitempty" json:"multipleAnd,omitempty"`
	Comparator             []SearchComparator         `bson:"comparator,omitempty" json:"comparator,omitempty"`
	Modifier               []SearchModifierCode       `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Chain                  []string                   `bson:"chain,omitempty" json:"chain,omitempty"`
	Component              []SearchParameterComponent `bson:"component,omitempty" json:"component,omitempty"`
}
type SearchParameterComponent struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Definition        string      `bson:"definition" json:"definition"`
	Expression        string      `bson:"expression" json:"expression"`
}
type OtherSearchParameter SearchParameter

// MarshalJSON marshals the given SearchParameter as JSON into a byte slice
func (r SearchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSearchParameter
		ResourceType string `json:"resourceType"`
	}{
		OtherSearchParameter: OtherSearchParameter(r),
		ResourceType:         "SearchParameter",
	})
}

// UnmarshalSearchParameter unmarshals a SearchParameter.
func UnmarshalSearchParameter(b []byte) (SearchParameter, error) {
	var searchParameter SearchParameter
	if err := json.Unmarshal(b, &searchParameter); err != nil {
		return searchParameter, err
	}
	return searchParameter, nil
}