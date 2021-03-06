// 练习 1.4： 修改dup2，出现重复的行时打印文件名称。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	files := os.Args[1:]

	// 从文件读入
	for _, arg := range files {
		// os.Open函数返回两个值。第一个值是被打开的文件(*os.File）
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup 2: %v\n", err)
			continue
		}
		countLines(f, arg)
		f.Close()
	}
}

func countLines(f *os.File, filename string)  {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	for _, n := range counts {
		if n > 1 {
			fmt.Printf("filename %s have repeat line\n", filename)
			break
		}
	}
}
