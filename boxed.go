package boxed

type Box[T any] interface {
	Get() T
	Put(T)
}

func Update[T any, B Box[T]](box B, f func(T) T) {
	box.Put(f(box.Get()))
}

func Swap[T any, B1 Box[T], B2 Box[T]](box1 B1, box2 B2) {
	tmp := box1.Get()
	box1.Put(box2.Get())
	box2.Put(tmp)
}

func FromTo[T any, B1 Box[T], B2 Box[T]](box1 B1, box2 B2) {
	box2.Put(box1.Get())
}

func ToFrom[T any, B1 Box[T], B2 Box[T]](box1 B1, box2 B2) {
	box1.Put(box2.Get())
}
