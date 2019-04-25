package main

import (
	"fmt"
	"strings"
)

func add(a, b int) (num int) {
	num = a + b
	return num
}

func sub(a, b int) (num int) {
	num = a - b
	return num
}

func main() {
	fmt.Println("请输入运算式:")
	var cal string
	fmt.Scan(&cal)
	s := strings.Split(cal, "")
	fmt.Println(s)
}
