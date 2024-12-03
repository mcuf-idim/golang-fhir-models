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

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/mcuf-idim/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ArtifactAssessmentWorkflowStatus is documented here http://hl7.org/fhir/ValueSet/artifactassessment-workflow-status
type ArtifactAssessmentWorkflowStatus int

const (
	ArtifactAssessmentWorkflowStatusSubmitted ArtifactAssessmentWorkflowStatus = iota
	ArtifactAssessmentWorkflowStatusTriaged
	ArtifactAssessmentWorkflowStatusWaitingForInput
	ArtifactAssessmentWorkflowStatusResolvedNoChange
	ArtifactAssessmentWorkflowStatusResolvedChangeRequired
	ArtifactAssessmentWorkflowStatusDeferred
	ArtifactAssessmentWorkflowStatusDuplicate
	ArtifactAssessmentWorkflowStatusApplied
	ArtifactAssessmentWorkflowStatusPublished
	ArtifactAssessmentWorkflowStatusEnteredInError
)

func (code ArtifactAssessmentWorkflowStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ArtifactAssessmentWorkflowStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "submitted":
		*code = ArtifactAssessmentWorkflowStatusSubmitted
	case "triaged":
		*code = ArtifactAssessmentWorkflowStatusTriaged
	case "waiting-for-input":
		*code = ArtifactAssessmentWorkflowStatusWaitingForInput
	case "resolved-no-change":
		*code = ArtifactAssessmentWorkflowStatusResolvedNoChange
	case "resolved-change-required":
		*code = ArtifactAssessmentWorkflowStatusResolvedChangeRequired
	case "deferred":
		*code = ArtifactAssessmentWorkflowStatusDeferred
	case "duplicate":
		*code = ArtifactAssessmentWorkflowStatusDuplicate
	case "applied":
		*code = ArtifactAssessmentWorkflowStatusApplied
	case "published":
		*code = ArtifactAssessmentWorkflowStatusPublished
	case "entered-in-error":
		*code = ArtifactAssessmentWorkflowStatusEnteredInError
	default:
		return fmt.Errorf("unknown ArtifactAssessmentWorkflowStatus code `%s`", s)
	}
	return nil
}
func (code ArtifactAssessmentWorkflowStatus) String() string {
	return code.Code()
}
func (code ArtifactAssessmentWorkflowStatus) Code() string {
	switch code {
	case ArtifactAssessmentWorkflowStatusSubmitted:
		return "submitted"
	case ArtifactAssessmentWorkflowStatusTriaged:
		return "triaged"
	case ArtifactAssessmentWorkflowStatusWaitingForInput:
		return "waiting-for-input"
	case ArtifactAssessmentWorkflowStatusResolvedNoChange:
		return "resolved-no-change"
	case ArtifactAssessmentWorkflowStatusResolvedChangeRequired:
		return "resolved-change-required"
	case ArtifactAssessmentWorkflowStatusDeferred:
		return "deferred"
	case ArtifactAssessmentWorkflowStatusDuplicate:
		return "duplicate"
	case ArtifactAssessmentWorkflowStatusApplied:
		return "applied"
	case ArtifactAssessmentWorkflowStatusPublished:
		return "published"
	case ArtifactAssessmentWorkflowStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code ArtifactAssessmentWorkflowStatus) Display() string {
	switch code {
	case ArtifactAssessmentWorkflowStatusSubmitted:
		return "Submitted"
	case ArtifactAssessmentWorkflowStatusTriaged:
		return "Triaged"
	case ArtifactAssessmentWorkflowStatusWaitingForInput:
		return "Waiting for Input"
	case ArtifactAssessmentWorkflowStatusResolvedNoChange:
		return "Resolved - No Change"
	case ArtifactAssessmentWorkflowStatusResolvedChangeRequired:
		return "Resolved - Change Required"
	case ArtifactAssessmentWorkflowStatusDeferred:
		return "Deferred"
	case ArtifactAssessmentWorkflowStatusDuplicate:
		return "Duplicate"
	case ArtifactAssessmentWorkflowStatusApplied:
		return "Applied"
	case ArtifactAssessmentWorkflowStatusPublished:
		return "Published"
	case ArtifactAssessmentWorkflowStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code ArtifactAssessmentWorkflowStatus) Definition() string {
	switch code {
	case ArtifactAssessmentWorkflowStatusSubmitted:
		return "The comment has been submitted, but the responsible party has not yet been determined, or the responsible party has not yet determined the next steps to be taken."
	case ArtifactAssessmentWorkflowStatusTriaged:
		return "The comment has been triaged, meaning the responsible party has been determined and next steps have been identified to address the comment."
	case ArtifactAssessmentWorkflowStatusWaitingForInput:
		return "The comment is waiting for input from a specific party before next steps can be taken."
	case ArtifactAssessmentWorkflowStatusResolvedNoChange:
		return "The comment has been resolved and no changes resulted from the resolution"
	case ArtifactAssessmentWorkflowStatusResolvedChangeRequired:
		return "The comment has been resolved and changes are required to address the comment"
	case ArtifactAssessmentWorkflowStatusDeferred:
		return "The comment is acceptable, but resolution of the comment and application of any associated changes have been deferred"
	case ArtifactAssessmentWorkflowStatusDuplicate:
		return "The comment is a duplicate of another comment already received"
	case ArtifactAssessmentWorkflowStatusApplied:
		return "The comment is resolved and any necessary changes have been applied"
	case ArtifactAssessmentWorkflowStatusPublished:
		return "The necessary changes to the artifact have been published in a new version of the artifact"
	case ArtifactAssessmentWorkflowStatusEnteredInError:
		return "The assessment was entered in error"
	}
	return "<unknown>"
}
