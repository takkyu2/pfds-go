package list

import "testing"

func TestList(t *testing.T) {
  var list *LinkedList[int]
  intStackTest(list, t)
  var Stream Stream[int]
  intStreamtest(Stream, t)
}

func intStackTest(list StackWithCat[int], t *testing.T) {
  // NOTE: list is NOT nil because of interface table!!
  // see The Go Programming Language, 7.5.1
  if !list.IsEmpty() {
    t.Error("Error in initial condition")
  }
  for i := 0; i < 10; i++ {
    list = list.Cons(i);
  }
  for i := 0; i < 10; i++ {
    elem, ok := list.Head();
    if !ok || elem != 9 - i {
      t.Error("Error in head")
    }
    list, ok = list.Tail();
    if !ok {
      t.Error("Error in tail")
    }
  }
  if !list.IsEmpty() {
    t.Error("Error in last condition")
  }
  list, ok := list.Tail();
  if ok {
    t.Error("tail should fail for nil")
  }
  _, ok = list.Head();
  if ok {
    t.Error("head should fail for nil")
  }
}

func intStreamtest(list Stream[int], t *testing.T) {
  // NOTE: list is NOT nil because of interface table!!
  // see The Go Programming Language, 7.5.1
  if !list.IsEmpty() {
    t.Error("Error in initial condition")
  }
  for i := 0; i < 10; i++ {
    list = list.Cons(i);
  }
  for i := 0; i < 10; i++ {
    elem, ok := list.Head();
    if !ok || elem != 9 - i {
      t.Error("Error in head")
    }
    list, ok = list.Tail();
    if !ok {
      t.Error("Error in tail")
    }
  }
  if !list.IsEmpty() {
    t.Error("Error in last condition")
  }
  list, ok := list.Tail();
  if ok {
    t.Error("tail should fail for nil")
  }
  _, ok = list.Head();
  if ok {
    t.Error("head should fail for nil")
  }
}
