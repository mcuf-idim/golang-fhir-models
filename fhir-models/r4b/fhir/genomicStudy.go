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

// GenomicStudy is documented here http://hl7.org/fhir/StructureDefinition/GenomicStudy
type GenomicStudy struct {
	Id                    *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                GenomicStudyStatus     `bson:"status" json:"status"`
	Type                  []CodeableConcept      `bson:"type,omitempty" json:"type,omitempty"`
	Subject               Reference              `bson:"subject" json:"subject"`
	Encounter             *Reference             `bson:"encounter,omitempty" json:"encounter,omitempty"`
	StartDate             *string                `bson:"startDate,omitempty" json:"startDate,omitempty"`
	BasedOn               []Reference            `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Referrer              *Reference             `bson:"referrer,omitempty" json:"referrer,omitempty"`
	Interpreter           []Reference            `bson:"interpreter,omitempty" json:"interpreter,omitempty"`
	Reason                []CodeableReference    `bson:"reason,omitempty" json:"reason,omitempty"`
	InstantiatesCanonical *string                `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       *string                `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	Note                  []Annotation           `bson:"note,omitempty" json:"note,omitempty"`
	Description           *string                `bson:"description,omitempty" json:"description,omitempty"`
	Analysis              []GenomicStudyAnalysis `bson:"analysis,omitempty" json:"analysis,omitempty"`
}
type GenomicStudyAnalysis struct {
	Id                    *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	MethodType            []CodeableConcept               `bson:"methodType,omitempty" json:"methodType,omitempty"`
	ChangeType            []CodeableConcept               `bson:"changeType,omitempty" json:"changeType,omitempty"`
	GenomeBuild           *CodeableConcept                `bson:"genomeBuild,omitempty" json:"genomeBuild,omitempty"`
	InstantiatesCanonical *string                         `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       *string                         `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	Title                 *string                         `bson:"title,omitempty" json:"title,omitempty"`
	Focus                 []Reference                     `bson:"focus,omitempty" json:"focus,omitempty"`
	Specimen              []Reference                     `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Date                  *string                         `bson:"date,omitempty" json:"date,omitempty"`
	Note                  []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
	ProtocolPerformed     *Reference                      `bson:"protocolPerformed,omitempty" json:"protocolPerformed,omitempty"`
	RegionsStudied        []Reference                     `bson:"regionsStudied,omitempty" json:"regionsStudied,omitempty"`
	RegionsCalled         []Reference                     `bson:"regionsCalled,omitempty" json:"regionsCalled,omitempty"`
	Input                 []GenomicStudyAnalysisInput     `bson:"input,omitempty" json:"input,omitempty"`
	Output                []GenomicStudyAnalysisOutput    `bson:"output,omitempty" json:"output,omitempty"`
	Performer             []GenomicStudyAnalysisPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	Device                []GenomicStudyAnalysisDevice    `bson:"device,omitempty" json:"device,omitempty"`
}
type GenomicStudyAnalysisInput struct {
	Id                    *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	File                  *Reference       `bson:"file,omitempty" json:"file,omitempty"`
	Type                  *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	GeneratedByIdentifier *Identifier      `bson:"generatedByIdentifier,omitempty" json:"generatedByIdentifier,omitempty"`
	GeneratedByReference  *Reference       `bson:"generatedByReference,omitempty" json:"generatedByReference,omitempty"`
}
type GenomicStudyAnalysisOutput struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	File              *Reference       `bson:"file,omitempty" json:"file,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}
type GenomicStudyAnalysisPerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Actor             *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}
type GenomicStudyAnalysisDevice struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Device            *Reference       `bson:"device,omitempty" json:"device,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
}
type OtherGenomicStudy GenomicStudy

// MarshalJSON marshals the given GenomicStudy as JSON into a byte slice
func (r GenomicStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGenomicStudy
		ResourceType string `json:"resourceType"`
	}{
		OtherGenomicStudy: OtherGenomicStudy(r),
		ResourceType:      "GenomicStudy",
	})
}

// UnmarshalGenomicStudy unmarshals a GenomicStudy.
func UnmarshalGenomicStudy(b []byte) (GenomicStudy, error) {
	var genomicStudy GenomicStudy
	if err := json.Unmarshal(b, &genomicStudy); err != nil {
		return genomicStudy, err
	}
	return genomicStudy, nil
}
