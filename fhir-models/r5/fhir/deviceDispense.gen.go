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

// DeviceDispense is documented here http://hl7.org/fhir/StructureDefinition/DeviceDispense
type DeviceDispense struct {
	Id                    *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                   `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                `bson:"text,omitempty" json:"text,omitempty"`
	Contained             []json.RawMessage         `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension             []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn               []Reference               `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf                []Reference               `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Status                DeviceDispenseStatusCodes `bson:"status" json:"status"`
	StatusReason          *CodeableReference        `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	Category              []CodeableConcept         `bson:"category,omitempty" json:"category,omitempty"`
	Device                CodeableReference         `bson:"device" json:"device"`
	Subject               Reference                 `bson:"subject" json:"subject"`
	Receiver              *Reference                `bson:"receiver,omitempty" json:"receiver,omitempty"`
	Encounter             *Reference                `bson:"encounter,omitempty" json:"encounter,omitempty"`
	SupportingInformation []Reference               `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Performer             []DeviceDispensePerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	Location              *Reference                `bson:"location,omitempty" json:"location,omitempty"`
	Type                  *CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	Quantity              *Quantity                 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	PreparedDate          *string                   `bson:"preparedDate,omitempty" json:"preparedDate,omitempty"`
	WhenHandedOver        *string                   `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	Destination           *Reference                `bson:"destination,omitempty" json:"destination,omitempty"`
	Note                  []Annotation              `bson:"note,omitempty" json:"note,omitempty"`
	UsageInstruction      *string                   `bson:"usageInstruction,omitempty" json:"usageInstruction,omitempty"`
	EventHistory          []Reference               `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}

func (r DeviceDispense) ContainedResources() []json.RawMessage {
	return r.Contained
}

type DeviceDispensePerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor             Reference        `bson:"actor" json:"actor"`
}
type OtherDeviceDispense DeviceDispense

// MarshalJSON marshals the given DeviceDispense as JSON into a byte slice
func (r DeviceDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceDispense
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceDispense: OtherDeviceDispense(r),
		ResourceType:        "DeviceDispense",
	})
}
func (r DeviceDispense) ResourceType() string {
	return "DeviceDispense"
}

// UnmarshalDeviceDispense unmarshals a DeviceDispense.
func UnmarshalDeviceDispense(b []byte) (DeviceDispense, error) {
	var deviceDispense DeviceDispense
	if err := json.Unmarshal(b, &deviceDispense); err != nil {
		return deviceDispense, err
	}
	return deviceDispense, nil
}
