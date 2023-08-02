// test/case1_test.go

package test

import (
	"bytes"
	"encoding/json"
	"my_test_framework/test"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// 定义要发送的JSON数据，模拟注册的用户信息
	jsonStr := `{}`

	// 创建一个POST请求，并将JSON数据作为请求体
	req, err := http.NewRequest("POST", "https://cmp-fe.ucloud.cn/api/gateway?Action=GetCaptcha", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// 设置请求的Content-Type为application/json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "gr_user_id=6a064cbf-90fd-455e-8be8-82fb82e1a064; das_session=84592504-430f-4083-b3a3-2f9b48768087; _hjSessionUser_1469719=eyJpZCI6ImI4ZjVhNTcyLTllZTEtNTIwNS05ZWU2LTc5Nzk1NzIxZWVhMyIsImNyZWF0ZWQiOjE2NzcwNTQxNDIwMTYsImV4aXN0aW5nIjpmYWxzZX0=; Hm_lvt_413fdc5943040809ed0703eabd01f173=1688818883,1690373331")

	// 创建一个模拟HTTP ResponseWriter，用于接收服务器响应
	rr := httptest.NewRecorder()

	// 在这里调用你的注册处理函数，将请求和响应记录下来
	// 注意：这里的代码需要依据你实际的服务器实现
	// 如：yourHandlerFunc(rr http.ResponseWriter, req *http.Request)

	// 在这里可以添加更多的断言，验证响应是否符合预期

	// 检查响应的状态码是否为200（成功）
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// 解析返回的JSON数据
	var responseData ResponseData
	err = json.NewDecoder(rr.Body).Decode(&responseData)
	if err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	// 断言 Message 字段的值为 "successful"
	expectedMessage := "Success"
	if responseData.Message != expectedMessage {
		t.Errorf("Expected 'Message' field to be '%s', but got '%s'", expectedMessage, responseData.Message)
	}

	// 最后，关闭请求和响应的连接
	req.Body.Close()
}
