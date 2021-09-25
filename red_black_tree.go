package main

type Color int

const (
  Black Color = iota
  Red
)

type redBlackTree[T ordered] struct {
  color Color
  elem T
  left, right *redBlackTree[T]
}

func (rbt *redBlackTree[T]) isEmpty() bool {
  return rbt == nil
}

func (rbt *redBlackTree[T]) member(t T) bool {
  if (rbt.isEmpty()) { return false }
  if (t < rbt.elem) {
    return rbt.left.member(t)
  } else if (rbt.elem < t) {
    return rbt.right.member(t)
  } else {
    return true
  }
}

func acceptRRXX[T ordered](z T, left, d *redBlackTree[T]) (*redBlackTree[T], bool) {
  if left == nil { return nil, false }
  lcolor, ll, y, c := left.color, left.left, left.elem, left.right
  if lcolor == Black || ll == nil { return nil, false }
  llcolor, a, x, b := ll.color, ll.left, ll.elem, ll.right
  if llcolor == Black { return nil, false }
  return &redBlackTree[T]{
    color:Red, 
    left:&redBlackTree[T]{color:Black, left:a, elem:x, right:b}, 
    elem:y, 
    right:&redBlackTree[T]{color:Black, left:c, elem:z, right:d},
  }, true
}

func acceptRXRX[T ordered](z T, left, d *redBlackTree[T]) (*redBlackTree[T], bool) {
  if left == nil { return nil, false }
  lcolor, a, x, lr := left.color, left.left, left.elem, left.right
  if lcolor == Black || lr == nil { return nil, false }
  lrcolor, b, y, c := lr.color, lr.left, lr.elem, lr.right
  if lrcolor == Black { return nil, false }
  return &redBlackTree[T]{
    color:Red, 
    left:&redBlackTree[T]{color:Black, left:a, elem:x, right:b}, 
    elem:y, 
    right:&redBlackTree[T]{color:Black, left:c, elem:z, right:d},
  }, true
}

func acceptXRRX[T ordered](x T, a, right *redBlackTree[T]) (*redBlackTree[T], bool) {
  if right == nil { return nil, false }
  rcolor, rl, z, d := right.color, right.left, right.elem, right.right
  if rcolor == Black || rl == nil { return nil, false }
  rlcolor, b, y, c := rl.color, rl.left, rl.elem, rl.right
  if rlcolor == Black { return nil, false }
  return &redBlackTree[T]{
    color:Red, 
    left:&redBlackTree[T]{color:Black, left:a, elem:x, right:b}, 
    elem:y, 
    right:&redBlackTree[T]{color:Black, left:c, elem:z, right:d},
  }, true
}

func acceptXRXR[T ordered](x T, a, right *redBlackTree[T]) (*redBlackTree[T], bool) {
  if right == nil { return nil, false }
  rcolor, b, y, rr := right.color, right.left, right.elem, right.right
  if rcolor == Black || rr == nil { return nil, false }
  rrcolor, c, z, d := rr.color, rr.left, rr.elem, rr.right
  if rrcolor == Black { return nil, false }
  return &redBlackTree[T]{
    color:Red, 
    left:&redBlackTree[T]{color:Black, left:a, elem:x, right:b}, 
    elem:y, 
    right:&redBlackTree[T]{color:Black, left:c, elem:z, right:d},
  }, true
}

func balanceRBT[T ordered](c Color, elem T, left, right *redBlackTree[T]) *redBlackTree[T] {
  if c == Red { return &redBlackTree[T]{color:c, elem:elem, left:left, right:right} }
  newTree, ok := acceptRRXX(elem, left, right)
  if ok { return newTree }
  newTree, ok = acceptRXRX(elem, left, right)
  if ok { return newTree }
  newTree, ok = acceptXRRX(elem, left, right)
  if ok { return newTree }
  newTree, ok = acceptXRXR(elem, left, right)
  if ok { return newTree }
  return &redBlackTree[T]{color:c, elem:elem, left:left, right:right}
}

func ins[T ordered](x T, s *redBlackTree[T]) *redBlackTree[T] {
  if s == nil { return &redBlackTree[T]{elem:x, color:Red} }
  color, a, y, b := s.color, s.left, s.elem, s.right
  if x < y { 
    return balanceRBT(color, y, ins(x, a), b) 
  } else if y < x {
    return balanceRBT(color, y, a, ins(x, b)) 
  } else {
    return s
  }
}

func (rbt *redBlackTree[T]) insertImpl(elem T) *redBlackTree[T] {
  inserted := ins(elem, rbt) // guaranteed to be non-nil
  a, y, b:= inserted.left, inserted.elem, inserted.right
  return &redBlackTree[T]{color:Black, left:a, elem: y, right: b}
}

func (rbt *redBlackTree[T]) insert(elem T) set[T] {
  return rbt.insertImpl(elem)
}
