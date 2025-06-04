package main

import "fmt"

func demoIfElse() {
	fmt.Println("=== if-else 演示 ===")

	if 7%2 == 0 {
		fmt.Println("7 是偶数")
	} else {
		fmt.Println("7 是奇数")
	}

	if 8%4 == 0 {

		fmt.Println("8 是 4 的倍数")
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d 是偶数\n", i)
		} else {
			fmt.Printf("%d 是奇数\n", i)
		}
	}
}
