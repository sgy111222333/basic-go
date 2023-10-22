package main

func Switch(status int) string {
	switch status {
	case 0:
		return "初始化"
	case 1:
		return "执行中"
	case 2:
		return "重试"
	default:
		return "未知状态"
	}
}

// SwitchV1 switch后面啥也不写, case后面写结果为bool值的表达式, 这时要尽量保证每一个条件都是互斥的
func SwitchV1(age int) string {
	switch {
	case age >= 18:
		return "成年"
	case age >= 36:
		return "中年" // ! 如果 status=40, 他会在匹配到成年时就break, 不会执行到中年
	case age < 18:
		return "小孩"
	default:
		return "年龄异常"
	}
}
