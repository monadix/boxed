package multibox

import (
	"fmt"
	"reflect"
)

type ErrUnboxedTargetType struct {
	gotType reflect.Type
}

func (e ErrUnboxedTargetType) Error() string {
	return fmt.Sprintf(
		"got unboxed target type %v on the first layer while initializing a MultiBox",
		e.gotType,
	)
}
