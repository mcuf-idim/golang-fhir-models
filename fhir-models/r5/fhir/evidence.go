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

// Evidence is documented here http://hl7.org/fhir/StructureDefinition/Evidence
type Evidence struct {
	Id                     *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string                      `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string                      `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString *string                      `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                      `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                      `bson:"name,omitempty" json:"name,omitempty"`
	Title                  *string                      `bson:"title,omitempty" json:"title,omitempty"`
	CiteAsReference        *Reference                   `bson:"citeAsReference,omitempty" json:"citeAsReference,omitempty"`
	CiteAsMarkdown         *string                      `bson:"citeAsMarkdown,omitempty" json:"citeAsMarkdown,omitempty"`
	Status                 PublicationStatus            `bson:"status" json:"status"`
	Experimental           *bool                        `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string                      `bson:"date,omitempty" json:"date,omitempty"`
	ApprovalDate           *string                      `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate         *string                      `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	Publisher              *string                      `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail              `bson:"contact,omitempty" json:"contact,omitempty"`
	Author                 []ContactDetail              `bson:"author,omitempty" json:"author,omitempty"`
	Editor                 []ContactDetail              `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer               []ContactDetail              `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser               []ContactDetail              `bson:"endorser,omitempty" json:"endorser,omitempty"`
	UseContext             []UsageContext               `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Purpose                *string                      `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright              *string                      `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel         *string                      `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	RelatedArtifact        []RelatedArtifact            `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Description            *string                      `bson:"description,omitempty" json:"description,omitempty"`
	Assertion              *string                      `bson:"assertion,omitempty" json:"assertion,omitempty"`
	Note                   []Annotation                 `bson:"note,omitempty" json:"note,omitempty"`
	VariableDefinition     []EvidenceVariableDefinition `bson:"variableDefinition" json:"variableDefinition"`
	SynthesisType          *CodeableConcept             `bson:"synthesisType,omitempty" json:"synthesisType,omitempty"`
	StudyDesign            []CodeableConcept            `bson:"studyDesign,omitempty" json:"studyDesign,omitempty"`
	Statistic              []EvidenceStatistic          `bson:"statistic,omitempty" json:"statistic,omitempty"`
	Certainty              []EvidenceCertainty          `bson:"certainty,omitempty" json:"certainty,omitempty"`
}
type EvidenceVariableDefinition struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string          `bson:"description,omitempty" json:"description,omitempty"`
	Note              []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
	VariableRole      CodeableConcept  `bson:"variableRole" json:"variableRole"`
	Observed          *Reference       `bson:"observed,omitempty" json:"observed,omitempty"`
	Intended          *Reference       `bson:"intended,omitempty" json:"intended,omitempty"`
	DirectnessMatch   *CodeableConcept `bson:"directnessMatch,omitempty" json:"directnessMatch,omitempty"`
}
type EvidenceStatistic struct {
	Id                  *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description         *string                                `bson:"description,omitempty" json:"description,omitempty"`
	Note                []Annotation                           `bson:"note,omitempty" json:"note,omitempty"`
	StatisticType       *CodeableConcept                       `bson:"statisticType,omitempty" json:"statisticType,omitempty"`
	Category            *CodeableConcept                       `bson:"category,omitempty" json:"category,omitempty"`
	Quantity            *Quantity                              `bson:"quantity,omitempty" json:"quantity,omitempty"`
	NumberOfEvents      *int                                   `bson:"numberOfEvents,omitempty" json:"numberOfEvents,omitempty"`
	NumberAffected      *int                                   `bson:"numberAffected,omitempty" json:"numberAffected,omitempty"`
	SampleSize          *EvidenceStatisticSampleSize           `bson:"sampleSize,omitempty" json:"sampleSize,omitempty"`
	AttributeEstimate   []EvidenceStatisticAttributeEstimate   `bson:"attributeEstimate,omitempty" json:"attributeEstimate,omitempty"`
	ModelCharacteristic []EvidenceStatisticModelCharacteristic `bson:"modelCharacteristic,omitempty" json:"modelCharacteristic,omitempty"`
}
type EvidenceStatisticSampleSize struct {
	Id                   *string      `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description          *string      `bson:"description,omitempty" json:"description,omitempty"`
	Note                 []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	NumberOfStudies      *int         `bson:"numberOfStudies,omitempty" json:"numberOfStudies,omitempty"`
	NumberOfParticipants *int         `bson:"numberOfParticipants,omitempty" json:"numberOfParticipants,omitempty"`
	KnownDataCount       *int         `bson:"knownDataCount,omitempty" json:"knownDataCount,omitempty"`
}
type EvidenceStatisticAttributeEstimate struct {
	Id                *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string                              `bson:"description,omitempty" json:"description,omitempty"`
	Note              []Annotation                         `bson:"note,omitempty" json:"note,omitempty"`
	Type              *CodeableConcept                     `bson:"type,omitempty" json:"type,omitempty"`
	Quantity          *Quantity                            `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Level             *json.Number                         `bson:"level,omitempty" json:"level,omitempty"`
	Range             *Range                               `bson:"range,omitempty" json:"range,omitempty"`
	AttributeEstimate []EvidenceStatisticAttributeEstimate `bson:"attributeEstimate,omitempty" json:"attributeEstimate,omitempty"`
}
type EvidenceStatisticModelCharacteristic struct {
	Id                *string                                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                `bson:"code" json:"code"`
	Value             *Quantity                                      `bson:"value,omitempty" json:"value,omitempty"`
	Variable          []EvidenceStatisticModelCharacteristicVariable `bson:"variable,omitempty" json:"variable,omitempty"`
	AttributeEstimate []EvidenceStatisticAttributeEstimate           `bson:"attributeEstimate,omitempty" json:"attributeEstimate,omitempty"`
}
type EvidenceStatisticModelCharacteristicVariable struct {
	Id                 *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	VariableDefinition Reference                 `bson:"variableDefinition" json:"variableDefinition"`
	Handling           *EvidenceVariableHandling `bson:"handling,omitempty" json:"handling,omitempty"`
	ValueCategory      []CodeableConcept         `bson:"valueCategory,omitempty" json:"valueCategory,omitempty"`
	ValueQuantity      []Quantity                `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange         []Range                   `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
}
type EvidenceCertainty struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string             `bson:"description,omitempty" json:"description,omitempty"`
	Note              []Annotation        `bson:"note,omitempty" json:"note,omitempty"`
	Type              *CodeableConcept    `bson:"type,omitempty" json:"type,omitempty"`
	Rating            *CodeableConcept    `bson:"rating,omitempty" json:"rating,omitempty"`
	Rater             *string             `bson:"rater,omitempty" json:"rater,omitempty"`
	Subcomponent      []EvidenceCertainty `bson:"subcomponent,omitempty" json:"subcomponent,omitempty"`
}
type OtherEvidence Evidence

// MarshalJSON marshals the given Evidence as JSON into a byte slice
func (r Evidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidence
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidence: OtherEvidence(r),
		ResourceType:  "Evidence",
	})
}

// UnmarshalEvidence unmarshals a Evidence.
func UnmarshalEvidence(b []byte) (Evidence, error) {
	var evidence Evidence
	if err := json.Unmarshal(b, &evidence); err != nil {
		return evidence, err
	}
	return evidence, nil
}
