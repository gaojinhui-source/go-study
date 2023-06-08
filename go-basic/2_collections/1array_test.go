package __collections

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var arr1 [5]int                                 // 第一种声明方式，全部是0值
	arr2 := [3]int{1, 3, 5}                         //第二种初始化，可以附带初始值
	arr3 := [...]int{231, 321, 3213, 123, 12312, 3} //使用3个点，可以创建不定长的数组，在运行时根据元素数量，但记住，数组是定长的
	var grid [4][5]int                              //二维数组 四行五列
	fmt.Println(arr1, arr2, arr3, grid)

	//遍历数组 方式1,普通for循环
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	//方式2 for range方式（推荐）
	for index, value := range arr2 {
		fmt.Println(index, value)
	}

	// 以下的话，仅遍历index
	//for i := range arr2 {
	//
	//}

	printArray(arr1)
	fmt.Println(arr1) //因为数组是值传递，所以即便在其他函数改变某个元素的值，实际改变的是拷贝的数组的值
}

// 1.数组是值传递
// 2.在go中，[5]int,[3]int是两个类型
func printArray(arr [5]int) {
	for index, value := range arr {
		fmt.Println(index, value)
	}
	arr[0] = 100 //在这里改变值不会影响到原来数组
}
