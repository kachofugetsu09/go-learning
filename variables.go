package main

import "fmt"

// 全局变量必须使用 var 声明
var globalVar = "这是全局变量"
var (
	// 批量声明全局变量
	maxSize     = 100
	appName     = "Go学习应用"
	isDebugMode = false
)

// demoVariables 演示 Go 语言中的变量声明和使用
// 展示不同类型的变量声明方式
func demoVariables() {
	fmt.Println("=== 变量声明演示 ===")

	// 演示全局变量的使用
	fmt.Println("全局变量 globalVar:", globalVar)
	fmt.Printf("应用配置 - 名称: %s, 最大尺寸: %d, 调试模式: %t\n", appName, maxSize, isDebugMode)

	// 1. 使用 var 声明并初始化变量
	var a = "initial value" // Go 自动推断 a 为 string 类型
	fmt.Println("变量 a:", a)

	// 1.1 对比：显式指定类型
	var aExplicit string = "initial value" // 明确指定为 string 类型
	fmt.Println("显式声明的变量 aExplicit:", aExplicit)

	// 1.2 查看变量类型的方法
	fmt.Printf("变量 a 的类型: %T\n", a)
	fmt.Printf("变量 aExplicit 的类型: %T\n", aExplicit)

	// 2. 同时声明多个变量
	var b, c int = 1, 2
	fmt.Println("变量 b, c:", b, c)

	// 3. 声明布尔类型变量
	var d = true
	fmt.Println("变量 d:", d)

	// 4. 声明但不初始化，使用默认值
	var e int
	fmt.Println("变量 e 的默认值:", e) // 输出变量 e 的默认值，0

	// 5. 使用短变量声明，自动推断类型
	f := "short"
	fmt.Println("短变量声明 f:", f) // 使用短变量声明，自动推断类型
	fmt.Println()
}

// demoVarVsShortDeclaration 详细演示 var 和 := 的使用场景
func demoVarVsShortDeclaration() {
	fmt.Println("=== var 与 := 使用场景对比 ===")

	// ========== 类型推断 vs 显式类型声明 ==========
	fmt.Println("--- 类型推断 vs 显式类型声明 ---")

	// var 的类型推断
	var autoInt = 42         // Go 推断为 int 类型
	var autoString = "hello" // Go 推断为 string 类型
	var autoFloat = 3.14     // Go 推断为 float64 类型（默认浮点类型）
	var autoBool = true      // Go 推断为 bool 类型

	// := 的类型推断
	shortInt := 42         // Go 推断为 int 类型
	shortString := "hello" // Go 推断为 string 类型
	shortFloat := 3.14     // Go 推断为 float64 类型
	shortBool := true      // Go 推断为 bool 类型

	fmt.Printf("var 推断类型 - int: %T, string: %T, float: %T, bool: %T\n",
		autoInt, autoString, autoFloat, autoBool)
	fmt.Printf(":= 推断类型 - int: %T, string: %T, float: %T, bool: %T\n",
		shortInt, shortString, shortFloat, shortBool)

	// 显式类型声明（当需要特定类型时）
	var explicitFloat32 float32 = 3.14 // 明确指定为 float32 而不是默认的 float64
	var explicitInt8 int8 = 100        // 明确指定为 int8 而不是默认的 int

	fmt.Printf("显式类型 - float32: %T, int8: %T\n", explicitFloat32, explicitInt8)

	// ========== 使用 var 的场景 ==========
	fmt.Println("\n--- 使用 var 的场景 ---")

	// 场景1: 需要显式指定类型时
	var age int = 25             // 明确指定为 int 类型
	var salary float64 = 50000.0 // 明确指定为 float64 类型
	var name string = "张三"       // 明确指定为 string 类型
	fmt.Printf("年龄: %d, 薪资: %.2f, 姓名: %s\n", age, salary, name)

	// 场景2: 声明变量但稍后赋值（零值初始化）
	var count int      // 默认值为 0
	var message string // 默认值为 ""
	var isActive bool  // 默认值为 false

	// 稍后赋值
	count = 10
	message = "Hello Go"
	isActive = true
	fmt.Printf("计数: %d, 消息: %s, 状态: %t\n", count, message, isActive)
	// 场景3: 批量声明变量
	var (
		username string = "admin"
		password string = "123456"
		port     int    = 8080
	)
	fmt.Printf("用户名: %s, 密码: %s, 端口: %d\n", username, password, port)

	// 场景4: 需要明确类型转换时
	var pi float32 = 3.14 // 明确指定为 float32
	var radius float32 = 5.0
	var area float32 = pi * radius * radius
	fmt.Printf("圆的面积: %.2f\n", area)

	// ========== 使用 := 的场景 ==========
	fmt.Println("\n--- 使用 := 的场景 ---")

	// 场景1: 类型显而易见，可以自动推断
	userID := 12345             // 自动推断为 int
	email := "user@example.com" // 自动推断为 string
	isLoggedIn := true          // 自动推断为 bool
	fmt.Printf("用户ID: %d, 邮箱: %s, 已登录: %t\n", userID, email, isLoggedIn)

	// 场景2: 函数返回值赋值
	result := calculateSum(10, 20) // 自动推断函数返回值类型
	fmt.Printf("计算结果: %d\n", result)

	// 场景3: 循环中的临时变量
	numbers := []int{1, 2, 3, 4, 5}
	for i, value := range numbers { // i 和 value 都使用短声明
		fmt.Printf("索引: %d, 值: %d\n", i, value)
	}

	// 场景4: if 语句中的临时变量
	if length := len("Go语言"); length > 0 {
		fmt.Printf("字符串长度: %d\n", length)
	}

	// 场景5: 错误处理模式
	if err := doSomething(); err != nil {
		fmt.Printf("发生错误: %v\n", err)
	}

	fmt.Println()
}

// calculateSum 计算两个整数的和
func calculateSum(a, b int) int {
	return a + b
}

// doSomething 模拟可能返回错误的函数
func doSomething() error {
	// 模拟操作，这里返回 nil 表示成功
	return nil
}

// demoCommonMistakes 演示常见错误和最佳实践
func demoCommonMistakes() {
	fmt.Println("=== 常见错误和最佳实践 ===")

	// ❌ 错误：重复声明
	// name := "第一次声明"
	// name := "第二次声明"  // 这会导致编译错误

	// ✅ 正确：重新赋值
	name := "第一次声明"
	name = "重新赋值" // 正确的重新赋值
	fmt.Println("姓名:", name)

	// ❌ 错误：在包级别使用短声明
	// packageVar := "错误"  // 这在包级别是不允许的

	// ✅ 正确：在包级别使用 var
	// var packageVar = "正确"

	// 最佳实践建议
	fmt.Println("\n--- 最佳实践建议 ---")
	fmt.Println("1. 函数内部优先使用 := (更简洁)")
	fmt.Println("2. 需要零值初始化时使用 var")
	fmt.Println("3. 需要明确类型时使用 var")
	fmt.Println("4. 全局变量必须使用 var")
	fmt.Println("5. 批量声明时使用 var")
	fmt.Println()
}
