package pfsenseapi

import "strings"

// StringArray is designed to unmarshal PFSense string arrays which look like
// "192.168.0.1,192.168.1.1"
type StringArray []string

// UnmarshalJSON implements the json.Unmarshaler interface.
func (sa *StringArray) UnmarshalJSON(data []byte) error {
	// Empty string is ""
	if len(data) <= 2 {
		*sa = make(StringArray, 0)
		return nil
	}

	// Remove quotes from string
	data = data[1 : len(data)-1]

	*sa = strings.Split(string(data), ",")
	return nil
}
