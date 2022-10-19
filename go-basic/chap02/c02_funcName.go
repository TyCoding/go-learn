package main

import (
	"fmt"
	"go-basic/chap02/c02"
)

func main() {
	c02.InnerM2()
	c02.OuterM2()

	param := c02.OuterParam{
		A: 1,
		B: "1",
	}
	fmt.Println(param)
	fmt.Printf("%T", c02.B)

}
