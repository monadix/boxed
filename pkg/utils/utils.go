package utils

func Zero[T any]() T {
	var t T
	return t
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func CastOrNil[T any](val any) T {
	if val == nil {
		return Zero[T]()
	}

	return val.(T)
}
