package main

import "sync"

type streamInterface[T any] interface {
  concat(streamInterface[T]) streamInterface[T]
  take(int) streamInterface[T]
  drop(int) streamInterface[T]
  reverse() streamInterface[T]
}

type lazy[T any] struct {
  getter func() T
}

func setLazy[T any](f func() T) *lazy[T] {
  var l lazy[T]
  var returnValue T
  var loadOnce sync.Once
  getter := func() T { 
    loadOnce.Do(func() { returnValue = f() })
    return returnValue
  }
  l.getter = getter
  return &l
}

func (l *lazy[T]) get() T {
  return l.getter()
}

type streamCell[T any] struct {
  elem T
  next stream[T]
}

type stream[T any] struct {
  cell *lazy[*streamCell[T]]
}

func (s stream[T]) isEmpty() bool {
  if s.cell == nil { return true }
  return s.cell.get() == nil
}

func (s stream[T]) head() (T, bool) {
  if s.isEmpty() {
    var zeroT T
    return zeroT, false
  }
  return s.cell.get().elem, true
}

func (s stream[T]) tail() (stream[T], bool) {
  if s.isEmpty() { return s, false }
  return s.cell.get().next, true
}

func (s stream[T]) cons(elem T) stream[T] {
  return stream[T]{cell: setLazy(func() *streamCell[T] {
    return &streamCell[T]{elem:elem, next:s}
  })}
}

func (s stream[T]) headTail() (T, stream[T], bool) {
  if s.isEmpty() { 
    var zeroT T
    return zeroT, s, false 
  }
  return s.cell.get().elem, s.cell.get().next, true
}

func (s stream[T]) concatImpl(t stream[T]) stream[T] {
  return stream[T]{cell: setLazy(func() *streamCell[T] {
    x, s, ok := s.headTail()
    if !ok { return t.cell.get() }
    return &streamCell[T] {elem:x, next:s.concatImpl(t)}
  })}
}

func (s stream[T]) concat(t streamInterface[T]) streamInterface[T] {
  switch t := t.(type) {
  case stream[T]:
    return s.concatImpl(t)
  default:
    panic("cannot happen")
  }
}

func (s stream[T]) takeImpl(i int) stream[T] {
  return stream[T]{cell: setLazy(func() *streamCell[T] {
    if i == 0 { return nil }
    x, s, ok := s.headTail()
    if !ok { return nil }
    return &streamCell[T] {elem:x, next:s.takeImpl(i-1)}
  })}
}

func (s stream[T]) take(i int) streamInterface[T] {
  return s.takeImpl(i)
}

func (s stream[T]) dropImpl(i int) stream[T] {
  return stream[T]{cell: setLazy(func() *streamCell[T] {
    if i == 0 { return s.cell.get() }
    _, s, ok := s.headTail()
    if !ok { return nil }
    return s.takeImpl(i-1).cell.get()
  })}
}

func (s stream[T]) drop(i int) streamInterface[T] {
  return s.dropImpl(i)
}

func (s stream[T]) revHelper(accum stream[T]) stream[T] {
  x, s2, ok := s.headTail()
  if !ok { return accum }
  return s2.revHelper(accum.cons(x))
}

func (s stream[T]) reverseImpl() stream[T] {
  return stream[T]{cell: setLazy(func() *streamCell[T] {
    var emptyStream stream[T]
    return s.revHelper(emptyStream).cell.get()
  })}
}

func (s stream[T]) reverse() streamInterface[T] {
  return s.reverseImpl()
}
