package main

import "fmt"

func main() {
	//1. for基本表达式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//2. for中初始语句忽略，分号必须写
	// i := 9
	// for ; i >= 0; i-- {
	// 	fmt.Println(i)
	// }

	// 3. for初始语句和结束语句均忽略
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//4.break
	// for i := 0; i < 4; i++ {
	// 	if i == 2 {
	// 		break
	// 	}
	// 	fmt.Println(i) //0,1
	// }

	//5.continue
	for i := 0; i < 4; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i) //0,1,3
	}
}
