package pfsenseapi

import (
	"encoding/json"
	"strconv"
)

// OptionalInt can unmarshal both JSON numbers and strings into an integer.
type OptionalJSONInt struct {
	Value *int
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (jsi *OptionalJSONInt) UnmarshalJSON(data []byte) error {
	// Try unmarshalling as int
	var intValue int
	if err := json.Unmarshal(data, &intValue); err != nil {
		var stringValue string
		if err := json.Unmarshal(data, &stringValue); err != nil {
			return err
		}

		if stringValue == "" {
			*jsi = OptionalJSONInt{}
			return nil
		} else {
			intValue, err = strconv.Atoi(stringValue)

			if err != nil {
				return err
			}
		}
	}

	*jsi = OptionalJSONInt{
		Value: &intValue,
	}

	return nil
}

type JSONInt int

// UnmarshalJSON implements the json.Unmarshaler interface.
func (jsi *JSONInt) UnmarshalJSON(data []byte) error {
	// Try unmarshalling as int
	var intValue int
	if err := json.Unmarshal(data, &intValue); err != nil {
		var stringValue string
		if err := json.Unmarshal(data, &stringValue); err != nil {
			return err
		}

		intValue, err = strconv.Atoi(stringValue)

		if err != nil {
			return err
		}
	}

	*jsi = JSONInt(intValue)

	return nil
}
