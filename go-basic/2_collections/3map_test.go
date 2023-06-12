package __collections

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	initMap()

	fmt.Println("-----------------------------------------------------------------")

	opMap()
}

func initMap() {

	//注意，map的key必须是可比较的类型，，除了slice,map,function的内建类型都可以作为,不包含不可比较类型的字段的struct也可以作为key
	m := map[string]string{
		"name":  "ccmouse",
		"name1": "ccmouse1",
		"name2": "ccmouse2",
	} // 声明并初始化一个map

	m1 := make(map[string]string) // 初始化一个空map
	m1["name"] = "ccmouse"        //map的赋值

	var m2 map[string]string
	//m2["name"] = "ccmouse" 注意,nil map不可操作

	m3 := make(map[string]string, 10) //初始化一个长度为10的map
	//m2["name"] = "ccmouse" 注意,nil map不可操作

	fmt.Println(m, m1, m2, m3)

	// 注意：使用for range遍历map时，map的顺序是无序的，因为hash本身就是无序，go为了增强这种感觉，对map的for range加了种子随机
	// 所以多次遍历同一个map，出现的顺序可能是不一样的
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func opMap() {
	// 取值以及判断值是否存在
	m := make(map[string]string) //初始化一个空map
	fmt.Println(m["zero_value"]) // go的map即使没有值，也可以去取，但是取的时候就是value类型的空值

	m["k1"] = "v1"   //添加一个元素
	v, ok := m["k1"] //可以通过这种方式取map，ok判断是否有这个值，v是取出来的值,当前案例为true
	fmt.Println(v, ok)

	_, ok = m["zero_value"] // 为false的案例
	fmt.Println(ok)

	//删除元素
	delete(m, "k1") //删除map的key
	_, ok = m["k1"] //可以通过这种方式取map，ok判断是否有这个值，v是取出来的值,当前案例为true
	fmt.Println(ok)

}
