//请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序
package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = [...]int{3, 7, 8, 9, 1}
	b := a[:]
	sort.Ints(b) //升序
	fmt.Println(b)
	sort.Sort(sort.Reverse(sort.IntSlice(b)))
	fmt.Println(b) //降序
}
