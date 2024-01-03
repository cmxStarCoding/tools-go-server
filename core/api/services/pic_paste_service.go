package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"tools/core/api/validator/pic"
)

type PicPasteService struct {
}

func (s PicPasteService) DoTask(request *pic.Request, TaskId any) {

	var requestMap = make(map[string]any)

	// 将结构体转换为 JSON
	jsonData, err := json.Marshal(request)
	//将json结构赋予map
	mapErr := json.Unmarshal(jsonData, &requestMap)
	if mapErr != nil {
		gin.DefaultWriter.Write([]byte(fmt.Sprintf("转化map错误，错误原因%s", mapErr.Error())))
	}
	requestMap["batch_no"] = TaskId
	requestMap["notify_url"] = "http://127.0.0.1:8080/api/v1/pic_paste_notify"

	//将字典转化为json
	requestJson, _ := json.Marshal(requestMap)

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 准备 HTTP 请求
	url := "http://127.0.0.1:8003/qrcode-replace/replace"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJson))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头为 JSON
	req.Header.Set("Content-Type", "application/json")

	// 发送 HTTP 请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	fmt.Println("Response Status:", resp.Status)

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应体数据
	fmt.Println("Response Body:", string(body))

}
