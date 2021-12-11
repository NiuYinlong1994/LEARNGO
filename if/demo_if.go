package main

import (
	"fmt"
)

func main() {
	score := 59
	if score >= 90 {
		fmt.Println("成绩优秀!")
	} else if score >= 60 {
		fmt.Println("成绩合格!")
	} else {
		fmt.Println("成绩不合格!")
	}
}
