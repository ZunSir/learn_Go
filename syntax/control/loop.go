package control

import "fmt"

// for 循环
func Loop1() {
	// i := 0 起始条件
	// i < 10 终止条件
	// i++ 循环体：可以理解为 i = i + 1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// for 循环的简化写法
func Loop2() {
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}
}

// for 循环- 死循环
func Loop3() {
	for {
		fmt.Println("I am a loop")
	}

	// for true {
	// 	fmt.Println("I am a loop")
	// }
}

// for range 遍历数组
func ForArr() {
	arr := [3]int{1, 2, 3}
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d \n", index, value)
	}
}

// for range 遍历切片
func ForSlice() {
	slice := []string{"h1", "h2", "h3"}
	for index, val := range slice {
		fmt.Printf("Slice Index: %d, Value: %s \n", index, val)
	}
}

// for range 遍历Map
func ForMap() {
	m := map[string]int{
		"key1": 100,
		"key2": 200,
	}
	for k, v := range m {
		fmt.Printf("Map Key: %s, Value: %d \n", k, v)
	}
}
