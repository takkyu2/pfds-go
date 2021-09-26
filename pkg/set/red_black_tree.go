package set

import "github.com/takkyu2/pfds-go/internal/constraints"

type Color int

const (
	Black Color = iota
	Red
)

type RedBlackTree[T constraints.Ordered] struct {
	color       Color
	elem        T
	left, right *RedBlackTree[T]
}

func (rbt *RedBlackTree[T]) IsEmpty() bool {
	return rbt == nil
}

func (rbt *RedBlackTree[T]) Member(t T) bool {
	if rbt.IsEmpty() {
		return false
	}
	if t < rbt.elem {
		return rbt.left.Member(t)
	} else if rbt.elem < t {
		return rbt.right.Member(t)
	} else {
		return true
	}
}

func acceptRRXX[T constraints.Ordered](z T, left, d *RedBlackTree[T]) (*RedBlackTree[T], bool) {
	if left == nil {
		return nil, false
	}
	lcolor, ll, y, c := left.color, left.left, left.elem, left.right
	if lcolor == Black || ll == nil {
		return nil, false
	}
	llcolor, a, x, b := ll.color, ll.left, ll.elem, ll.right
	if llcolor == Black {
		return nil, false
	}
	return &RedBlackTree[T]{
		color: Red,
		left:  &RedBlackTree[T]{color: Black, left: a, elem: x, right: b},
		elem:  y,
		right: &RedBlackTree[T]{color: Black, left: c, elem: z, right: d},
	}, true
}

func acceptRXRX[T constraints.Ordered](z T, left, d *RedBlackTree[T]) (*RedBlackTree[T], bool) {
	if left == nil {
		return nil, false
	}
	lcolor, a, x, lr := left.color, left.left, left.elem, left.right
	if lcolor == Black || lr == nil {
		return nil, false
	}
	lrcolor, b, y, c := lr.color, lr.left, lr.elem, lr.right
	if lrcolor == Black {
		return nil, false
	}
	return &RedBlackTree[T]{
		color: Red,
		left:  &RedBlackTree[T]{color: Black, left: a, elem: x, right: b},
		elem:  y,
		right: &RedBlackTree[T]{color: Black, left: c, elem: z, right: d},
	}, true
}

func acceptXRRX[T constraints.Ordered](x T, a, right *RedBlackTree[T]) (*RedBlackTree[T], bool) {
	if right == nil {
		return nil, false
	}
	rcolor, rl, z, d := right.color, right.left, right.elem, right.right
	if rcolor == Black || rl == nil {
		return nil, false
	}
	rlcolor, b, y, c := rl.color, rl.left, rl.elem, rl.right
	if rlcolor == Black {
		return nil, false
	}
	return &RedBlackTree[T]{
		color: Red,
		left:  &RedBlackTree[T]{color: Black, left: a, elem: x, right: b},
		elem:  y,
		right: &RedBlackTree[T]{color: Black, left: c, elem: z, right: d},
	}, true
}

func acceptXRXR[T constraints.Ordered](x T, a, right *RedBlackTree[T]) (*RedBlackTree[T], bool) {
	if right == nil {
		return nil, false
	}
	rcolor, b, y, rr := right.color, right.left, right.elem, right.right
	if rcolor == Black || rr == nil {
		return nil, false
	}
	rrcolor, c, z, d := rr.color, rr.left, rr.elem, rr.right
	if rrcolor == Black {
		return nil, false
	}
	return &RedBlackTree[T]{
		color: Red,
		left:  &RedBlackTree[T]{color: Black, left: a, elem: x, right: b},
		elem:  y,
		right: &RedBlackTree[T]{color: Black, left: c, elem: z, right: d},
	}, true
}

func balanceRBT[T constraints.Ordered](c Color, elem T, left, right *RedBlackTree[T]) *RedBlackTree[T] {
	if c == Red {
		return &RedBlackTree[T]{color: c, elem: elem, left: left, right: right}
	}
	newTree, ok := acceptRRXX(elem, left, right)
	if ok {
		return newTree
	}
	newTree, ok = acceptRXRX(elem, left, right)
	if ok {
		return newTree
	}
	newTree, ok = acceptXRRX(elem, left, right)
	if ok {
		return newTree
	}
	newTree, ok = acceptXRXR(elem, left, right)
	if ok {
		return newTree
	}
	return &RedBlackTree[T]{color: c, elem: elem, left: left, right: right}
}

func ins[T constraints.Ordered](x T, s *RedBlackTree[T]) *RedBlackTree[T] {
	if s == nil {
		return &RedBlackTree[T]{elem: x, color: Red}
	}
	color, a, y, b := s.color, s.left, s.elem, s.right
	if x < y {
		return balanceRBT(color, y, ins(x, a), b)
	} else if y < x {
		return balanceRBT(color, y, a, ins(x, b))
	} else {
		return s
	}
}

func (rbt *RedBlackTree[T]) InsertImpl(elem T) *RedBlackTree[T] {
	inserted := ins(elem, rbt) // guaranteed to be non-nil
	a, y, b := inserted.left, inserted.elem, inserted.right
	return &RedBlackTree[T]{color: Black, left: a, elem: y, right: b}
}

func (rbt *RedBlackTree[T]) Insert(elem T) Set[T] {
	return rbt.InsertImpl(elem)
}
