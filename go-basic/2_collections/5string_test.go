package __collections

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestString(t *testing.T) {
	s := "yes我爱慕课网!"    //utf-8，可变长的编码格式
	fmt.Println(len(s)) //这里打印的是19，因为当前是以byte数组形式打印的，但实际上中文字符占用3个字节，1*4+5*3+1 = 19

	//可以打印出
	for _, val := range s {
		fmt.Print(val, " ")
	}

	fmt.Println()

	for i, val := range s {
		//根据下标可以看出，3结束后是6
		fmt.Printf("(%d,%d)", i, val) // (0,79)(1,65)(2,73)(3,6211)(6,7231)(9,6155)(12,8BFE)(15,7F51)(18,21)
	}
	fmt.Println()

	fmt.Println(utf8.RuneCountInString(s)) // go内置的utf8操作的库，这个可以返回一个字符串真正rune长度
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) //拿到一个rune就返回，直到最后一个byte
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	for _, val := range s {
		fmt.Printf("%c ", val)
	}

	// 如何中文字符呢?

}
