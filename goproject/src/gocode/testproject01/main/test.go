package main //声明文件所在的包,每个go文件必须有归属的包

import (
	"fmt"
	"unsafe"
) //引入需要的包,引用后必须要使用

func main() { // 主函数,程序的入口
	fmt.Println("Hello Golang") // 不需要以;
	//  一次声明多变量的三种方式
	var n1, n2, n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4, name, n5 = 10, "golang", 7.8
	fmt.Println(n4, name, n5)

	n6, height := 6.9, 100
	fmt.Println(n6, height)

	fmt.Println(n7)
	fmt.Println(n8)

	//基本数据类型
	// 有符号整数
	var aaa int8 = 127
	aaa = aaa + 2
	fmt.Println(aaa) // 输出 -127
	// 无符号整数
	var uuu uint8 = 255
	fmt.Println(uuu) //输出 255
	uuu = uuu + 2
	fmt.Println(uuu) // 输出1
	//查看 占用字节数
	fmt.Println("uint8占用字节数:", unsafe.Sizeof(uuu)) //输出:uint8占用字节数: 1

	//浮点数
	var num1 float32
	num1 = 3.14
	fmt.Println(num1) //3.14
	var num2 = -3.14
	fmt.Println(num2) //-3.14
	var num3 = 3.14e2
	fmt.Println(num3) //314
	var num4 = 3.14e-2
	fmt.Println(num4) //0.0314

}

// 全局变量
var n7 = 255
var n8 = 22
var (
	n9  = 10
	n10 = 20
)
