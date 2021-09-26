package list

type LinkedList[T any] struct {
	elem T
	next *LinkedList[T]
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l == nil
}

func (l *LinkedList[T]) ConsImpl(elem T) *LinkedList[T] {
	return &LinkedList[T]{elem: elem, next: l}
}

func (l *LinkedList[T]) Cons(elem T) StackWithCat[T] {
	return l.ConsImpl(elem)
}

func (l *LinkedList[T]) Head() (T, bool) {
	if l == nil {
		var zeroT T
		return zeroT, false
	}
	return l.elem, true
}

func (l *LinkedList[T]) TailImpl() (*LinkedList[T], bool) {
	if l == nil {
		return l, false
	}
	return l.next, true
}

func (l *LinkedList[T]) Tail() (StackWithCat[T], bool) {
	return l.TailImpl()
}

func (l *LinkedList[T]) HeadTail() (T, *LinkedList[T], bool) {
	if l == nil {
		var zeroT T
		return zeroT, l, false
	}
	return l.elem, l.next, true
}

func (l *LinkedList[T]) ConcatImpl(r *LinkedList[T]) *LinkedList[T] {
	head, tail, ok := l.HeadTail()
	if !ok {
		return r
	}
	return tail.ConcatImpl(r).ConsImpl(head)
}

// Apparently we cannot enforce that r is *LinkedList[T] by using Go's interface, so
// check the concrete type of r at runtime
func (l *LinkedList[T]) Concat(r StackWithCat[T]) StackWithCat[T] {
	switch t := r.(type) {
	case *LinkedList[T]:
		return l.ConcatImpl(t)
	default:
		panic("Should not happen")
	}
}

func (l *LinkedList[T]) Update(i int, elem T) (StackWithCat[T], bool) {
	if i < 0 {
		return l, false
	}

	head, tail, ok := l.HeadTail()
	if !ok {
		return l, false
	}

	if i == 0 {
		return tail.Cons(elem), true
	} else {
		newtail, ok := tail.Update(i-1, elem)
		if !ok {
			return l, false
		}
		return newtail.Cons(head), true
	}
}

func (l *LinkedList[T]) Rev() *LinkedList[T] {
	var newList *LinkedList[T]
	for !l.IsEmpty() {
		head, tail, ok := l.HeadTail()
		if !ok {
			break
		}
		newList = newList.ConsImpl(head)
		l = tail
	}
	return newList
}
