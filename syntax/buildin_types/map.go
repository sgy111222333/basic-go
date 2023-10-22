package main

import "fmt"

func Map() {
	// 初始化
	m1 := map[string]string{
		"key1": "value1",
	}
	fmt.Println(m1)
	m2 := make(map[string]string, 4)
	m2["key2"] = "value2"

	// 读取
	val1, ok := m1["key1"]
	if ok {
		println(val1)
	}
	for k := range m1 {
		println(k, m1[k])
	}
}
