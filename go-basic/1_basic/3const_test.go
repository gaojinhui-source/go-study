package main

import (
	"fmt"
	"math"
	"testing"
)

const page = "page"
const (
	packageConst = "packageConst"
)

func TestConst(t *testing.T) {
	constInit()
	enums()
}

func constInit() {
	const filename string = "a.txt"
	const a, b = 1, 2 //不声明类型时，当没有类型要求的话，
	// 常量就相当于一个文本，在什么地方用到，就会将他尝试转为什么类型,这个不同于变量，变量短声明也是有类型的
	// 但是一旦给类型的花，那么常量就有类型了
	c := int(math.Sqrt(a*a + b*b)) //常量的计算相当于文本替换，在计算中既可以
	fmt.Println(c)
}

// go没有真正的枚举类型
func enums() {
	const (
		cpp = iota
		_
		java
		python
	)

	fmt.Println(cpp, java, python)

	// iota可以参与运算
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
	)
	fmt.Println(b, kb, mb, gb)

}
