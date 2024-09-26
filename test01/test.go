package main

import (
	"fmt"
	"time"
)

// compareStrings 判断两个字符串是否相等
func compareStrings(s1, s2 string) bool {
	return s1 == s2
}

func main() {
	//str1 := "hello"
	// 生成长度为 5000 的字符串
	//baseStr := "hello"
	//longStr := strings.Repeat(baseStr, 5000/len(baseStr))

	baseStr := "9月23日晚，有人在網上發帖稱鬆江車墩發生爆炸。鬆江警方核查証實，該事件為當晚8時許，一輛途經G15沈海高速鬆江段的貨車車頭部位發生自燃，未造成人員受傷。\n\n經查，25歲女子王某某在家中陽台上看到路邊有火光和冒煙后，拍攝了視頻。隨后，她在未核實事實的情況下，為博人眼球，在所拍攝的視頻中自行添加了爆炸音效，並發布網帖稱“鬆江車墩發生爆炸”。\n\n目前，1王某某已被警方依法處以行政拘留1。"
	baseStr2 := "9月23日晚，有人在網上發帖稱鬆江車墩發生爆炸。鬆江警方核查証實，該事件為當晚8時許，一輛途經G15沈海高速鬆江段的貨車車頭部位發生自燃，未造成人員受傷。\n\n經查，25歲女子王某某在家中陽台上看到路邊有火光和冒煙后，拍攝了視頻。隨后，她在未核實事實的情況下，為博人眼球，在所拍攝的視頻中自行添加了爆炸音效，並發布網帖稱“鬆江車墩發生爆炸”。\n\n目前，2王某某已被警方依法處以行政拘留1。"

	// 开始计时
	startTime := time.Now()

	timestampMs := startTime.UnixMilli()
	fmt.Println("Unix 时间戳（毫秒）:", timestampMs)

	// 进行字符串比较
	result := compareStrings(baseStr, baseStr2)

	//time.Sleep(100)

	a := 10000

	for a > 0 {
		//time.Sleep(100)

		result = compareStrings(baseStr, baseStr2)

		a--

	}

	endTime := time.Now()

	fmt.Println("Unix 时间戳（毫秒）:", endTime.UnixMilli())

	// 结束计时
	//duration := time.Since(startTime)

	// 将持续时间转换为毫秒
	//milliseconds := duration.Milliseconds()

	fmt.Printf("str1 和 str2 是否相等？%t\n", result) // 输出 false
	//fmt.Printf("比较耗时: %d\n", duration)
	fmt.Printf("比较耗时（毫秒）: %d\n", endTime.UnixMilli()-startTime.UnixMilli())
}

/*
main 函数
//  */
// func main() {
// 	println("Hello, world!")

// 	// 基本数据类型
// 	var a int = 10;

// 	println(a)

// 	// 浮点型
// 	var b float64 = 3.14
// 	println(b)

// 	//

// 	test1()

// }
