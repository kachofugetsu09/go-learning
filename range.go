package main

import "fmt"

func rangeDemo(){

	nums := []int{1, 2, 3, 4, 5}
	sum :=0
	for _,num := range nums{ // 使用 _ 忽略索引
		sum += num
	}

	fmt.Println("Sum of nums:", sum)

	for i, num := range nums{
		fmt.Println("Index:", i, "Value:", num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cherry"}
	for k, v := range kvs{
		fmt.Println("Key:", k, "Value:", v)
	}

	for k := range kvs {
		fmt.Println("Key:", k) // 只遍历键
	}

	for _,v := range kvs {
		fmt.Println("Value:", v) // 只遍历值
	}

	for i,c :=range "go"{ //range 在字符串中迭代 unicode 码点(code point)。 第一个返回值是字符的起始字节位置，然后第二个是字符本身。
		fmt.Println("Index:", i, "Rune:", c) // 遍历字符串中的每个字符
	}
}