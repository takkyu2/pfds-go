package randomaccesslist

type RandomAccessListInterface[T any] interface {
	IsEmpty() bool
	Cons(T) RandomAccessListInterface[T]
	Head() (T, bool)
	Tail() (RandomAccessListInterface[T], bool)
	Lookup(int) (T, bool)
	Update(int, T) (RandomAccessListInterface[T], bool)
}
