/*
观察下面代码，写出最终的打印结果
func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
*/
package main

import (
	"fmt"
)

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) //[1 2 3]
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)    //1 3
	fmt.Printf("%+v\n", s)         //[1 3]
	fmt.Printf("%+v\n", m["q1mi"]) //[1 3 3]
}
