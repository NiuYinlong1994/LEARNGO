//找出数组中和为指定值的两个元素的下标，
//比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
package main

import (
	"fmt"
)

func main() {
	element1 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(element1); i++ {
		//fmt.Println(element1[i])
		for j := i + 1; j < len(element1); j++ {
			//fmt.Println(element1[j])
			if element1[i]+element1[j] == 8 {
				fmt.Printf("两个和为8的元素的下标:%d,%d\n", i, j)
				fmt.Printf("两个和为8的元素:%d+%d=%d\n", element1[i], element1[j], element1[i]+element1[j])
			}
		}
	}

}
