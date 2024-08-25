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
