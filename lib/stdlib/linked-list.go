package stdlib

type List[T comparable] struct {
	next *List[T]
	val  T
}

func genericFunction[T comparable](a T, b T) bool {
	return a == b
}

func CreateListNode() {
	pew := genericFunction[interface{}](1, "02")

	head := List[any]{nil, 5}

	foo := List[any]{&head, "sid"}
}
