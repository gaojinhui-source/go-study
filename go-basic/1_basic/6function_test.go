package main

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

// 单一返回值
func oneReturn() int {
	return -1
}

// go支持单函数多返回值
func multiReturn() (int, int) {
	return -1, 0
}

// 多返回有名返回值，不建议使用，可读性较差，逻辑简单时使用  ps：defer时不太一样
func multiNameReturn() (q, r int) {
	q = -1
	r = 0
	return
}

func haveDivError(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b not 0")
	}
	return a / b, nil
}

// 函数式编程,go的函数可以作为参数传入
func funcEvalProgram(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	name := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args %d,%d\n", name, a, b)
	return op(a, b)
}

func TestFunction(t *testing.T) {
	a := oneReturn()
	fmt.Println(a)
	b, c := multiReturn() //多个返回值使用逗号隔开
	fmt.Println(b, c)
	q, _ := multiNameReturn() //可以使用下划线不接收返回值
	fmt.Println(q)

	div, err := haveDivError(3, 1) //多数情况下都是一个返回值+error的形式，注意，平时在开发中，不要省略error，不处理就抛出，不可不接收
	if err != nil {
		panic(err)
	}
	fmt.Println(div)

	res := funcEvalProgram(func(i int, i2 int) int {
		return i + i2
	}, 1, 2)

	fmt.Println(res)
}
