// Copyright 2019 The Samply Development Community
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

// InsurancePlan is documented here http://hl7.org/fhir/StructureDefinition/InsurancePlan
type InsurancePlan struct {
	Id                *string                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                 `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative              `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            *PublicationStatus      `bson:"status,omitempty" json:"status,omitempty"`
	Type              []CodeableConcept       `bson:"type,omitempty" json:"type,omitempty"`
	Name              *string                 `bson:"name,omitempty" json:"name,omitempty"`
	Alias             []string                `bson:"alias,omitempty" json:"alias,omitempty"`
	Period            *Period                 `bson:"period,omitempty" json:"period,omitempty"`
	OwnedBy           *Reference              `bson:"ownedBy,omitempty" json:"ownedBy,omitempty"`
	AdministeredBy    *Reference              `bson:"administeredBy,omitempty" json:"administeredBy,omitempty"`
	CoverageArea      []Reference             `bson:"coverageArea,omitempty" json:"coverageArea,omitempty"`
	Contact           []InsurancePlanContact  `bson:"contact,omitempty" json:"contact,omitempty"`
	Endpoint          []Reference             `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	Network           []Reference             `bson:"network,omitempty" json:"network,omitempty"`
	Coverage          []InsurancePlanCoverage `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Plan              []InsurancePlanPlan     `bson:"plan,omitempty" json:"plan,omitempty"`
}
type InsurancePlanContact struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Purpose           *CodeableConcept `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Name              *HumanName       `bson:"name,omitempty" json:"name,omitempty"`
	Telecom           []ContactPoint   `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Address           *Address         `bson:"address,omitempty" json:"address,omitempty"`
}
type InsurancePlanCoverage struct {
	Id                *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept                `bson:"type" json:"type"`
	Network           []Reference                    `bson:"network,omitempty" json:"network,omitempty"`
	Benefit           []InsurancePlanCoverageBenefit `bson:"benefit" json:"benefit"`
}
type InsurancePlanCoverageBenefit struct {
	Id                *string                             `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                         `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                         `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept                     `bson:"type" json:"type"`
	Requirement       *string                             `bson:"requirement,omitempty" json:"requirement,omitempty"`
	Limit             []InsurancePlanCoverageBenefitLimit `bson:"limit,omitempty" json:"limit,omitempty"`
}
type InsurancePlanCoverageBenefitLimit struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Value             *Quantity        `bson:"value,omitempty" json:"value,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}
type InsurancePlanPlan struct {
	Id                *string                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type              *CodeableConcept                `bson:"type,omitempty" json:"type,omitempty"`
	CoverageArea      []Reference                     `bson:"coverageArea,omitempty" json:"coverageArea,omitempty"`
	Network           []Reference                     `bson:"network,omitempty" json:"network,omitempty"`
	GeneralCost       []InsurancePlanPlanGeneralCost  `bson:"generalCost,omitempty" json:"generalCost,omitempty"`
	SpecificCost      []InsurancePlanPlanSpecificCost `bson:"specificCost,omitempty" json:"specificCost,omitempty"`
}
type InsurancePlanPlanGeneralCost struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	GroupSize         *int             `bson:"groupSize,omitempty" json:"groupSize,omitempty"`
	Cost              *Money           `bson:"cost,omitempty" json:"cost,omitempty"`
	Comment           *string          `bson:"comment,omitempty" json:"comment,omitempty"`
}
type InsurancePlanPlanSpecificCost struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          CodeableConcept                        `bson:"category" json:"category"`
	Benefit           []InsurancePlanPlanSpecificCostBenefit `bson:"benefit,omitempty" json:"benefit,omitempty"`
}
type InsurancePlanPlanSpecificCostBenefit struct {
	Id                *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept                            `bson:"type" json:"type"`
	Cost              []InsurancePlanPlanSpecificCostBenefitCost `bson:"cost,omitempty" json:"cost,omitempty"`
}
type InsurancePlanPlanSpecificCostBenefitCost struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept   `bson:"type" json:"type"`
	Applicability     *CodeableConcept  `bson:"applicability,omitempty" json:"applicability,omitempty"`
	Qualifiers        []CodeableConcept `bson:"qualifiers,omitempty" json:"qualifiers,omitempty"`
	Value             *Quantity         `bson:"value,omitempty" json:"value,omitempty"`
}
type OtherInsurancePlan InsurancePlan

// MarshalJSON marshals the given InsurancePlan as JSON into a byte slice
func (r InsurancePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInsurancePlan
		ResourceType string `json:"resourceType"`
	}{
		OtherInsurancePlan: OtherInsurancePlan(r),
		ResourceType:       "InsurancePlan",
	})
}

// UnmarshalInsurancePlan unmarshals a InsurancePlan.
func UnmarshalInsurancePlan(b []byte) (InsurancePlan, error) {
	var insurancePlan InsurancePlan
	if err := json.Unmarshal(b, &insurancePlan); err != nil {
		return insurancePlan, err
	}
	return insurancePlan, nil
}
