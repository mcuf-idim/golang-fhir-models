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

// Appointment is documented here http://hl7.org/fhir/StructureDefinition/Appointment
type Appointment struct {
	Id                     *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                 AppointmentStatus               `bson:"status" json:"status"`
	CancellationReason     *CodeableConcept                `bson:"cancellationReason,omitempty" json:"cancellationReason,omitempty"`
	Class                  []CodeableConcept               `bson:"class,omitempty" json:"class,omitempty"`
	ServiceCategory        []CodeableConcept               `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	ServiceType            []CodeableReference             `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	Specialty              []CodeableConcept               `bson:"specialty,omitempty" json:"specialty,omitempty"`
	AppointmentType        *CodeableConcept                `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	Reason                 []CodeableReference             `bson:"reason,omitempty" json:"reason,omitempty"`
	Priority               *CodeableConcept                `bson:"priority,omitempty" json:"priority,omitempty"`
	Description            *string                         `bson:"description,omitempty" json:"description,omitempty"`
	Replaces               []Reference                     `bson:"replaces,omitempty" json:"replaces,omitempty"`
	VirtualService         []VirtualServiceDetail          `bson:"virtualService,omitempty" json:"virtualService,omitempty"`
	SupportingInformation  []Reference                     `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	PreviousAppointment    *Reference                      `bson:"previousAppointment,omitempty" json:"previousAppointment,omitempty"`
	OriginatingAppointment *Reference                      `bson:"originatingAppointment,omitempty" json:"originatingAppointment,omitempty"`
	Start                  *string                         `bson:"start,omitempty" json:"start,omitempty"`
	End                    *string                         `bson:"end,omitempty" json:"end,omitempty"`
	MinutesDuration        *int                            `bson:"minutesDuration,omitempty" json:"minutesDuration,omitempty"`
	RequestedPeriod        []Period                        `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty"`
	Slot                   []Reference                     `bson:"slot,omitempty" json:"slot,omitempty"`
	Account                []Reference                     `bson:"account,omitempty" json:"account,omitempty"`
	Created                *string                         `bson:"created,omitempty" json:"created,omitempty"`
	CancellationDate       *string                         `bson:"cancellationDate,omitempty" json:"cancellationDate,omitempty"`
	Note                   []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
	PatientInstruction     []CodeableReference             `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	BasedOn                []Reference                     `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Subject                *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Participant            []AppointmentParticipant        `bson:"participant" json:"participant"`
	RecurrenceId           *int                            `bson:"recurrenceId,omitempty" json:"recurrenceId,omitempty"`
	OccurrenceChanged      *bool                           `bson:"occurrenceChanged,omitempty" json:"occurrenceChanged,omitempty"`
	RecurrenceTemplate     []AppointmentRecurrenceTemplate `bson:"recurrenceTemplate,omitempty" json:"recurrenceTemplate,omitempty"`
}
type AppointmentParticipant struct {
	Id                *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              []CodeableConcept   `bson:"type,omitempty" json:"type,omitempty"`
	Period            *Period             `bson:"period,omitempty" json:"period,omitempty"`
	Actor             *Reference          `bson:"actor,omitempty" json:"actor,omitempty"`
	Required          *bool               `bson:"required,omitempty" json:"required,omitempty"`
	Status            ParticipationStatus `bson:"status" json:"status"`
}
type AppointmentRecurrenceTemplate struct {
	Id                    *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Timezone              *CodeableConcept                              `bson:"timezone,omitempty" json:"timezone,omitempty"`
	RecurrenceType        CodeableConcept                               `bson:"recurrenceType" json:"recurrenceType"`
	LastOccurrenceDate    *string                                       `bson:"lastOccurrenceDate,omitempty" json:"lastOccurrenceDate,omitempty"`
	OccurrenceCount       *int                                          `bson:"occurrenceCount,omitempty" json:"occurrenceCount,omitempty"`
	OccurrenceDate        []string                                      `bson:"occurrenceDate,omitempty" json:"occurrenceDate,omitempty"`
	WeeklyTemplate        *AppointmentRecurrenceTemplateWeeklyTemplate  `bson:"weeklyTemplate,omitempty" json:"weeklyTemplate,omitempty"`
	MonthlyTemplate       *AppointmentRecurrenceTemplateMonthlyTemplate `bson:"monthlyTemplate,omitempty" json:"monthlyTemplate,omitempty"`
	YearlyTemplate        *AppointmentRecurrenceTemplateYearlyTemplate  `bson:"yearlyTemplate,omitempty" json:"yearlyTemplate,omitempty"`
	ExcludingDate         []string                                      `bson:"excludingDate,omitempty" json:"excludingDate,omitempty"`
	ExcludingRecurrenceId []int                                         `bson:"excludingRecurrenceId,omitempty" json:"excludingRecurrenceId,omitempty"`
}
type AppointmentRecurrenceTemplateWeeklyTemplate struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Monday            *bool       `bson:"monday,omitempty" json:"monday,omitempty"`
	Tuesday           *bool       `bson:"tuesday,omitempty" json:"tuesday,omitempty"`
	Wednesday         *bool       `bson:"wednesday,omitempty" json:"wednesday,omitempty"`
	Thursday          *bool       `bson:"thursday,omitempty" json:"thursday,omitempty"`
	Friday            *bool       `bson:"friday,omitempty" json:"friday,omitempty"`
	Saturday          *bool       `bson:"saturday,omitempty" json:"saturday,omitempty"`
	Sunday            *bool       `bson:"sunday,omitempty" json:"sunday,omitempty"`
	WeekInterval      *int        `bson:"weekInterval,omitempty" json:"weekInterval,omitempty"`
}
type AppointmentRecurrenceTemplateMonthlyTemplate struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DayOfMonth        *int        `bson:"dayOfMonth,omitempty" json:"dayOfMonth,omitempty"`
	NthWeekOfMonth    *Coding     `bson:"nthWeekOfMonth,omitempty" json:"nthWeekOfMonth,omitempty"`
	DayOfWeek         *Coding     `bson:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty"`
	MonthInterval     int         `bson:"monthInterval" json:"monthInterval"`
}
type AppointmentRecurrenceTemplateYearlyTemplate struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	YearInterval      int         `bson:"yearInterval" json:"yearInterval"`
}
type OtherAppointment Appointment

// MarshalJSON marshals the given Appointment as JSON into a byte slice
func (r Appointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointment
		ResourceType string `json:"resourceType"`
	}{
		OtherAppointment: OtherAppointment(r),
		ResourceType:     "Appointment",
	})
}
func (r Appointment) ResourceType() string {
	return "Appointment"
}

// UnmarshalAppointment unmarshals a Appointment.
func UnmarshalAppointment(b []byte) (Appointment, error) {
	var appointment Appointment
	if err := json.Unmarshal(b, &appointment); err != nil {
		return appointment, err
	}
	return appointment, nil
}
