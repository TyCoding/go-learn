package main

import (
	"fmt"
)

// 变量基本定义
func m1() {
	// 完整定义一个变量
	var v string = "test"
	fmt.Printf("参数值：%v, 参数类型：%T\n", v, v)

	// 省略类型定义变量，go将根据初始值自动推算类型
	var v2 = "test"
	fmt.Printf("参数值：%v, 参数类型：%T\n", v2, v2)

	// 省略初始值，go将自动给一个零值，例如数值类型是0，字符串类型是""，map等对象类型是nil
	// 不过在go中即使是零值，也是可以直接参与运算的，因此不用担心空指针问题
	var v3 int
	var v4 string
	var v5 map[string]string
	fmt.Println(v3, v4, v5)

	// 多变量定义
	var a, b, c int
	fmt.Println(a, b, c)

	// 多变量定义，省略类型；Go将根据初始值自动推算变量类型
	var e, f, g = 1, "test", true
	fmt.Println(e, f, g)

	// 简短变量声明，也就是声明变量并赋初始值
	v6 := "test"
	v6 = "test2" //通过`:=`声明的变量不可再用`:=`，也就是此形式对于同一变量只能用一次
	fmt.Println(v6)
}

// 指针，拿到一个变量的指针：`&变量名`；
// 拿到一个指针变量的实际值：`*变量名`；
// 定义一个指针类型变量：`变量名 *类型`
func m2() {
	a, b := 1, 1

	c := make(map[string]string)
	c["1"] = "test1"
	c["2"] = "test2"

	swap(&a, b, c)
	fmt.Printf("修改值：a=%v, b=%v, c=%v\n", a, b, c)
}

func swap(a *int, b int, c map[string]string) {
	fmt.Printf("原始值：a=%v, b=%v, c=%v\n", *a, b, c)
	*a = 10
	b = 10
	c["3"] = "test3"
}

func main() {
	m1()
	m2()
}
