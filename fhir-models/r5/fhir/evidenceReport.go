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

// EvidenceReport is documented here http://hl7.org/fhir/StructureDefinition/EvidenceReport
type EvidenceReport struct {
	Id                *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string                   `bson:"url,omitempty" json:"url,omitempty"`
	Status            PublicationStatus         `bson:"status" json:"status"`
	UseContext        []UsageContext            `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Identifier        []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	RelatedIdentifier []Identifier              `bson:"relatedIdentifier,omitempty" json:"relatedIdentifier,omitempty"`
	CiteAsReference   *Reference                `bson:"citeAsReference,omitempty" json:"citeAsReference,omitempty"`
	CiteAsMarkdown    *string                   `bson:"citeAsMarkdown,omitempty" json:"citeAsMarkdown,omitempty"`
	Type              *CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	Note              []Annotation              `bson:"note,omitempty" json:"note,omitempty"`
	RelatedArtifact   []RelatedArtifact         `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	Subject           EvidenceReportSubject     `bson:"subject" json:"subject"`
	Publisher         *string                   `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact           []ContactDetail           `bson:"contact,omitempty" json:"contact,omitempty"`
	Author            []ContactDetail           `bson:"author,omitempty" json:"author,omitempty"`
	Editor            []ContactDetail           `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer          []ContactDetail           `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser          []ContactDetail           `bson:"endorser,omitempty" json:"endorser,omitempty"`
	RelatesTo         []EvidenceReportRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Section           []EvidenceReportSection   `bson:"section,omitempty" json:"section,omitempty"`
}
type EvidenceReportSubject struct {
	Id                *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Characteristic    []EvidenceReportSubjectCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	Note              []Annotation                          `bson:"note,omitempty" json:"note,omitempty"`
}
type EvidenceReportSubjectCharacteristic struct {
	Id                   *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code                 CodeableConcept `bson:"code" json:"code"`
	ValueReference       Reference       `bson:"valueReference" json:"valueReference"`
	ValueCodeableConcept CodeableConcept `bson:"valueCodeableConcept" json:"valueCodeableConcept"`
	ValueBoolean         bool            `bson:"valueBoolean" json:"valueBoolean"`
	ValueQuantity        Quantity        `bson:"valueQuantity" json:"valueQuantity"`
	ValueRange           Range           `bson:"valueRange" json:"valueRange"`
	Exclude              *bool           `bson:"exclude,omitempty" json:"exclude,omitempty"`
	Period               *Period         `bson:"period,omitempty" json:"period,omitempty"`
}
type EvidenceReportRelatesTo struct {
	Id                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              ReportRelationshipType        `bson:"code" json:"code"`
	Target            EvidenceReportRelatesToTarget `bson:"target" json:"target"`
}
type EvidenceReportRelatesToTarget struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url               *string     `bson:"url,omitempty" json:"url,omitempty"`
	Identifier        *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Display           *string     `bson:"display,omitempty" json:"display,omitempty"`
	Resource          *Reference  `bson:"resource,omitempty" json:"resource,omitempty"`
}
type EvidenceReportSection struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Title             *string                 `bson:"title,omitempty" json:"title,omitempty"`
	Focus             *CodeableConcept        `bson:"focus,omitempty" json:"focus,omitempty"`
	FocusReference    *Reference              `bson:"focusReference,omitempty" json:"focusReference,omitempty"`
	Author            []Reference             `bson:"author,omitempty" json:"author,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Mode              *ListMode               `bson:"mode,omitempty" json:"mode,omitempty"`
	OrderedBy         *CodeableConcept        `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	EntryClassifier   []CodeableConcept       `bson:"entryClassifier,omitempty" json:"entryClassifier,omitempty"`
	EntryReference    []Reference             `bson:"entryReference,omitempty" json:"entryReference,omitempty"`
	EntryQuantity     []Quantity              `bson:"entryQuantity,omitempty" json:"entryQuantity,omitempty"`
	EmptyReason       *CodeableConcept        `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
	Section           []EvidenceReportSection `bson:"section,omitempty" json:"section,omitempty"`
}
type OtherEvidenceReport EvidenceReport

// MarshalJSON marshals the given EvidenceReport as JSON into a byte slice
func (r EvidenceReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidenceReport
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidenceReport: OtherEvidenceReport(r),
		ResourceType:        "EvidenceReport",
	})
}

// UnmarshalEvidenceReport unmarshals a EvidenceReport.
func UnmarshalEvidenceReport(b []byte) (EvidenceReport, error) {
	var evidenceReport EvidenceReport
	if err := json.Unmarshal(b, &evidenceReport); err != nil {
		return evidenceReport, err
	}
	return evidenceReport, nil
}
