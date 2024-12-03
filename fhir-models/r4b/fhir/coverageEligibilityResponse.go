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

// CoverageEligibilityResponse is documented here http://hl7.org/fhir/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponse struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                                  `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                                `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                                `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                             `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes           `bson:"status" json:"status"`
	Purpose           []EligibilityResponsePurpose           `bson:"purpose" json:"purpose"`
	Patient           Reference                              `bson:"patient" json:"patient"`
	Event             []CoverageEligibilityResponseEvent     `bson:"event,omitempty" json:"event,omitempty"`
	ServicedDate      *string                                `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod    *Period                                `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	Created           string                                 `bson:"created" json:"created"`
	Requestor         *Reference                             `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Request           Reference                              `bson:"request" json:"request"`
	Outcome           EligibilityOutcome                     `bson:"outcome" json:"outcome"`
	Disposition       *string                                `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Insurer           Reference                              `bson:"insurer" json:"insurer"`
	Insurance         []CoverageEligibilityResponseInsurance `bson:"insurance,omitempty" json:"insurance,omitempty"`
	PreAuthRef        *string                                `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	Form              *CodeableConcept                       `bson:"form,omitempty" json:"form,omitempty"`
	Error             []CoverageEligibilityResponseError     `bson:"error,omitempty" json:"error,omitempty"`
}
type CoverageEligibilityResponseEvent struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	WhenDateTime      string          `bson:"whenDateTime" json:"whenDateTime"`
	WhenPeriod        Period          `bson:"whenPeriod" json:"whenPeriod"`
}
type CoverageEligibilityResponseInsurance struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Coverage          Reference                                  `bson:"coverage" json:"coverage"`
	Inforce           *bool                                      `bson:"inforce,omitempty" json:"inforce,omitempty"`
	BenefitPeriod     *Period                                    `bson:"benefitPeriod,omitempty" json:"benefitPeriod,omitempty"`
	Item              []CoverageEligibilityResponseInsuranceItem `bson:"item,omitempty" json:"item,omitempty"`
}
type CoverageEligibilityResponseInsuranceItem struct {
	Id                      *string                                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []Extension                                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension                                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category                *CodeableConcept                                  `bson:"category,omitempty" json:"category,omitempty"`
	ProductOrService        *CodeableConcept                                  `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	Modifier                []CodeableConcept                                 `bson:"modifier,omitempty" json:"modifier,omitempty"`
	Provider                *Reference                                        `bson:"provider,omitempty" json:"provider,omitempty"`
	Excluded                *bool                                             `bson:"excluded,omitempty" json:"excluded,omitempty"`
	Name                    *string                                           `bson:"name,omitempty" json:"name,omitempty"`
	Description             *string                                           `bson:"description,omitempty" json:"description,omitempty"`
	Network                 *CodeableConcept                                  `bson:"network,omitempty" json:"network,omitempty"`
	Unit                    *CodeableConcept                                  `bson:"unit,omitempty" json:"unit,omitempty"`
	Term                    *CodeableConcept                                  `bson:"term,omitempty" json:"term,omitempty"`
	Benefit                 []CoverageEligibilityResponseInsuranceItemBenefit `bson:"benefit,omitempty" json:"benefit,omitempty"`
	AuthorizationRequired   *bool                                             `bson:"authorizationRequired,omitempty" json:"authorizationRequired,omitempty"`
	AuthorizationSupporting []CodeableConcept                                 `bson:"authorizationSupporting,omitempty" json:"authorizationSupporting,omitempty"`
	AuthorizationUrl        *string                                           `bson:"authorizationUrl,omitempty" json:"authorizationUrl,omitempty"`
}
type CoverageEligibilityResponseInsuranceItemBenefit struct {
	Id                 *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type               CodeableConcept `bson:"type" json:"type"`
	AllowedUnsignedInt *int            `bson:"allowedUnsignedInt,omitempty" json:"allowedUnsignedInt,omitempty"`
	AllowedString      *string         `bson:"allowedString,omitempty" json:"allowedString,omitempty"`
	AllowedMoney       *Money          `bson:"allowedMoney,omitempty" json:"allowedMoney,omitempty"`
	UsedUnsignedInt    *int            `bson:"usedUnsignedInt,omitempty" json:"usedUnsignedInt,omitempty"`
	UsedString         *string         `bson:"usedString,omitempty" json:"usedString,omitempty"`
	UsedMoney          *Money          `bson:"usedMoney,omitempty" json:"usedMoney,omitempty"`
}
type CoverageEligibilityResponseError struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Expression        []string        `bson:"expression,omitempty" json:"expression,omitempty"`
}
type OtherCoverageEligibilityResponse CoverageEligibilityResponse

// MarshalJSON marshals the given CoverageEligibilityResponse as JSON into a byte slice
func (r CoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverageEligibilityResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverageEligibilityResponse: OtherCoverageEligibilityResponse(r),
		ResourceType:                     "CoverageEligibilityResponse",
	})
}

// UnmarshalCoverageEligibilityResponse unmarshals a CoverageEligibilityResponse.
func UnmarshalCoverageEligibilityResponse(b []byte) (CoverageEligibilityResponse, error) {
	var coverageEligibilityResponse CoverageEligibilityResponse
	if err := json.Unmarshal(b, &coverageEligibilityResponse); err != nil {
		return coverageEligibilityResponse, err
	}
	return coverageEligibilityResponse, nil
}
