package main

import "fmt"

func ForArray() {
	println("遍历数组")
	arr := [3]string{"A", "B", "C"}
	for idx, val := range arr {
		println(idx, val)
	}
}

func ForSlice() {
	println("遍历切片")
	arr := [3]string{"A", "B", "C"}
	for idx, val := range arr {
		println(idx, val)
	}
}
func ForMap() {
	println("遍历map")
	m := map[string]string{
		"1": "A",
		"2": "B",
	}
	for key, value := range m {
		println(key, value)
	}
}

// * for还可以用于channel的遍历

func LoopBug() {
	users := []User{
		{name: "Tom"}, {name: "Jerry"},
	}
	m := make(map[string]*User)
	// ! for range里的u是个临时变量, 每次从users里取一个元素存入u, 但u的地址始终不变
	for _, u := range users {
		m[u.name] = &u
	}
	for key, val := range m {
		fmt.Printf("%s,%p\n", key, val)
	}
}

type User struct {
	name string
}
