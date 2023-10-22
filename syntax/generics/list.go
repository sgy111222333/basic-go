package main

type ListV1[T any] interface {
	Add(index int, val T)
	Append(val T) error
	Delete(index int) error
}

type LinkedList[T any] struct {
	head *nodeV1[T]
}

func (l *LinkedList[T]) Add(index int, val T) {

}

type nodeV1[T any] struct {
	data T
}

func UseList() {
	l := &LinkedList[int]{}
	l.Add(1, 123)
}
