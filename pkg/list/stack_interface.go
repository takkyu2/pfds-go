package list

type Stack[T any] interface {
  // empty() stack[T] Not necessary in Go
  IsEmpty() bool
  Cons(T) Stack[T]
  Head() (T, bool)
  Tail() (Stack[T], bool)
}

type StackWithCat[T any] interface {
  // empty() stack[T] Not necessary in Go
  IsEmpty() bool
  Cons(T) StackWithCat[T]
  Head() (T, bool)
  Tail() (StackWithCat[T], bool)
  Update(i int, elem T) (StackWithCat[T], bool)
  Concat(StackWithCat[T]) StackWithCat[T]
}
