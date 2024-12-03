package types

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
