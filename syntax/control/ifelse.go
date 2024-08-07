package control

import "fmt"

// if
func IfOnly(age int) {
	if age >= 18 {
		fmt.Println("已成年")
	}
}

// if-else
func IfElse(age int) {
	if age >= 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}
}

// if-else if
func IfElseIf(age int) {
	if age >= 18 {
		fmt.Println("已成年")
	} else if age >= 12 {
		fmt.Println("青年")
	} else {
		fmt.Println("小朋友")
	}
}

// 新写法
func IfNewVariable(start int, end int) string {
	if distance := end - start; distance > 100 {
		return "距离太远"
	} else if distance > 60 {
		return "距离适中"
	} else {
		return "距离近"
	}
}
