package randomaccesslist

import "github.com/takkyu2/pfds-go/internal/genericdata"

type node[T any] struct {
  rk int
  left,right *tree[T]
}

func (nd node[T]) toTree() tree[T] {
  var node genericdata.Either[T, node[T]]
  node = node.ToRight(nd)
  return tree[T]{node:node}
}
