package boxutils

import (
	"reflect"

	"github.com/monadix/boxed"
	"github.com/monadix/boxed/boxes/funcbox"
	"github.com/monadix/boxed/pkg/reflection"
)

func MagicAsBox[T any](val any) (boxed.Box[T], error) {
	methGet, err := reflection.GetMethodWithTypes(val, "Get",
		[]reflect.Type{},
		[]reflect.Type{reflect.TypeFor[T](), reflect.TypeFor[error]()},
	)
	if err != nil {
		return nil, err
	}

	methPut, err := reflection.GetMethodWithTypes(val, "Put",
		[]reflect.Type{reflect.TypeFor[T]()},
		[]reflect.Type{reflect.TypeFor[error]()},
	)
	if err != nil {
		return nil, err
	}

	return funcbox.New(
		func() (T, error) {
			res := methGet.Call([]reflect.Value{})
			return res[0].Interface().(T), res[1].Interface().(error)
		},
		func(v T) error {
			res := methPut.Call([]reflect.Value{reflect.ValueOf(v)})
			return res[0].Interface().(error)
		},
	), nil
}
