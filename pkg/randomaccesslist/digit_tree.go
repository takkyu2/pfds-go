package randomaccesslist

import "github.com/takkyu2/pfds-go/internal/genericdata"

type digitTree[T any] struct {
	tr genericdata.Optional[tree[T]] // ZERO if nil, One otherwise
}

func (dt digitTree[T]) isZero() bool {
	return dt.tr.IsNil()
}

func (dt digitTree[T]) getOne() (tree[T], bool) {
	return dt.tr.FromOpt()
}
