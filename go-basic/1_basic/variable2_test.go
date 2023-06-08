package main

// go的内建变量
/*
	1. bool string
	2. (u)int (u)int8 (u)int16 (u)int32 (u)int64 uintptr  int根据操作系统，如果操作系统是32位，int就是int32
	3. byte rune     ps:rune是int32的别名，因为go的默认编码是utf-8，utf8的是变长编码，一般是3个字节，int32是四个字节，rune就是其他语言的char
	4. float32 float64 complex64 complex128

	注意:在go中所有类型转换都是强制转换
*/
