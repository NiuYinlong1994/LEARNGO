package main

import "fmt"

func main() {
	s1 := "Hello,World!"
	s2 := "你好，世界！"
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(s1 + s2)                //字符串的拼接“+”
	s3 := fmt.Sprintf("%s\n%s", s1, s2) //字符串的拼接“fmt.Sprintf”
	println(s3)
}
