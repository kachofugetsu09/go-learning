package main

import (
	"fmt"
)

// demoPointer 综合指针演示 - 从基础到深入
func demoPointer() {
	fmt.Println("🔍 ========== Go 指针全面解析 ==========")
	fmt.Println("作为 Java 开发者，让我们深入理解 Go 的指针概念")
	fmt.Println()

	// 1. 基础指针概念
	demoBasicPointerConcepts()

	// 2. 值传递 vs 指针传递
	demoValueVsPointerPassing()

	// 3. 核心问题：为什么是 *iptr = 0 而不是 iptr = 0
	demoWhyStarPtrNotPtr()

	// 4. 内存布局可视化
	demoMemoryLayout()

	// 5. 指针操作对比
	demoPointerOperations()

	// 6. 常见陷阱和最佳实践
	demoPointerPitfalls()

	// 7. Go vs Java 对比
	demoGoVsJavaPointers()

	fmt.Println("✅ 指针演示完成! 现在你应该完全理解 Go 指针了。")
	fmt.Println("=" + fmt.Sprintf("%*s", 50, "="))
}

// 1. 基础指针概念
func demoBasicPointerConcepts() {
	fmt.Println("📚 1. 基础指针概念")
	fmt.Println("─────────────────────")

	// 声明一个普通变量
	i := 42
	fmt.Printf("变量 i 的值: %d\n", i)
	fmt.Printf("变量 i 的地址: %p\n", &i)

	// 声明一个指针变量
	var p *int // 声明一个指向 int 类型的指针
	p = &i     // 将 i 的地址赋给指针 p
	fmt.Printf("指针 p 的值（存储的地址）: %p\n", p)
	fmt.Printf("指针 p 本身的地址: %p\n", &p)
	fmt.Printf("指针 p 指向的值: %d\n", *p)

	// 通过指针修改值
	*p = 100
	fmt.Printf("通过指针修改后，i 的值: %d\n", i)
	fmt.Printf("验证 p 和 &i 是同一个地址: %t\n", p == &i)

	fmt.Println()
}

// zeroval 通过值传递修改参数（不会影响原变量）
func zeroval(ival int) {
	fmt.Printf("  📍 zeroval函数内部 - 修改前参数值: %d, 参数地址: %p\n", ival, &ival)
	ival = 0 // 只修改了参数的副本
	fmt.Printf("  📍 zeroval函数内部 - 修改后参数值: %d, 参数地址: %p\n", ival, &ival)
}

// zerptr 通过指针传递修改参数（会影响原变量）
func zerptr(iptr *int) {
	fmt.Printf("  📍 zerptr函数内部 - 指针地址: %p, 指针指向的值: %d\n", iptr, *iptr)
	*iptr = 0 // 通过指针修改原变量的值
	fmt.Printf("  📍 zerptr函数内部 - 修改后指针指向的值: %d\n", *iptr)
}

// 2. 值传递 vs 指针传递
func demoValueVsPointerPassing() {
	fmt.Println("🔄 2. 值传递 vs 指针传递对比")
	fmt.Println("────────────────────────────")

	i := 1
	fmt.Printf("初始值: %d, 地址: %p\n", i, &i)

	fmt.Println("\n🟡 值传递演示:")
	zeroval(i)
	fmt.Printf("调用 zeroval 后，原变量 i 的值: %d (没有改变!)\n", i)

	fmt.Println("\n🟢 指针传递演示:")
	zerptr(&i)
	fmt.Printf("调用 zerptr 后，原变量 i 的值: %d (被修改了!)\n", i)

	fmt.Println()
}

