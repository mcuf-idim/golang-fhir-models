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

// SequenceType is documented here http://hl7.org/fhir/ValueSet/sequence-type
type SequenceType int

const (
	SequenceTypeAa SequenceType = iota
	SequenceTypeDna
	SequenceTypeRna
)

func (code SequenceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SequenceType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "aa":
		*code = SequenceTypeAa
	case "dna":
		*code = SequenceTypeDna
	case "rna":
		*code = SequenceTypeRna
	default:
		return fmt.Errorf("unknown SequenceType code `%s`", s)
	}
	return nil
}
func (code SequenceType) String() string {
	return code.Code()
}
func (code SequenceType) Code() string {
	switch code {
	case SequenceTypeAa:
		return "aa"
	case SequenceTypeDna:
		return "dna"
	case SequenceTypeRna:
		return "rna"
	}
	return "<unknown>"
}
func (code SequenceType) Display() string {
	switch code {
	case SequenceTypeAa:
		return "AA Sequence"
	case SequenceTypeDna:
		return "DNA Sequence"
	case SequenceTypeRna:
		return "RNA Sequence"
	}
	return "<unknown>"
}
func (code SequenceType) Definition() string {
	switch code {
	case SequenceTypeAa:
		return "Amino acid sequence."
	case SequenceTypeDna:
		return "DNA Sequence."
	case SequenceTypeRna:
		return "RNA Sequence."
	}
	return "<unknown>"
}
