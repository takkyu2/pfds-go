package main

import (
	"testing"
)

func TestRAL(t *testing.T) {
  var ral binaryRAL[int]
  intRALTest(ral, t)
}

func intRALTest(ral randomAccessListInterface[int], t *testing.T) {
  for i := 0; i < 10; i++ {
    _, ok := ral.lookup(i)
    if ok {
      t.Errorf("%d should not be a member", i)
    }
    ral = ral.cons(i)
  }
  for i := 0; i < 10; i++ {
    val, ok := ral.lookup(i)
    if !ok {
      t.Errorf("%d should be a member", i)
      continue
    }
    if val != 9 - i {
      t.Errorf("%d should be %d", val, 9 - i)
    }
  }
  ral2 := ral
  for i := 0; i < 10; i++ {
    val, ok := ral.head()
    if !ok {
      t.Errorf("%d should be a member", i)
    }
    if val != 9 - i {
      t.Errorf("%d should be %d", val, 9 - i)
    }
    tail, ok := ral.tail()
    if !ok {
      t.Errorf("error in tail")
    }
    ral = tail
  }
  for i := 0; i < 10; i++ {
    ral2, ok := ral2.update(i, i*100)
    if !ok {
      t.Errorf("update failed")
    }
    val, ok := ral2.lookup(i)
    if val != i*100 {
      t.Errorf("%d should be %d", val, i*100)
    }
  }
}

