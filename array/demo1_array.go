//数组的定义:
//var 数组变量名 [元素数量]T

package main

import (
	"fmt"
)

func main() {
	//初始化数组时可以使用初始化列表来设置数组元素的值
	// var cityarry = [3]string{"北京", "上海", "广州"}
	// fmt.Println(cityarry)

	//按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度
	// var cityarry1 = [...]int{1, 2, 3, 4}
	// fmt.Println(cityarry1)
	// fmt.Printf("type of cityarry1:%T\n", cityarry1)

	//使用指定索引值的方式来初始化数组
	var cityarry2 = [...]string{0: "北京", 2: "上海", 10: "广州"}
	fmt.Println(cityarry2)
	fmt.Printf("type of cityarry2:%T\n", cityarry2)
}
