package main

import "fmt"

func f3(p *int) {
	*p = 10 //解引用
}

func f2(v int) {
	v = 8
}

func main() {
	var v = 12
	fmt.Println(v)
	f2(v) //复制了一份v,不是复制指向同一个对象的指针
	fmt.Println(v)
	f3(&v) //复制了一个指向同一个对象的指针
	fmt.Println(v)
	/**
	输出:
	12
	12
	10
	*/
}
