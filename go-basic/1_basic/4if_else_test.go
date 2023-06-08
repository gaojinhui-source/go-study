package main

import (
	"fmt"
	"os"
	"testing"
)

func TestIfElse(t *testing.T) {
	//readFile()
	eval(1, 2, "+")
	eval(1, 2, "testfallthrough")
}

func readFile() {
	const filename = "file.txt"

	// go中的if不需要括号，且{和}不能写在同一行
	//content, err := os.ReadFile(filename)
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println(string(content))
	//}

	//go可以在if的条件中调用函数并返回结果，然后用分号分隔判断，注意，此时的file是局部变量，出了这个if就不能使用了
	if file, err := os.ReadFile(filename); err != nil {
		panic(err)
	} else {
		fmt.Println(string(file))
	}
}

// 在go中的switch会自动break，因为其实大多数时候时不需要继续向下执行的，除非使用fallthrough关键字,才会继续向下执行
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "testfallthrough":
		fmt.Println("test")
		fallthrough
	case "testfallthrough2":
		fmt.Println("fallthrough")
	default:
		// 当所有没匹配到时，会进入到default
		panic("unknown op")
	}
	return result
}
