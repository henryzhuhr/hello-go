package json_test

import (
	"encoding/json"
	"os"
	"testing"
)

// 将json数据写入到文件中
func TestWriteJsonToFile(t *testing.T) {
	const J_FILE = "/tmp/json_demo.json"
	// 打开文件
	// file, err := os.OpenFile(J_FILE, os.O_CREATE|os.O_RDWR, 0666)
	file, err := os.Create(J_FILE)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	// 延迟关闭文件
	defer file.Close()

	// 创建一个json编码器
	encoder := json.NewEncoder(file)

	// 创建一个map
	jsonData := map[string]interface{}{
		"name":  "zhangsan",
		"age":   18,
		"email": "xxx@xxx.com",
	}

	// 编码map数据到文件中
	err = encoder.Encode(jsonData)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("run 'cat %s' to show json data", jFile)
}

// 从文件中读取json数据
func TestReadJsonFromFile(t *testing.T) {
	// 打开文件
	file, err := os.OpenFile("./demo.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	// 延迟关闭文件
	defer file.Close()

	// 创建一个json编码器
	decoder := json.NewDecoder(file)

	// 解码json数据，存储到map中
	jsonData := make(map[string]interface{})
	err = decoder.Decode(&jsonData)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("jsonData: %v", jsonData)
}
