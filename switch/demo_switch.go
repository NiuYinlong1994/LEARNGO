//Go语言规定每个switch只能有一个default分支

package main

import (
	"fmt"
)

func main() {
	// key := 3
	// switch key {
	// case 1:
	// 	fmt.Println("足球")
	// case 2:
	// 	fmt.Println("篮球")
	// case 3:
	// 	fmt.Println("排球")
	// case 4:
	// 	fmt.Println("羽毛球")
	// case 5:
	// 	fmt.Println("棒球")
	// default:
	// 	fmt.Println("无效输入")
	// }

	//一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	// switch key2 := 10; key2 {
	// case 1, 3, 5, 7, 9:
	// 	fmt.Println("奇数")
	// case 2, 4, 6, 8, 10:
	// 	fmt.Println("偶数")
	// default:
	// fmt.Println(n)
	// }

	//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
	// key3 := -1
	// switch {
	// case key3 >= 95:
	// 	fmt.Println("超常发挥")
	// case key3 >= 85 && key3 < 95:
	// 	fmt.Println("稳定发挥")
	// case key3 >= 60 && key3 < 85:
	// 	fmt.Println("刚刚合格")
	// default:
	// 	fmt.Println("发挥失误")
	// }

	//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	key4 := 'a'
	switch {
	case key4 == 'a':
		fmt.Println("a")
		fallthrough
	case key4 == 'b':
		fmt.Println("b")
	case key4 == 'b':
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
