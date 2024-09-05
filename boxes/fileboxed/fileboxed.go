package fileboxed

import (
	"os"

	unmarshal "github.com/monadix/boxed/pkg/un-marshal"
	"github.com/monadix/boxed/pkg/utils"
)

type FileBox[T unmarshal.UnMarshaler] struct {
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

	err = result.Unmarshal(data)
	return result, err
}

func (b FileBox[T]) Put(t T) error {
	data, err := t.Marshal()
	if err != nil {
		return err
	}

	return os.WriteFile(b.path, data, 0644)
}

func New[T unmarshal.UnMarshaler](path string) (FileBox[T], error) {
	_, err := os.Stat(path)
	return FileBox[T]{path: path}, err
}

func NewWith[T unmarshal.UnMarshaler](path string, t T) (FileBox[T], error) {
	data, err := t.Marshal()
	if err != nil {
		return utils.Zero[FileBox[T]](), err
	}

	return FileBox[T]{path: path}, os.WriteFile(path, data, 0644)
}

func NewWithDefault[T unmarshal.UnMarshaler](path string, t T) (FileBox[T], error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return NewWith(path, t)
	}

	return FileBox[T]{path: path}, err
}
