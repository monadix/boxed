package fileboxed

import (
	"bytes"
	"encoding/gob"
	"os"

	"github.com/monadix/boxed/pkg/utils"
)

type FileBox[T any] struct {
	path string
}

func (b FileBox[T]) Path() string {
	return b.path
}

func (b FileBox[T]) Get() (T, error) {
	var result T

	data, err := os.ReadFile(b.path)
	if err != nil {
		return utils.Zero[T](), err
	}

	err = gob.NewDecoder(bytes.NewBuffer(data)).Decode(&result)
	return result, err
}

func (b FileBox[T]) Put(t T) error {
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(t)
	if err != nil {
		return err
	}

	return os.WriteFile(b.path, buffer.Bytes(), 0644)
}

func New[T any](path string) (FileBox[T], error) {
	_, err := os.Stat(path)
	return FileBox[T]{path: path}, err
}

func NewWith[T any](path string, t T) (FileBox[T], error) {
	var buffer bytes.Buffer

	err := gob.NewEncoder(&buffer).Encode(t)
	if err != nil {
		return utils.Zero[FileBox[T]](), err
	}

	return FileBox[T]{path: path}, os.WriteFile(path, buffer.Bytes(), 0644)
}

func NewWithDefault[T any](path string, t T) (FileBox[T], error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return NewWith(path, t)
	}

	return FileBox[T]{path: path}, err
}
