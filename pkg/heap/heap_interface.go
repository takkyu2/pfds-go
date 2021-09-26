package heap

import "github.com/takkyu2/pfds-go/internal/constraints"

type Heap[T constraints.Ordered] interface {
  IsEmpty() bool
  Insert(T) Heap[T]
  Merge(Heap[T]) Heap[T]
  FindMin() (T, bool)
  DeleteMin() (Heap[T], bool)
}
