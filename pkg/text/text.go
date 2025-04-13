package text

import (
	"encoding/json"
	"strings"
)

type TrimmedString string

func (s *TrimmedString) UnmarshalJSON(b []byte) error {
	var raw string
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	*s = TrimmedString(strings.TrimSpace(raw))
	return nil
}
