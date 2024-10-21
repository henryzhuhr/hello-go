package json_test

import (
	"encoding/json"
	"testing"
)

// 将json字符串转换为go结构体
func TestJsonStringToGoStruct(t *testing.T) {
	// 定义一个json字符串
	jsonStr := `{"name":"zhangsan","age":18,"email":"xxx@xxx.com"}`
	// 定义一个map
	jsonData := make(map[string]interface{})
	// 将json字符串解码到map中
	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("jsonData: %v", jsonData)
}

// 将go结构体转换为json字符串
func TestGoStructToJsonString(t *testing.T) {
	// 创建一个结构体
	user := User{
		Name: "zhangsan",
		Age:  18,
	}
	// 将结构体编码为json字符串
	jsonStr, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("jsonStr: %s", string(jsonStr))
}
