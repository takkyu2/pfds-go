package set

import "github.com/takkyu2/pfds-go/internal/constraints"

type Set[T constraints.Ordered] interface {
	Insert(T) Set[T]
	Member(T) bool
}

type FiniteMap[K constraints.Ordered, V any] interface {
	Bind(K, V) FiniteMap[K, V]
	Lookup(K) (V, bool)
}
