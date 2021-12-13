//注意： 遍历map时的元素顺序与添加键值对的顺序无关
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// scoreMap := make(map[string]int)
	// scoreMap["张三"] = 80
	// scoreMap["李四"] = 65
	// scoreMap["王五"] = 100
	// //遍历key,value
	// for key, value := range scoreMap {
	// 	fmt.Println(key, value)
	// }

	//只想遍历key
	// for key := range scoreMap {
	// 	fmt.Println(key)
	// }

	//只想遍历vlaue
	// for _, value := range scoreMap {
	// 	fmt.Println(value)
	// }

	//按照指定顺序遍历map

	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0-99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的Key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
