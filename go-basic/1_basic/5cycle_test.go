package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func forTest() {

	// go中的for不需要括号
	// for可以省略初始条件，终止条件，递增表达式
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
}

func convertToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile() {
	file, err := os.Open("a.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	// Scan()返回的是bool，当只有一个的时候根while()没区别
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// 啥都没有就是死循环
	for {
		fmt.Println("forever")
	}
}

func TestLoop(t *testing.T) {
	fmt.Println(convertToBinary(123213))
}
