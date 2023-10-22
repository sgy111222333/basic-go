// ! Go没有构造函数
package main

import "fmt"

func main() {
	u1 := &User{}
	println(u1)
	u1 = new(User)
	println(u1)

	u2 := User{}
	fmt.Printf("%+v\n", u2)
	u2.Name = "Jerry"
	println(u2.Name)

	// 自动初始化
	var u3 User
	fmt.Printf("%+v\n", u3)

	var u4 *User
	fmt.Printf("%+v\n", u4) // nil

	u5 := User{Name: "Jerry", Age: 18}
	fmt.Printf("%+v\n", u5)

	ChangeUser()
	Components()
}

func UseList() {
	l1 := LinkedList{}
	l1Ptr := &l1
	var l2 LinkedList = *l1Ptr
	fmt.Printf("%+v\n", l2)

	var l3 *LinkedList
	fmt.Printf("%+v\n", l3) // nil
}
