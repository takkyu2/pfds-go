package main

import "testing"

func TestStack(t *testing.T) {
  var list *linkedList[int]
  intStackTest(list, t)
  var stream stream[int]
  intStreamtest(stream, t)
}

func intStackTest(list stackWithCat[int], t *testing.T) {
  // NOTE: list is NOT nil because of interface table!!
  // see The Go Programming Language, 7.5.1
  if !list.isEmpty() {
    t.Error("Error in initial condition")
  }
  for i := 0; i < 10; i++ {
    list = list.cons(i);
  }
  for i := 0; i < 10; i++ {
    elem, ok := list.head();
    if !ok || elem != 9 - i {
      t.Error("Error in head")
    }
    list, ok = list.tail();
    if !ok {
      t.Error("Error in tail")
    }
  }
  if !list.isEmpty() {
    t.Error("Error in last condition")
  }
  list, ok := list.tail();
  if ok {
    t.Error("tail should fail for nil")
  }
  _, ok = list.head();
  if ok {
    t.Error("head should fail for nil")
  }
}

func intStreamtest(list stream[int], t *testing.T) {
  // NOTE: list is NOT nil because of interface table!!
  // see The Go Programming Language, 7.5.1
  if !list.isEmpty() {
    t.Error("Error in initial condition")
  }
  for i := 0; i < 10; i++ {
    list = list.cons(i);
  }
  for i := 0; i < 10; i++ {
    elem, ok := list.head();
    if !ok || elem != 9 - i {
      t.Error("Error in head")
    }
    list, ok = list.tail();
    if !ok {
      t.Error("Error in tail")
    }
  }
  if !list.isEmpty() {
    t.Error("Error in last condition")
  }
  list, ok := list.tail();
  if ok {
    t.Error("tail should fail for nil")
  }
  _, ok = list.head();
  if ok {
    t.Error("head should fail for nil")
  }
}
