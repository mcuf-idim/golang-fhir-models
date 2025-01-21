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

// ImagingSelection is documented here http://hl7.org/fhir/StructureDefinition/ImagingSelection
type ImagingSelection struct {
	Id                  *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Contained           []json.RawMessage           `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension           []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier          []Identifier                `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status              ImagingSelectionStatus      `bson:"status" json:"status"`
	Subject             *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Issued              *string                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Performer           []ImagingSelectionPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	BasedOn             []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	Category            []CodeableConcept           `bson:"category,omitempty" json:"category,omitempty"`
	Code                CodeableConcept             `bson:"code" json:"code"`
	StudyUid            *string                     `bson:"studyUid,omitempty" json:"studyUid,omitempty"`
	DerivedFrom         []Reference                 `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	Endpoint            []Reference                 `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	SeriesUid           *string                     `bson:"seriesUid,omitempty" json:"seriesUid,omitempty"`
	SeriesNumber        *int                        `bson:"seriesNumber,omitempty" json:"seriesNumber,omitempty"`
	FrameOfReferenceUid *string                     `bson:"frameOfReferenceUid,omitempty" json:"frameOfReferenceUid,omitempty"`
	BodySite            *CodeableReference          `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Focus               []Reference                 `bson:"focus,omitempty" json:"focus,omitempty"`
	Instance            []ImagingSelectionInstance  `bson:"instance,omitempty" json:"instance,omitempty"`
}

func (r ImagingSelection) ContainedResources() []json.RawMessage {
	return r.Contained
}

type ImagingSelectionPerformer struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	Actor             *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
}
type ImagingSelectionInstance struct {
	Id                *string                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Uid               string                                  `bson:"uid" json:"uid"`
	Number            *int                                    `bson:"number,omitempty" json:"number,omitempty"`
	SopClass          *Coding                                 `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	Subset            []string                                `bson:"subset,omitempty" json:"subset,omitempty"`
	ImageRegion2D     []ImagingSelectionInstanceImageRegion2D `bson:"imageRegion2D,omitempty" json:"imageRegion2D,omitempty"`
	ImageRegion3D     []ImagingSelectionInstanceImageRegion3D `bson:"imageRegion3D,omitempty" json:"imageRegion3D,omitempty"`
}
type ImagingSelectionInstanceImageRegion2D struct {
	Id                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RegionType        ImagingSelection2DGraphicType `bson:"regionType" json:"regionType"`
	Coordinate        []json.Number                 `bson:"coordinate" json:"coordinate"`
}
type ImagingSelectionInstanceImageRegion3D struct {
	Id                *string                       `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                   `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                   `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	RegionType        ImagingSelection3DGraphicType `bson:"regionType" json:"regionType"`
	Coordinate        []json.Number                 `bson:"coordinate" json:"coordinate"`
}
type OtherImagingSelection ImagingSelection

// MarshalJSON marshals the given ImagingSelection as JSON into a byte slice
func (r ImagingSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImagingSelection
		ResourceType string `json:"resourceType"`
	}{
		OtherImagingSelection: OtherImagingSelection(r),
		ResourceType:          "ImagingSelection",
	})
}
func (r ImagingSelection) ResourceType() string {
	return "ImagingSelection"
}
func (r ImagingSelection) ResourceIdentifier() string {
	if r.Id != nil {
		return *r.Id
	}
	return ""
}

// UnmarshalImagingSelection unmarshals a ImagingSelection.
func UnmarshalImagingSelection(b []byte) (ImagingSelection, error) {
	var imagingSelection ImagingSelection
	if err := json.Unmarshal(b, &imagingSelection); err != nil {
		return imagingSelection, err
	}
	return imagingSelection, nil
}
