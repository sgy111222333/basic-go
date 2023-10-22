package main

func Defer() {
	defer func() {
		println("defer NO.1")
	}()
	defer func() {
		println("defer NO.2")
	}()
}

func main() {
	//Defer()
	//DeferClosure()           // 1
	//DeferClosureV1()         // 0
	//println(DeferReturn())   // 0
	//println(DeferReturnV1()) // 1
	//println(DeferReturnV2().name)
	DeferClosureLoopV1()
	DeferClosureLoopV2()
	DeferClosureLoopV3()

}

// ! defer与闭包

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	// ! 若没有传惨, 打印当defer真正被调用时i的值, i=1之后defer才被调用
	i = 1
}

func DeferClosureV1() {
	j := 0
	defer func(val int) {
		println(val)
	}(j)
	// ! 括号里的j表示将此时j的值作为参数传入, 也就是0
	j = 1
}

// DeferReturn 这样的defer不能修改a
func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

// DeferReturnV1 这样的defer能修改a, 给返回值起名,相当于提前声明了返回值
func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV2() *MyStruct {
	res := &MyStruct{
		name: "Tom",
	}
	defer func() {
		res.name = "Jack"
	}()
	return res
}

type MyStruct struct {
	name string
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
