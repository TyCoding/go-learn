package main

import "fmt"

func main() {
	fmt.Printf("万能占位符 %v\n", 1)
	fmt.Printf("十进制整数 %d\n", 111)
	fmt.Printf("十六进制 %x， 二进制 %b\n", 111, 111)
	fmt.Printf("浮点数 %f\n", 1.1)
	fmt.Printf("布尔类型 %t\n", true)
	fmt.Printf("字符 %c\n", 's')
	fmt.Printf("字符串 %s\n", "s")
	fmt.Printf("带引号的字符 %q\n", "s")
	fmt.Printf("变量类型 %T\n", "s")
}
