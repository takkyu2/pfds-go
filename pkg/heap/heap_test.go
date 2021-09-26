package heap

import "testing"

func TestHeap(t *testing.T) {
  var lheap *LeftistHeap[int]
  intHeapTest(lheap, t)
  var bheap = newBinomialHeap[int]()
  intHeapTest(bheap, t)
}

func intHeapTest(heap Heap[int], t *testing.T) {
  heap2 := heap
  if !heap.IsEmpty() {
    t.Error("Error in initial condition")
  }
  for i := 0; i < 10; i++ {
    heap = heap.Insert(i)
    m, ok := heap.FindMin()
    if !ok || m != 0 {
      t.Error("findMin error")
    }
  }
  for i := 100; i > 90; i-- {
    heap2 = heap2.Insert(i)
    m, ok := heap2.FindMin()
    if !ok || m != i {
      t.Error("findMin error")
    }
  }
  heap2, ok := heap2.DeleteMin()
  if !ok {
    t.Error("delteMin error")
  }
  heap3 := heap.Merge(heap2)
  m, ok := heap3.FindMin()
  if !ok  {
    t.Error("findMin error ok")
  }
  if m != 0 {
    t.Errorf("findMin %d should be equal to %d", m, 0)
  }
  for i := 0; i < 10; i++ {
    heap3, ok = heap3.DeleteMin()
    if !ok {
      t.Error("deleteMin error")
    }
    m, ok2 := heap3.FindMin()
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
