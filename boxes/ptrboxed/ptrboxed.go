package ptrboxed

type PtrBox[T any] struct {
	ptr *T
}

func (b PtrBox[T]) Get() T {
	return *b.ptr
}

func (b PtrBox[T]) Put(t T) {
	*b.ptr = t
}

func New[T any](t T) PtrBox[T] {
	return PtrBox[T]{ptr: &t}
}
