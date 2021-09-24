package main

import "testing"

func TestHeap(t *testing.T) {
  var lheap *leftiestHeap[int]
  intHeapTest(lheap, t)
  var bheap = newBinomialHeap[int]()
  intHeapTest(bheap, t)
}

func intHeapTest(heap heap[int], t *testing.T) {
  heap2 := heap
  if !heap.isEmpty() {
    t.Error("Error in initial condition")
  }
  for i := 0; i < 10; i++ {
    heap = heap.insert(i)
    m, ok := heap.findMin()
    if !ok || m != 0 {
      t.Error("findMin error")
    }
  }
  for i := 100; i > 90; i-- {
    heap2 = heap2.insert(i)
    m, ok := heap2.findMin()
    if !ok || m != i {
      t.Error("findMin error")
    }
  }
  heap2, ok := heap2.deleteMin()
  if !ok {
    t.Error("delteMin error")
  }
  heap3 := heap.merge(heap2)
  m, ok := heap3.findMin()
  if !ok  {
    t.Error("findMin error ok")
  }
  if m != 0 {
    t.Errorf("findMin %d should be equal to %d", m, 0)
  }
  for i := 0; i < 10; i++ {
    heap3, ok = heap3.deleteMin()
    if !ok {
      t.Error("deleteMin error")
    }
    m, ok2 := heap3.findMin()
    if !ok2 {
      t.Error("findMin error")
    }
    mexp := i+1
    if i == 9 { mexp = 92 }
    if m != mexp {
      t.Errorf("findMin %d should be equal to %d", m, mexp)
    }
  }
}
