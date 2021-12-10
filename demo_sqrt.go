/*Calculate the oblique side length of a right triangle
比如计算直角三角形的斜边长时使用math包的Sqrt()函数，
该函数接收的是float64类型的参数，而变量a和b都是int类型的，
这个时候就需要将a和b强制类型转换为float64类型。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //math.sqrt()接收的参数是float64类型，需要强制转换成int类型
	fmt.Println(c)
}
