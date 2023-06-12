package __oop

import (
	"fmt"
	"testing"
)

// go仅支持封装，不支持继承和多态
// go只有struct，没有class

type treeNode struct {
	value       int
	left, right *treeNode
}

// 注意，此时返回的是局部变量的地址，局部变量的内存是分配在栈上，但是当将一个结构体的指针当做返回值返回时，这个变量就会内存到堆上，这就是内存逃逸
func newTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}

// go的方法定义 以下为指针接收者，其实和下面写法是一样的
// 只有使用指针才可以改变结构体内容

func (s *treeNode) print() {
	print(s.value, "\n")
}

//方法其实和下面的写法是一样的，更多是用于实现接口
//func printNode(s *treeNode) {
//	print(s.value, "\n")
//}

// 如果 该方法是值接收者，那么set的实际上还是node的副本，所以需要改为指针接收者，大多数情况下也都是用指针接收者
// nil指针也可以调用方法
func (s *treeNode) setValue(val int) {
	if s == nil {
		fmt.Println("setting to nil node ,ignored")
		return
	}
	s.value = val
}

// 中序遍历
func (s *treeNode) traverse() {
	if s == nil {
		return
	}
	s.left.traverse()
	s.print()
	s.right.traverse()

}

func TestTreeNode(t *testing.T) {
	root := treeNode{
		value: 3,
	}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = newTreeNode(2)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
	n1 := newTreeNode(3)
	n1.print()
	n1.setValue(4)
	n1.print()

	var pRoot *treeNode
	pRoot.setValue(100)
}
