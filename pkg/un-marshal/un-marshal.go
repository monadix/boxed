package unmarshal

type Marshaler interface {
	Marshal() ([]byte, error)
}

type Unmarshaler interface {
	Unmarshal([]byte) error
}

type UnMarshaler interface {
	Marshaler
	Unmarshaler
}
