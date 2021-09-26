package list

import "github.com/takkyu2/pfds-go/internal/lazy"

type StreamInterface[T any] interface {
	Concat(StreamInterface[T]) StreamInterface[T]
	Take(int) StreamInterface[T]
	Drop(int) StreamInterface[T]
	Reverse() StreamInterface[T]
}

type streamCell[T any] struct {
	elem T
	next Stream[T]
}

type Stream[T any] struct {
	cell *lazy.Lazy[*streamCell[T]]
}

func (s Stream[T]) IsEmpty() bool {
	if s.cell == nil {
		return true
	}
	return s.cell.Get() == nil
}

func (s Stream[T]) Head() (T, bool) {
	if s.IsEmpty() {
		var zeroT T
		return zeroT, false
	}
	return s.cell.Get().elem, true
}

func (s Stream[T]) Tail() (Stream[T], bool) {
	if s.IsEmpty() {
		return s, false
	}
	return s.cell.Get().next, true
}

func (s Stream[T]) Cons(elem T) Stream[T] {
	return Stream[T]{cell: lazy.SetLazy(func() *streamCell[T] {
		return &streamCell[T]{elem: elem, next: s}
	})}
}

func (s Stream[T]) HeadTail() (T, Stream[T], bool) {
	if s.IsEmpty() {
		var zeroT T
		return zeroT, s, false
	}
	return s.cell.Get().elem, s.cell.Get().next, true
}

func (s Stream[T]) ConcatImpl(t Stream[T]) Stream[T] {
	return Stream[T]{cell: lazy.SetLazy(func() *streamCell[T] {
		x, s, ok := s.HeadTail()
		if !ok {
			return t.cell.Get()
		}
		return &streamCell[T]{elem: x, next: s.ConcatImpl(t)}
	})}
}

func (s Stream[T]) Concat(t StreamInterface[T]) StreamInterface[T] {
	switch t := t.(type) {
	case Stream[T]:
		return s.ConcatImpl(t)
	default:
		panic("cannot happen")
	}
}

func (s Stream[T]) TakeImpl(i int) Stream[T] {
	return Stream[T]{cell: lazy.SetLazy(func() *streamCell[T] {
		if i == 0 {
			return nil
		}
		x, s, ok := s.HeadTail()
		if !ok {
			return nil
		}
		return &streamCell[T]{elem: x, next: s.TakeImpl(i - 1)}
	})}
}

func (s Stream[T]) Take(i int) StreamInterface[T] {
	return s.TakeImpl(i)
}

func (s Stream[T]) DropImpl(i int) Stream[T] {
	return Stream[T]{cell: lazy.SetLazy(func() *streamCell[T] {
		if i == 0 {
			return s.cell.Get()
		}
		_, s, ok := s.HeadTail()
		if !ok {
			return nil
		}
		return s.TakeImpl(i - 1).cell.Get()
	})}
}

func (s Stream[T]) Drop(i int) StreamInterface[T] {
	return s.DropImpl(i)
}

func (s Stream[T]) revHelper(accum Stream[T]) Stream[T] {
	x, s2, ok := s.HeadTail()
	if !ok {
		return accum
	}
	return s2.revHelper(accum.Cons(x))
}

func (s Stream[T]) ReverseImpl() Stream[T] {
	return Stream[T]{cell: lazy.SetLazy(func() *streamCell[T] {
		var emptyStream Stream[T]
		return s.revHelper(emptyStream).cell.Get()
	})}
}

func (s Stream[T]) Reverse() StreamInterface[T] {
	return s.ReverseImpl()
}
