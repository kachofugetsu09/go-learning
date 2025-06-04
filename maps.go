package main

import "fmt"

// demoMaps 演示 Go 语言中的 map 基本操作
func demoMaps() {
	m := make(map[string]int) // 创建一个空的 map

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m) // 打印 map

	v1 := m["k1"]           // 访问 map 中的元素
	fmt.Println("v1: ", v1) // 打印 k1 的值

	fmt.Println("len:", len(m)) // 打印 map 的长度

	delete(m, "k2")        // 删除 k2
	fmt.Println("map:", m) // 打印删除后的 map

	_, prs := m["k2"] // 检查 k2 是否存在 之所以用_是用来消除键不存在和键值为零产生的歧义
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} // 创建并初始化一个 map
	fmt.Println("map:", n)                  // 打印初始化的 map
}

// demoMapVsJava 对比 Go Map 和 Java HashMap 的差异
// 重点解释为什么返回 bool 而不是 null
func demoMapVsJava() {
	fmt.Println("=== Go Map vs Java HashMap 差异演示 ===")

	// 1. Go 语言的设计哲学：明确性和类型安全
	m := make(map[string]int)
	m["existing"] = 42
	m["zero"] = 0 // 特意设置一个值为 0 的键

	fmt.Println("--- Go 的双返回值机制 ---")

	// 情况1: 键存在，值不为零
	value1, exists1 := m["existing"]
	fmt.Printf("键 'existing': 值=%d, 存在=%t\n", value1, exists1)

	// 情况2: 键存在，但值为零（这就是关键！）
	value2, exists2 := m["zero"]
	fmt.Printf("键 'zero': 值=%d, 存在=%t\n", value2, exists2)

	// 情况3: 键不存在
	value3, exists3 := m["notfound"]
	fmt.Printf("键 'notfound': 值=%d, 存在=%t\n", value3, exists3)

	fmt.Println("\n--- 为什么不能用 null？---")
	fmt.Println("问题分析：")
	fmt.Println("1. 如果键 'zero' 存在且值为 0")
	fmt.Println("2. 如果键 'notfound' 不存在，Go 返回零值也是 0")
	fmt.Println("3. 单纯看值，无法区分这两种情况！")

	// 4. 演示单返回值的问题
	fmt.Println("\n--- 单返回值的歧义演示 ---")
	singleValue1 := m["zero"]     // 存在的键，值为 0
	singleValue2 := m["notfound"] // 不存在的键，返回零值 0

	fmt.Printf("仅获取值 - 'zero': %d\n", singleValue1)
	fmt.Printf("仅获取值 - 'notfound': %d\n", singleValue2)
	fmt.Println("看！两个都是 0，无法区分是否存在！")

	fmt.Println("\n--- Go 的解决方案：双返回值 ---")
	fmt.Println("第一个返回值：键对应的值（如果不存在返回零值）")
	fmt.Println("第二个返回值：bool 类型，true=存在，false=不存在")
	fmt.Println("这样就完美解决了歧义问题！")

	fmt.Println()
}

// demoJavaStyleComparison 用 Go 模拟 Java 的 HashMap 行为来对比
func demoJavaStyleComparison() {
	fmt.Println("=== 与 Java HashMap 的对比 ===")

	// Java 中的 HashMap 行为模拟
	fmt.Println("--- Java HashMap 的行为 ---")
	fmt.Println("HashMap<String, Integer> map = new HashMap<>();")
	fmt.Println("map.put(\"existing\", 42);")
	fmt.Println("map.put(\"zero\", 0);")
	fmt.Println()
	fmt.Println("// Java 中检查键是否存在")
	fmt.Println("Integer value1 = map.get(\"existing\");  // 返回 42")
	fmt.Println("Integer value2 = map.get(\"zero\");     // 返回 0")
	fmt.Println("Integer value3 = map.get(\"notfound\"); // 返回 null")
	fmt.Println()
	fmt.Println("// Java 需要额外的方法来检查")
	fmt.Println("boolean exists = map.containsKey(\"notfound\"); // false")

	fmt.Println("\n--- Go 的优势 ---")
	m := make(map[string]int)
	m["existing"] = 42
	m["zero"] = 0

	// Go 一步到位，同时获取值和存在性
	_, exists := m["notfound"]
	fmt.Printf("Go: value, exists := m[\"notfound\"]  // exists = %t\n", exists)
	fmt.Println("优势：一次操作同时获得值和存在性，更高效！")

	fmt.Println("\n--- 类型安全对比 ---")
	fmt.Println("Java: 可能返回 null，需要空指针检查")
	fmt.Println("Go: 永远不会有 null，类型安全，零值 + bool 标志")

	fmt.Println()
}

