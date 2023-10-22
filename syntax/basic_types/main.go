package main

import (
	"math"
	"strconv"
	"unicode/utf8"
)

func main() {
	println(math.Abs(-1.23))
	ExtremeNum()
	Strings()
	Bytes()
}

func ExtremeNum() {
	println(math.MaxInt8)
	println("float64最小正数", math.SmallestNonzeroFloat64)
}

func Strings() {
	// 双引号 建议把要打印的内容粘贴进ide, 这样可以自动转译
	// 反引号 所见即所得
	println("hello \" go")
	println(`hello,
          换行了
				gogo`)
	// 字符串拼接数字
	println("hello" + strconv.Itoa(123))

	println(len("中国人"))                    // len返回字节长度
	println(utf8.RuneCountInString("中国人")) // 返回字符个数
}

func Bytes() {
	var A byte = 'A' // byte就是uint8(0~255), C语言里的char
	println(A)
	// string可以和[]byte互转
	var s string = "hello"
	println(s)
	var bs []byte = []byte(s)
	println(bs)
	var s1 string = string(bs)
	println(s1)
}

func Bools() {
	var a bool = true
	var b bool = false
	println(a && b)
	println(a || b)
	println(!a)
	// * 	!(a&&b) == !a || !b
	// * 	!(a||b) == !a && !b
}
