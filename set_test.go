package main

import (
	"testing"
)

func TestSet(t *testing.T) {
  var set *unbalancedSet[int]
  intSetTest(set, t)
  toListTest(set, t)
  var finiteMap *unbalancedFiniteMap[rune, int]
  intMapTest(finiteMap, t)
  var rbt *redBlackTree[int]
  intSetTest(rbt, t)
  testRBTInvariant(t)
}

func intSetTest(set set[int], t *testing.T) {
  for i := 0; i < 10; i++ {
    if set.member(i) {
      t.Errorf("%d should not be a member", i)
    }
    set = set.insert(i)
    if !set.member(i) {
      t.Errorf("%d should be a member", i)
    }
  }
}

func intMapTest(finiteMap finiteMap[rune, int], t *testing.T) {
  for i, ch := range "こんにちは！" {
    _, ok := finiteMap.lookup(ch)
    if ok {
      t.Errorf("%c should not be a member", ch)
    }
    finiteMap = finiteMap.bind(ch, i)
    v, ok := finiteMap.lookup(ch)
    if !ok || v != i {
      t.Errorf("%c should be a member", ch)
    }
  }
}

func toListTest(set *unbalancedSet[int], t *testing.T) {
  for i := 99; i >= 0; i-- {
    set = set.insertImpl(i)
    if !set.member(i) {
      t.Errorf("%d should be a member", i)
    }
  }
  list := set.toList()
  for i := 0; i < 100; i++ {
    head, tail, ok := list.headTail()
    if !ok {
      t.Errorf("Error in tolist")
    }
    if head != i {
      t.Errorf("Error, head %d should be %d", head, i)
    }
    list = tail
  }
}

func testRBTInvariant(t *testing.T) {
  var set *redBlackTree[int]
  for i := 99999; i >= 0; i-- {
    set = set.insertImpl(i)
  }
  verifyRBT(set, 100000, t)
}

func checkChild[T ordered](rbt *redBlackTree[T]) bool {
  if rbt.isEmpty() { return true }
  color, left, right := rbt.color, rbt.left, rbt.right
  if color == Black { return true }
  if left.isEmpty() && right.isEmpty() { return true }
  if left.isEmpty() { return right.color == Black }
  if right.isEmpty() { return left.color == Black }
  return left.color == Black && right.color == Black
}

func checkNode[T ordered](rbt *redBlackTree[T], blackNum int, ch chan<-int, t *testing.T) {
  if rbt.isEmpty() { return }
  if rbt.color == Black {
    blackNum++
  }
  if rbt.left.isEmpty() && rbt.right.isEmpty() {
    ch <- blackNum
    ch <- -1
    return
  }
  ch <- -1
  ok := checkChild(rbt)
  if !ok {
    t.Errorf("Error in RBT, red should have only black children")
  }
  go checkNode(rbt.left, blackNum, ch, t)
  checkNode(rbt.right, blackNum, ch, t)
}

// not really effcient, since the main goroutine must do loop over remSignals
// times; this is not needed if we know the number of leafs in rbt in advance.
func verifyRBT[T ordered](rbt *redBlackTree[T], rbtSz int, t *testing.T) {
  ch := make(chan int)

  remSignals := rbtSz
  var set *unbalancedSet[int]
  go checkNode(rbt, 0, ch, t)
  for remSignals > 0 {
    val := <- ch
    if val == -1 {
      remSignals--
      continue
    }
    set = set.insertImpl(val)
  }
  list := set.toList()
  head, tail, ok := list.headTail()
  if ok && !tail.isEmpty() {
    t.Errorf("Error in RBT, Every path to leaf should have the same number of blacks")
    return
  }
  if !ok && rbtSz != 0 {
    t.Errorf("Error in RBT, Something went wrong with insert or toList")
    return
  }
  t.Logf("The black node rank is %d for the rbt with size %d", head, rbtSz)
}
