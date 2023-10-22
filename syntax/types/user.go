package main

import "fmt"

func ChangeUser() {
	u1 := User{Name: "Tom", Age: 18}
	fmt.Printf("%+v\n", u1)
	fmt.Printf("u1地址: %p\n", &u1)

	u1.ChangeName("Jerry") // 此时发生了复制, 值传递
	u1.ChangeAge(35)
	fmt.Printf("%+v\n", u1)

	u2 := &User{Name: "小明", Age: 18}
	fmt.Printf("%+v\n", u2)
	fmt.Printf("u2地址: %p\n", &u2)

	u2.ChangeName("Jerry")
	u2.ChangeAge(35)
	fmt.Printf("%+v\n", u2)

	// * 猜测结果: Tom 35 小明 35
}

type User struct {
	Age      int
	Name     string
	NickName string
}

func (u User) ChangeName(name string) {
	u.NickName = name
}

func (u *User) ChangeAge(age int) {
	u.Age = age
}
