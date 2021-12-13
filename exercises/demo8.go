//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "how do you do"
	strSlice := strings.Split(str, " ") //切分字符串数组
	fmt.Println(strSlice)
	countMap := make(map[string]int, 3) //初始化map
	for _, value := range strSlice {
		countMap[value]++
	}
	fmt.Println(countMap)
}
