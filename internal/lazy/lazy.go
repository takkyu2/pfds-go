package lazy

import "sync"

type Lazy[T any] struct {
	getter func() T
}

func SetLazy[T any](f func() T) *Lazy[T] {
	var l Lazy[T]
	var returnValue T
	var loadOnce sync.Once
	getter := func() T {
		loadOnce.Do(func() { returnValue = f() })
		return returnValue
	}
	l.getter = getter
	return &l
}

func (l *Lazy[T]) Get() T {
	return l.getter()
}
