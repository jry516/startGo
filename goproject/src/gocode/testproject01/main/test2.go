package main //声明文件所在的包,每个go文件必须有归属的包

import (
	"fmt"
	"strconv"
) //引入需要的包,引用后必须要使用

func main() { // 主函数,程序的入口
	//字符类型,只能使用单引号, 默认使用UTF-8,UTF-8低位与ASCII码相同
	var c1 byte = 'a'
	fmt.Println(c1) // 64

	var c2 int = '汉' // 超ASCII码部分可以使用int类型保存
	fmt.Println(c2)

	var c3 = "字" // 使用双引号为string类型
	fmt.Println(c3)

	var n1 int = 23

	// int 转 string
	var s1 string = fmt.Sprintf("%d", n1)

	fmt.Printf("s1 = %v \n", s1)
	var s2 string = "true"
	var b2 bool
	b2, _ = strconv.ParseBool(s2)
	fmt.Printf("b2 = %t \n", b2)

	var i3 int32 = 32

	//& 取地址符
	var p1 *int32 = &i3
	fmt.Println("p1 = ", p1)

	fmt.Println("*p1 = ", *p1) // * 取值

	// 1 可以通过指针改变指向值
	var i11 int = 23
	var p2 *int = &i11
	*p2 = 66
	fmt.Println("i11 = ", i11) // 66
	// 2 指针变量接收的一定是地址值
	// 3 指针变量的值不可以不匹配
	// 4 基本数据类型都有对应的指针类型

	// _ 代表忽略,代表此值接收但不使用
	parseBool, _ := strconv.ParseBool("true")
	fmt.Println("parseBool = ", parseBool)

	// 包名package名字尽量与目录一致
	// 程序入口是main包下的main方法
	// 和标准库不要冲突
	// 如果函数名,变量名,常量名首字母大写,其它包可访问(public)
	// 如果首字母小写,其它包不可访问(private)
	// 导入包的地址是从$GOPATH/src/后开始

}
