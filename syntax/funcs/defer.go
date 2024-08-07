package funcs

import "fmt"

// defer: 延迟调用
// 后进先出，先进后出
func Defer() {
	defer func() {
		fmt.Println("第一个 defer")
	}()

	defer func() {
		fmt.Println("第二个 defer")
	}()

	defer func() {
		fmt.Println("第三个 defer")
	}()
}

func DeferClosure() {
	i := 0

	defer func() {
		fmt.Println("i 的值为:", i)
	}()

	i = 1
	// 最后输出结果为：1
}

func DeferClosure2() {
	i := 0
	defer func(val int) {
		fmt.Println("i 的值为:", val)
	}(i)
	i = 1
	// 最后输出结果为：0
}

// 练习题1
func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

// 练习题2
func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Println(val)
		}(i)
	}
}

// 练习题3
func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Println(j)
		}()
	}
}
