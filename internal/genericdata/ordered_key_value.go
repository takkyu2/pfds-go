package genericdata

import "github.com/takkyu2/pfds-go/internal/constraints"

type OrderedKeyValue[K constraints.Ordered, V any] struct {
	Key   K
	Value V
}
