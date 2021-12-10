/*因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。
字符串底层是一个byte数组，所以可以和[]byte类型相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
*/
package main

import (
	"fmt"
)

func main() {
	s1 := "jeremy,你好!"
	for i := 0; i < len(s1); i++ { //byte
		fmt.Printf("%c", s1[i])
	}

	for _, j := range s1 { //rune
		fmt.Printf("%c", j)
	}
}
