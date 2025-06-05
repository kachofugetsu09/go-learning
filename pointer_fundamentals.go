package main

import "fmt"

// demoPointerFundamentals 演示指针的基本运算规则
func demoPointerFundamentals() {
	fmt.Println("🔍 ========== 指针基本运算规则 ==========")
	fmt.Println("解答：*(&i) == i 吗？")
	fmt.Println()

	// 1. 基本概念验证
	demoBasicPointerMath()

	// 2. 详细步骤分解  
	demoStepByStepAnalysis()

	// 3. 多种类型验证
	demoMultipleTypes()

	// 4. 内存地址分析
	demoMemoryAddressAnalysis()

	// 5. 常见误区澄清
	demoClarifyMisconceptions()

	fmt.Println("✅ 结论：*(&i) == i 在 Go 中永远成立！")
	fmt.Println("=" + fmt.Sprintf("%*s", 50, "="))
}

// 1. 基本概念验证
func demoBasicPointerMath() {
	fmt.Println("📚 1. 基本概念验证")
	fmt.Println("─────────────────")

	i := 42
	fmt.Printf("变量 i 的值: %d\n", i)
	fmt.Printf("变量 i 的地址: %p\n", &i)

	// 关键验证：*(&i) == i
	result := *(&i)
	fmt.Printf("*(&i) 的值: %d\n", result)
	fmt.Printf("*(&i) == i 的结果: %t\n", *(&i) == i)

	fmt.Println("\n🔍 运算过程分解:")
	fmt.Println("1. &i     → 获取 i 的内存地址")
	fmt.Println("2. *(&i)  → 对地址解引用，获取该地址存储的值")
	fmt.Println("3. 结果   → 回到原始值 i")

	fmt.Println()
}

// 2. 详细步骤分解
func demoStepByStepAnalysis() {
	fmt.Println("🔬 2. 详细步骤分解")
	fmt.Println("─────────────────")

	i := 100
	fmt.Printf("步骤0: i = %d (原始值)\n", i)

	// 步骤1：取地址
	addr := &i
	fmt.Printf("步骤1: &i = %p (取地址操作)\n", addr)

	// 步骤2：解引用
	value := *addr
	fmt.Printf("步骤2: *(&i) = %d (解引用操作)\n", value)

	// 步骤3：比较
	fmt.Printf("步骤3: *(&i) == i 结果: %t\n", value == i)

	fmt.Println("\n📋 运算规则总结:")
	fmt.Println("& (取地址操作符) 和 * (解引用操作符) 是互逆操作")
	fmt.Println("就像数学中的 +5 和 -5，或者 ×2 和 ÷2")
	fmt.Println("&i 得到地址，*(&i) 通过地址回到原值")

	fmt.Println()
}

// 3. 多种类型验证
func demoMultipleTypes() {
	fmt.Println("🧪 3. 多种类型验证")
	fmt.Println("─────────────────")

	// int 类型
	intVal := 42
	fmt.Printf("int:    %d, *(&intVal) == intVal: %t\n", intVal, *(&intVal) == intVal)

	// string 类型
	strVal := "Hello Go"
	fmt.Printf("string: %s, *(&strVal) == strVal: %t\n", strVal, *(&strVal) == strVal)

	// float64 类型
	floatVal := 3.14159
	fmt.Printf("float:  %.5f, *(&floatVal) == floatVal: %t\n", floatVal, *(&floatVal) == floatVal)

	// bool 类型
	boolVal := true
	fmt.Printf("bool:   %t, *(&boolVal) == boolVal: %t\n", boolVal, *(&boolVal) == boolVal)

	// 切片类型
	sliceVal := []int{1, 2, 3}
	fmt.Printf("slice:  %v, *(&sliceVal) 相等: %t\n", sliceVal, 
		fmt.Sprintf("%v", *(&sliceVal)) == fmt.Sprintf("%v", sliceVal))

	fmt.Println("\n💡 结论：对于所有类型，*(&variable) == variable 都成立")

	fmt.Println()
}

// 4. 内存地址分析
func demoMemoryAddressAnalysis() {
	fmt.Println("🧠 4. 内存地址分析")
	fmt.Println("─────────────────")

	i := 777
	fmt.Printf("变量 i 的值: %d\n", i)
	fmt.Printf("变量 i 的地址: %p\n", &i)

	// 创建指针变量
	ptr := &i
	fmt.Printf("指针 ptr 的值: %p (存储 i 的地址)\n", ptr)
	fmt.Printf("指针 ptr 的地址: %p\n", &ptr)

	fmt.Println("\n🎯 关键理解:")
	fmt.Printf("• &i 返回地址: %p\n", &i)
	fmt.Printf("• ptr 存储地址: %p\n", ptr)
	fmt.Printf("• &i == ptr: %t (同一个地址)\n", &i == ptr)
	fmt.Printf("• *ptr == i: %t (同一个值)\n", *ptr == i)
	fmt.Printf("• *(&i) == i: %t (关键等式)\n", *(&i) == i)

	fmt.Println("\n📐 内存模型:")
	fmt.Println("┌─────────────┬─────────┬──────────┐")
	fmt.Println("│   操作      │  地址   │    值    │")
	fmt.Println("├─────────────┼─────────┼──────────┤")
	fmt.Printf("│     i       │   ---   │   %d   │\n", i)
	fmt.Printf("│    &i       │ %p │   ---    │\n", &i)
	fmt.Printf("│   *(&i)     │   ---   │   %d   │\n", *(&i))
	fmt.Println("└─────────────┴─────────┴──────────┘")

	fmt.Println()
}

// 5. 常见误区澄清
func demoClarifyMisconceptions() {
	fmt.Println("❓ 5. 常见误区澄清")
	fmt.Println("─────────────────")

	i := 50

	fmt.Println("🚫 常见误解:")
	fmt.Println("误解1: '取地址后就变成了不同的东西'")
	fmt.Printf("事实: &i 只是获取地址 %p，i 本身还是 %d\n", &i, i)

	fmt.Println("\n误解2: '指针运算会改变原始值'")
	fmt.Printf("事实: *(&i) 只是读取操作，i 还是 %d\n", i)

	fmt.Println("\n误解3: '&和*不是完全的逆操作'")
	fmt.Println("事实: 对于同一个变量，&和*确实是逆操作")

	fmt.Println("\n✅ 正确理解:")
	fmt.Println("1. & 操作: 变量 → 地址")
	fmt.Println("2. * 操作: 地址 → 变量")
	fmt.Println("3. 组合: *(&变量) = 变量")

	fmt.Println("\n🏠 生活类比:")
	fmt.Println("• i 就像你家里的电视")
	fmt.Println("• &i 就像你家的地址")
	fmt.Println("• *(&i) 就像根据地址找到你家，再看到电视")
	fmt.Println("• 最终看到的还是同一台电视!")

	fmt.Println("\n🧮 数学类比:")
	fmt.Println("就像数学中:")
	fmt.Println("• f(x) = x + 5")
	fmt.Println("• g(x) = x - 5") 
	fmt.Println("• g(f(x)) = g(x + 5) = (x + 5) - 5 = x")
	fmt.Println("• 同样: *(&i) = i")

	// 实际验证
	fmt.Println("\n🔬 最终验证:")
	for j := 0; j < 5; j++ {
		testVal := j * 10
		fmt.Printf("测试 %d: *(&%d) == %d 结果: %t\n", 
			j+1, testVal, testVal, *(&testVal) == testVal)
	}

	fmt.Println()
}
