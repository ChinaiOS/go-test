package main

import "fmt"

/**
方法声明与调用
1. 函数可以有多个返回值

2. 函数式编程：
1) 函数作为参数
2) 函数作为返回值
3) 匿名方法
4) 闭包
5) defer



*/

func test3() {

	//a, _ := test33(1, 2)
	//
	//fmt.Printf("%d\n", a)
	////fmt.Printf("%d\n", b)
	//
	//test341 := test34
	//
	//test341()

	test351 := test35(1, 2)

	//test351()

	fmt.Printf("%s\n", test351())

}

func test33(a int, b int) (int, int) {

	fmt.Printf("%d\n", a)
	//println(b)

	//test341 := test34()
	//
	//test341()

	return a + 1, b + 2

}

func test34() func() {
	fmt.Println("test34")

	return func() {
		fmt.Println("test34  方法作为返回值")
	}

}

func test35(a int, b int) func() string {
	return func() string {
		return fmt.Sprintf("%d %d", a, b)
	}
}
