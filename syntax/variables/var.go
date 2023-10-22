package main

import (
	"fmt"
	"github.com/sgy111222333/basic-go/syntax/variables/demo"
)

func main() {
	var a int = 123
	println(a)
	var b = 456
	println(b)
	var c uint = 15 //这里的uint不可以省略, 因为不加uint会被判断成int
	println(c)
	println(demo.Global)
	println(demo.External)
	//demo.External = "哈哈哈" // const常量不能修改
	//println(demo.internal)
	//println(demo.internalV1)
	f := 3.14 // 可以直接冒等, 等价于 var f float64 = 64
	fmt.Printf("%f, %T", f, f)
}
