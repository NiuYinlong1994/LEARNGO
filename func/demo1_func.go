//定义了函数之后，我们可以通过函数名()的方式调用函数
package main

import (
	"fmt"
)

func main() {
	sayHello()
	z := intSum(10, 20)
	fmt.Println(z)
}

func sayHello() {
	fmt.Println("Hello,沙河")
}

func intSum(x int, y int) int {
	return x + y
}
