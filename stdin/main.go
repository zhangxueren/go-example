package main

import "fmt"

func main() {
	no := 0
	fmt.Print("请输入姓名和年龄，使用空格分隔：\n")
	for {
		var (
			name string
			age int
		)
		fmt.Scanln(&name, &age)
		if name == "" {
			break
		}
		no ++
		fmt.Printf("NO:%d name:%s age:%d\n", no, name, age)
	}

}
