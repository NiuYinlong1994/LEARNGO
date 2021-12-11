package main

import (
	"fmt"
)

func main() {
	//十进制
	var a int = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)

	//八进制
	b := 077
	fmt.Printf("%o \n", b)

	//十六进制
	var c = 0xff
	fmt.Printf("%x \n", c)
}
