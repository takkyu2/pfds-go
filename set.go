package main

type unbalancedSet[T ordered] struct {
  elem T
  left *unbalancedSet[T]
  right *unbalancedSet[T]
}

func (s *unbalancedSet[T]) isEmpty() bool {
  return s == nil;
}

func (s *unbalancedSet[T]) member(t T) bool {
  if s.isEmpty() { return false }

  elem, left, right := s.elem, s.left, s.right
  if t < elem {
    return left.member(t)
  } else if elem < t {
    return right.member(t)
  } else {
    return true
  }
}

func (s *unbalancedSet[T]) insertImpl(t T) *unbalancedSet[T] {
  if s.isEmpty() { return &unbalancedSet[T]{elem: t} }

  elem, left, right := s.elem, s.left, s.right
  if t < elem {
    return &unbalancedSet[T]{elem:elem, left:left.insertImpl(t), right:right}
  } else if elem < t {
    return &unbalancedSet[T]{elem:elem, left:left, right:right.insertImpl(t)}
  } else {
    return s
  }
}

func (s *unbalancedSet[T]) insert(t T) set[T] {
  return s.insertImpl(t);
}


type unbalancedFiniteMap[K ordered, V any] struct {
  elem orderedKeyValue[K, V]
  left *unbalancedFiniteMap[K, V]
  right *unbalancedFiniteMap[K, V]
}

func (s *unbalancedFiniteMap[K, V]) isEmpty() bool {
  return s == nil;
}

func (s *unbalancedFiniteMap[K, V]) lookup(key K) (V, bool) {
  if s.isEmpty() {
    var zeroV V
    return zeroV, false 
  }

  elem, left, right := s.elem, s.left, s.right
  if key < elem.key {
    return left.lookup(key)
  } else if elem.key < key {
    return right.lookup(key)
  } else {
    return elem.value, true
  }
}

func (s *unbalancedFiniteMap[K, V]) bindImpl(key K, value V) *unbalancedFiniteMap[K, V] {
  if s.isEmpty() { return &unbalancedFiniteMap[K, V]{elem: orderedKeyValue[K,V]{key:key, value:value}} }

  elem, left, right := s.elem, s.left, s.right
  if key < elem.key {
    return &unbalancedFiniteMap[K, V]{elem:elem, left:left.bindImpl(key, value), right:right}
  } else if elem.key < key {
    return &unbalancedFiniteMap[K, V]{elem:elem, left:left, right:right.bindImpl(key, value)}
  } else {
    return s
  }
}

func (s *unbalancedFiniteMap[K, V]) bind(key K, value V) finiteMap[K, V] {
  return s.bindImpl(key, value);
}
