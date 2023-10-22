package main

func Func4() {
	myFunc3 := Func3 // ! 将函数赋值给变量, 只写函数名不写括号
	_, _ = myFunc3(1, 2)
}

// Func6 方法作为返回值
func Func6() func(name string) string {
	// ?只能返回匿名函数?
	return func(name string) string {
		return "hello, " + name
	}
}

func Func6Invoke() {
	fn := Func6()
	str := fn("sgy")
	println(str)
}

func Func7() {
	fn1 := func(name string) string {
		return "hello, " + name
	}("广广") // ! 这样fn1是string类型, 因为执行了匿名函数
	println(fn1)
	fn2 := func(name string) string {
		return "hello, " + name
	} // ! 这样fn2是函数类型, 并没执行匿名函数
	println(fn2)
}
