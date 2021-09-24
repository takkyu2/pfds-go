package main

type stack[T any] interface {
  // empty() stack[T] Not necessary in Go
  isEmpty() bool
  cons(T) stack[T]
  head() (T, bool)
  tail() (stack[T], bool)
}

type stackWithCat[T any] interface {
  // empty() stack[T] Not necessary in Go
  isEmpty() bool
  cons(T) stackWithCat[T]
  head() (T, bool)
  tail() (stackWithCat[T], bool)
  update(i int, elem T) (stackWithCat[T], bool)
  concat(stackWithCat[T]) stackWithCat[T]
}
