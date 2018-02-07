package main

import (
	"fmt"
	"strconv"
	//"strings"
)

func main() {
	//字符串的使用
	str_sample()
	fmt.Println("--------------")
	//string包函数的使用
	//string_function_sample()
}

func str_sample() {
	str := "Hello"
	str += "world"
	fmt.Println("字符串add:", str)
	fmt.Println("count string length:", len(str))
	fmt.Println(" get the array: %c %d \n", str[0], str[1])

	//遍历
	r := []rune(str + "中国")
	fmt.Println("兼容unicde 字符遍历:")
	for i := 0; i < len(str); i++ {
		fmt.Println(string(r[i]))
	}

	//字符串转整数
	p, err := strconv.Atoi("12")
	if err == nil {
		fmt.Printf("字符串转整数：%d \n", p)
	}

	//整数转字符串
	fmt.Println("整数转字符串：", strconv.Itoa(12345))

	//字符串转byte
	fmt.Println("字符串转byte:", []byte("hello go"))

	//byte转字符串
	fmt.Println("[]BYTE转字符串", string([]byte{51, 52, 53, 54, 55, 56}))

	//十六进制
	fmt.Println("十六进制：", strconv.FormatInt(int64(28), 16))

	//十六进制转整数
	n2, _ := strconv.ParseInt("be", 16, 32)
	fmt.Printf("十六进制字符串转整数: %d  \n", int(n2))

}
