package main

import "fmt"

func main() {
	fmt.Println("=== Go 语言演示程序 ===")

	fmt.Println("helloworld演示")
	demoHelloWorld()

	fmt.Println("变量演示")
	demoVariables()

	fmt.Println("var 与 := 详细对比演示")
	demoVarVsShortDeclaration()

	fmt.Println("常见错误和最佳实践")
	demoCommonMistakes()

	fmt.Println("循环演示")
	demoFor()

	fmt.Println("if-else演示")
	demoIfElse()

	fmt.Println("switch演示")
	switchDemo()

	fmt.Println("数组演示")
	demoArrays()

	fmt.Println("切片演示")
	demoSlice()

	fmt.Println("map演示")
	demoMaps()

	fmt.Println("Go Map vs Java HashMap 对比")
	demoMapVsJava()

	fmt.Println("Java 风格对比")
	demoJavaStyleComparison()

	fmt.Println("Go Map 常用模式")
	demoCommonPatterns()

	fmt.Println("Go 零值概念")
	demoZeroValues()

	fmt.Println("Go Range 演示")
	rangeDemo()

	fmt.Println("Go 函数演示")
	// 函数演示
	fmt.Println("plus(1, 2) =", plus(1, 2))
	fmt.Println("plusAndPlus(1, 2, 3) =", plusAndPlus(1, 2, 3))

	fmt.Println("变参函数演示")
	sum(1, 2, 3, 4, 5)

	fmt.Println("闭包演示")
	closureDemo()

	fmt.Println("递归演示")
	factDemo()

	fmt.Println("指针全面演示")
	demoPointer()

	fmt.Println("结构体演示")
	personDemo()

	fmt.Println("方法演示")
	methodDemo()

	fmt.Println("接口演示")
	interfaceDemo()

	fmt.Println("embeddingDemo 演示")
	embeddingDemo()
}
