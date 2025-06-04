package main
import "fmt"

func demoSlice(){
	s :=make([]string, 3) // 创建一个长度为3的切片
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	s = append(s, "d", "e") // 使用 append 添加元素
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c) // 复制切片

	l := s[2:5]
	fmt.Println("slc:", l) // 切片操作，获取索引2到5的元素

	l = s[:5] // 获取前5个元素
	fmt.Println("slc2:", l)

	l = s[2:] // 获取索引2到末尾的元素
	fmt.Println("slc3:", l)

	t := []string{"a", "b", "c", "d", "e"}
	fmt.Println("dcl:", t) // 声明并初始化一个字符串切片

	twoD := make([][]int, 3) // 创建一个二维切片
	for i := 0; i < 3; i++ {
		innerLen := i +1
		twoD[i] = make([]int, innerLen) // 每个内层切片的长度递增
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD) // 打印二维切片
}