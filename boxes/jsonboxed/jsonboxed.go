package jsonboxed

import "encoding/json"

type JsonBox[T any] struct {
	value *T
}

func (b JsonBox[T]) Get() (T, error) {
	return *b.value, nil
}

func (b JsonBox[T]) Put(t T) error {
	*b.value = t
	return nil
}

func (b JsonBox[T]) Marshal() ([]byte, error) {
	return json.Marshal(b.value)
}

func (b JsonBox[T]) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &b.value)
}

func New[T any](val T) JsonBox[T] {
	return JsonBox[T]{value: &val}
}
