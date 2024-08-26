package multibox

import (
	"fmt"
	"reflect"
)

type ErrMultiBoxInvalidType struct {
	missingMethod string
	gotType       reflect.Type
	expectedType  reflect.Type
}

func (e ErrMultiBoxInvalidType) Error() string {
	return fmt.Sprintf(
		"while working with MultiBox got type %v, but expected either a boxed.Box (missing method: %v) or a %v",
		e.gotType,
		e.missingMethod,
		e.expectedType,
	)
}

type ErrMultiBoxInvalidMethod struct {
	invalidMethod string
	gotIns        []reflect.Type
	gotOuts       []reflect.Type
	expectedIns   []reflect.Type
	expectedOuts  []reflect.Type
}

func (e ErrMultiBoxInvalidMethod) Error() string {
	return fmt.Sprintf(
		"while working with MultiBox got method %v with ins %v and outs %v, but expected ins %v and outs %v",
		e.invalidMethod,
		e.gotIns,
		e.gotOuts,
		e.expectedIns,
		e.expectedOuts,
	)
}
