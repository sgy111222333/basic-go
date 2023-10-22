package main

// Sum 	计算一个int切片的和
func Sum(vals []int) int {
	var res int
	for _, v := range vals {
		res += v
	}
	return res
}
