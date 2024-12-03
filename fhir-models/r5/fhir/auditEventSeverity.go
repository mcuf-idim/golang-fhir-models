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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// AuditEventSeverity is documented here http://hl7.org/fhir/ValueSet/audit-event-severity
type AuditEventSeverity int

const (
	AuditEventSeverityEmergency AuditEventSeverity = iota
	AuditEventSeverityAlert
	AuditEventSeverityCritical
	AuditEventSeverityError
	AuditEventSeverityWarning
	AuditEventSeverityNotice
	AuditEventSeverityInformational
	AuditEventSeverityDebug
)

func (code AuditEventSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AuditEventSeverity) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "emergency":
		*code = AuditEventSeverityEmergency
	case "alert":
		*code = AuditEventSeverityAlert
	case "critical":
		*code = AuditEventSeverityCritical
	case "error":
		*code = AuditEventSeverityError
	case "warning":
		*code = AuditEventSeverityWarning
	case "notice":
		*code = AuditEventSeverityNotice
	case "informational":
		*code = AuditEventSeverityInformational
	case "debug":
		*code = AuditEventSeverityDebug
	default:
		return fmt.Errorf("unknown AuditEventSeverity code `%s`", s)
	}
	return nil
}
func (code AuditEventSeverity) String() string {
	return code.Code()
}
func (code AuditEventSeverity) Code() string {
	switch code {
	case AuditEventSeverityEmergency:
		return "emergency"
	case AuditEventSeverityAlert:
		return "alert"
	case AuditEventSeverityCritical:
		return "critical"
	case AuditEventSeverityError:
		return "error"
	case AuditEventSeverityWarning:
		return "warning"
	case AuditEventSeverityNotice:
		return "notice"
	case AuditEventSeverityInformational:
		return "informational"
	case AuditEventSeverityDebug:
		return "debug"
	}
	return "<unknown>"
}
func (code AuditEventSeverity) Display() string {
	switch code {
	case AuditEventSeverityEmergency:
		return "Emergency"
	case AuditEventSeverityAlert:
		return "Alert"
	case AuditEventSeverityCritical:
		return "Critical"
	case AuditEventSeverityError:
		return "Error"
	case AuditEventSeverityWarning:
		return "Warning"
	case AuditEventSeverityNotice:
		return "Notice"
	case AuditEventSeverityInformational:
		return "Informational"
	case AuditEventSeverityDebug:
		return "Debug"
	}
	return "<unknown>"
}
func (code AuditEventSeverity) Definition() string {
	switch code {
	case AuditEventSeverityEmergency:
		return "System is unusable. e.g., This level should only be reported by infrastructure and should not be used by applications."
	case AuditEventSeverityAlert:
		return "Notification should be sent to trigger action be taken. e.g., Loss of the primary network connection needing attention."
	case AuditEventSeverityCritical:
		return "Critical conditions. e.g., A failure in the system's primary application that will reset automatically."
	case AuditEventSeverityError:
		return "Error conditions. e.g., An application has exceeded its file storage limit and attempts to write are failing. "
	case AuditEventSeverityWarning:
		return "Warning conditions. May indicate that an error will occur if action is not taken. e.g., A non-root file system has only 2GB remaining."
	case AuditEventSeverityNotice:
		return "Notice messages. Normal but significant condition. Events that are unusual, but not error conditions."
	case AuditEventSeverityInformational:
		return "Normal operational messages that require no action. e.g., An application has started, paused, or ended successfully."
	case AuditEventSeverityDebug:
		return "Debug-level messages. Information useful to developers for debugging the application."
	}
	return "<unknown>"
}
