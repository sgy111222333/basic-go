package main

// 不定参数: 最后一个参数可以传入任意多的值, 但只能放在最后

func YourName(name string, alias ...string) {
	for i, s := range alias {
		println(i, s)
	}
}

func YourNameInvoke() {
	YourName("孙广宇")
	YourName("孙广宇", "sgy")
	YourName("孙广宇", "sgy", "gg", "广广")
}

// ! 不定长参数的个数上限, 等于切片的参数上限, 等于Int的最大值
// * Option模式大量应用了不定长参数
func main() {
	YourNameInvoke()
}
