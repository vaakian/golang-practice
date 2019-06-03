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
	// 遍历map
	for key, val := range myMap {
		fmt.Printf("%s: %d\n", key, val)
	}
}