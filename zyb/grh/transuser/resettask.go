package transuser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ResetTransUserTask() {
	// 重置转化用户任务
	for {
		ret := request()
		if ret == "" {
			break
		}
	}
}

func request() string {
	// 目标URL
	url := "https://wxtools.zuoyebang.cc/wxwork/tool/main?op=resetTask"

	// 设置请求头
	headers := map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7",
		"Cache-Control":             "no-cache",
		"Content-Type":              "multipart/form-data; boundary=----WebKitFormBoundaryu72ZetdZjLDdS2sh",
		"Cookie":                    "uid=zhangxueren; Hm_lvt_c33960c712441eec1b994580263ccb1a=1721738711; fp=b1185116131254e0f9d7cca202f46803; RANGERS_WEB_ID=5d80359c-cb85-4c27-82d0-45da799bfe41; RANGERS_SAMPLE=0.4995041547977497; ZYBIPSCAS=IPS_148ed96d3ef99c0a58ca780fec79604f1729395310; ZYBIPSUN=7a68616e6778756572656e",
		"Origin":                    "https://wxtools.zuoyebang.cc",
		"Pragma":                    "no-cache",
		"Priority":                  "u=0, i",
		"Referer":                   "https://wxtools.zuoyebang.cc/wxwork/tool/main?op=resetTask",
		"Sec-Ch-Ua":                 `"Chromium";v="130", "Google Chrome";v="130", "Not?A_Brand";v="99"`,
		"Sec-Ch-Ua-Mobile":          "?0",
		"Sec-Ch-Ua-Platform":        `"macOS"`,
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "same-origin",
		"Sec-Fetch-User":            "?1",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
	}

	// 构建表单数据
	formData := `------WebKitFormBoundaryu72ZetdZjLDdS2sh
Content-Disposition: form-data; name="menuId"

53
------WebKitFormBoundaryu72ZetdZjLDdS2sh
Content-Disposition: form-data; name="referUrl"

------WebKitFormBoundaryu72ZetdZjLDdS2sh
Content-Disposition: form-data; name="taskIds"

------WebKitFormBoundaryu72ZetdZjLDdS2sh
Content-Disposition: form-data; name="page"

form
------WebKitFormBoundaryu72ZetdZjLDdS2sh
Content-Disposition: form-data; name="status"

4
------WebKitFormBoundaryu72ZetdZjLDdS2sh
Content-Disposition: form-data; name="maxTransDate"

2024-10-12
------WebKitFormBoundaryu72ZetdZjLDdS2sh
Content-Disposition: form-data; name="limit"

10000
------WebKitFormBoundaryu72ZetdZjLDdS2sh
Content-Disposition: form-data; name="page"

form
------WebKitFormBoundaryu72ZetdZjLDdS2sh--`

	body := bytes.NewBufferString(formData)

	// 创建HTTP请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// 添加请求头
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resp.Body.Close()

	// 处理响应
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return ""
	}

	fmt.Println("Response status:", resp.Status)
	if strings.Contains(string(respBody), "3000条") {
		fmt.Println("ResetTransUserTask success 3000条")
		return "3000条"
	} else {
		fmt.Println("ResetTransUserTask failed")
	}

	return ""
}
