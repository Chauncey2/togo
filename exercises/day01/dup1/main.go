package main

import (
	"bufio"
	"fmt"
	"os"
	// "io/ioutil" // 已弃用
	"strings"
)

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			} else {
				countLines(f, counts)
				f.Close()
			}
		}

		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

// 一次性读取文本，并进行处理
func dup3(){
	counts:=make(map[string]int)
	for _ ,filename :=range os.Args[1:]{
		//input:=io.ReadFile(filename)
		data,err := os.ReadFile(filename) //ReadFile函数返回一个字节切片（byte slice
		if err !=nil{
			fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
			continue
		}
		for _,line:= range strings.Split(string(data),"\n"){
			counts[line]++
		}
	}
	for line,n:=range counts{
		if n>1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}

func main() {
	// 1 2 方法都是流式读入，读一行处理一行
	// dup1()
	// dup2()
	// 一次性加载所有输入到内存进行处理
	dup3()
}
