//使用delete()内建函数从map中删除一组键值对
//delete(map, key)
//map:表示要删除键值对的map;  key:表示要删除的键值对的键
package main

import (
	"fmt"
)

func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 65
	scoreMap["王五"] = 100
	delete(scoreMap, "李四") //将李四：65从map中删除
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
}
