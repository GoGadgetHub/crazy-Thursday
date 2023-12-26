package cmd

import (
	"fmt"
	"log"
	"os"
)

func Print(res string) {
	// 创建文件
	file, err := os.Create("./output/output.txt")
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// 将内容写入文件
	_, err = file.WriteString(res)
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}

	fmt.Println("内容已成功写入文件 output.txt")
}
