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

// ActivityDefinition is documented here http://hl7.org/fhir/StructureDefinition/ActivityDefinition
type ActivityDefinition struct {
	Id                           *string                          `bson:"id,omitempty" json:"id,omitempty"`
	Meta                         *Meta                            `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules                *string                          `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                     *string                          `bson:"language,omitempty" json:"language,omitempty"`
	Text                         *Narrative                       `bson:"text,omitempty" json:"text,omitempty"`
	Extension                    []Extension                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension            []Extension                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                          *string                          `bson:"url,omitempty" json:"url,omitempty"`
	Identifier                   []Identifier                     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                      *string                          `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString       *string                          `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding       *Coding                          `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                         *string                          `bson:"name,omitempty" json:"name,omitempty"`
	Title                        *string                          `bson:"title,omitempty" json:"title,omitempty"`
	Subtitle                     *string                          `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	Status                       PublicationStatus                `bson:"status" json:"status"`
	Experimental                 *bool                            `bson:"experimental,omitempty" json:"experimental,omitempty"`
	SubjectCodeableConcept       *CodeableConcept                 `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty"`
	SubjectReference             *Reference                       `bson:"subjectReference,omitempty" json:"subjectReference,omitempty"`
	SubjectCanonical             *string                          `bson:"subjectCanonical,omitempty" json:"subjectCanonical,omitempty"`
	Date                         *string                          `bson:"date,omitempty" json:"date,omitempty"`
	Publisher                    *string                          `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                      []ContactDetail                  `bson:"contact,omitempty" json:"contact,omitempty"`
	Description                  *string                          `bson:"description,omitempty" json:"description,omitempty"`
	UseContext                   []UsageContext                   `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction                 []CodeableConcept                `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose                      *string                          `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Usage                        *string                          `bson:"usage,omitempty" json:"usage,omitempty"`
	Copyright                    *string                          `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel               *string                          `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	ApprovalDate                 *string                          `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate               *string                          `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod              *Period                          `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Topic                        []CodeableConcept                `bson:"topic,omitempty" json:"topic,omitempty"`
	Author                       []ContactDetail                  `bson:"author,omitempty" json:"author,omitempty"`
	Editor                       []ContactDetail                  `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer                     []ContactDetail                  `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser                     []ContactDetail                  `bson:"endorser,omitempty" json:"endorser,omitempty"`
	RelatedArtifact              []RelatedArtifact                `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Library                      []string                         `bson:"library,omitempty" json:"library,omitempty"`
	Kind                         *RequestResourceTypes            `bson:"kind,omitempty" json:"kind,omitempty"`
	Profile                      *string                          `bson:"profile,omitempty" json:"profile,omitempty"`
	Code                         *CodeableConcept                 `bson:"code,omitempty" json:"code,omitempty"`
	Intent                       *RequestIntent                   `bson:"intent,omitempty" json:"intent,omitempty"`
	Priority                     *RequestPriority                 `bson:"priority,omitempty" json:"priority,omitempty"`
	DoNotPerform                 *bool                            `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	TimingTiming                 *Timing                          `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	TimingAge                    *Age                             `bson:"timingAge,omitempty" json:"timingAge,omitempty"`
	TimingRange                  *Range                           `bson:"timingRange,omitempty" json:"timingRange,omitempty"`
	TimingDuration               *Duration                        `bson:"timingDuration,omitempty" json:"timingDuration,omitempty"`
	AsNeededBoolean              *bool                            `bson:"asNeededBoolean,omitempty" json:"asNeededBoolean,omitempty"`
	AsNeededCodeableConcept      *CodeableConcept                 `bson:"asNeededCodeableConcept,omitempty" json:"asNeededCodeableConcept,omitempty"`
	Location                     *CodeableReference               `bson:"location,omitempty" json:"location,omitempty"`
	Participant                  []ActivityDefinitionParticipant  `bson:"participant,omitempty" json:"participant,omitempty"`
	ProductReference             *Reference                       `bson:"productReference,omitempty" json:"productReference,omitempty"`
	ProductCodeableConcept       *CodeableConcept                 `bson:"productCodeableConcept,omitempty" json:"productCodeableConcept,omitempty"`
	Quantity                     *Quantity                        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Dosage                       []Dosage                         `bson:"dosage,omitempty" json:"dosage,omitempty"`
	BodySite                     []CodeableConcept                `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SpecimenRequirement          []string                         `bson:"specimenRequirement,omitempty" json:"specimenRequirement,omitempty"`
	ObservationRequirement       []string                         `bson:"observationRequirement,omitempty" json:"observationRequirement,omitempty"`
	ObservationResultRequirement []string                         `bson:"observationResultRequirement,omitempty" json:"observationResultRequirement,omitempty"`
	Transform                    *string                          `bson:"transform,omitempty" json:"transform,omitempty"`
	DynamicValue                 []ActivityDefinitionDynamicValue `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty"`
}
type ActivityDefinitionParticipant struct {
	Id                *string                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *ActionParticipantType `bson:"type,omitempty" json:"type,omitempty"`
	TypeCanonical     *string                `bson:"typeCanonical,omitempty" json:"typeCanonical,omitempty"`
	TypeReference     *Reference             `bson:"typeReference,omitempty" json:"typeReference,omitempty"`
	Role              *CodeableConcept       `bson:"role,omitempty" json:"role,omitempty"`
	Function          *CodeableConcept       `bson:"function,omitempty" json:"function,omitempty"`
}
type ActivityDefinitionDynamicValue struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Path              string      `bson:"path" json:"path"`
	Expression        Expression  `bson:"expression" json:"expression"`
}
type OtherActivityDefinition ActivityDefinition

// MarshalJSON marshals the given ActivityDefinition as JSON into a byte slice
func (r ActivityDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherActivityDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherActivityDefinition: OtherActivityDefinition(r),
		ResourceType:            "ActivityDefinition",
	})
}
func (r ActivityDefinition) ResourceType() string {
	return "ActivityDefinition"
}

// UnmarshalActivityDefinition unmarshals a ActivityDefinition.
func UnmarshalActivityDefinition(b []byte) (ActivityDefinition, error) {
	var activityDefinition ActivityDefinition
	if err := json.Unmarshal(b, &activityDefinition); err != nil {
		return activityDefinition, err
	}
	return activityDefinition, nil
}
