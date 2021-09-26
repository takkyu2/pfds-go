package genericdata

type Either[L any, R any] struct {
	left    L
	right   R
	isRight bool
}

func (e Either[L, R]) GetLeft() (L, bool) {
	return e.left, !e.isRight
}

func (e Either[L, R]) GetRight() (R, bool) {
	return e.right, e.isRight
}

func (e Either[L, R]) ToLeft(left L) Either[L, R] {
	return Either[L, R]{left: left, right: e.right, isRight: false}
}

func (e Either[L, R]) ToRight(right R) Either[L, R] {
	return Either[L, R]{left: e.left, right: right, isRight: true}
}
