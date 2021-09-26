package set

import (
	"github.com/takkyu2/pfds-go/internal/constraints"
	"github.com/takkyu2/pfds-go/internal/genericdata"
)

// NOTE: It would be better to implement generic UnbalancedSet and adapt it to map, 
// but I can't figure out how to do that; It seems visitor pattern does not work...
type UnbalancedFiniteMap[K constraints.Ordered, V any] struct {
  elem genericdata.OrderedKeyValue[K, V]
  left *UnbalancedFiniteMap[K, V]
  right *UnbalancedFiniteMap[K, V]
}

func (s *UnbalancedFiniteMap[K, V]) IsEmpty() bool {
  return s == nil;
}

func (s *UnbalancedFiniteMap[K, V]) Lookup(key K) (V, bool) {
  if s.IsEmpty() {
    var zeroV V
    return zeroV, false 
  }

  elem, left, right := s.elem, s.left, s.right
  if key < elem.Key {
    return left.Lookup(key)
  } else if elem.Key < key {
    return right.Lookup(key)
  } else {
    return elem.Value, true
  }
}

func (s *UnbalancedFiniteMap[K, V]) BindImpl(key K, value V) *UnbalancedFiniteMap[K, V] {
  if s.IsEmpty() { return &UnbalancedFiniteMap[K, V]{elem: genericdata.OrderedKeyValue[K,V]{Key:key, Value:value}} }

  elem, left, right := s.elem, s.left, s.right
  if key < elem.Key {
    return &UnbalancedFiniteMap[K, V]{elem:elem, left:left.BindImpl(key, value), right:right}
  } else if elem.Key < key {
    return &UnbalancedFiniteMap[K, V]{elem:elem, left:left, right:right.BindImpl(key, value)}
  } else {
    return s
  }
}

func (s *UnbalancedFiniteMap[K, V]) Bind(key K, value V) FiniteMap[K, V] {
  return s.BindImpl(key, value);
}
