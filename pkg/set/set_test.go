package set

import (
	"testing"
	"github.com/takkyu2/pfds-go/internal/constraints"
)

func TestSet(t *testing.T) {
  var set *UnbalancedSet[int]
  intSetTest(set, t)
  toListTest(set, t)
  var finiteMap *UnbalancedFiniteMap[rune, int]
  intMapTest(finiteMap, t)
  var rbt *RedBlackTree[int]
  intSetTest(rbt, t)
  testRBTInvariant(t)
}

func intSetTest(set Set[int], t *testing.T) {
  for i := 0; i < 10; i++ {
    if set.Member(i) {
      t.Errorf("%d should not be a member", i)
    }
    set = set.Insert(i)
    if !set.Member(i) {
      t.Errorf("%d should be a member", i)
    }
  }
}

func intMapTest(finiteMap FiniteMap[rune, int], t *testing.T) {
  for i, ch := range "こんにちは！" {
    _, ok := finiteMap.Lookup(ch)
    if ok {
      t.Errorf("%c should not be a member", ch)
    }
    finiteMap = finiteMap.Bind(ch, i)
    v, ok := finiteMap.Lookup(ch)
    if !ok || v != i {
      t.Errorf("%c should be a member", ch)
    }
  }
}

func toListTest(set *UnbalancedSet[int], t *testing.T) {
  for i := 99; i >= 0; i-- {
    set = set.InsertImpl(i)
    if !set.Member(i) {
      t.Errorf("%d should be a member", i)
    }
  }
  list := set.ToList()
  for i := 0; i < 100; i++ {
    head, tail, ok := list.HeadTail()
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
  var set *RedBlackTree[int]
  for i := 99999; i >= 0; i-- {
    set = set.InsertImpl(i)
  }
  verifyRBT(set, 100000, t)
}

func checkChild[T constraints.Ordered](rbt *RedBlackTree[T]) bool {
  if rbt.IsEmpty() { return true }
  color, left, right := rbt.color, rbt.left, rbt.right
  if color == Black { return true }
  if left.IsEmpty() && right.IsEmpty() { return true }
  if left.IsEmpty() { return right.color == Black }
  if right.IsEmpty() { return left.color == Black }
  return left.color == Black && right.color == Black
}

func checkNode[T constraints.Ordered](rbt *RedBlackTree[T], blackNum int, ch chan<-int, t *testing.T) {
  if rbt.IsEmpty() { return }
  if rbt.color == Black {
    blackNum++
  }
  if rbt.left.IsEmpty() && rbt.right.IsEmpty() {
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
func verifyRBT[T constraints.Ordered](rbt *RedBlackTree[T], rbtSz int, t *testing.T) {
  ch := make(chan int)

  remSignals := rbtSz
  var set *UnbalancedSet[int]
  go checkNode(rbt, 0, ch, t)
  for remSignals > 0 {
    val := <- ch
    if val == -1 {
      remSignals--
      continue
    }
    set = set.InsertImpl(val)
  }
  list := set.ToList()
  head, tail, ok := list.HeadTail()
  if ok && !tail.IsEmpty() {
    t.Errorf("Error in RBT, Every path to leaf should have the same number of blacks")
    return
  }
  if !ok && rbtSz != 0 {
    t.Errorf("Error in RBT, Something went wrong with insert or toList")
    return
  }
  t.Logf("The black node rank is %d for the rbt with size %d", head, rbtSz)
}
