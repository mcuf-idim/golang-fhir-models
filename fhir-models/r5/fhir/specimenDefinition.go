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

// SpecimenDefinition is documented here http://hl7.org/fhir/StructureDefinition/SpecimenDefinition
type SpecimenDefinition struct {
	Id                     *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string                        `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             *Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string                        `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString *string                        `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                        `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                        `bson:"name,omitempty" json:"name,omitempty"`
	Title                  *string                        `bson:"title,omitempty" json:"title,omitempty"`
	DerivedFromCanonical   []string                       `bson:"derivedFromCanonical,omitempty" json:"derivedFromCanonical,omitempty"`
	DerivedFromUri         []string                       `bson:"derivedFromUri,omitempty" json:"derivedFromUri,omitempty"`
	Status                 PublicationStatus              `bson:"status" json:"status"`
	Experimental           *bool                          `bson:"experimental,omitempty" json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept               `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                     `bson:"subjectReference,omitempty" json:"subjectReference,omitempty"`
	Date                   *string                        `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string                        `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail                `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            *string                        `bson:"description,omitempty" json:"description,omitempty"`
	UseContext             []UsageContext                 `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept              `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose                *string                        `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright              *string                        `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel         *string                        `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                        `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate         *string                        `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                        `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	TypeCollected          *CodeableConcept               `bson:"typeCollected,omitempty" json:"typeCollected,omitempty"`
	PatientPreparation     []CodeableConcept              `bson:"patientPreparation,omitempty" json:"patientPreparation,omitempty"`
	TimeAspect             *string                        `bson:"timeAspect,omitempty" json:"timeAspect,omitempty"`
	Collection             []CodeableConcept              `bson:"collection,omitempty" json:"collection,omitempty"`
	TypeTested             []SpecimenDefinitionTypeTested `bson:"typeTested,omitempty" json:"typeTested,omitempty"`
}
type SpecimenDefinitionTypeTested struct {
	Id                 *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	IsDerived          *bool                                  `bson:"isDerived,omitempty" json:"isDerived,omitempty"`
	Type               *CodeableConcept                       `bson:"type,omitempty" json:"type,omitempty"`
	Preference         SpecimenContainedPreference            `bson:"preference" json:"preference"`
	Container          *SpecimenDefinitionTypeTestedContainer `bson:"container,omitempty" json:"container,omitempty"`
	Requirement        *string                                `bson:"requirement,omitempty" json:"requirement,omitempty"`
	RetentionTime      *Duration                              `bson:"retentionTime,omitempty" json:"retentionTime,omitempty"`
	SingleUse          *bool                                  `bson:"singleUse,omitempty" json:"singleUse,omitempty"`
	RejectionCriterion []CodeableConcept                      `bson:"rejectionCriterion,omitempty" json:"rejectionCriterion,omitempty"`
	Handling           []SpecimenDefinitionTypeTestedHandling `bson:"handling,omitempty" json:"handling,omitempty"`
	TestingDestination []CodeableConcept                      `bson:"testingDestination,omitempty" json:"testingDestination,omitempty"`
}
type SpecimenDefinitionTypeTestedContainer struct {
	Id                    *string                                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension                                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Material              *CodeableConcept                                `bson:"material,omitempty" json:"material,omitempty"`
	Type                  *CodeableConcept                                `bson:"type,omitempty" json:"type,omitempty"`
	Cap                   *CodeableConcept                                `bson:"cap,omitempty" json:"cap,omitempty"`
	Description           *string                                         `bson:"description,omitempty" json:"description,omitempty"`
	Capacity              *Quantity                                       `bson:"capacity,omitempty" json:"capacity,omitempty"`
	MinimumVolumeQuantity *Quantity                                       `bson:"minimumVolumeQuantity,omitempty" json:"minimumVolumeQuantity,omitempty"`
	MinimumVolumeString   *string                                         `bson:"minimumVolumeString,omitempty" json:"minimumVolumeString,omitempty"`
	Additive              []SpecimenDefinitionTypeTestedContainerAdditive `bson:"additive,omitempty" json:"additive,omitempty"`
	Preparation           *string                                         `bson:"preparation,omitempty" json:"preparation,omitempty"`
}
type SpecimenDefinitionTypeTestedContainerAdditive struct {
	Id                      *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	AdditiveCodeableConcept CodeableConcept `bson:"additiveCodeableConcept" json:"additiveCodeableConcept"`
	AdditiveReference       Reference       `bson:"additiveReference" json:"additiveReference"`
}
type SpecimenDefinitionTypeTestedHandling struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	TemperatureQualifier *CodeableConcept `bson:"temperatureQualifier,omitempty" json:"temperatureQualifier,omitempty"`
	TemperatureRange     *Range           `bson:"temperatureRange,omitempty" json:"temperatureRange,omitempty"`
	MaxDuration          *Duration        `bson:"maxDuration,omitempty" json:"maxDuration,omitempty"`
	Instruction          *string          `bson:"instruction,omitempty" json:"instruction,omitempty"`
}
type OtherSpecimenDefinition SpecimenDefinition

// MarshalJSON marshals the given SpecimenDefinition as JSON into a byte slice
func (r SpecimenDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSpecimenDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherSpecimenDefinition: OtherSpecimenDefinition(r),
		ResourceType:            "SpecimenDefinition",
	})
}

// UnmarshalSpecimenDefinition unmarshals a SpecimenDefinition.
func UnmarshalSpecimenDefinition(b []byte) (SpecimenDefinition, error) {
	var specimenDefinition SpecimenDefinition
	if err := json.Unmarshal(b, &specimenDefinition); err != nil {
		return specimenDefinition, err
	}
	return specimenDefinition, nil
}
