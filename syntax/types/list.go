package main

// List 是个接口
type List interface {
	Add(index int, val any)
	Append(val any)
	Delete(index int) error
}

// LinkedList 是个结构体
type LinkedList struct {
	head node
}

// ! 当一个结构体具备接口的所有方法, 它就实现了这个接口

func (l *LinkedList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(index int) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Add(index int, val any) {

}

type node struct {
	next *node
}

func UseListV1() {
	l := &LinkedList{}
	l.Add(1, 123)
	l.Add(1, "123")
}
