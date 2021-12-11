package main

import (
	"fmt"
)

func main() {
	array1 := [...]int{1, 1, 2, 3, 2, 3, 4}
	array2 := 0
	for _, i := range array1 {
		array2 = i ^ array2
	}
	fmt.Println(array2)
}
