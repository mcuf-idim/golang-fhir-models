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

// ConditionDefinition is documented here http://hl7.org/fhir/StructureDefinition/ConditionDefinition
type ConditionDefinition struct {
	Id                     *string                            `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                              `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                            `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                            `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                         `bson:"text,omitempty" json:"text,omitempty"`
	Contained              []json.RawMessage                  `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension              []Extension                        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string                            `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string                            `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString *string                            `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                            `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                            `bson:"name,omitempty" json:"name,omitempty"`
	Title                  *string                            `bson:"title,omitempty" json:"title,omitempty"`
	Subtitle               *string                            `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	Status                 PublicationStatus                  `bson:"status" json:"status"`
	Experimental           *bool                              `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string                            `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string                            `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail                    `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            *string                            `bson:"description,omitempty" json:"description,omitempty"`
	UseContext             []UsageContext                     `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                  `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Code                   CodeableConcept                    `bson:"code" json:"code"`
	Severity               *CodeableConcept                   `bson:"severity,omitempty" json:"severity,omitempty"`
	BodySite               *CodeableConcept                   `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Stage                  *CodeableConcept                   `bson:"stage,omitempty" json:"stage,omitempty"`
	HasSeverity            *bool                              `bson:"hasSeverity,omitempty" json:"hasSeverity,omitempty"`
	HasBodySite            *bool                              `bson:"hasBodySite,omitempty" json:"hasBodySite,omitempty"`
	HasStage               *bool                              `bson:"hasStage,omitempty" json:"hasStage,omitempty"`
	Definition             []string                           `bson:"definition,omitempty" json:"definition,omitempty"`
	Observation            []ConditionDefinitionObservation   `bson:"observation,omitempty" json:"observation,omitempty"`
	Medication             []ConditionDefinitionMedication    `bson:"medication,omitempty" json:"medication,omitempty"`
	Precondition           []ConditionDefinitionPrecondition  `bson:"precondition,omitempty" json:"precondition,omitempty"`
	Team                   []Reference                        `bson:"team,omitempty" json:"team,omitempty"`
	Questionnaire          []ConditionDefinitionQuestionnaire `bson:"questionnaire,omitempty" json:"questionnaire,omitempty"`
	Plan                   []ConditionDefinitionPlan          `bson:"plan,omitempty" json:"plan,omitempty"`
}

func (r ConditionDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type ConditionDefinitionObservation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}
type ConditionDefinitionMedication struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}
type ConditionDefinitionPrecondition struct {
	Id                   *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type                 ConditionPreconditionType `bson:"type" json:"type"`
	Code                 CodeableConcept           `bson:"code" json:"code"`
	ValueCodeableConcept *CodeableConcept          `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity                 `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
}
type ConditionDefinitionQuestionnaire struct {
	Id                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Purpose           ConditionQuestionnairePurpose `bson:"purpose" json:"purpose"`
	Reference         Reference                     `bson:"reference" json:"reference"`
}
type ConditionDefinitionPlan struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Reference         Reference        `bson:"reference" json:"reference"`
}
type OtherConditionDefinition ConditionDefinition

// MarshalJSON marshals the given ConditionDefinition as JSON into a byte slice
func (r ConditionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConditionDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherConditionDefinition: OtherConditionDefinition(r),
		ResourceType:             "ConditionDefinition",
	})
}
func (r ConditionDefinition) ResourceType() string {
	return "ConditionDefinition"
}

// UnmarshalConditionDefinition unmarshals a ConditionDefinition.
func UnmarshalConditionDefinition(b []byte) (ConditionDefinition, error) {
	var conditionDefinition ConditionDefinition
	if err := json.Unmarshal(b, &conditionDefinition); err != nil {
		return conditionDefinition, err
	}
	return conditionDefinition, nil
}
