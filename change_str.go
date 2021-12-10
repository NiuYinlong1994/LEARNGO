/*要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。*/
package main

import (
	"fmt"
)

func main() {
	s1 := "neremy"
	byteS1 := []byte(s1)               //string转换成byte
	byteS1[0] = 'J'                    //修改第一个byte的内容
	fmt.Printf("%T\n%T\n", byteS1, s1) //打印转换类型
	fmt.Println(string(byteS1))        //转换回string并打印

	s2 := "帅弟"
	runeS2 := []rune(s2)               //string转换成rune
	runeS2[1] = '哥'                    //修改第二个rune的内容
	fmt.Printf("%T\n%T\n", runeS2, s2) //打印转换类型
	fmt.Println(string(runeS2))        //转换回string并打印
}
