package randomaccesslist

import (
	"github.com/takkyu2/pfds-go/internal/genericdata"
	"github.com/takkyu2/pfds-go/pkg/list"
)

type BinaryRAL[T any] struct {
  tList *list.LinkedList[digitTree[T]]
}

func (ral BinaryRAL[T]) IsEmpty() bool {
  return ral.tList == nil
}

func (ral BinaryRAL[T]) listCons(dt digitTree[T]) BinaryRAL[T] {
  return BinaryRAL[T]{tList: ral.tList.ConsImpl(dt)}
}

func (ral BinaryRAL[T]) listHead() (digitTree[T], bool) {
  return ral.tList.Head()
}

func (ral BinaryRAL[T]) listTail() (*list.LinkedList[digitTree[T]], bool) {
  return ral.tList.TailImpl()
}

func (ral BinaryRAL[T]) listHeadTail() (digitTree[T], *list.LinkedList[digitTree[T]], bool) {
  return ral.tList.HeadTail()
}

func (t tree[T]) size() int {
  w, ok := t.getNode()
  if ok { return w.rk }
  return 1
}

func (t1 tree[T]) link(t2 tree[T]) tree[T] {
  return node[T]{rk: t1.size() + t2.size(), left: &t1, right: &t2}.toTree()
}

func (ral BinaryRAL[T]) consTree(tr tree[T]) BinaryRAL[T] {
  head, tail, ok := ral.listHeadTail()
  if !ok {
    return BinaryRAL[T]{tList: ral.tList.ConsImpl(tr.toDigit())}
  }
  one, ok := head.getOne()
  if !ok {
    return BinaryRAL[T]{tList:tail}.listCons(tr.toDigit())
  }
  return BinaryRAL[T]{tList:tail}.consTree(tr.link(one)).listCons(digitTree[T]{})
}

func (ral BinaryRAL[T]) unconsTree() (tree[T], BinaryRAL[T], bool) {
  head, tail, ok := ral.listHeadTail()
  if !ok { return tree[T]{}, ral, false }
  if head.isZero() {
    tr, tsp, _ := BinaryRAL[T]{tList:tail}.unconsTree()
    nd, _ := tr.getNode()
    return *nd.left, tsp.listCons(nd.right.toDigit()), true
  }
  one, _ := head.getOne()
  if tail.IsEmpty() {
    return one, BinaryRAL[T]{}, true 
  }
  return one, BinaryRAL[T]{tList:tail.ConsImpl(digitTree[T]{})}, true
}

func (ral BinaryRAL[T]) ConsImpl(x T) BinaryRAL[T] {
  var node genericdata.Either[T, node[T]]
  node = node.ToLeft(x)
  return ral.consTree(tree[T]{node:node})
}

func (ral BinaryRAL[T]) Cons(x T) RandomAccessListInterface[T] {
  return ral.ConsImpl(x)
}

func (ral BinaryRAL[T]) Head() (T, bool) {
  tr, _, ok := ral.unconsTree()
  if !ok {
    var zeroT T
    return zeroT, false
  }
  lf, _ := tr.getLeaf()
  return lf, true
}

func (ral BinaryRAL[T]) TailImpl() (BinaryRAL[T], bool) {
  _, tsp, ok := ral.unconsTree()
  if !ok {
    return ral, false
  }
  return tsp, true
}

func (ral BinaryRAL[T]) Tail() (RandomAccessListInterface[T], bool) {
  return ral.TailImpl()
}

func (tr tree[T]) lookupTree(i int) (T, bool) {
  node, ok := tr.getNode()
  if !ok {
    leaf, _ := tr.getLeaf()
    if i == 0 {
      return leaf, true
    } else {
      return leaf, false
    }
  }
  if i < node.rk / 2 {
    return node.left.lookupTree(i)
  } else {
    return node.right.lookupTree(i - node.rk / 2)
  }
}

func (tr tree[T]) updateTree(i int, elem T) (tree[T], bool) {
  nd, ok := tr.getNode()
  if !ok {
    if i == 0 {
      return tree[T]{node:genericdata.Either[T, node[T]]{}.ToLeft(elem)}, true
    } else {
      return tr, false
    }
  }
  if i < nd.rk / 2 {
    newLeft, ok := nd.left.updateTree(i, elem)
    if !ok {
      return tr, false
    }
    newNd := node[T]{rk:nd.rk, left:&newLeft, right:nd.right}
    return tree[T]{node:genericdata.Either[T, node[T]]{}.ToRight(newNd)}, true
  } else {
    newRight, ok := nd.right.updateTree(i - nd.rk / 2, elem)
    if !ok {
      return tr, false
    }
    newNd := node[T]{rk:nd.rk, left:nd.left, right:&newRight}
    return tree[T]{node:genericdata.Either[T, node[T]]{}.ToRight(newNd)}, true
  }
}

func (ral BinaryRAL[T]) Lookup(i int) (T, bool) {
  head, tail, ok := ral.listHeadTail()
  if !ok { 
    var zeroT T
    return zeroT, false
  }
  one, ok := head.getOne()
  if !ok {
    return BinaryRAL[T]{tList:tail}.Lookup(i)
  }
  if i < one.size() {
    return one.lookupTree(i)
  } else {
    return BinaryRAL[T]{tList:tail}.Lookup(i - one.size())
  }
}

func (ral BinaryRAL[T]) UpdateImpl(i int, elem T) (BinaryRAL[T], bool) {
  head, tail, ok := ral.listHeadTail()
  if !ok { return ral, false }
  one, ok := head.getOne()
  if !ok {
    ral2, ok := BinaryRAL[T]{tList:tail}.UpdateImpl(i, elem)
    return ral2.listCons(digitTree[T]{}), ok
  }
  if i < one.size() {
    newtr, ok := one.updateTree(i, elem)
    if !ok {
      return ral, false
    }
    return BinaryRAL[T]{tList:tail}.listCons(digitTree[T]{tr:genericdata.ToOpt(newtr)}), true
  } else {
    newral, ok := BinaryRAL[T]{tList:tail}.UpdateImpl(i-one.size(),elem)
    if !ok {
      return ral, false
    }
    return newral.listCons(digitTree[T]{tr:genericdata.ToOpt(one)}), true
  }
}

func (ral BinaryRAL[T]) Update(i int, elem T) (RandomAccessListInterface[T], bool) {
  return ral.UpdateImpl(i, elem)
}
