package go_error

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func GetFileString(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("[error] Failed to open file: %v\n", err)
		// 异常退出 该怎么写
		return "", err
	}
	defer file.Close() // 关闭文件

	var fileContent string
	scanner := bufio.NewScanner(file) // 创建一个新的 Scanner 用来读取文件
	for scanner.Scan() {              // 逐行读取文件内容
		line := scanner.Text()
		fileContent += line + " "
	}

	// 检查读取过程中是否发生错误
	if err := scanner.Err(); err != nil {
		fmt.Printf("[error] Failed to read file: %v\n", err)
		return "", err
	}
	return fileContent, nil
}

func TestOpenFile(t *testing.T) {
	var filename, fileContent string
	var err error

	filename = "test.txt"
	fileContent, err = GetFileString(filename)
	if err != nil {
		fmt.Printf("[error] GetFileString(\"%s\") : %v\n", filename, err)
	}
	fmt.Printf("file %s : %s\n", filename, fileContent)

	filename = "test1.txt"
	fileContent, err = GetFileString(filename)
	if err != nil {
		fmt.Printf("[error] GetFileString(\"%s\") : %v\n", filename, err)
	}
	fmt.Printf("file %s : %s\n", filename, fileContent)
}
