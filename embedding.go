package main

import "fmt"

type base struct{
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("Base number: %d", b.num)
}

type container struct{
	base
	str string
}

func embeddingDemo(){
	co := container{
		base: base{num: 42},
		str:  "Hello, World!",
	}
	

	fmt.Printf("Container: %s, %s\n", co.describe(), co.str) 

	//我们可以直接在 co 上访问 base 中定义的字段, 例如： co.num. 或者可以使用嵌入类型名拼写出完整路径
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface{
		describe() string
	}

	//由于嵌入了base 所以base的方法也变成container的方法 所以可以直接调用 所以我们这里视作container也实现了describer接口
	var d describer = co
	fmt.Println("Describer output:", d.describe())
}