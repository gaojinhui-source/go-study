package __dependency_manage

/*
	1. gopath模式 的依赖管理方式其实就是没有依赖管理，go的env里面有一个GOPATH变量，在这个下面会有一个src，所有的依赖和project都放在src目录下
	当go运行时，找一个包的时候，会先在GOROOT下找，如果找不到，会去GOPATH/src目录下找，还找不到就会抛出错误,存在的问题
		* 多个项目依赖同一个包，但依赖版本不一致时会有问题
		* 全部项目都在一个path下，会非常的大
		* 没有版本管理，非常不智能
	2. go vendor模式 在GOPATH基础下，每个项目会有一个vendor目录来单独存放依赖库，这样可以解决GOPATH下多个项目依赖不同版本包的问题，但其他问题仍然存在

	3. go module 模式（重点）

	以下文档详解go module 各个参数 以及命令
	https://blog.csdn.net/u011069013/article/details/110114319


*/
