package main

func main() {
	//Array(100) // ! 如果编译器没发现越界, 运行之后会panic
	//Map()
	println(Sum([]int{123, 123})) // * 但不支持int64或int32的切片, 只能重新定义函数, 除非用范型

}
