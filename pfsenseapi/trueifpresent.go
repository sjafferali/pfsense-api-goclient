package pfsenseapi

// TrueIfPresent is designed to unmarshal PFSense boolean values that can indicate
// truth by having an empty string as the value of the property
type TrueIfPresent bool

// UnmarshalJSON implements the json.Unmarshaler interface.
func (tip *TrueIfPresent) UnmarshalJSON(data []byte) error {
	// If it has any value at all it's true
	*tip = true

	return nil
}
