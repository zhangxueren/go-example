package wxgroup

import (
	"fmt"
	"net/http"
	"strings"
)

func ResetRecoverTask() {
	groupIds := ""
	groupIdList := strings.Split(groupIds, ",")
	for _, groupId := range groupIdList {
		resetRecoverTask(groupId)
	}

}

func resetRecoverTask(groupId string) {
	url := "https://wxtools.zuoyebang.cc/wxqk-go/tools/resetrecovergrouptask?wxGroupId=" + groupId

	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", "RANGERS_WEB_ID=6430599a-1243-4e7d-9d30-4de92cf0df82; RANGERS_SAMPLE=0.07367398890598964; uid=zhangxueren; ZYBIPSCAS=IPS_1d59147c193f2fc2a3e3909dae0dca171715597013; Hm_lvt_c33960c712441eec1b994580263ccb1a=1715756176; Hm_lpvt_c33960c712441eec1b994580263ccb1a=1716189999; ZYBUSS=PwzWvo3l9NUKJIBEqGfFWPG56knpG0J71FHfn_eZZ5eDC8qqgojR-vJLf_QVeWqq")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("priority", "u=0, i")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"125\", \"Chromium\";v=\"125\", \"Not.A/Brand\";v=\"24\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	// 发送 HTTP 请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 输出响应结果
	fmt.Println("resetRecoverTask groupId:", groupId, " response status:", resp.Status)
}
