package multibox

import (
	"reflect"

	"github.com/monadix/boxed"
	"github.com/monadix/boxed/pkg/boxutils"
	"github.com/monadix/boxed/pkg/reflection"
	"github.com/monadix/boxed/pkg/utils"
)

type MultiBox[T any] struct {
	value any
}

func New[T any](val any) (MultiBox[T], error) {
	if reflect.TypeOf(val) == reflect.TypeFor[T]() {
		return utils.Zero[MultiBox[T]](), ErrUnboxedTargetType{
			gotType: reflect.TypeOf(val),
		}
	}

	res := MultiBox[T]{
		value: val,
	}

	for reflect.TypeOf(val) != reflect.TypeFor[T]() {
		box, err := boxutils.MagicAsBox[any](val)
		if err != nil {
			return utils.Zero[MultiBox[T]](), err
		}

		val, err = box.Get()
		if err != nil {
			return utils.Zero[MultiBox[T]](), err
		}

		res.value = val
	}

	return res, nil
}

func (b MultiBox[T]) Get() (T, error) {
	res := b.value

	for reflect.TypeOf(res) != reflect.TypeFor[T]() {
		methGet, err := reflection.GetMethodWithTypes(res, "Get",
			[]reflect.Type{},
			[]reflect.Type{reflect.TypeFor[any](), reflect.TypeFor[error]()},
		)
		if err != nil {
			return utils.Zero[T](), err
		}

		vals := methGet.Call([]reflect.Value{})
		if err := utils.CastOrNil[error](vals[1].Interface()); err != nil {
			return utils.Zero[T](), err
		}

		res = vals[0].Interface()
	}

	return res.(T), nil
}

func (b MultiBox[T]) Put(val T) error {
	if box, isBox := b.value.(boxed.Box[T]); isBox {
		box.Put(val)
	}

	box, err := boxutils.MagicAsBox[any](b.value)
	if err != nil {
		return err
	}

	err = boxed.Update(box, func(value any) (any, error) {
		return value, MultiBox[T]{value}.Put(val)
	})
	if err != nil {
		return err
	}

	return nil
}
