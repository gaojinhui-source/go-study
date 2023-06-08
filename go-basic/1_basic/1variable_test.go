package main

import (
	"fmt"
	"testing"
)

// 声明包变量,作用域为该包
var oo int     //通过var + 类型
var aa int = 3 //通过var + 类型 + 初始值
// 可以将多个声明放在一个var下
var (
	cc bool = true
	bb      = "kkk" //声明包变量也可以省略类型
)

// 声明变量
func variableZeroValue() {
	var a int
	var b string
	// 当仅用var声明时，变量的值为0值
	fmt.Printf("%d,%q \n", a, b) //格式化输出，%d打印int型，%s为打印字符串，但是%q可以打印引号
}

// 声明并初始化变量
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	// 当用var声明时，变量的值为0值
	fmt.Println(a, b, s)
}

// 去类型声明并初始化变量
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def" // go可以直接去掉类型声明变量，因为通过变量的值可以值就可以推断出类型
	fmt.Println(a, b, c, s)
}

// 短声明，在上面的基础上去掉var 关键字，但是注意，包变量的声明必须用var 不可以用段声明的方式
func variableShort() {
	a, b, c, s := 3, 4, true, "def" //去掉var关键字，直接用 := 来声明并初始化
	b = 5                           // 当要修改变量值的时候，不能带冒号，只有创建时可以带
	fmt.Println(a, b, c, s)

}

func TestName(t *testing.T) {
	//variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShort()
}
