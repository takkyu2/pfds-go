package main

type ordered interface {
  ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

type orderedKeyValue[K ordered, V any] struct {
  key K
  value V
}

type set[T ordered] interface {
  insert(T) set[T]
  member(T) bool
}

type finiteMap[K ordered, V any] interface {
  bind(K, V) finiteMap[K,V]
  lookup(K) (V, bool)
}
