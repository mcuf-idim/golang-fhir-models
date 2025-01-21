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

// GraphDefinition is documented here http://hl7.org/fhir/StructureDefinition/GraphDefinition
type GraphDefinition struct {
	Id                     *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Contained              []json.RawMessage     `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension              []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string               `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string               `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString *string               `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding               `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                   string                `bson:"name" json:"name"`
	Title                  *string               `bson:"title,omitempty" json:"title,omitempty"`
	Status                 PublicationStatus     `bson:"status" json:"status"`
	Experimental           *bool                 `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string               `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string               `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail       `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            *string               `bson:"description,omitempty" json:"description,omitempty"`
	UseContext             []UsageContext        `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept     `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose                *string               `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright              *string               `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel         *string               `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	Start                  *string               `bson:"start,omitempty" json:"start,omitempty"`
	Node                   []GraphDefinitionNode `bson:"node,omitempty" json:"node,omitempty"`
	Link                   []GraphDefinitionLink `bson:"link,omitempty" json:"link,omitempty"`
}

func (r GraphDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type GraphDefinitionNode struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	NodeId            string      `bson:"nodeId" json:"nodeId"`
	Description       *string     `bson:"description,omitempty" json:"description,omitempty"`
	Type              string      `bson:"type" json:"type"`
	Profile           *string     `bson:"profile,omitempty" json:"profile,omitempty"`
}
type GraphDefinitionLink struct {
	Id                *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string                          `bson:"description,omitempty" json:"description,omitempty"`
	Min               *int                             `bson:"min,omitempty" json:"min,omitempty"`
	Max               *string                          `bson:"max,omitempty" json:"max,omitempty"`
	SourceId          string                           `bson:"sourceId" json:"sourceId"`
	Path              *string                          `bson:"path,omitempty" json:"path,omitempty"`
	SliceName         *string                          `bson:"sliceName,omitempty" json:"sliceName,omitempty"`
	TargetId          string                           `bson:"targetId" json:"targetId"`
	Params            *string                          `bson:"params,omitempty" json:"params,omitempty"`
	Compartment       []GraphDefinitionLinkCompartment `bson:"compartment,omitempty" json:"compartment,omitempty"`
}
type GraphDefinitionLinkCompartment struct {
	Id                *string              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Use               GraphCompartmentUse  `bson:"use" json:"use"`
	Rule              GraphCompartmentRule `bson:"rule" json:"rule"`
	Code              CompartmentType      `bson:"code" json:"code"`
	Expression        *string              `bson:"expression,omitempty" json:"expression,omitempty"`
	Description       *string              `bson:"description,omitempty" json:"description,omitempty"`
}
type OtherGraphDefinition GraphDefinition

// MarshalJSON marshals the given GraphDefinition as JSON into a byte slice
func (r GraphDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGraphDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherGraphDefinition: OtherGraphDefinition(r),
		ResourceType:         "GraphDefinition",
	})
}
func (r GraphDefinition) ResourceType() string {
	return "GraphDefinition"
}
func (r GraphDefinition) ResourceIdentifier() string {
	if r.Id != nil {
		return *r.Id
	}
	return ""
}

// UnmarshalGraphDefinition unmarshals a GraphDefinition.
func UnmarshalGraphDefinition(b []byte) (GraphDefinition, error) {
	var graphDefinition GraphDefinition
	if err := json.Unmarshal(b, &graphDefinition); err != nil {
		return graphDefinition, err
	}
	return graphDefinition, nil
}
