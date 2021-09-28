package main

import "fmt"

func main() {
	var input int
	fmt.Print("输入一个数字:")
	fmt.Scanln(&input) //必须为内存空间地址
	fmt.Println("输入的数字为:", input)

	var inputBool bool

	fmt.Print("输入一个bool")
	fmt.Scanln(&inputBool)
	fmt.Println("inputBool = ", inputBool)

	fmt.Println("输入一个数字,bool")
	fmt.Scanf("%d,%t", &input, &inputBool)

	fmt.Printf("input = %d,inputBool = %t", input, inputBool)
}
