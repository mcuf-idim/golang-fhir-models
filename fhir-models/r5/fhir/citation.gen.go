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

// Citation is documented here http://hl7.org/fhir/StructureDefinition/Citation
type Citation struct {
	Id                     *string                  `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                    `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                  `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                  `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative               `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                    *string                  `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             []Identifier             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                *string                  `bson:"version,omitempty" json:"version,omitempty"`
	VersionAlgorithmString *string                  `bson:"versionAlgorithmString,omitempty" json:"versionAlgorithmString,omitempty"`
	VersionAlgorithmCoding *Coding                  `bson:"versionAlgorithmCoding,omitempty" json:"versionAlgorithmCoding,omitempty"`
	Name                   *string                  `bson:"name,omitempty" json:"name,omitempty"`
	Title                  *string                  `bson:"title,omitempty" json:"title,omitempty"`
	Status                 PublicationStatus        `bson:"status" json:"status"`
	Experimental           *bool                    `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Date                   *string                  `bson:"date,omitempty" json:"date,omitempty"`
	Publisher              *string                  `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ContactDetail          `bson:"contact,omitempty" json:"contact,omitempty"`
	Description            *string                  `bson:"description,omitempty" json:"description,omitempty"`
	UseContext             []UsageContext           `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept        `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Purpose                *string                  `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Copyright              *string                  `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CopyrightLabel         *string                  `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	ApprovalDate           *string                  `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate         *string                  `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                  `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Author                 []ContactDetail          `bson:"author,omitempty" json:"author,omitempty"`
	Editor                 []ContactDetail          `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer               []ContactDetail          `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser               []ContactDetail          `bson:"endorser,omitempty" json:"endorser,omitempty"`
	Summary                []CitationSummary        `bson:"summary,omitempty" json:"summary,omitempty"`
	Classification         []CitationClassification `bson:"classification,omitempty" json:"classification,omitempty"`
	Note                   []Annotation             `bson:"note,omitempty" json:"note,omitempty"`
	CurrentState           []CodeableConcept        `bson:"currentState,omitempty" json:"currentState,omitempty"`
	StatusDate             []CitationStatusDate     `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	RelatedArtifact        []RelatedArtifact        `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	CitedArtifact          *CitationCitedArtifact   `bson:"citedArtifact,omitempty" json:"citedArtifact,omitempty"`
}
type CitationSummary struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Style             *CodeableConcept `bson:"style,omitempty" json:"style,omitempty"`
	Text              string           `bson:"text" json:"text"`
}
type CitationClassification struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Classifier        []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
}
type CitationStatusDate struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Activity          CodeableConcept `bson:"activity" json:"activity"`
	Actual            *bool           `bson:"actual,omitempty" json:"actual,omitempty"`
	Period            Period          `bson:"period" json:"period"`
}
type CitationCitedArtifact struct {
	Id                *string                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	RelatedIdentifier []Identifier                           `bson:"relatedIdentifier,omitempty" json:"relatedIdentifier,omitempty"`
	DateAccessed      *string                                `bson:"dateAccessed,omitempty" json:"dateAccessed,omitempty"`
	Version           *CitationCitedArtifactVersion          `bson:"version,omitempty" json:"version,omitempty"`
	CurrentState      []CodeableConcept                      `bson:"currentState,omitempty" json:"currentState,omitempty"`
	StatusDate        []CitationCitedArtifactStatusDate      `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	Title             []CitationCitedArtifactTitle           `bson:"title,omitempty" json:"title,omitempty"`
	Abstract          []CitationCitedArtifactAbstract        `bson:"abstract,omitempty" json:"abstract,omitempty"`
	Part              *CitationCitedArtifactPart             `bson:"part,omitempty" json:"part,omitempty"`
	RelatesTo         []CitationCitedArtifactRelatesTo       `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	PublicationForm   []CitationCitedArtifactPublicationForm `bson:"publicationForm,omitempty" json:"publicationForm,omitempty"`
	WebLocation       []CitationCitedArtifactWebLocation     `bson:"webLocation,omitempty" json:"webLocation,omitempty"`
	Classification    []CitationCitedArtifactClassification  `bson:"classification,omitempty" json:"classification,omitempty"`
	Contributorship   *CitationCitedArtifactContributorship  `bson:"contributorship,omitempty" json:"contributorship,omitempty"`
	Note              []Annotation                           `bson:"note,omitempty" json:"note,omitempty"`
}
type CitationCitedArtifactVersion struct {
	Id                *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Value             string      `bson:"value" json:"value"`
	BaseCitation      *Reference  `bson:"baseCitation,omitempty" json:"baseCitation,omitempty"`
}
type CitationCitedArtifactStatusDate struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Activity          CodeableConcept `bson:"activity" json:"activity"`
	Actual            *bool           `bson:"actual,omitempty" json:"actual,omitempty"`
	Period            Period          `bson:"period" json:"period"`
}
type CitationCitedArtifactTitle struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Language          *CodeableConcept  `bson:"language,omitempty" json:"language,omitempty"`
	Text              string            `bson:"text" json:"text"`
}
type CitationCitedArtifactAbstract struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Language          *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
	Text              string           `bson:"text" json:"text"`
	Copyright         *string          `bson:"copyright,omitempty" json:"copyright,omitempty"`
}
type CitationCitedArtifactPart struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Value             *string          `bson:"value,omitempty" json:"value,omitempty"`
	BaseCitation      *Reference       `bson:"baseCitation,omitempty" json:"baseCitation,omitempty"`
}
type CitationCitedArtifactRelatesTo struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              string            `bson:"type" json:"type"`
	Classifier        []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
	Label             *string           `bson:"label,omitempty" json:"label,omitempty"`
	Display           *string           `bson:"display,omitempty" json:"display,omitempty"`
	Citation          *string           `bson:"citation,omitempty" json:"citation,omitempty"`
	Document          *Attachment       `bson:"document,omitempty" json:"document,omitempty"`
	Resource          *string           `bson:"resource,omitempty" json:"resource,omitempty"`
	ResourceReference *Reference        `bson:"resourceReference,omitempty" json:"resourceReference,omitempty"`
}
type CitationCitedArtifactPublicationForm struct {
	Id                    *string                                          `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension                                      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                                      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	PublishedIn           *CitationCitedArtifactPublicationFormPublishedIn `bson:"publishedIn,omitempty" json:"publishedIn,omitempty"`
	CitedMedium           *CodeableConcept                                 `bson:"citedMedium,omitempty" json:"citedMedium,omitempty"`
	Volume                *string                                          `bson:"volume,omitempty" json:"volume,omitempty"`
	Issue                 *string                                          `bson:"issue,omitempty" json:"issue,omitempty"`
	ArticleDate           *string                                          `bson:"articleDate,omitempty" json:"articleDate,omitempty"`
	PublicationDateText   *string                                          `bson:"publicationDateText,omitempty" json:"publicationDateText,omitempty"`
	PublicationDateSeason *string                                          `bson:"publicationDateSeason,omitempty" json:"publicationDateSeason,omitempty"`
	LastRevisionDate      *string                                          `bson:"lastRevisionDate,omitempty" json:"lastRevisionDate,omitempty"`
	Language              []CodeableConcept                                `bson:"language,omitempty" json:"language,omitempty"`
	AccessionNumber       *string                                          `bson:"accessionNumber,omitempty" json:"accessionNumber,omitempty"`
	PageString            *string                                          `bson:"pageString,omitempty" json:"pageString,omitempty"`
	FirstPage             *string                                          `bson:"firstPage,omitempty" json:"firstPage,omitempty"`
	LastPage              *string                                          `bson:"lastPage,omitempty" json:"lastPage,omitempty"`
	PageCount             *string                                          `bson:"pageCount,omitempty" json:"pageCount,omitempty"`
	Copyright             *string                                          `bson:"copyright,omitempty" json:"copyright,omitempty"`
}
type CitationCitedArtifactPublicationFormPublishedIn struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Identifier        []Identifier     `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Title             *string          `bson:"title,omitempty" json:"title,omitempty"`
	Publisher         *Reference       `bson:"publisher,omitempty" json:"publisher,omitempty"`
	PublisherLocation *string          `bson:"publisherLocation,omitempty" json:"publisherLocation,omitempty"`
}
type CitationCitedArtifactWebLocation struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Classifier        []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
	Url               *string           `bson:"url,omitempty" json:"url,omitempty"`
}
type CitationCitedArtifactClassification struct {
	Id                 *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type               *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Classifier         []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
	ArtifactAssessment []Reference       `bson:"artifactAssessment,omitempty" json:"artifactAssessment,omitempty"`
}
type CitationCitedArtifactContributorship struct {
	Id                *string                                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Complete          *bool                                         `bson:"complete,omitempty" json:"complete,omitempty"`
	Entry             []CitationCitedArtifactContributorshipEntry   `bson:"entry,omitempty" json:"entry,omitempty"`
	Summary           []CitationCitedArtifactContributorshipSummary `bson:"summary,omitempty" json:"summary,omitempty"`
}
type CitationCitedArtifactContributorshipEntry struct {
	Id                   *string                                                         `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension                                                     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                                                     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Contributor          Reference                                                       `bson:"contributor" json:"contributor"`
	ForenameInitials     *string                                                         `bson:"forenameInitials,omitempty" json:"forenameInitials,omitempty"`
	Affiliation          []Reference                                                     `bson:"affiliation,omitempty" json:"affiliation,omitempty"`
	ContributionType     []CodeableConcept                                               `bson:"contributionType,omitempty" json:"contributionType,omitempty"`
	Role                 *CodeableConcept                                                `bson:"role,omitempty" json:"role,omitempty"`
	ContributionInstance []CitationCitedArtifactContributorshipEntryContributionInstance `bson:"contributionInstance,omitempty" json:"contributionInstance,omitempty"`
	CorrespondingContact *bool                                                           `bson:"correspondingContact,omitempty" json:"correspondingContact,omitempty"`
	RankingOrder         *int                                                            `bson:"rankingOrder,omitempty" json:"rankingOrder,omitempty"`
}
type CitationCitedArtifactContributorshipEntryContributionInstance struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              CodeableConcept `bson:"type" json:"type"`
	Time              *string         `bson:"time,omitempty" json:"time,omitempty"`
}
type CitationCitedArtifactContributorshipSummary struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Style             *CodeableConcept `bson:"style,omitempty" json:"style,omitempty"`
	Source            *CodeableConcept `bson:"source,omitempty" json:"source,omitempty"`
	Value             string           `bson:"value" json:"value"`
}
type OtherCitation Citation

// MarshalJSON marshals the given Citation as JSON into a byte slice
func (r Citation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCitation
		ResourceType string `json:"resourceType"`
	}{
		OtherCitation: OtherCitation(r),
		ResourceType:  "Citation",
	})
}
func (r Citation) ResourceType() string {
	return "Citation"
}

// UnmarshalCitation unmarshals a Citation.
func UnmarshalCitation(b []byte) (Citation, error) {
	var citation Citation
	if err := json.Unmarshal(b, &citation); err != nil {
		return citation, err
	}
	return citation, nil
}