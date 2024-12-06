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

// Contract is documented here http://hl7.org/fhir/StructureDefinition/Contract
type Contract struct {
	Id                       *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta                     *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules            *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                 *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text                     *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension                []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier               []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Url                      *string                      `bson:"url,omitempty" json:"url,omitempty"`
	Version                  *string                      `bson:"version,omitempty" json:"version,omitempty"`
	Status                   *ContractResourceStatusCodes `bson:"status,omitempty" json:"status,omitempty"`
	LegalState               *CodeableConcept             `bson:"legalState,omitempty" json:"legalState,omitempty"`
	InstantiatesCanonical    *Reference                   `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri          *string                      `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	ContentDerivative        *CodeableConcept             `bson:"contentDerivative,omitempty" json:"contentDerivative,omitempty"`
	Issued                   *string                      `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies                  *Period                      `bson:"applies,omitempty" json:"applies,omitempty"`
	ExpirationType           *CodeableConcept             `bson:"expirationType,omitempty" json:"expirationType,omitempty"`
	Subject                  []Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Authority                []Reference                  `bson:"authority,omitempty" json:"authority,omitempty"`
	Domain                   []Reference                  `bson:"domain,omitempty" json:"domain,omitempty"`
	Site                     []Reference                  `bson:"site,omitempty" json:"site,omitempty"`
	Name                     *string                      `bson:"name,omitempty" json:"name,omitempty"`
	Title                    *string                      `bson:"title,omitempty" json:"title,omitempty"`
	Subtitle                 *string                      `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	Alias                    []string                     `bson:"alias,omitempty" json:"alias,omitempty"`
	Author                   *Reference                   `bson:"author,omitempty" json:"author,omitempty"`
	Scope                    *CodeableConcept             `bson:"scope,omitempty" json:"scope,omitempty"`
	TopicCodeableConcept     *CodeableConcept             `bson:"topicCodeableConcept,omitempty" json:"topicCodeableConcept,omitempty"`
	TopicReference           *Reference                   `bson:"topicReference,omitempty" json:"topicReference,omitempty"`
	Type                     *CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	SubType                  []CodeableConcept            `bson:"subType,omitempty" json:"subType,omitempty"`
	ContentDefinition        *ContractContentDefinition   `bson:"contentDefinition,omitempty" json:"contentDefinition,omitempty"`
	Term                     []ContractTerm               `bson:"term,omitempty" json:"term,omitempty"`
	SupportingInfo           []Reference                  `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	RelevantHistory          []Reference                  `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
	Signer                   []ContractSigner             `bson:"signer,omitempty" json:"signer,omitempty"`
	Friendly                 []ContractFriendly           `bson:"friendly,omitempty" json:"friendly,omitempty"`
	Legal                    []ContractLegal              `bson:"legal,omitempty" json:"legal,omitempty"`
	Rule                     []ContractRule               `bson:"rule,omitempty" json:"rule,omitempty"`
	LegallyBindingAttachment *Attachment                  `bson:"legallyBindingAttachment,omitempty" json:"legallyBindingAttachment,omitempty"`
	LegallyBindingReference  *Reference                   `bson:"legallyBindingReference,omitempty" json:"legallyBindingReference,omitempty"`
}
type ContractContentDefinition struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept                        `bson:"type" json:"type"`
	SubType           *CodeableConcept                       `bson:"subType,omitempty" json:"subType,omitempty"`
	Publisher         *Reference                             `bson:"publisher,omitempty" json:"publisher,omitempty"`
	PublicationDate   *string                                `bson:"publicationDate,omitempty" json:"publicationDate,omitempty"`
	PublicationStatus ContractResourcePublicationStatusCodes `bson:"publicationStatus" json:"publicationStatus"`
	Copyright         *string                                `bson:"copyright,omitempty" json:"copyright,omitempty"`
}
type ContractTerm struct {
	Id                   *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Issued               *string                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies              *Period                     `bson:"applies,omitempty" json:"applies,omitempty"`
	TopicCodeableConcept *CodeableConcept            `bson:"topicCodeableConcept,omitempty" json:"topicCodeableConcept,omitempty"`
	TopicReference       *Reference                  `bson:"topicReference,omitempty" json:"topicReference,omitempty"`
	Type                 *CodeableConcept            `bson:"type,omitempty" json:"type,omitempty"`
	SubType              *CodeableConcept            `bson:"subType,omitempty" json:"subType,omitempty"`
	Text                 *string                     `bson:"text,omitempty" json:"text,omitempty"`
	SecurityLabel        []ContractTermSecurityLabel `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Offer                ContractTermOffer           `bson:"offer" json:"offer"`
	Asset                []ContractTermAsset         `bson:"asset,omitempty" json:"asset,omitempty"`
	Action               []ContractTermAction        `bson:"action,omitempty" json:"action,omitempty"`
	Group                []ContractTerm              `bson:"group,omitempty" json:"group,omitempty"`
}
type ContractTermSecurityLabel struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Number            []int       `bson:"number,omitempty" json:"number,omitempty"`
	Classification    Coding      `bson:"classification" json:"classification"`
	Category          []Coding    `bson:"category,omitempty" json:"category,omitempty"`
	Control           []Coding    `bson:"control,omitempty" json:"control,omitempty"`
}
type ContractTermOffer struct {
	Id                  *string                   `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Party               []ContractTermOfferParty  `bson:"party,omitempty" json:"party,omitempty"`
	Topic               *Reference                `bson:"topic,omitempty" json:"topic,omitempty"`
	Type                *CodeableConcept          `bson:"type,omitempty" json:"type,omitempty"`
	Decision            *CodeableConcept          `bson:"decision,omitempty" json:"decision,omitempty"`
	DecisionMode        []CodeableConcept         `bson:"decisionMode,omitempty" json:"decisionMode,omitempty"`
	Answer              []ContractTermOfferAnswer `bson:"answer,omitempty" json:"answer,omitempty"`
	Text                *string                   `bson:"text,omitempty" json:"text,omitempty"`
	LinkId              []string                  `bson:"linkId,omitempty" json:"linkId,omitempty"`
	SecurityLabelNumber []int                     `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
}
type ContractTermOfferParty struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         []Reference     `bson:"reference" json:"reference"`
	Role              CodeableConcept `bson:"role" json:"role"`
}
type ContractTermOfferAnswer struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ValueBoolean      bool        `bson:"valueBoolean" json:"valueBoolean"`
	ValueDecimal      json.Number `bson:"valueDecimal" json:"valueDecimal"`
	ValueInteger      int         `bson:"valueInteger" json:"valueInteger"`
	ValueDate         string      `bson:"valueDate" json:"valueDate"`
	ValueDateTime     string      `bson:"valueDateTime" json:"valueDateTime"`
	ValueTime         string      `bson:"valueTime" json:"valueTime"`
	ValueString       string      `bson:"valueString" json:"valueString"`
	ValueUri          string      `bson:"valueUri" json:"valueUri"`
	ValueAttachment   Attachment  `bson:"valueAttachment" json:"valueAttachment"`
	ValueCoding       Coding      `bson:"valueCoding" json:"valueCoding"`
	ValueQuantity     Quantity    `bson:"valueQuantity" json:"valueQuantity"`
	ValueReference    Reference   `bson:"valueReference" json:"valueReference"`
}
type ContractTermAsset struct {
	Id                  *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Scope               *CodeableConcept              `bson:"scope,omitempty" json:"scope,omitempty"`
	Type                []CodeableConcept             `bson:"type,omitempty" json:"type,omitempty"`
	TypeReference       []Reference                   `bson:"typeReference,omitempty" json:"typeReference,omitempty"`
	Subtype             []CodeableConcept             `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Relationship        *Coding                       `bson:"relationship,omitempty" json:"relationship,omitempty"`
	Context             []ContractTermAssetContext    `bson:"context,omitempty" json:"context,omitempty"`
	Condition           *string                       `bson:"condition,omitempty" json:"condition,omitempty"`
	PeriodType          []CodeableConcept             `bson:"periodType,omitempty" json:"periodType,omitempty"`
	Period              []Period                      `bson:"period,omitempty" json:"period,omitempty"`
	UsePeriod           []Period                      `bson:"usePeriod,omitempty" json:"usePeriod,omitempty"`
	Text                *string                       `bson:"text,omitempty" json:"text,omitempty"`
	LinkId              []string                      `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Answer              []ContractTermOfferAnswer     `bson:"answer,omitempty" json:"answer,omitempty"`
	SecurityLabelNumber []int                         `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
	ValuedItem          []ContractTermAssetValuedItem `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
}
type ContractTermAssetContext struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         *Reference        `bson:"reference,omitempty" json:"reference,omitempty"`
	Code              []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Text              *string           `bson:"text,omitempty" json:"text,omitempty"`
}
type ContractTermAssetValuedItem struct {
	Id                    *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime         *string          `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice             *Money           `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                *json.Number     `bson:"factor,omitempty" json:"factor,omitempty"`
	Points                *json.Number     `bson:"points,omitempty" json:"points,omitempty"`
	Net                   *Money           `bson:"net,omitempty" json:"net,omitempty"`
	Payment               *string          `bson:"payment,omitempty" json:"payment,omitempty"`
	PaymentDate           *string          `bson:"paymentDate,omitempty" json:"paymentDate,omitempty"`
	Responsible           *Reference       `bson:"responsible,omitempty" json:"responsible,omitempty"`
	Recipient             *Reference       `bson:"recipient,omitempty" json:"recipient,omitempty"`
	LinkId                []string         `bson:"linkId,omitempty" json:"linkId,omitempty"`
	SecurityLabelNumber   []int            `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
}
type ContractTermAction struct {
	Id                  *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	DoNotPerform        *bool                       `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	Type                CodeableConcept             `bson:"type" json:"type"`
	Subject             []ContractTermActionSubject `bson:"subject,omitempty" json:"subject,omitempty"`
	Intent              CodeableConcept             `bson:"intent" json:"intent"`
	LinkId              []string                    `bson:"linkId,omitempty" json:"linkId,omitempty"`
	Status              CodeableConcept             `bson:"status" json:"status"`
	Context             *Reference                  `bson:"context,omitempty" json:"context,omitempty"`
	ContextLinkId       []string                    `bson:"contextLinkId,omitempty" json:"contextLinkId,omitempty"`
	OccurrenceDateTime  *string                     `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *Period                     `bson:"occurrencePeriod,omitempty" json:"occurrencePeriod,omitempty"`
	OccurrenceTiming    *Timing                     `bson:"occurrenceTiming,omitempty" json:"occurrenceTiming,omitempty"`
	Requester           []Reference                 `bson:"requester,omitempty" json:"requester,omitempty"`
	RequesterLinkId     []string                    `bson:"requesterLinkId,omitempty" json:"requesterLinkId,omitempty"`
	PerformerType       []CodeableConcept           `bson:"performerType,omitempty" json:"performerType,omitempty"`
	PerformerRole       *CodeableConcept            `bson:"performerRole,omitempty" json:"performerRole,omitempty"`
	Performer           *Reference                  `bson:"performer,omitempty" json:"performer,omitempty"`
	PerformerLinkId     []string                    `bson:"performerLinkId,omitempty" json:"performerLinkId,omitempty"`
	Reason              []CodeableReference         `bson:"reason,omitempty" json:"reason,omitempty"`
	ReasonLinkId        []string                    `bson:"reasonLinkId,omitempty" json:"reasonLinkId,omitempty"`
	Note                []Annotation                `bson:"note,omitempty" json:"note,omitempty"`
	SecurityLabelNumber []int                       `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
}
type ContractTermActionSubject struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Reference         []Reference      `bson:"reference" json:"reference"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}
type ContractSigner struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              Coding      `bson:"type" json:"type"`
	Party             Reference   `bson:"party" json:"party"`
	Signature         []Signature `bson:"signature" json:"signature"`
}
type ContractFriendly struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `bson:"contentAttachment" json:"contentAttachment"`
	ContentReference  Reference   `bson:"contentReference" json:"contentReference"`
}
type ContractLegal struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `bson:"contentAttachment" json:"contentAttachment"`
	ContentReference  Reference   `bson:"contentReference" json:"contentReference"`
}
type ContractRule struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	ContentAttachment Attachment  `bson:"contentAttachment" json:"contentAttachment"`
	ContentReference  Reference   `bson:"contentReference" json:"contentReference"`
}
type OtherContract Contract

// MarshalJSON marshals the given Contract as JSON into a byte slice
func (r Contract) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherContract
		ResourceType string `json:"resourceType"`
	}{
		OtherContract: OtherContract(r),
		ResourceType:  "Contract",
	})
}
func (r Contract) ResourceType() string {
	return "Contract"
}

// UnmarshalContract unmarshals a Contract.
func UnmarshalContract(b []byte) (Contract, error) {
	var contract Contract
	if err := json.Unmarshal(b, &contract); err != nil {
		return contract, err
	}
	return contract, nil
}