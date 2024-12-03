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

// Encounter is documented here http://hl7.org/fhir/StructureDefinition/Encounter
type Encounter struct {
	Id                 *string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta               *Meta                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules      *string                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language           *string                `bson:"language,omitempty" json:"language,omitempty"`
	Text               *Narrative             `bson:"text,omitempty" json:"text,omitempty"`
	Extension          []Extension            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier         []Identifier           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status             EncounterStatus        `bson:"status" json:"status"`
	Class              []CodeableConcept      `bson:"class,omitempty" json:"class,omitempty"`
	Priority           *CodeableConcept       `bson:"priority,omitempty" json:"priority,omitempty"`
	Type               []CodeableConcept      `bson:"type,omitempty" json:"type,omitempty"`
	ServiceType        []CodeableReference    `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Subject            *Reference             `bson:"subject,omitempty" json:"subject,omitempty"`
	SubjectStatus      *CodeableConcept       `bson:"subjectStatus,omitempty" json:"subjectStatus,omitempty"`
	EpisodeOfCare      []Reference            `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	BasedOn            []Reference            `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	CareTeam           []Reference            `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	PartOf             *Reference             `bson:"partOf,omitempty" json:"partOf,omitempty"`
	ServiceProvider    *Reference             `bson:"serviceProvider,omitempty" json:"serviceProvider,omitempty"`
	Participant        []EncounterParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	Appointment        []Reference            `bson:"appointment,omitempty" json:"appointment,omitempty"`
	VirtualService     []VirtualServiceDetail `bson:"virtualService,omitempty" json:"virtualService,omitempty"`
	ActualPeriod       *Period                `bson:"actualPeriod,omitempty" json:"actualPeriod,omitempty"`
	PlannedStartDate   *string                `bson:"plannedStartDate,omitempty" json:"plannedStartDate,omitempty"`
	PlannedEndDate     *string                `bson:"plannedEndDate,omitempty" json:"plannedEndDate,omitempty"`
	Length             *Duration              `bson:"length,omitempty" json:"length,omitempty"`
	Reason             []EncounterReason      `bson:"reason,omitempty" json:"reason,omitempty"`
	Diagnosis          []EncounterDiagnosis   `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	Account            []Reference            `bson:"account,omitempty" json:"account,omitempty"`
	DietPreference     []CodeableConcept      `bson:"dietPreference,omitempty" json:"dietPreference,omitempty"`
	SpecialArrangement []CodeableConcept      `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty"`
	SpecialCourtesy    []CodeableConcept      `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty"`
	Admission          *EncounterAdmission    `bson:"admission,omitempty" json:"admission,omitempty"`
	Location           []EncounterLocation    `bson:"location,omitempty" json:"location,omitempty"`
}
type EncounterParticipant struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Period            *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Actor             *Reference        `bson:"actor,omitempty" json:"actor,omitempty"`
}
type EncounterReason struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Use               []CodeableConcept   `bson:"use,omitempty" json:"use,omitempty"`
	Value             []CodeableReference `bson:"value,omitempty" json:"value,omitempty"`
}
type EncounterDiagnosis struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Condition         []CodeableReference `bson:"condition,omitempty" json:"condition,omitempty"`
	Use               []CodeableConcept   `bson:"use,omitempty" json:"use,omitempty"`
}
type EncounterAdmission struct {
	Id                     *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	PreAdmissionIdentifier *Identifier      `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty"`
	Origin                 *Reference       `bson:"origin,omitempty" json:"origin,omitempty"`
	AdmitSource            *CodeableConcept `bson:"admitSource,omitempty" json:"admitSource,omitempty"`
	ReAdmission            *CodeableConcept `bson:"reAdmission,omitempty" json:"reAdmission,omitempty"`
	Destination            *Reference       `bson:"destination,omitempty" json:"destination,omitempty"`
	DischargeDisposition   *CodeableConcept `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty"`
}
type EncounterLocation struct {
	Id                *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Location          Reference                `bson:"location" json:"location"`
	Status            *EncounterLocationStatus `bson:"status,omitempty" json:"status,omitempty"`
	Form              *CodeableConcept         `bson:"form,omitempty" json:"form,omitempty"`
	Period            *Period                  `bson:"period,omitempty" json:"period,omitempty"`
}
type OtherEncounter Encounter

// MarshalJSON marshals the given Encounter as JSON into a byte slice
func (r Encounter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEncounter
		ResourceType string `json:"resourceType"`
	}{
		OtherEncounter: OtherEncounter(r),
		ResourceType:   "Encounter",
	})
}

// UnmarshalEncounter unmarshals a Encounter.
func UnmarshalEncounter(b []byte) (Encounter, error) {
	var encounter Encounter
	if err := json.Unmarshal(b, &encounter); err != nil {
		return encounter, err
	}
	return encounter, nil
}
