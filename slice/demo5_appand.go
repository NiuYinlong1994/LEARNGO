//Go语言的内建函数append()可以为切片动态添加元素。
//可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）
package main

import (
	"fmt"
)

func main() {
	// var s []int
	// s = append(s, 1)
	// s = append(s, 2, 3, 4, 5)
	// s2 := []int{6, 7, 8}
	// s = append(s, s2...)
	// fmt.Println(s)

	//每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
	//当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
	//“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值
	// var numSlice []int
	// for i := 0; i < 10; i++ {
	// 	numSlice = append(numSlice, i)
	// 	fmt.Printf("index:%d value:%v len:%d cap:%d ptr:%p\n", i, numSlice, len(numSlice), cap(numSlice), numSlice)
	// }

	//append()函数还支持一次性追加多个元素
	var citySlice []string
	citySlice = append(citySlice, "北京")             //一次追加一个元素
	citySlice = append(citySlice, "上海", "广州", "深圳") //一次追加多个元素
	citySlice2 := []string{"重庆", "成都"}
	citySlice = append(citySlice, citySlice2...) //追加切片
	fmt.Println(citySlice)
}
