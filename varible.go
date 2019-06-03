package main
import (
	"fmt"
)
func main() {
	// map 类似JavaScript对象，python字典dict
	// 使用make建立map，make(map[索引类型]目标类型)
	
	myMap := make(map[string]int)
	myMap["age"] = 20
	myMap["weight"] = 70
	myMap["tall"] = 175
	// 遍历map 此时的key, val是临时的
	for key, val := range myMap {
		fmt.Printf("%s: %d\n", key, val)
	}
	// 指向map的map
	myMap2Map := make(map[string]map[string]int)
	myMap2Map["innerMap"] = myMap
	for key, val := range myMap2Map["innerMap"] {
		fmt.Printf("%s: %d\n", key, val)
	}
	// fmt.Print(myMap2Map["innerMap"])

	// 定义数组 1. 确定数组类型， 2. 不使用make就要确定大小
	// 类似C语言，不可改变数组容量
	var arr1 [5]int
	// arr1 := []int{1,2,3}
	arr1[3] = 5
	for key, val := range arr1 {
		fmt.Printf("%s: %d\n", key, val)
	}
	// make数组切片
	slice := make([]int, 0, 5)
	// 复制数组
	slice = arr1[0:5]
	fmt.Print(slice)
	// 数组长度
	fmt.Print(len(slice))


	// 字符串
	var myStr string //未初始化时为空字符串
	var myStr2 = "123"
	myStr3 := "456"
	fmt.Print(myStr, myStr2, myStr3)
}