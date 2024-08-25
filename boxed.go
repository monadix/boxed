package boxed

type Box[T any] interface {
	Get() (T, error)
	Put(T) error
}

func Update[T any, B Box[T]](box B, f func(T) T) error {
	val, err := box.Get()
	if err != nil {
		return err
	}

	return box.Put(f(val))
}

func Swap[T any, B1 Box[T], B2 Box[T]](box1 B1, box2 B2) error {
	val1, err := box1.Get()
	if err != nil {
		return err
	}

	val2, err := box2.Get()
	if err != nil {
		return err
	}

	err = box1.Put(val2)
	if err != nil {
		return err
	}

	return box2.Put(val1)
}

func FromTo[T any, B1 Box[T], B2 Box[T]](box1 B1, box2 B2) error {
	tmp, err := box1.Get()
	if err != nil {
		return err
	}

	return box2.Put(tmp)
}
