package main

import (
	"testing"
)

func TestSet(t *testing.T) {
  var set *unbalancedSet[int]
  intSetTest(set, t)
  var finiteMap *unbalancedFiniteMap[rune, int]
  intMapTest(finiteMap, t)
}

func intSetTest(set set[int], t *testing.T) {
  // NOTE: list is NOT nil because of interface table!!
  // see The Go Programming Language, 7.5.1
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
  // NOTE: list is NOT nil because of interface table!!
  // see The Go Programming Language, 7.5.1
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
