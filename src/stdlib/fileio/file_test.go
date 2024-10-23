package fileio_test

import (
	"os"
	"testing"
)

const DEFAULT_FILE_PATH = "/tmp/hello.txt"

// 打开文件，读取文件内容
func TestOpenFile(t *testing.T) {
	// 预先写入文件
	err := os.WriteFile(DEFAULT_FILE_PATH, []byte("hello world!\n"), 0666)
	if err != nil {
		t.Fatal(err)
	}

	// 打开文件
	file, err := os.Open(DEFAULT_FILE_PATH)
	if err != nil {
		t.Fatal(err)
	}
	file.Close()
	file, err = os.OpenFile(DEFAULT_FILE_PATH, os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	// 关闭文件
	defer file.Close()

	t.Log("file:", file)
}
