package heap

import "github.com/takkyu2/pfds-go/internal/constraints"

type LeftistHeap[T constraints.Ordered] struct {
  rk int
  elem T
  left, right *LeftistHeap[T]
}

func makeHeap[T constraints.Ordered](elem T, left, right *LeftistHeap[T]) *LeftistHeap[T] {
  if left.rank() >= right.rank() {
    return &LeftistHeap[T]{rk: right.rank()+1, elem:elem, left:left, right:right}
  } else {
    return &LeftistHeap[T]{rk: left.rank()+1, elem:elem, left:right, right:left}
  }
}

func (h *LeftistHeap[T]) IsEmpty() bool {
  return h == nil 
}

func (h *LeftistHeap[T]) rank() int {
  if h.IsEmpty() { return 0 }
  return h.rk
}

func (lhs *LeftistHeap[T]) MergeImpl(rhs *LeftistHeap[T]) *LeftistHeap[T] {
  if lhs.IsEmpty() { return rhs }
  if rhs.IsEmpty() { return lhs }
  elem1, left1, right1 := lhs.elem, lhs.left, lhs.right
  elem2, left2, right2 := rhs.elem, rhs.left, rhs.right
  if elem1 < elem2 {
    return makeHeap(elem1, left1, right1.MergeImpl(rhs))
  } else {
    return makeHeap(elem2, left2, lhs.MergeImpl(right2))
  }
}

func (lhs *LeftistHeap[T]) Merge(rhs Heap[T]) Heap[T] {
  switch rhs := rhs.(type) {
  case *LeftistHeap[T]:
    return lhs.MergeImpl(rhs)
  default:
    panic("Should not happen")
  }
}

func (h *LeftistHeap[T]) InsertImpl(elem T) *LeftistHeap[T] {
  return (&LeftistHeap[T]{rk:1, elem:elem}).MergeImpl(h)
}

func (h *LeftistHeap[T]) Insert(elem T) Heap[T] {
  return h.InsertImpl(elem)
}

func (h *LeftistHeap[T]) FindMin() (T, bool) {
  if h.IsEmpty() {
    var zeroT T
    return zeroT, false 
  }
  return h.elem, true
}

func (h *LeftistHeap[T]) DeleteMinImpl() (*LeftistHeap[T], bool) {
  if h.IsEmpty() {
    return h, false 
  }
  return h.left.MergeImpl(h.right), true
}

func (h *LeftistHeap[T]) DeleteMin() (Heap[T], bool) {
  return h.DeleteMinImpl()
}

