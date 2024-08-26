package multibox

import (
	"reflect"

	"github.com/monadix/boxed"
	"github.com/monadix/boxed/pkg/utils"
)

type MultiBox[T any] struct {
	box boxed.Box[any]
}

func deepUnbox[T any](v any) (T, error) {
	if reflect.TypeOf(v) == reflect.TypeFor[T]() {
		return v.(T), nil
	}

	methodGet := reflect.ValueOf(v).MethodByName("Get")

	if !methodGet.IsValid() {
		return utils.Zero[T](), ErrMultiBoxInvalidType{
			missingMethod: "Get",
			gotType:       reflect.TypeOf(v),
			expectedType:  reflect.TypeFor[T](),
		}
	}

	methGetType := methodGet.Type()
	if methGetType.NumIn() != 0 || methGetType.NumOut() != 2 || !methGetType.Out(1).Implements(reflect.TypeFor[error]()) {
		ins := make([]reflect.Type, 0, methGetType.NumIn())
		outs := make([]reflect.Type, 0, methGetType.NumOut())

		for i := 0; i < methGetType.NumIn(); i++ {
			ins = append(ins, methGetType.In(i))
		}

		for i := 0; i < methGetType.NumOut(); i++ {
			outs = append(outs, methGetType.Out(i))
		}

		return utils.Zero[T](), ErrMultiBoxInvalidMethod{
			invalidMethod: "Get",
			gotIns:        ins,
			gotOuts:       outs,
			expectedIns:   []reflect.Type{},
			expectedOuts:  []reflect.Type{reflect.TypeFor[any](), reflect.TypeFor[error]()},
		}
	}

	return utils.Zero[T](), nil //TODO
}
