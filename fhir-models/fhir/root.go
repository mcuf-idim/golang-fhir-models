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

//go:generate fhir-models-gen gen-resources .

package fhir

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Integer64 int64

func (code Integer64) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatInt(int64(code), 10))
}

func (code *Integer64) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse %s as integer64: %w", json, err)
	}
	*code = Integer64(i)
	return nil
}

func (code *Integer64) Int64() int64 {
	return int64(*code)
}
