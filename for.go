package main

import "fmt"

func demoFor() {

	i := 1
	for i <= 3 {
		fmt.Printf("第 %d 次循环\n", i)
		i++
	}

	for j := 1; j <= 3; j++ {
		fmt.Printf("第 %d 次循环\n", j)
	}

	var a int = 2
	for {
		a++
		fmt.Println("无限循环")
		if a > 3 {
			fmt.Println("a 大于 3，退出循环")
			break // 使用 break 退出循环
		}
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // 如果 n 是偶数，跳过本次循环
		}
		fmt.Printf("奇数: %d\n", n) // 输出奇数
	}
}
