package main

type linkedList[T any] struct {
  elem T
  next *linkedList[T]
}

func (l *linkedList[T]) isEmpty() bool {
  return l == nil;
}

func (l *linkedList[T]) cons(elem T) stackWithCat[T] {
  return &linkedList[T] {elem:elem, next:l}
}

func (l *linkedList[T]) head() (T, bool) {
  if l == nil {
    var zeroT T
    return zeroT, false
  }
  return l.elem, true;
}

func (l *linkedList[T]) tail() (stackWithCat[T], bool) {
  if l == nil {
    return l, false
  }
  return l.next, true;
}

func (l *linkedList[T]) headTail() (T, *linkedList[T], bool) {
  if l == nil {
    var zeroT T
    return zeroT, l, false
  }
  return l.elem, l.next, true
}

func (l *linkedList[T]) concat(r stackWithCat[T]) stackWithCat[T] {
  head, tail, ok := l.headTail()
  if !ok { return r }
  return tail.concat(r).cons(head)
}

func (l *linkedList[T]) update(i int, elem T) (stackWithCat[T], bool) {
  if i < 0 { return l, false; }

  head, tail, ok := l.headTail()
  if !ok { return l, false; }

  if i == 0 {
    return tail.cons(elem), true
  } else {
    newtail, ok := tail.update(i-1, elem)
    if !ok {
      return l, false;
    }
    return newtail.cons(head), true
  }
}
