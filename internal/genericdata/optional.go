package genericdata

type Optional[T any] struct {
	hasValue bool
	value    T
}

func ToOpt[T any](elem T) Optional[T] {
	return Optional[T]{hasValue: true, value: elem}
}

func (opt Optional[T]) FromOpt() (T, bool) {
	return opt.value, opt.hasValue
}

func (opt Optional[T]) IsNil() bool {
	return !opt.hasValue
}
