package main

type leftiestHeap[T ordered] struct {
  rk int
  elem T
  left, right *leftiestHeap[T]
}

func makeHeap[T ordered](elem T, left, right *leftiestHeap[T]) *leftiestHeap[T] {
  if left.rank() >= right.rank() {
    return &leftiestHeap[T]{rk: right.rank()+1, elem:elem, left:left, right:right}
  } else {
    return &leftiestHeap[T]{rk: left.rank()+1, elem:elem, left:right, right:left}
  }
}

func (h *leftiestHeap[T]) isEmpty() bool {
  return h == nil 
}

func (h *leftiestHeap[T]) rank() int {
  if h.isEmpty() { return 0 }
  return h.rk
}

func (lhs *leftiestHeap[T]) mergeImpl(rhs *leftiestHeap[T]) *leftiestHeap[T] {
  if lhs.isEmpty() { return rhs }
  if rhs.isEmpty() { return lhs }
  elem1, left1, right1 := lhs.elem, lhs.left, lhs.right
  elem2, left2, right2 := rhs.elem, rhs.left, rhs.right
  if elem1 < elem2 {
    return makeHeap(elem1, left1, right1.mergeImpl(rhs))
  } else {
    return makeHeap(elem2, left2, lhs.mergeImpl(right2))
  }
}

func (lhs *leftiestHeap[T]) merge(rhs heap[T]) heap[T] {
  switch rhs := rhs.(type) {
  case *leftiestHeap[T]:
    return lhs.mergeImpl(rhs)
  default:
    panic("Should not happen")
  }
}

func (h *leftiestHeap[T]) insert(elem T) heap[T] {
  return (&leftiestHeap[T]{rk:1, elem:elem}).merge(h)
}

func (h *leftiestHeap[T]) findMin() (T, bool) {
  if h.isEmpty() {
    var zeroT T
    return zeroT, false 
  }
  return h.elem, true
}

func (h *leftiestHeap[T]) deleteMin() (heap[T], bool) {
  if h.isEmpty() {
    return h, false 
  }
  return h.left.merge(h.right), true
}

type binomialHeapTree[T ordered] struct {
  rk int
  elem T
  tList *linkedList[*binomialHeapTree[T]]
}

type binomialHeap[T ordered] struct {
  tList *linkedList[*binomialHeapTree[T]]
}

func newBinomialHeap[T ordered]() *binomialHeap[T] {
  return &binomialHeap[T]{}
}

func (bh *binomialHeap[T]) isEmpty() bool {
  return bh.tList.isEmpty()
}

func (bt binomialHeapTree[T]) rank() int {
  return bt.rk
}

func (bt binomialHeapTree[T]) root() T {
  return bt.elem
}

func (bt1 binomialHeapTree[T]) link(bt2 binomialHeapTree[T]) binomialHeapTree[T] {
  if bt1.elem <= bt2.elem {
    return binomialHeapTree[T]{rk:bt1.rk+1, elem:bt1.elem, tList:bt1.tList.consImpl(&bt2)}
  } else {
    return binomialHeapTree[T]{rk:bt1.rk+1, elem:bt2.elem, tList:bt2.tList.consImpl(&bt1)}
  }
}

func (bt binomialHeapTree[T]) insTree(tList *linkedList[*binomialHeapTree[T]]) *linkedList[*binomialHeapTree[T]] {
  head, tail, ok := tList.headTail()
  if !ok {
    var newTList *linkedList[*binomialHeapTree[T]]
    return newTList.consImpl(&bt)
  }
  if bt.rank() < head.rank() {
    return tList.consImpl(&bt)
  } else {
    return bt.link(*head).insTree(tail)
  }
}

func (bh *binomialHeap[T]) insertImpl(t T) *binomialHeap[T] {
  tr := binomialHeapTree[T]{rk: 0, elem:t}
  newTList := tr.insTree(bh.tList)
  return &binomialHeap[T]{tList:newTList}
}

func (bh *binomialHeap[T]) insert(t T) heap[T] {
  return bh.insertImpl(t)
}

func (lhs *binomialHeap[T]) mergeImpl(rhs *binomialHeap[T]) *binomialHeap[T] {
  if (rhs.isEmpty()) { return lhs }
  if (lhs.isEmpty()) { return rhs }
  head1, tail1, _ := lhs.tList.headTail()
  head2, tail2, _ := rhs.tList.headTail()
  if head1.rank() < head2.rank() {
    bh := &binomialHeap[T]{tList: tail1}
    return &binomialHeap[T]{tList: bh.mergeImpl(rhs).tList.consImpl(head1)}
  } else if head2.rank() < head1.rank() {
    bh := &binomialHeap[T]{tList: tail2}
    return &binomialHeap[T]{tList: lhs.mergeImpl(bh).tList.consImpl(head2)}
  } else {
    tail1 := &binomialHeap[T]{tList: tail1}
    tail2 := &binomialHeap[T]{tList: tail2}
    newTList := head1.link(*head2).insTree(tail1.mergeImpl(tail2).tList)
    return &binomialHeap[T]{tList: newTList}
  }
}

func (lhs *binomialHeap[T]) merge(rhs heap[T]) heap[T] {
  switch rhs := rhs.(type) {
  case *binomialHeap[T]:
    return lhs.mergeImpl(rhs)
  default:
    panic("Should not happen")
  }
}

type Pair[T any, S any] struct {
  fi T
  se S
}

func (bh *binomialHeap[T]) removeMinTree() (Pair[*binomialHeapTree[T], *linkedList[*binomialHeapTree[T]]], bool)  {
  head, tail, ok := bh.tList.headTail()
  if !ok { return Pair[*binomialHeapTree[T], *linkedList[*binomialHeapTree[T]]]{}, false }
  if tail.isEmpty() {
    return Pair[*binomialHeapTree[T], *linkedList[*binomialHeapTree[T]]]{fi: head}, true
  }
  bh2 := &(binomialHeap[T]{tList:tail})
  headTail2, _ := bh2.removeMinTree()
  head2, tail2 := headTail2.fi, headTail2.se
  if head.elem < head2.elem {
    return Pair[*binomialHeapTree[T], *linkedList[*binomialHeapTree[T]]]{head, tail}, true
  } else {
    return Pair[*binomialHeapTree[T], *linkedList[*binomialHeapTree[T]]]{head2, tail2.consImpl(head)}, true
  }
}

func (bh *binomialHeap[T]) findMin() (T, bool) {
  t, ok := bh.removeMinTree()
  if !ok {
    var zeroT T
    return zeroT, false
  }
  return t.fi.elem, true
}

func (bh *binomialHeap[T]) deleteMinImpl() (*binomialHeap[T], bool) {
  t, ok := bh.removeMinTree()
  if !ok { return bh, false }
  node, ts2 := t.fi, t.se
  ts1 := node.tList
  ts3 := &binomialHeap[T]{tList:ts1.rev()}
  ts4 := &binomialHeap[T]{tList:ts2}
  return ts3.mergeImpl(ts4), true
}

func (bh *binomialHeap[T]) deleteMin() (heap[T], bool) {
  return bh.deleteMinImpl()
}
