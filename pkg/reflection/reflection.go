package reflection

import (
	"reflect"

	"github.com/monadix/boxed/pkg/utils"
)

func GetInsAndOuts(typ reflect.Type) ([]reflect.Type, []reflect.Type) {
	ins := make([]reflect.Type, 0, typ.NumIn())
	outs := make([]reflect.Type, 0, typ.NumOut())

	for i := 0; i < typ.NumIn(); i++ {
		ins = append(ins, typ.In(i))
	}

	for i := 0; i < typ.NumOut(); i++ {
		outs = append(outs, typ.Out(i))
	}

	return ins, outs
}

func GetMethodWithTypes(val any, name string, ins []reflect.Type, outs []reflect.Type) (reflect.Value, error) {
	meth := reflect.ValueOf(val).MethodByName(name)

	if !meth.IsValid() {
		return utils.Zero[reflect.Value](), ErrNoSuchMethod{
			missingMethod: name,
			gotType:       reflect.TypeOf(val),
		}
	}

	gotIns, gotOuts := GetInsAndOuts(meth.Type())

	errInvalidMethod := ErrInvalidMethod{
		gotType:       reflect.TypeOf(val),
		invalidMethod: name,
		gotIns:        gotIns,
		gotOuts:       gotOuts,
		expectedIns:   ins,
		expectedOuts:  outs,
	}

	if meth.Type().NumIn() != len(ins) ||
		meth.Type().NumOut() != len(outs) {

		return utils.Zero[reflect.Value](), errInvalidMethod
	}

	for i := range ins {
		if !gotIns[i].AssignableTo(ins[i]) {
			return utils.Zero[reflect.Value](), errInvalidMethod
		}
	}

	for i := range outs {
		if !gotOuts[i].AssignableTo(outs[i]) {
			return utils.Zero[reflect.Value](), errInvalidMethod
		}
	}

	return meth, nil
}

func CallMethodWithTypes(val any, name string, ins []reflect.Value, outTypes []reflect.Type) ([]reflect.Value, error) {
	inTypes := make([]reflect.Type, 0, len(ins))

	for _, in := range ins {
		inTypes = append(inTypes, in.Type())
	}

	meth, err := GetMethodWithTypes(val, name, inTypes, outTypes)
	if err != nil {
		return nil, err
	}

	return meth.Call(ins), nil
}
