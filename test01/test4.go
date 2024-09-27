package main

import "fmt"

/**
控制结构
1）if else
2）for
3）switch
4）select
*/

func test4() {
	fmt.Println("test4")
	test41(15)
	test42(3)
}

func test41(age int) {
	if age >= 18 {
		fmt.Println("age >= 18")
	} else {
		fmt.Println("age < 18")
	}
}

func test42(num int) {
	for i := num; i > 0; i-- {
		fmt.Println(i)
	}
}