// 3. 核心问题：为什么是 *iptr = 0 而不是 iptr = 0
func demoWhyStarPtrNotPtr() {
	fmt.Println("❓ 3. 为什么是 *iptr = 0 而不是 iptr = 0？")
	fmt.Println("─────────────────────────────────────")

	i := 1
	iptr := &i
	fmt.Printf("变量 i 的值: %d, 地址: %p\n", i, &i)
	fmt.Printf("指针 iptr 的值(存储的地址): %p\n", iptr)
	fmt.Printf("指针 iptr 指向的值: %d\n", *iptr)

	fmt.Println("\n🎯 执行 *iptr = 0 (修改指针指向的值):")
	*iptr = 0
	fmt.Printf("结果: i = %d, iptr = %p, *iptr = %d\n", i, iptr, *iptr)
	fmt.Println("✅ 指针地址没变，但指向的值变了!")

	// 演示如果要修改指针本身
	fmt.Println("\n🎯 如果要修改指针本身(让它指向别的地址):")
	j := 99
	oldPtr := iptr
	iptr = &j // 修改指针本身
	fmt.Printf("新变量 j = %d, 地址 = %p\n", j, &j)
	fmt.Printf("原指针值: %p\n", oldPtr)
	fmt.Printf("新指针值: %p (指向了不同的地址!)\n", iptr)
	fmt.Printf("新指向的值: %d\n", *iptr)

	fmt.Println("\n🚫 为什么不能写 iptr = 0？")
	fmt.Println("   因为 0 不是一个有效的内存地址!")
	fmt.Println("   内存地址 0 通常是系统保留的，程序无法访问")
	fmt.Println("   在 Go 中，nil 是指针的零值，表示不指向任何地址")

	var nilPtr *int
	fmt.Printf("   nil 指针的值: %v\n", nilPtr)
	fmt.Printf("   nil 指针是否为 nil: %t\n", nilPtr == nil)
	fmt.Println("   ⚠️  对 nil 指针进行解引用(*nilPtr)会导致程序崩溃!")

	fmt.Println()
}

// 4. 内存布局可视化
func demoMemoryLayout() {
	fmt.Println("🧠 4. 内存布局可视化")
	fmt.Println("──────────────────")

	i := 42
	ptr := &i

	fmt.Println("内存布局概念图:")
	fmt.Println("┌─────────────────┬──────────┬─────────────┐")
	fmt.Println("│   内存地址      │   存储值  │    变量名   │")
	fmt.Println("├─────────────────┼──────────┼─────────────┤")
	fmt.Printf("│ %p │    %d    │      i      │\n", &i, i)
	fmt.Printf("│ %p │ %p │     ptr     │\n", &ptr, ptr)
	fmt.Println("└─────────────────┴──────────┴─────────────┘")

	fmt.Println("\n📝 解释:")
	fmt.Printf("• 变量 i 存储在地址 %p，值是 %d\n", &i, i)
	fmt.Printf("• 指针 ptr 存储在地址 %p，它的值是 %p (指向 i 的地址)\n", &ptr, ptr)
	fmt.Printf("• ptr 和 &i 的值相同: %t\n", ptr == &i)

	fmt.Println("\n🏠 房子类比:")
	fmt.Println("• 把指针想象成门牌号码")
	fmt.Printf("• 门牌号码本身: ptr = %p\n", ptr)
	fmt.Println("• 门牌号码指向的房子里的东西: *ptr = " + fmt.Sprintf("%d", *ptr))
	fmt.Println("• *ptr = 0 意思是: 去到这个门牌号，把房子里的东西换成0")
	fmt.Println("• ptr = &other 意思是: 换一个门牌号")

	fmt.Println()
}

// 5. 指针操作对比
func demoPointerOperations() {
	fmt.Println("⚖️  5. 指针操作对比表")
	fmt.Println("─────────────────────")

	fmt.Println("┌─────────────────────┬─────────────────────┐")
	fmt.Println("│      *ptr = 0       │      ptr = &other   │")
	fmt.Println("├─────────────────────┼─────────────────────┤")
	fmt.Println("│   修改指向的值       │     修改指针本身     │")
	fmt.Println("│   (解引用赋值)       │     (指针赋值)      │")
	fmt.Println("│   指针地址不变       │     指针地址改变     │")
	fmt.Println("│   目标地址的值改变   │     指向新的地址     │")
	fmt.Println("│   原变量被修改       │     指向新变量      │")
	fmt.Println("└─────────────────────┴─────────────────────┘")

	// 实际演示
	fmt.Println("\n🧪 实际演示:")
	x := 10
	y := 20
	ptr := &x

	fmt.Printf("初始状态: x=%d, y=%d, ptr指向x\n", x, y)

	// 操作1: *ptr = 0
	*ptr = 0
	fmt.Printf("执行 *ptr = 0 后: x=%d, y=%d, ptr仍指向x\n", x, y)

	// 操作2: ptr = &y
	ptr = &y
	fmt.Printf("执行 ptr = &y 后: x=%d, y=%d, ptr现在指向y\n", x, y)

	fmt.Println()
}

