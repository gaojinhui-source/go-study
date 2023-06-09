package __collections

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 5, 6, 767, 8, 9, 90}
	// 实际只要带下标就会把数组转为切片，切片就是数组的封装，或者说视图
	fmt.Println("arr[2:6] = ", arr[2:6]) //取下标从2到6，左开右闭的切片，实际取得是2,3,4,5
	fmt.Println("arr[:6] = ", arr[:6])   //左边没有值时代表从下标0位置到下标为6，也是左开右闭，1，2，3，4，5，5
	fmt.Println("arr[2:] = ", arr[2:])   // 右边没有值代表从下标为2的一直取到最后一个元素
	fmt.Println("arr[:] = ", arr[:])     // 左右都没有，就是全部，其实就是把一个数组转为一个切片了

	//slice不再是值类型了，可以理解成slice其实就是一个操作数据的指针，所以就算将slice当作参数传递，仍然会改变slice的值，实际复制的就是操作数组的指针
	s1 := arr[2:]
	fmt.Println(s1)
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	//slice也可以基于slice再创建
	s2 := arr[:]
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	//extend
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	ss1 := arr2[2:6] // 2，3，4，5
	fmt.Println(ss1) //这个正常可以直接取出
	//fmt.Println(ss1[4])// 会报错
	ss2 := ss1[3:5] //注意，此时ss1 只有4个元素，也就是说最多取到下标为3的位置，取下标为4都会报错，但是此时取[3,5)并不会报错，结果为5,6
	//当取的时候越界的时候，会往底层的拓展数组中去取，也就是2,3,4,5,6,7中取所以此时取到的就是5，6
	//slice一般有三个元素，就是ptr(指向slice的第一个元素)，len(slice的长度),cap(slice的容量，或者说底层数组的大小)
	//只要不超过cap，实际上slice是可以向后扩展的，但不可以向前扩展，因为指针是指向第一个元素的，可以往后取，但不可以向前取
	// 取下标补可以超过len,向后扩展不可以超过cap
	fmt.Println(ss2)
	fmt.Println(len(ss1), cap(ss1))

	fmt.Println("------------------------------------------------------------------------------------------")

	appendSlice()

	fmt.Println("------------------------------------------------------------------------------------------")

	initSlice()

	fmt.Println("------------------------------------------------------------------------------------------")

	copySlice()

	fmt.Println("------------------------------------------------------------------------------------------")
	//如果删除元素，直接将这个元素跳过并且通过append赋值给一个新的slice
	s := []int{1, 2, 3, 4}
	//删除下标为2,此时实际就将2的元素删掉
	s = append(s[:2], s[3:]...) // s[3:]...代表3以后的全部元素都追加到2之后，就是下标2以后的全部前移，而且这样可以复用2的空间

}

func updateSlice(s []int) {
	s[0] = 100
}

// 当需要在slice添加元素时，除了直接修改下标时，go内置了append函数专门为切片添加元素
// s1 := append(s, 10)
// 当添加元素的数量超过cap(arr)时，slice会自动扩容，开启一个新的arr将旧arr的元素拷贝到新的数组
// 原先的数组如果没人用的话会被垃圾回收
func appendSlice() {
	arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[:len(arr)-1]
	s1 := append(s, 10)
	fmt.Println(s1)
	s2 := append(s1, 11)
	fmt.Println(s2)
	fmt.Println(arr)
	// 为什么append要一个slice接收
	//因为go是值传递，而slice其实并不是数组的指针，只是存储了一部分数组的信息，所以当append时，会将原slice拷贝一份传入，
	//然后通过这个slice的信息，去修改原数组的信息，但是并没有修改原slice的信息，所以需要一个新的slice来接收返回的信息
}

func initSlice() {
	var s []int // 这样可以声明slice，但是不可以直接访问下标，因为此时slice是nil，会panic,但是可以通过append来添加元素
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 6} // 这样是声明加初始化，可以直接访问下标，并且可以append
	fmt.Println(s1)

	s2 := make([]int, 16) // slice可以通过make命令来创建，如果只有一个参数，代表创建一个len和cap都为16的数组
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 0, 16) // 如果有两个参数，第一个代表len，第二个代表cap，以下的意思为创建一个len为0，cap为16的切片
	fmt.Println(s3, len(s3), cap(s3))

	//相对而言，使用make初始化数组较为规范，一来他全部是初始值，而且可以手动设置底层数组的长度
}

func printSlice(s []int) {
	fmt.Println(len(s), cap(s))
}

// 复制数组，注意，目标切片必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致，copy() 函数的返回值表示实际发生复制的元素个数。
func copySlice() {
	s1 := []int{1, 2, 34, 5, 6, 78, 9}
	s2 := make([]int, 32)
	s3 := make([]int, 0)
	copy(s2, s1)
	copy(s3, s1)
	fmt.Println(s1)
	fmt.Println(s2) // s2复制成功
	fmt.Println(s3) // s3复制失败
}
