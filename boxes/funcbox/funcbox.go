package funcbox

type FuncBox[T any] struct {
	get func() (T, error)
	put func(T) error
}

func New[T any](get func() (T, error), put func(T) error) FuncBox[T] {
	return FuncBox[T]{get: get, put: put}
}

func (b FuncBox[T]) Get() (T, error) {
	return b.get()
}

func (b FuncBox[T]) Put(t T) error {
	return b.put(t)
}
