package main

type heap[T ordered] interface {
  isEmpty() bool
  insert(T) heap[T]
  merge(heap[T]) heap[T]
  findMin() (T, bool)
  deleteMin() (heap[T], bool)
}
