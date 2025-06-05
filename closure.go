package main
import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		fmt.Println("Current value of i:", i)
		return i
	}
}

func closureDemo(){
	nextInt := intSeq()

	fmt.Println(nextInt()) // 输出 1
	fmt.Println(nextInt()) // 输出 2
	fmt.Println(nextInt()) // 输出 3

	newInt := intSeq() // 创建一个新的闭包
	fmt.Println(newInt()) // 输出 1 (新的闭包从 0 开始)
}