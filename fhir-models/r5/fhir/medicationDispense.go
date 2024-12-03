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

// MedicationDispense is documented here http://hl7.org/fhir/StructureDefinition/MedicationDispense
type MedicationDispense struct {
	Id                        *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Meta                      *Meta                           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules             *string                         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                  *string                         `bson:"language,omitempty" json:"language,omitempty"`
	Text                      *Narrative                      `bson:"text,omitempty" json:"text,omitempty"`
	Extension                 []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn                   []Reference                     `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf                    []Reference                     `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                    MedicationDispenseStatusCodes   `bson:"status" json:"status"`
	NotPerformedReason        *CodeableReference              `bson:"notPerformedReason,omitempty" json:"notPerformedReason,omitempty"`
	StatusChanged             *string                         `bson:"statusChanged,omitempty" json:"statusChanged,omitempty"`
	Category                  []CodeableConcept               `bson:"category,omitempty" json:"category,omitempty"`
	Medication                CodeableReference               `bson:"medication" json:"medication"`
	Subject                   Reference                       `bson:"subject" json:"subject"`
	Encounter                 *Reference                      `bson:"encounter,omitempty" json:"encounter,omitempty"`
	SupportingInformation     []Reference                     `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Performer                 []MedicationDispensePerformer   `bson:"performer,omitempty" json:"performer,omitempty"`
	Location                  *Reference                      `bson:"location,omitempty" json:"location,omitempty"`
	AuthorizingPrescription   []Reference                     `bson:"authorizingPrescription,omitempty" json:"authorizingPrescription,omitempty"`
	Type                      *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	Quantity                  *Quantity                       `bson:"quantity,omitempty" json:"quantity,omitempty"`
	DaysSupply                *Quantity                       `bson:"daysSupply,omitempty" json:"daysSupply,omitempty"`
	Recorded                  *string                         `bson:"recorded,omitempty" json:"recorded,omitempty"`
	WhenPrepared              *string                         `bson:"whenPrepared,omitempty" json:"whenPrepared,omitempty"`
	WhenHandedOver            *string                         `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	Destination               *Reference                      `bson:"destination,omitempty" json:"destination,omitempty"`
	Receiver                  []Reference                     `bson:"receiver,omitempty" json:"receiver,omitempty"`
	Note                      []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
	RenderedDosageInstruction *string                         `bson:"renderedDosageInstruction,omitempty" json:"renderedDosageInstruction,omitempty"`
	DosageInstruction         []Dosage                        `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	Substitution              *MedicationDispenseSubstitution `bson:"substitution,omitempty" json:"substitution,omitempty"`
	EventHistory              []Reference                     `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}
type MedicationDispensePerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}
type MedicationDispenseSubstitution struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	WasSubstituted    bool              `bson:"wasSubstituted" json:"wasSubstituted"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Reason            []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	ResponsibleParty  *Reference        `bson:"responsibleParty,omitempty" json:"responsibleParty,omitempty"`
}
type OtherMedicationDispense MedicationDispense

// MarshalJSON marshals the given MedicationDispense as JSON into a byte slice
func (r MedicationDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationDispense
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicationDispense: OtherMedicationDispense(r),
		ResourceType:            "MedicationDispense",
	})
}

// UnmarshalMedicationDispense unmarshals a MedicationDispense.
func UnmarshalMedicationDispense(b []byte) (MedicationDispense, error) {
	var medicationDispense MedicationDispense
	if err := json.Unmarshal(b, &medicationDispense); err != nil {
		return medicationDispense, err
	}
	return medicationDispense, nil
}