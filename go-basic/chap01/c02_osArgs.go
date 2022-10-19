package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("索引：[%v]， 参数：[%v]\n", i, arg)
	}
}
