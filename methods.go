package main

import "fmt"

type rect struct{
	width,height int
}

type counter struct{
	count int
}

func(c counter) incVal(){
	fmt.Println("当前计数:", c.count)
	c.count++
	fmt.Println("增加后的计数:", c.count)
}

func(c *counter) incPtr(){
	fmt.Println("当前计数:", c.count)
	c.count++
	fmt.Println("增加后的计数:", c.count)
}

func (r *rect) area() int {
	return r.width * r.height
}
// area 方法使用指针接收者，允许修改 rect 的值
// 这样可以避免在每次调用时都复制整个结构体，提高性能


func(r rect) perim() int {
	return 2 * (r.width + r.height)
}
//值接收者会复制一份值，然后操作原值不会改变

func methodDemo(){
	r := rect{width: 10, height: 5}
	fmt.Println("矩形的面积:", r.area()) // 调用指针接收者方法
	fmt.Println("矩形的周长:", r.perim()) // 调用值接收者方法

	rp := &r // 创建指向 rect 的指针
	fmt.Println("通过指针调用矩形的面积:", rp.area()) // 仍然可以通过指针调用指针接收者方法
	fmt.Println("通过指针调用矩形的周长:", rp.perim()) // 仍然可以通过指针调用值接收者方法

	count := counter{count: 0}
	fmt.Println("使用值接收者调用 incVal 方法:")
	count.incVal() // 调用值接收者方法，计数不会改变
	fmt.Println(count)
	fmt.Println("使用指针接收者调用 incPtr 方法:")
	count.incPtr() // 调用指针接收者方法，计数会改变
	fmt.Println(count) // 输出计数器的当前状态
	count.incPtr()
	fmt.Println(count) // 输出计数器的当前状态
	count.incVal() // 再次调用值接收者方法，计数不会改变
	fmt.Println("再次调用值接收者方法后，计数器状态:", count) // 输出计数器的当前状态
}
