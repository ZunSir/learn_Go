package funcs

/**
递归
	递归使用不当会造成 stack overflow 错误。
	一个goruniter的栈大小默认是2KB，因此递归调用层数过多会导致栈溢出。
**/

// func Recursive() {
// Recursive() // Error：Stack overflow
// }

// 解决 Error：Stack overflow
func Recursive(n int) {
	if n > 10 {
		return
	}
	Recursive(n + 1)
}
