package main

// Closure 闭包
func Closure(name string) func() string {
	// ! 闭包 = name:上下文变量,局部变量不算 + 函数本身
	return func() string {
		return "hello, " + name
	}
}
func ClosureInvoke() {
	c := Closure("广广")
	println(c())
}

func Closure_test() {
	ClosureInvoke()
}

// ! 闭包如果使用不当会引起内存泄漏的问题, 因为一个对象被闭包引用的话, 它不会被GC
