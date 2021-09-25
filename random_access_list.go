package main

type either[L any, R any] struct {
  left L
  right R
  isRight bool
}

func (e either[L, R]) getLeft() (L, bool) {
  return e.left, !e.isRight
}

func (e either[L, R]) getRight() (R, bool) {
  return e.right, e.isRight
}

func (e either[L, R]) toLeft(left L) either[L, R] {
  return either[L,R]{left:left, right:e.right, isRight:false}
}

func (e either[L, R]) toRight(right R) either[L, R] {
  return either[L,R]{left:e.left, right:right, isRight:true}
}

type optional[T any] struct {
  hasValue bool
  value T
}

func toOpt[T any](elem T) optional[T] {
  return optional[T] {hasValue: true, value:elem}
}

func (opt optional[T]) fromOpt() (T, bool) {
  return opt.value, opt.hasValue
}

func (opt optional[T]) isNil() bool {
  return !opt.hasValue
}

type randomAccessListInterface[T any] interface {
  isEmpty() bool
  cons(T) randomAccessListInterface[T]
  head() (T, bool)
  tail() (randomAccessListInterface[T], bool)
  lookup(int) (T, bool)
  update(int, T) (randomAccessListInterface[T], bool)
}

type tree[T any] struct {
  node either[T, node[T]]
}

func (t tree[T]) getLeaf() (T, bool) {
  return t.node.getLeft()
}

func (t tree[T]) getNode() (node[T], bool) {
  return t.node.getRight()
}

func (t tree[T]) toDigit() digitTree[T] {
  return digitTree[T]{tr: toOpt(t)}
}

type node[T any] struct {
  rk int
  left,right *tree[T]
}

func (nd node[T]) toTree() tree[T] {
  var node either[T, node[T]]
  node = node.toRight(nd)
  return tree[T]{node:node}
}

type digitTree[T any] struct {
  tr optional[tree[T]] // ZERO if nil, One otherwise
}

func (dt digitTree[T]) isZero() bool {
  return dt.tr.isNil()
}

func (dt digitTree[T]) getOne() (tree[T], bool) {
  return dt.tr.fromOpt()
}

type binaryRAL[T any] struct {
  tList *linkedList[digitTree[T]]
}

func (ral binaryRAL[T]) isEmpty() bool {
  return ral.tList == nil
}

func (ral binaryRAL[T]) listCons(dt digitTree[T]) binaryRAL[T] {
  return binaryRAL[T]{tList: ral.tList.consImpl(dt)}
}

func (ral binaryRAL[T]) listHead() (digitTree[T], bool) {
  return ral.tList.head()
}

func (ral binaryRAL[T]) listTail() (*linkedList[digitTree[T]], bool) {
  return ral.tList.tailImpl()
}

func (ral binaryRAL[T]) listHeadTail() (digitTree[T], *linkedList[digitTree[T]], bool) {
  return ral.tList.headTail()
}

func (t tree[T]) size() int {
  w, ok := t.getNode()
  if ok { return w.rk }
  return 1
}

func (t1 tree[T]) link(t2 tree[T]) tree[T] {
  return node[T]{rk: t1.size() + t2.size(), left: &t1, right: &t2}.toTree()
}

func (ral binaryRAL[T]) consTree(tr tree[T]) binaryRAL[T] {
  head, tail, ok := ral.listHeadTail()
  if !ok {
    return binaryRAL[T]{tList: ral.tList.consImpl(tr.toDigit())}
  }
  one, ok := head.getOne()
  if !ok {
    return binaryRAL[T]{tList:tail}.listCons(tr.toDigit())
  }
  return binaryRAL[T]{tList:tail}.consTree(tr.link(one)).listCons(digitTree[T]{})
}

func (ral binaryRAL[T]) unconsTree() (tree[T], binaryRAL[T], bool) {
  head, tail, ok := ral.listHeadTail()
  if !ok { return tree[T]{}, ral, false }
  if head.isZero() {
    tr, tsp, _ := binaryRAL[T]{tList:tail}.unconsTree()
    nd, _ := tr.getNode()
    return *nd.left, tsp.listCons(nd.right.toDigit()), true
  }
  one, _ := head.getOne()
  if tail.isEmpty() {
    return one, binaryRAL[T]{}, true 
  }
  return one, binaryRAL[T]{tList:tail.consImpl(digitTree[T]{})}, true
}

func (ral binaryRAL[T]) consImpl(x T) binaryRAL[T] {
  var node either[T, node[T]]
  node = node.toLeft(x)
  return ral.consTree(tree[T]{node:node})
}

func (ral binaryRAL[T]) cons(x T) randomAccessListInterface[T] {
  return ral.consImpl(x)
}

func (ral binaryRAL[T]) head() (T, bool) {
  tr, _, ok := ral.unconsTree()
  if !ok {
    var zeroT T
    return zeroT, false
  }
  lf, _ := tr.getLeaf()
  return lf, true
}

func (ral binaryRAL[T]) tailImpl() (binaryRAL[T], bool) {
  _, tsp, ok := ral.unconsTree()
  if !ok {
    return ral, false
  }
  return tsp, true
}

func (ral binaryRAL[T]) tail() (randomAccessListInterface[T], bool) {
  return ral.tailImpl()
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
      return tree[T]{node:either[T, node[T]]{}.toLeft(elem)}, true
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
    return tree[T]{node:either[T, node[T]]{}.toRight(newNd)}, true
  } else {
    newRight, ok := nd.right.updateTree(i - nd.rk / 2, elem)
    if !ok {
      return tr, false
    }
    newNd := node[T]{rk:nd.rk, left:nd.left, right:&newRight}
    return tree[T]{node:either[T, node[T]]{}.toRight(newNd)}, true
  }
}

func (ral binaryRAL[T]) lookup(i int) (T, bool) {
  head, tail, ok := ral.listHeadTail()
  if !ok { 
    var zeroT T
    return zeroT, false
  }
  one, ok := head.getOne()
  if !ok {
    return binaryRAL[T]{tList:tail}.lookup(i)
  }
  if i < one.size() {
    return one.lookupTree(i)
  } else {
    return binaryRAL[T]{tList:tail}.lookup(i - one.size())
  }
}

func (ral binaryRAL[T]) updateImpl(i int, elem T) (binaryRAL[T], bool) {
  head, tail, ok := ral.listHeadTail()
  if !ok { return ral, false }
  one, ok := head.getOne()
  if !ok {
    ral2, ok := binaryRAL[T]{tList:tail}.updateImpl(i, elem)
    return ral2.listCons(digitTree[T]{}), ok
  }
  if i < one.size() {
    newtr, ok := one.updateTree(i, elem)
    if !ok {
      return ral, false
    }
    return binaryRAL[T]{tList:tail}.listCons(digitTree[T]{tr:toOpt(newtr)}), true
  } else {
    newral, ok := binaryRAL[T]{tList:tail}.updateImpl(i-one.size(),elem)
    if !ok {
      return ral, false
    }
    return newral.listCons(digitTree[T]{tr:toOpt(one)}), true
  }
}

func (ral binaryRAL[T]) update(i int, elem T) (randomAccessListInterface[T], bool) {
  return ral.updateImpl(i, elem)
}
