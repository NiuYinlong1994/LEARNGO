//求数组[1, 3, 5, 7, 8]所有元素的和

package main

import (
	"fmt"
)

func main() {
	var elementSum int
	elementArray := [...]int{1, 3, 5, 7, 8}
	for _, element := range elementArray {
		elementSum += element
	}
	fmt.Printf("数组中所有元素的和=%d\n", elementSum)
}
