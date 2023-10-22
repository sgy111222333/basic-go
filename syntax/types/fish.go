package main

type Fish struct {
}

func (f Fish) Swim() {
}

type FakeFish Fish // 相当于一个新类型
type Yu = Fish     // 相当于起别名

func (f FakeFish) FakeSwim() {
}

func UseFish() {
	f1 := Fish{}
	f1.Swim()

	f2 := FakeFish{}
	//	f2.FakeSwim()// ! 不可以这么用
	f2.FakeSwim()

	f3 := Fish(f2) // ! 类型转换
	f3.Swim()

	yu := Yu{}
	yu.Swim()
}
