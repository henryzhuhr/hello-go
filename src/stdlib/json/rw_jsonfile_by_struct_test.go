package json_test

import (
	"encoding/json"
	"os"
	"testing"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// 将结构体数据写入到文件中
func TestWriteByStructJsonToFile(t *testing.T) {
	jFile := "/tmp/json_by_struct_demo.json"
	// 打开文件
	file, err := os.OpenFile(jFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	// 延迟关闭文件
	defer file.Close()

	// 创建一个json编码器
	encoder := json.NewEncoder(file)

	// 创建一个结构体
	user := User{
		Name: "zhangsan",
		Age:  18,
	}

	// 编码结构体数据到文件中
	err = encoder.Encode(user)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("run 'cat %s' to show json data", jFile)
}

// 从文件中读取json数据到结构体
func TestReadJsonByStructFromFile(t *testing.T) {
	// 打开文件
	file, err := os.OpenFile("demo.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	// 延迟关闭文件
	defer file.Close()

	// 创建一个json编码器
	decoder := json.NewDecoder(file)

	// 解码json数据，存储到结构体中
	user := User{}
	err = decoder.Decode(&user)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("user.Name : %s", user.Name)
	t.Logf("user.Age  : %d", user.Age)
	t.Logf("user.Email: %s", user.Email)
}
