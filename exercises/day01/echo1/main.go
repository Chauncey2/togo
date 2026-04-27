package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(){
	s, sep := "", ""
	fmt.Println(os.Args[0]) // 程序名
	for tmp, arg := range os.Args[1:] {
		fmt.Println(tmp, arg) // 索引和参数
		s += sep + arg
		sep = " "
	}
	fmt.Println(s) // 参数拼接
}

func echo3(){
	fmt.Println(strings.Join(os.Args[1:]," "))
}

func main() {
	// echo1()
	// echo2()
	echo3()
}
