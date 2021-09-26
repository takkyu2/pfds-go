package set

import (
	"github.com/takkyu2/pfds-go/internal/constraints"
	"github.com/takkyu2/pfds-go/pkg/list"
)

type UnbalancedSet[T constraints.Ordered] struct {
  elem T
  left *UnbalancedSet[T]
  right *UnbalancedSet[T]
}

func (s *UnbalancedSet[T]) IsEmpty() bool {
  return s == nil;
}

func (s *UnbalancedSet[T]) Member(t T) bool {
  if s.IsEmpty() { return false }

  elem, left, right := s.elem, s.left, s.right
  if t < elem {
    return left.Member(t)
  } else if elem < t {
    return right.Member(t)
  } else {
    return true
  }
}

func (s *UnbalancedSet[T]) InsertImpl(t T) *UnbalancedSet[T] {
  if s.IsEmpty() { return &UnbalancedSet[T]{elem: t} }

  elem, left, right := s.elem, s.left, s.right
  if t < elem {
    return &UnbalancedSet[T]{elem:elem, left:left.InsertImpl(t), right:right}
  } else if elem < t {
    return &UnbalancedSet[T]{elem:elem, left:left, right:right.InsertImpl(t)}
  } else {
    return s
  }
}

func (s *UnbalancedSet[T]) Insert(t T) Set[T] {
  return s.InsertImpl(t);
}

func toListHelper[T constraints.Ordered](list *list.LinkedList[T], set *UnbalancedSet[T]) *list.LinkedList[T] {
  if set.IsEmpty() { return list }
  list = toListHelper(list, set.left)
  list = list.ConsImpl(set.elem)
  list = toListHelper(list, set.right)
  return list
}

func (s *UnbalancedSet[T]) ToList() *list.LinkedList[T] {
  var list *list.LinkedList[T]
  return toListHelper(list, s).Rev()
}