// demoCommonPatterns 演示 Go Map 常用模式
func demoCommonPatterns() {
	fmt.Println("=== Go Map 常用模式演示 ===")

	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 0 // 特意设置为 0

	// 模式1: 检查键是否存在
	fmt.Println("--- 模式1: 检查键是否存在 ---")
	if _, exists := m["apple"]; exists {
		fmt.Println("苹果存在于 map 中")
	} else {
		fmt.Println("苹果不存在于 map 中")
	}

	// 模式2: 获取值，如果不存在使用默认值
	fmt.Println("\n--- 模式2: 获取值或默认值 ---")
	getValue := func(key string, defaultValue int) int {
		if value, exists := m[key]; exists {
			return value
		}
		return defaultValue
	}

	fmt.Printf("苹果数量: %d\n", getValue("apple", -1))  // 存在，返回 5
	fmt.Printf("香蕉数量: %d\n", getValue("banana", -1)) // 存在但为 0，返回 0
	fmt.Printf("橘子数量: %d\n", getValue("orange", -1)) // 不存在，返回默认值 -1

	// 模式3: 安全删除（Java 风格 vs Go 风格）
	fmt.Println("\n--- 模式3: 安全删除 ---")
	// Java 风格：先检查再删除
	if _, exists := m["apple"]; exists {
		delete(m, "apple")
		fmt.Println("Java 风格：先检查，再删除")
	}

	// Go 风格：直接删除（删除不存在的键是安全的）
	delete(m, "notexist") // 这是安全的，不会报错
	fmt.Println("Go 风格：直接删除，即使键不存在也安全")

	// 模式4: 计数器模式
	fmt.Println("\n--- 模式4: 计数器模式 ---")
	counter := make(map[string]int)
	words := []string{"go", "java", "go", "python", "go"}

	for _, word := range words {
		// Go 的零值特性让计数器模式非常简洁
		counter[word]++ // 如果键不存在，零值是 0，0+1=1
	}

	for word, count := range counter {
		fmt.Printf("%s: %d 次\n", word, count)
	}

	fmt.Println()
}

// demoZeroValues 演示 Go 的零值概念
func demoZeroValues() {
	fmt.Println("=== Go 零值概念演示 ===")

	fmt.Println("--- 不同类型的零值 ---")
	intMap := make(map[string]int)
	stringMap := make(map[string]string)
	boolMap := make(map[string]bool)

	// 访问不存在的键，返回对应类型的零值
	fmt.Printf("int 零值: %d\n", intMap["notexist"])         // 0
	fmt.Printf("string 零值: '%s'\n", stringMap["notexist"]) // ""
	fmt.Printf("bool 零值: %t\n", boolMap["notexist"])       // false

	// 这就是为什么需要第二个返回值来区分
	fmt.Println("\n--- 零值 vs 不存在的区别 ---")
	intMap["zero"] = 0
	intMap["empty"] = 0 // 假设这是有意设置的

	// 仅通过值无法区分
	fmt.Printf("intMap[\"zero\"]: %d\n", intMap["zero"])
	fmt.Printf("intMap[\"notexist\"]: %d\n", intMap["notexist"])
	fmt.Println("两者都是 0，但含义不同！")

	// 通过第二个返回值区分
	value1, exists1 := intMap["zero"]
	value2, exists2 := intMap["notexist"]
	fmt.Printf("\"zero\": 值=%d, 存在=%t\n", value1, exists1)
	fmt.Printf("\"notexist\": 值=%d, 存在=%t\n", value2, exists2)

	fmt.Println()
}
