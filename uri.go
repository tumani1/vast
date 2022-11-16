package vast

import "bytes"

// URI is a string that allows for stripping of whitespace when Unmarshalled
type URI string

// MarshalText implements the encoding.TextMarshaler interface.
func (s URI) MarshalText() ([]byte, error) {
	return []byte(s), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (s *URI) UnmarshalText(data []byte) error {
	*s = URI(bytes.TrimSpace(data))
	return nil
}

// String implements Stringer interface
func (s URI) String() string {
	return string(s)
}
