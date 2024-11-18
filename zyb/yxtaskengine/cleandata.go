package yxtaskengine

import (
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

func CleanAddMemberCtrl() {
	url := "http://10.109.18.239:8099/yxtaskengine/tool/cleanaddmemberctrl?maxId=6898863&limit=1000&status=2"
	//url := "http://10.109.18.239:8099/yxtaskengine/tool/cleanaddmemberctrl?maxId=7830950&limit=1000&status=1"
	//url := "http://10.109.18.239:8099/yxtaskengine/tool/cleanaddmemberctrl?maxId=7830950&limit=1000&status=3"
	//url := "http://10.109.18.239:8099/yxtaskengine/tool/cleanaddmemberctrl?maxId=7830950&limit=1000&status=0"

	for i := 0; i < 10000; i++ {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error on request", i+1, ":", err)
			continue
		}
		defer resp.Body.Close()

		// 这里可以处理响应数据，例如打印状态码
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
		}

		fmt.Printf("Request %d, Status Code: %d count:%d\n", i+1, resp.StatusCode, len(gjson.Get(string(body), "data.remIds").Array()))

		// 避免请求过于频繁，可以添加适当的延迟
		//time.Sleep(1 * time.Second)
	}
}
