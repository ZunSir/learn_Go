package funcs

import "fmt"

/*
函数式编程
*/

func Functional1() {
	fmt.Println("Hello Functional 1")
}

func Functional2(age int) {
	fmt.Println("Functional2 age ", age)
}

func UserFunctional1() {
	// 方法可以赋值给变量
	myFunc := Functional1
	// 调用
	myFunc()
	// 带参数的函数
	myFunc2 := Functional2
	myFunc2(18)
}

// 2. 局部方法
func Functional3() {
	// 新定义一个方法，赋值给了fn变量
	fn := func() string {
		return "hello"
	}
	// 直接调用
	fn()
}

// 3. 方法作为返回值
// 返回一个，返回 string 的无参方法
func Functional4() func() string {
	return func() string {
		return "hello,functional4"
	}
}

// 4. 匿名函数: 立刻发起调用
func Functional5() {
	fn := func() string {
		return "Hello,Functional5"
	}()
	fmt.Println(fn)
}

// 5. 闭包
// 方法 + 它绑定的上下文
func Functional6(name string) func() string {
	return func() string {
		return "Hello, " + name
	}
}

// 6. 不定参数：最后一个参数可以传任意多个值
func YourName(name string, alases ...string) {
	// alases 是一个切片
	if len(alases) > 0 {
		fmt.Println("Your name is ", name, " and your aliases are ", alases)
	}

}
func CallYourName() {
	YourName("Tom")
	YourName("Jerry", "Jerry1", "Jerry2")
	// 切片传入
	alases := []string{"Tom1", "Tom2"}
	YourName("Tom3", alases...)
}
