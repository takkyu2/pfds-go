package main

type unbalancedMap[K ordered, V any] struct {
  unbalancedSet[K]
}

func (m *unbalancedMap[K, V]) bind(key K, value V) {

}
