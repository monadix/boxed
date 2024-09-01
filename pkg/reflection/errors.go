package reflection

import (
	"fmt"
	"reflect"
)

type ErrNoSuchMethod struct {
	missingMethod string
	gotType       reflect.Type
}

func (e ErrNoSuchMethod) Error() string {
	return fmt.Sprintf("no such method %v on type %v", e.missingMethod, e.gotType)
}

type ErrInvalidMethod struct {
	gotType       reflect.Type
	invalidMethod string
	gotIns        []reflect.Type
	gotOuts       []reflect.Type
	expectedIns   []reflect.Type
	expectedOuts  []reflect.Type
}

func (e ErrInvalidMethod) Error() string {
	return fmt.Sprintf(
		"method %v on type %v has inputs %v and outputs %v, but expected inputs %v and outputs %v",
		e.invalidMethod,
		e.gotType,
		e.gotIns,
		e.gotOuts,
		e.expectedIns,
		e.expectedOuts,
	)
}
