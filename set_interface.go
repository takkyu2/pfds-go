package main

type ordered interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// We use visitor pattern here
type orderedNode[T ordered, V any] interface {
  ltElem(orderedElem[T]) bool
  ltKeyValue(orderedKeyValue[T, V]) bool
  accept(orderedNode[T, V]) bool
}

type orderedElem[T ordered] struct {
  elem T
}

func (lhs orderedElem[T]) ltElem(rhs orderedElem[T]) bool {
  return lhs.elem < rhs.elem
}

func (lhs orderedElem[T]) ltKeyValue(rhs orderedKeyValue[T, any]) bool {
  panic("This function should not be called")
}

func (rhs orderedElem[T]) accept(lhs orderedNode[T, any]) bool {
  return lhs.ltElem(rhs);
}

type orderedKeyValue[K ordered, V any] struct {
  key K
  value V
}

func (lhs orderedKeyValue[T, any]) ltElem(rhs orderedElem[T]) bool {
  panic("This function should not be called")
}

func (lhs orderedKeyValue[T, any]) ltKeyValue(rhs orderedKeyValue[T, any]) bool {
  return lhs.key < rhs.key
}

func (rhs orderedKeyValue[T, any]) accept(lhs orderedNode[T, any]) bool {
  return lhs.ltKeyValue(rhs);
}

func genericLt[T ordered, V any](lhs orderedNode[T, V], rhs orderedNode[T, V]) bool{
  return rhs.accept(lhs);
}

type set[T ordered] interface {
  insert(T) set[T]
  member(T) bool
}

type finiteMap[K ordered, V any] interface {
  bind(K, V) finiteMap[K,V]
  lookup(K) (V, bool)
}
