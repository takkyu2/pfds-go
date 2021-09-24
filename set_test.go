package main

import (
	"testing"
)

func TestSet(t *testing.T) {
  var set *unbalancedSet[int]
  intSetTest(set, t)
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
