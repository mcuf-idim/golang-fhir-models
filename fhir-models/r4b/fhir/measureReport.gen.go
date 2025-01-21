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

// MeasureReport is documented here http://hl7.org/fhir/StructureDefinition/MeasureReport
type MeasureReport struct {
	Id                  *string               `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string               `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string               `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative            `bson:"text,omitempty" json:"text,omitempty"`
	Contained           []json.RawMessage     `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension           []Extension           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              MeasureReportStatus   `bson:"status" json:"status"`
	Type                MeasureReportType     `bson:"type" json:"type"`
	DataUpdateType      *SubmitDataUpdateType `bson:"dataUpdateType,omitempty" json:"dataUpdateType,omitempty"`
	Measure             *string               `bson:"measure,omitempty" json:"measure,omitempty"`
	Subject             *Reference            `bson:"subject,omitempty" json:"subject,omitempty"`
	Date                *string               `bson:"date,omitempty" json:"date,omitempty"`
	Reporter            *Reference            `bson:"reporter,omitempty" json:"reporter,omitempty"`
	ReportingVendor     *Reference            `bson:"reportingVendor,omitempty" json:"reportingVendor,omitempty"`
	Location            *Reference            `bson:"location,omitempty" json:"location,omitempty"`
	Period              Period                `bson:"period" json:"period"`
	InputParameters     *Reference            `bson:"inputParameters,omitempty" json:"inputParameters,omitempty"`
	Scoring             *CodeableConcept      `bson:"scoring,omitempty" json:"scoring,omitempty"`
	ImprovementNotation *CodeableConcept      `bson:"improvementNotation,omitempty" json:"improvementNotation,omitempty"`
	Group               []MeasureReportGroup  `bson:"group,omitempty" json:"group,omitempty"`
	SupplementalData    []Reference           `bson:"supplementalData,omitempty" json:"supplementalData,omitempty"`
	EvaluatedResource   []Reference           `bson:"evaluatedResource,omitempty" json:"evaluatedResource,omitempty"`
}

func (r MeasureReport) ContainedResources() []json.RawMessage {
	return r.Contained
}

type MeasureReportGroup struct {
	Id                          *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension                   []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension           []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId                      *string                        `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Code                        *CodeableConcept               `bson:"code,omitempty" json:"code,omitempty"`
	Subject                     *Reference                     `bson:"subject,omitempty" json:"subject,omitempty"`
	Population                  []MeasureReportGroupPopulation `bson:"population,omitempty" json:"population,omitempty"`
	MeasureScoreQuantity        *Quantity                      `bson:"measureScoreQuantity,omitempty" json:"measureScoreQuantity,omitempty"`
	MeasureScoreDateTime        *string                        `bson:"measureScoreDateTime,omitempty" json:"measureScoreDateTime,omitempty"`
	MeasureScoreCodeableConcept *CodeableConcept               `bson:"measureScoreCodeableConcept,omitempty" json:"measureScoreCodeableConcept,omitempty"`
	MeasureScorePeriod          *Period                        `bson:"measureScorePeriod,omitempty" json:"measureScorePeriod,omitempty"`
	MeasureScoreRange           *Range                         `bson:"measureScoreRange,omitempty" json:"measureScoreRange,omitempty"`
	MeasureScoreDuration        *Duration                      `bson:"measureScoreDuration,omitempty" json:"measureScoreDuration,omitempty"`
	Stratifier                  []MeasureReportGroupStratifier `bson:"stratifier,omitempty" json:"stratifier,omitempty"`
}
type MeasureReportGroupPopulation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId            *string          `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Count             *int             `bson:"count,omitempty" json:"count,omitempty"`
	SubjectResults    *Reference       `bson:"subjectResults,omitempty" json:"subjectResults,omitempty"`
	SubjectReport     []Reference      `bson:"subjectReport,omitempty" json:"subjectReport,omitempty"`
	Subjects          *Reference       `bson:"subjects,omitempty" json:"subjects,omitempty"`
}
type MeasureReportGroupStratifier struct {
	Id                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId            *string                               `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Code              *CodeableConcept                      `bson:"code,omitempty" json:"code,omitempty"`
	Stratum           []MeasureReportGroupStratifierStratum `bson:"stratum,omitempty" json:"stratum,omitempty"`
}
type MeasureReportGroupStratifierStratum struct {
	Id                          *string                                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension                   []Extension                                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension           []Extension                                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValueCodeableConcept        *CodeableConcept                                `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueBoolean                *bool                                           `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueQuantity               *Quantity                                       `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange                  *Range                                          `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueReference              *Reference                                      `bson:"valueReference,omitempty" json:"valueReference,omitempty"`
	Component                   []MeasureReportGroupStratifierStratumComponent  `bson:"component,omitempty" json:"component,omitempty"`
	Population                  []MeasureReportGroupStratifierStratumPopulation `bson:"population,omitempty" json:"population,omitempty"`
	MeasureScoreQuantity        *Quantity                                       `bson:"measureScoreQuantity,omitempty" json:"measureScoreQuantity,omitempty"`
	MeasureScoreDateTime        *string                                         `bson:"measureScoreDateTime,omitempty" json:"measureScoreDateTime,omitempty"`
	MeasureScoreCodeableConcept *CodeableConcept                                `bson:"measureScoreCodeableConcept,omitempty" json:"measureScoreCodeableConcept,omitempty"`
	MeasureScorePeriod          *Period                                         `bson:"measureScorePeriod,omitempty" json:"measureScorePeriod,omitempty"`
	MeasureScoreRange           *Range                                          `bson:"measureScoreRange,omitempty" json:"measureScoreRange,omitempty"`
	MeasureScoreDuration        *Duration                                       `bson:"measureScoreDuration,omitempty" json:"measureScoreDuration,omitempty"`
}
type MeasureReportGroupStratifierStratumComponent struct {
	Id                   *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId               *string         `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Code                 CodeableConcept `bson:"code" json:"code"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange"`
	ValueReference       Reference       `bson:"valueReference" json:"valueReference"`
}
type MeasureReportGroupStratifierStratumPopulation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId            *string          `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Count             *int             `bson:"count,omitempty" json:"count,omitempty"`
	SubjectResults    *Reference       `bson:"subjectResults,omitempty" json:"subjectResults,omitempty"`
	SubjectReport     []Reference      `bson:"subjectReport,omitempty" json:"subjectReport,omitempty"`
	Subjects          *Reference       `bson:"subjects,omitempty" json:"subjects,omitempty"`
}
type OtherMeasureReport MeasureReport

// MarshalJSON marshals the given MeasureReport as JSON into a byte slice
func (r MeasureReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeasureReport
		ResourceType string `json:"resourceType"`
	}{
		OtherMeasureReport: OtherMeasureReport(r),
		ResourceType:       "MeasureReport",
	})
}
func (r MeasureReport) ResourceType() string {
	return "MeasureReport"
}
func (r MeasureReport) ResourceIdentifier() string {
	if r.Id != nil {
		return *r.Id
	}
	return ""
}

// UnmarshalMeasureReport unmarshals a MeasureReport.
func UnmarshalMeasureReport(b []byte) (MeasureReport, error) {
	var measureReport MeasureReport
	if err := json.Unmarshal(b, &measureReport); err != nil {
		return measureReport, err
	}
	return measureReport, nil
}
