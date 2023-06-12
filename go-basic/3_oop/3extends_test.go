package __oop

/*
	go语言没有继承，可以通过实现组合或别名实现继承
	一般内嵌比较类似其他语言的继承
*/

// 组合
type myTreeNode struct {
	node *treeNode
}

// 别名
type queue []int

// 内嵌的做法
// 我们可以不用调用内部字段，直接调用该类型的字段或方法,其实就是语法糖，有类似于重写的功能，但是他可以通过直接调用内嵌字段来调用子结构体的方法
// 而且，内嵌不可以将子类赋值给父类
// 归根结底，内嵌就是组合的语法糖而已，可以减少很多代码，没有什么真正特殊的地方
type inMyTreeNode struct {
	*treeNode
	value int
}
