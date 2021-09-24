package main

type unbalancedSet[V ordered, T orderedNode[V, interface{}]] struct {
  elem T
  left *unbalancedSet[V, T]
  right *unbalancedSet[V, T]
}

func (s *unbalancedSet[V, T]) isEmpty() bool {
  return s == nil;
}

func (s *unbalancedSet[V, T]) member(t T) bool {
  if s.isEmpty() { return false }

  elem, left, right := s.elem, s.left, s.right

  if genericLt(t,elem) {
    return left.member(t)
  } else if elem2.lt(t2) {
    return right.member(t)
  } else {
    return true
  }
}

func (s *unbalancedSet[V, T]) insertImpl(t T) *unbalancedSet[V, T] {
  if s.isEmpty() { return &unbalancedSet[V, T]{elem: t} }

  elem, left, right := s.elem, s.left, s.right
  t2, elem2 := orderedElem[T]{t}, orderedElem[T]{elem}
  if t2.lt(elem2) {
    return &unbalancedSet[V, T]{elem:elem, left:left.insertImpl(t), right:right}
  } else if elem2.lt(t2) {
    return &unbalancedSet[V, T]{elem:elem, left:left, right:right.insertImpl(t)}
  } else {
    return s
  }
}

func (s *unbalancedSet[V, T]) insert(t T) set[T] {
  return s.insertImpl(t);
}
