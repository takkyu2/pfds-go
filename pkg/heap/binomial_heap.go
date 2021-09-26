package heap

import (
	"github.com/takkyu2/pfds-go/internal/constraints"
	"github.com/takkyu2/pfds-go/internal/genericdata"
	"github.com/takkyu2/pfds-go/pkg/list"
)

type binomialHeapTree[T constraints.Ordered] struct {
  rk int
  elem T
  tList *list.LinkedList[*binomialHeapTree[T]]
}

func (bt binomialHeapTree[T]) rank() int {
  return bt.rk
}

func (bt binomialHeapTree[T]) root() T {
  return bt.elem
}

func (bt1 binomialHeapTree[T]) link(bt2 binomialHeapTree[T]) binomialHeapTree[T] {
  if bt1.elem <= bt2.elem {
    return binomialHeapTree[T]{rk:bt1.rk+1, elem:bt1.elem, tList:bt1.tList.ConsImpl(&bt2)}
  } else {
    return binomialHeapTree[T]{rk:bt1.rk+1, elem:bt2.elem, tList:bt2.tList.ConsImpl(&bt1)}
  }
}

type BinomialHeap[T constraints.Ordered] struct {
  tList *list.LinkedList[*binomialHeapTree[T]]
}

func newBinomialHeap[T constraints.Ordered]() *BinomialHeap[T] {
  return &BinomialHeap[T]{}
}

func (bh *BinomialHeap[T]) IsEmpty() bool {
  return bh.tList.IsEmpty()
}

func (bt binomialHeapTree[T]) insTree(tList *list.LinkedList[*binomialHeapTree[T]]) *list.LinkedList[*binomialHeapTree[T]] {
  head, tail, ok := tList.HeadTail()
  if !ok {
    var newTList *list.LinkedList[*binomialHeapTree[T]]
    return newTList.ConsImpl(&bt)
  }
  if bt.rank() < head.rank() {
    return tList.ConsImpl(&bt)
  } else {
    return bt.link(*head).insTree(tail)
  }
}

func (bh *BinomialHeap[T]) InsertImpl(t T) *BinomialHeap[T] {
  tr := binomialHeapTree[T]{rk: 0, elem:t}
  newTList := tr.insTree(bh.tList)
  return &BinomialHeap[T]{tList:newTList}
}

func (bh *BinomialHeap[T]) Insert(t T) Heap[T] {
  return bh.InsertImpl(t)
}

func (lhs *BinomialHeap[T]) MergeImpl(rhs *BinomialHeap[T]) *BinomialHeap[T] {
  if (rhs.IsEmpty()) { return lhs }
  if (lhs.IsEmpty()) { return rhs }
  head1, tail1, _ := lhs.tList.HeadTail()
  head2, tail2, _ := rhs.tList.HeadTail()
  if head1.rank() < head2.rank() {
    bh := &BinomialHeap[T]{tList: tail1}
    return &BinomialHeap[T]{tList: bh.MergeImpl(rhs).tList.ConsImpl(head1)}
  } else if head2.rank() < head1.rank() {
    bh := &BinomialHeap[T]{tList: tail2}
    return &BinomialHeap[T]{tList: lhs.MergeImpl(bh).tList.ConsImpl(head2)}
  } else {
    tail1 := &BinomialHeap[T]{tList: tail1}
    tail2 := &BinomialHeap[T]{tList: tail2}
    newTList := head1.link(*head2).insTree(tail1.MergeImpl(tail2).tList)
    return &BinomialHeap[T]{tList: newTList}
  }
}

func (lhs *BinomialHeap[T]) Merge(rhs Heap[T]) Heap[T] {
  switch rhs := rhs.(type) {
  case *BinomialHeap[T]:
    return lhs.MergeImpl(rhs)
  default:
    panic("Should not happen")
  }
}

func (bh *BinomialHeap[T]) removeMinTree() (genericdata.Pair[*binomialHeapTree[T], *list.LinkedList[*binomialHeapTree[T]]], bool)  {
  head, tail, ok := bh.tList.HeadTail()
  if !ok { return genericdata.Pair[*binomialHeapTree[T], *list.LinkedList[*binomialHeapTree[T]]]{}, false }
  if tail.IsEmpty() {
    return genericdata.Pair[*binomialHeapTree[T], *list.LinkedList[*binomialHeapTree[T]]]{First: head}, true
  }
  bh2 := &(BinomialHeap[T]{tList:tail})
  HeadTail2, _ := bh2.removeMinTree()
  head2, tail2 := HeadTail2.First, HeadTail2.Second
  if head.elem < head2.elem {
    return genericdata.Pair[*binomialHeapTree[T], *list.LinkedList[*binomialHeapTree[T]]]{First:head, Second:tail}, true
  } else {
    return genericdata.Pair[*binomialHeapTree[T], *list.LinkedList[*binomialHeapTree[T]]]{First:head2, Second:tail2.ConsImpl(head)}, true
  }
}

func (bh *BinomialHeap[T]) FindMin() (T, bool) {
  t, ok := bh.removeMinTree()
  if !ok {
    var zeroT T
    return zeroT, false
  }
  return t.First.elem, true
}

func (bh *BinomialHeap[T]) DeleteMinImpl() (*BinomialHeap[T], bool) {
  t, ok := bh.removeMinTree()
  if !ok { return bh, false }
  node, ts2 := t.First, t.Second
  ts1 := node.tList
  ts3 := &BinomialHeap[T]{tList:ts1.Rev()}
  ts4 := &BinomialHeap[T]{tList:ts2}
  return ts3.MergeImpl(ts4), true
}

func (bh *BinomialHeap[T]) DeleteMin() (Heap[T], bool) {
  return bh.DeleteMinImpl()
}
