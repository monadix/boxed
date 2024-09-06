package boxutils

import (
	"reflect"

	"github.com/monadix/boxed"
	"github.com/monadix/boxed/boxes/funcbox"
	"github.com/monadix/boxed/pkg/reflection"
	"github.com/monadix/boxed/pkg/utils"
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
			return utils.CastOrNil[T](res[0].Interface()), utils.CastOrNil[error](res[1].Interface())
		},
		func(v T) error {
			res := methPut.Call([]reflect.Value{reflect.ValueOf(v)})
			return utils.CastOrNil[error](res[0].Interface())
		},
	), nil
}
