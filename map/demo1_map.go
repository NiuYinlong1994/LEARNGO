//map[KeyType]ValueType
//KeyType:表示键的类型; ValueType:表示键对应的值的类型
//map类型的变量默认初始值为nil，需要使用make()函数来分配内存
//make(map[KeyType]ValueType, [cap])

package main

import (
	"fmt"
)

func main() {
	// scoreMap := make(map[string]int, 8)
	// scoreMap["张三"] = 90
	// scoreMap["小明"] = 100
	// fmt.Println(scoreMap)
	// fmt.Println(scoreMap["小明"])
	// fmt.Printf("type of a:%T\n", scoreMap)

	// userInfo := map[string]string{
	// 	"username": "沙河小王子",
	// 	"password": "123456",
	// }
	// fmt.Println(userInfo)

	//判断某个Key是否存在？
	//value, ok := map[key]
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	//如果key存在ok为true,value为对应的值;
	//不存在ok为false,value为值类型的零值;
	value, ok := scoreMap["张三"]
	if ok {
		fmt.Println(value)
		fmt.Printf("存在key值\n")
	} else {
		fmt.Println("查无此人")
	}
}
