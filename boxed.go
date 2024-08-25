package boxed

type Box[T any] interface {
	Get() T
	Put(T)
}

func Update[T any, B Box[T]](box B, f func(T) T) {
	box.Put(f(box.Get()))
}
