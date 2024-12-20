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

// CanonicalResource is documented here http://hl7.org/fhir/StructureDefinition/CanonicalResource
type CanonicalResource struct {
	Id                     *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Contained              []json.RawMessage `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension              []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string           `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string           `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString *string           `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding           `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                   *string           `bson:"name,omitempty" json:"name,omitempty"`
	Title                  *string           `bson:"title,omitempty" json:"title,omitempty"`
	Status                 PublicationStatus `bson:"status" json:"status"`
	Experimental           *bool             `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string           `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string           `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail   `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            *string           `bson:"description,omitempty" json:"description,omitempty"`
	UseContext             []UsageContext    `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose                *string           `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright              *string           `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel         *string           `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
}

func (r CanonicalResource) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCanonicalResource CanonicalResource

// MarshalJSON marshals the given CanonicalResource as JSON into a byte slice
func (r CanonicalResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCanonicalResource
		ResourceType string `json:"resourceType"`
	}{
		OtherCanonicalResource: OtherCanonicalResource(r),
		ResourceType:           "CanonicalResource",
	})
}
func (r CanonicalResource) ResourceType() string {
	return "CanonicalResource"
}

// UnmarshalCanonicalResource unmarshals a CanonicalResource.
func UnmarshalCanonicalResource(b []byte) (CanonicalResource, error) {
	var canonicalResource CanonicalResource
	if err := json.Unmarshal(b, &canonicalResource); err != nil {
		return canonicalResource, err
	}
	return canonicalResource, nil
}
