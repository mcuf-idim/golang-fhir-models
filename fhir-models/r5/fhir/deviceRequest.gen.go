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

// DeviceRequest is documented here http://hl7.org/fhir/StructureDefinition/DeviceRequest
type DeviceRequest struct {
	Id                    *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	Contained             []json.RawMessage        `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension             []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string                 `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                 `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	BasedOn               []Reference              `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Replaces              []Reference              `bson:"replaces,omitempty" json:"replaces,omitempty"`
	GroupIdentifier       *Identifier              `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	Status                *RequestStatus           `bson:"status,omitempty" json:"status,omitempty"`
	Intent                RequestIntent            `bson:"intent" json:"intent"`
	Priority              *RequestPriority         `bson:"priority,omitempty" json:"priority,omitempty"`
	DoNotPerform          *bool                    `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	Code                  CodeableReference        `bson:"code" json:"code"`
	Quantity              *int                     `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Parameter             []DeviceRequestParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	Subject               Reference                `bson:"subject" json:"subject"`
	Encounter             *Reference               `bson:"encounter,omitempty" json:"encounter,omitempty"`
	OccurrenceDateTime    *string                  `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod      *Period                  `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming      *Timing                  `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	AuthoredOn            *string                  `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	Requester             *Reference               `bson:"requester,omitempty" json:"requester,omitempty"`
	Performer             *CodeableReference       `bson:"performer,omitempty" json:"performer,omitempty"`
	Reason                []CodeableReference      `bson:"reason,omitempty" json:"reason,omitempty"`
	AsNeeded              *bool                    `bson:"asNeeded,omitempty" json:"asNeeded,omitempty"`
	AsNeededFor           *CodeableConcept         `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty"`
	Insurance             []Reference              `bson:"insurance,omitempty" json:"insurance,omitempty"`
	SupportingInfo        []Reference              `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	Note                  []Annotation             `bson:"note,omitempty" json:"note,omitempty"`
	RelevantHistory       []Reference              `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}

func (r DeviceRequest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type DeviceRequestParameter struct {
	Id                   *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code                 *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
}
type OtherDeviceRequest DeviceRequest

// MarshalJSON marshals the given DeviceRequest as JSON into a byte slice
func (r DeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceRequest: OtherDeviceRequest(r),
		ResourceType:       "DeviceRequest",
	})
}
func (r DeviceRequest) ResourceType() string {
	return "DeviceRequest"
}

// UnmarshalDeviceRequest unmarshals a DeviceRequest.
func UnmarshalDeviceRequest(b []byte) (DeviceRequest, error) {
	var deviceRequest DeviceRequest
	if err := json.Unmarshal(b, &deviceRequest); err != nil {
		return deviceRequest, err
	}
	return deviceRequest, nil
}
