package main

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	var a int = 2
	var pa *int = &a //一个类型前加一个*号代表是这个类型的指针类型
	//在一个变量前加&,代表取这个类型的指针类型，在一个指针变量前加*,代表取这个指针的值
	*pa = 3
	fmt.Println(a)
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}

// 使用指针类型，将指针对应的值交换，但实际中这样使用很少，通常都是使用值传递，并将值返回出去
func swap(a, b *int) {
	fmt.Println(*a, *b)
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

/*
	TODO 补充引用传递和值传递概念
	ps: go语言全部都是值传递,多数时候我们可以利用指针来实现引用传递的效果，我们可以直接用指针，也可以用结构体封装指针，
	复制结构体，但实际的内存还是只有一份
	所以我们在创建结构体的时候，要思考字段时需要指针类型还是值类型
*/
