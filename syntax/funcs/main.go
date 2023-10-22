package main

func Invoke() {
	println(Func0("广广"))
	str1, err := Func2(12, 13)
	println(str1, err)
	_, err = Func3(1, 2)
}

// Func0 单一返回值
func Func0(name string) string {
	return "hello" + name
}

// Func1 多个返回值
func Func1(a, b, c int, name string) (string, error) {
	return "hello", nil
}

// Func2 带名字的返回值
func Func2(a, b int) (str string, err error) {
	// * 指定了返回值的名字, 相当于声明了它, 可以直接用
	str = "aaa"
	// * 指定返回值名字后, 可以只写return
	return
}
func Func3(a, b int) (str string, err error) {
	res := "hello"
	// * 也可以不用定义了的返回值名字, 而返回别的变量
	return res, nil
}

// ! 左边至少有一个新变量的时候用 " := ", 修改变量的值用 " = "
// ! 堆是goroutine共享的, 栈是goroutine私有的
