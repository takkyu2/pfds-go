package randomaccesslist

import "github.com/takkyu2/pfds-go/internal/genericdata"

type tree[T any] struct {
	node genericdata.Either[T, node[T]]
}

func (t tree[T]) getLeaf() (T, bool) {
	return t.node.GetLeft()
}

func (t tree[T]) getNode() (node[T], bool) {
	return t.node.GetRight()
}

func (t tree[T]) toDigit() digitTree[T] {
	return digitTree[T]{tr: genericdata.ToOpt(t)}
}
