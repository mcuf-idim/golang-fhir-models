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

// MeasureReportStatus is documented here http://hl7.org/fhir/ValueSet/measure-report-status
type MeasureReportStatus int

const (
	MeasureReportStatusComplete MeasureReportStatus = iota
	MeasureReportStatusPending
	MeasureReportStatusError
)

func (code MeasureReportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MeasureReportStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "complete":
		*code = MeasureReportStatusComplete
	case "pending":
		*code = MeasureReportStatusPending
	case "error":
		*code = MeasureReportStatusError
	default:
		return fmt.Errorf("unknown MeasureReportStatus code `%s`", s)
	}
	return nil
}
func (code MeasureReportStatus) String() string {
	return code.Code()
}
func (code MeasureReportStatus) Code() string {
	switch code {
	case MeasureReportStatusComplete:
		return "complete"
	case MeasureReportStatusPending:
		return "pending"
	case MeasureReportStatusError:
		return "error"
	}
	return "<unknown>"
}
func (code MeasureReportStatus) Display() string {
	switch code {
	case MeasureReportStatusComplete:
		return "Complete"
	case MeasureReportStatusPending:
		return "Pending"
	case MeasureReportStatusError:
		return "Error"
	}
	return "<unknown>"
}
func (code MeasureReportStatus) Definition() string {
	switch code {
	case MeasureReportStatusComplete:
		return "The report is complete and ready for use."
	case MeasureReportStatusPending:
		return "The report is currently being generated."
	case MeasureReportStatusError:
		return "An error occurred attempting to generate the report."
	}
	return "<unknown>"
}
