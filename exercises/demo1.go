/*编写代码统计出字符串"hello沙河小王子"中汉字的数量*/
package main

import (
	"fmt"
)

func main() {
	str := "hello沙河小王子"
	var count int
	for _, i := range str {
		if i > 256 {
			count++             //如果字符的十进制ASCII码值大于256，计数加一
			fmt.Printf("%c", i) //打印出汉字
		}
	}
	fmt.Println("\n汉字数量:", count) //打印出汉字数量
}
