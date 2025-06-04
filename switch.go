package main

import (
	"fmt"
	"time"
)

func switchDemo() {

	i := 2
	fmt.Println("=== switch-case 演示 ===")
	switch i {
	case 1:
		fmt.Println("i 是 1")
	case 2:
		fmt.Println("i 是 2")
	default:
		fmt.Println("i 不是 1 或 2")
	}

	switch time.Now().Weekday() {
	case time.Sunday:
		fmt.Println("今天是星期日")
	case time.Saturday:
		fmt.Println("今天是星期六")
	default:
		fmt.Println("今天不是星期五或星期六")
	}

	t := time.Now()
	switch{
	case t.Hour() <12 :
		fmt.Println("现在是上午")
	default:
		fmt.Println("现在是下午或晚上")
	}

	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
	whatAmI(true)
	whatAmI(42)
	whatAmI("Hello, Go!")

}