// 6. 常见陷阱和最佳实践
func demoPointerPitfalls() {
	fmt.Println("⚠️  6. 常见陷阱和最佳实践")
	fmt.Println("──────────────────────────")

	fmt.Println("🚫 常见陷阱:")
	fmt.Println("1. 对 nil 指针解引用")
	fmt.Println("   var p *int")
	fmt.Println("   // fmt.Println(*p)  ← 这会 panic!")

	fmt.Println("\n2. 返回局部变量的指针 (Go 中是安全的，但要理解)")
	fmt.Println("   Go 的垃圾回收器会处理这种情况")

	fmt.Println("\n3. 指针和切片/映射的混淆")
	fmt.Println("   切片和映射本身就是引用类型，通常不需要指针")

	fmt.Println("\n✅ 最佳实践:")
	fmt.Println("1. 使用指针传递大型结构体以避免复制")
	fmt.Println("2. 需要修改函数参数时使用指针")
	fmt.Println("3. 总是检查指针是否为 nil")
	fmt.Println("4. 使用 new() 或 &Type{} 创建指针")

	// 演示安全的指针检查
	fmt.Println("\n🛡️ 安全的指针使用:")
	var safePtr *int
	if safePtr != nil {
		fmt.Println("指针安全，可以解引用")
	} else {
		fmt.Println("指针为 nil，不能解引用")
	}

	// 使用 new 创建指针
	newPtr := new(int)
	*newPtr = 42
	fmt.Printf("使用 new() 创建的指针: %p, 值: %d\n", newPtr, *newPtr)

	fmt.Println()
}

// 7. Go vs Java 对比
func demoGoVsJavaPointers() {
	fmt.Println("☕ 7. Go vs Java 对比")
	fmt.Println("────────────────────")

	fmt.Println("Java (引用):")
	fmt.Println("```java")
	fmt.Println("Integer obj = new Integer(42);")
	fmt.Println("Integer ref = obj;  // ref 指向同一个对象")
	fmt.Println("obj = null;         // obj 不再指向对象")
	fmt.Println("// ref 仍然指向原对象")
	fmt.Println("```")

	fmt.Println("\nGo (指针):")
	fmt.Println("```go")
	fmt.Println("i := 42")
	fmt.Println("ptr := &i           // ptr 指向 i 的内存地址")
	fmt.Println("*ptr = 0            // 修改 i 的值")
	fmt.Println("ptr = nil           // ptr 不再指向任何地址")
	fmt.Println("// i 仍然存在，值为 0")
	fmt.Println("```")

	fmt.Println("\n🔍 关键区别:")
	fmt.Println("┌─────────────────┬─────────────────┬─────────────────┐")
	fmt.Println("│      特性       │      Java       │       Go        │")
	fmt.Println("├─────────────────┼─────────────────┼─────────────────┤")
	fmt.Println("│   指针可见性     │      隐藏        │      显式       │")
	fmt.Println("│   内存管理      │    垃圾回收      │    垃圾回收     │")
	fmt.Println("│   指针运算      │     不允许       │     不允许      │")
	fmt.Println("│   空指针        │      null       │      nil        │")
	fmt.Println("│   解引用        │      自动        │   手动 (*)     │")
	fmt.Println("│   基本类型      │   对象包装       │    直接指针     │")
	fmt.Println("└─────────────────┴─────────────────┴─────────────────┘")

	fmt.Println("\n💡 理解要点:")
	fmt.Println("• Java 的对象引用类似于 Go 的指针，但语法更隐式")
	fmt.Println("• Go 的指针更接近 C/C++，但有垃圾回收")
	fmt.Println("• Java 中 obj.field，Go 中可以是 (*ptr).field 或 ptr.field")
	fmt.Println("• Go 的指针让内存操作更透明和可控")
	fmt.Println()
}
