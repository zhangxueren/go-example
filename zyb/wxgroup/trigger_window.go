package wxgroup

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func TriggerWindow() {
	// wxIds1 := []string{"robotWxId:wwd583450d95c39236_wkwx_qw230906000192", "robotWxId:wwd583450d95c39236_wkwx_qw230906000655", "robotWxId:wwd583450d95c39236_wkwx_qw220602000200", "robotWxId:wwd583450d95c39236_wkwx_qw220811000513", "robotWxId:wwd583450d95c39236_wkwx_qw220808000471", "robotWxId:ww7809deb62d31506e_wkwx_qw240122000463", "robotWxId:wwd583450d95c39236_wkwx_qw220811000337", "robotWxId:wwd583450d95c39236_wkwx_qw220811000304", "robotWxId:ww7809deb62d31506e_wkwx_qw230802001292", "robotWxId:wwd583450d95c39236_wkwx_qw220811000390", "robotWxId:wwd583450d95c39236_wkwx_qw220811000553", "robotWxId:wwd583450d95c39236_wkwx_qw230906000669", "robotWxId:wwd583450d95c39236_wkwx_qw220811000373", "robotWxId:wwd583450d95c39236_wkwx_qw220602000167", "robotWxId:wwd583450d95c39236_wkwx_qw240122000755", "robotWxId:wwd583450d95c39236_wkwx_qw240202000638", "robotWxId:wwb5f29f00ef255b16_wkwx_qw210727000059", "robotWxId:wwd583450d95c39236_wkwx_qw230906000711", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000143", "robotWxId:ww7809deb62d31506e_wkwx_qw230904000448", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000178", "robotWxId:wwd583450d95c39236_wkwx_qw230830000773", "robotWxId:wwd583450d95c39236_wkwx_qw230906000701", "robotWxId:wwd583450d95c39236_wkwx_qw240202000616", "robotWxId:ww7809deb62d31506e_wkwx_qw231226000108", "robotWxId:wwd583450d95c39236_wkwx_qw230516000046", "robotWxId:ww7809deb62d31506e_wkwx_qw220529000118", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001341", "robotWxId:wwd583450d95c39236_wkwx_qw230906000391", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001338", "robotWxId:ww7809deb62d31506e_wkwx_BJ1-202106170732", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000139", "robotWxId:ww7809deb62d31506e_wkwx_qw230808000905", "robotWxId:ww7809deb62d31506e_wkwx_qw230829000127", "robotWxId:wwd583450d95c39236_wkwx_qw240122000749", "robotWxId:ww7809deb62d31506e_wkwx_qw230808000871", "robotWxId:wwd583450d95c39236_wkwx_qw230605000252", "robotWxId:ww7809deb62d31506e_wkwx_qw240416000417", "robotWxId:ww7809deb62d31506e_wkwx_qw230810000085", "robotWxId:wwd583450d95c39236_wkwx_qw220602000184", "robotWxId:ww7809deb62d31506e_wkwx_qw230802001432", "robotWxId:wwd583450d95c39236_wkwx_qw220811000340", "robotWxId:wwd583450d95c39236_wkwx_qw240122000760", "robotWxId:wwb5f29f00ef255b16_wkwx_qw210727000060", "robotWxId:ww7809deb62d31506e_wkwx_qw231226000108", "robotWxId:wwd583450d95c39236_wkwx_qw230906000655", "robotWxId:wwd583450d95c39236_wkwx_qw220811000340", "robotWxId:ww7809deb62d31506e_wkwx_qw230828000559", "robotWxId:wwd583450d95c39236_wkwx_qw240122000760", "robotWxId:wwb5f29f00ef255b16_wkwx_qw210727000059", "robotWxId:ww7809deb62d31506e_wkwx_qw230810000085", "robotWxId:wwd583450d95c39236_wkwx_qw240122000749", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001338", "robotWxId:wwd583450d95c39236_wkwx_qw230906000391", "robotWxId:wwd583450d95c39236_wkwx_qw220811000390", "robotWxId:wwd583450d95c39236_wkwx_qw220808000471", "robotWxId:wwd583450d95c39236_wkwx_qw220811000489", "robotWxId:wwd583450d95c39236_wkwx_qw220811000373", "robotWxId:wwd583450d95c39236_wkwx_qw230906000669", "robotWxId:wwd583450d95c39236_wkwx_qw220811000553", "robotWxId:wwb5f29f00ef255b16_wkwx_qw211019000140", "robotWxId:wwd583450d95c39236_wkwx_qw230516000046", "robotWxId:ww7809deb62d31506e_wkwx_qw230802001292", "robotWxId:wwd583450d95c39236_wkwx_qw220602000167", "robotWxId:wwd583450d95c39236_wkwx_qw220811000513", "robotWxId:wwd583450d95c39236_wkwx_qw230906000192", "robotWxId:ww7809deb62d31506e_wkwx_qw240122000463", "robotWxId:wwb5f29f00ef255b16_wkwx_qw211018001346", "robotWxId:ww7809deb62d31506e_wkwx_qw230904000448", "robotWxId:wwb5f29f00ef255b16_wkwx_qw221115000328", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001357", "robotWxId:wwd583450d95c39236_wkwx_qw240122000755", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001341", "robotWxId:wwd583450d95c39236_wkwx_qw220811000337", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000193", "robotWxId:ww7809deb62d31506e_wkwx_qw230802001432", "robotWxId:wwd583450d95c39236_wkwx_qw230906000711", "robotWxId:ww7809deb62d31506e_wkwx_qw230111000307", "robotWxId:wwd583450d95c39236_wkwx_qw220602000207", "robotWxId:ww7809deb62d31506e_wkwx_qw230808000866", "robotWxId:wwd583450d95c39236_wkwx_qw230830000773", "robotWxId:wwd583450d95c39236_wkwx_qw220602000200", "robotWxId:ww7809deb62d31506e_wkwx_qw230901000652", "robotWxId:ww7809deb62d31506e_wkwx_qw230824000155", "robotWxId:ww7809deb62d31506e_wkwx_qw230808000865", "robotWxId:wwd583450d95c39236_wkwx_qw240202000616", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000171", "robotWxId:wwd583450d95c39236_wkwx_qw220811000304", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001460", "robotWxId:ww7809deb62d31506e_wkwx_BJ1-202106170732", "robotWxId:ww7809deb62d31506e_wkwx_qw230808000871", "robotWxId:wwd583450d95c39236_wkwx_qw230906000701", "robotWxId:ww7809deb62d31506e_wkwx_qw240202000415", "robotWxId:ww7809deb62d31506e_wkwx_qw220529000126", "robotWxId:wwd583450d95c39236_wkwx_qw230516000046", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001341", "robotWxId:wwd583450d95c39236_wkwx_qw230906000711", "robotWxId:ww7809deb62d31506e_wkwx_qw231226000108", "robotWxId:wwb5f29f00ef255b16_wkwx_qw221115000328", "robotWxId:wwd583450d95c39236_wkwx_qw220602000184", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001338", "robotWxId:wwd583450d95c39236_wkwx_qw230906000655", "robotWxId:ww7809deb62d31506e_wkwx_BJ1-202106170732", "robotWxId:wwd583450d95c39236_wkwx_qw220602000167", "robotWxId:ww7809deb62d31506e_wkwx_qw230808000865", "robotWxId:wwd583450d95c39236_wkwx_qw240122000755", "robotWxId:wwd583450d95c39236_wkwx_qw220811000489", "robotWxId:ww7809deb62d31506e_wkwx_qw230808000871", "robotWxId:wwd583450d95c39236_wkwx_qw220602000200", "robotWxId:wwd583450d95c39236_wkwx_qw220811000373", "robotWxId:wwb5f29f00ef255b16_wkwx_qw210629001547", "robotWxId:wwd583450d95c39236_wkwx_qw230906000701", "robotWxId:wwd583450d95c39236_wkwx_qw230906000669", "robotWxId:wwd583450d95c39236_wkwx_qw230830000773", "robotWxId:wwd583450d95c39236_wkwx_qw220811000340", "robotWxId:wwd583450d95c39236_wkwx_qw220811000513", "robotWxId:wwd583450d95c39236_wkwx_qw220808000471", "robotWxId:ww7809deb62d31506e_wkwx_qw240122000463", "robotWxId:wwd583450d95c39236_wkwx_qw220602000207", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000140", "robotWxId:wwb5f29f00ef255b16_wkwx_qw211101000079", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000132", "robotWxId:ww7809deb62d31506e_wkwx_qw230802001292", "robotWxId:wwd583450d95c39236_wkwx_qw220811000304", "robotWxId:ww7809deb62d31506e_wkwx_qw240202000415", "robotWxId:ww7809deb62d31506e_wkwx_qw230904000448", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000201", "robotWxId:wwd583450d95c39236_wkwx_qw220811000390", "robotWxId:ww7809deb62d31506e_wkwx_qw230802001432", "robotWxId:wwd583450d95c39236_wkwx_qw230906000192", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000125", "robotWxId:ww7809deb62d31506e_wkwx_qw230904000421", "robotWxId:ww7809deb62d31506e_wkwx_qw230810000085", "robotWxId:wwd583450d95c39236_wkwx_qw230906000391", "robotWxId:wwd583450d95c39236_wkwx_qw240122000749"}
	// wxIds2 := []string{"robotWxId:ww7809deb62d31506e_wkwx_qw230802001292", "robotWxId:wwd583450d95c39236_wkwx_qw220811000373", "robotWxId:ww7809deb62d31506e_wkwx_qw230908000252", "robotWxId:ww7809deb62d31506e_wkwx_qw230907000295", "robotWxId:wwd583450d95c39236_wkwx_qw230830000773", "robotWxId:wwd583450d95c39236_wkwx_qw220602000167", "robotWxId:ww7809deb62d31506e_wkwx_BJ1-202106170732", "robotWxId:wwb5f29f00ef255b16_wkwx_qw211104000388", "robotWxId:wwd583450d95c39236_wkwx_qw240122000760", "robotWxId:ww7809deb62d31506e_wkwx_qw240117000191", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001378", "robotWxId:wwd583450d95c39236_wkwx_qw220811000304", "robotWxId:ww7809deb62d31506e_wkwx_qw231226000108", "robotWxId:wwd583450d95c39236_wkwx_qw220811000340", "robotWxId:wwbd04cf92469dbec9_wkwx_qw240204000009", "robotWxId:wwd583450d95c39236_wkwx_qw240122000749", "robotWxId:wwf68f93e668aae3b3_wkwx_qw211210000218", "robotWxId:ww7809deb62d31506e_wkwx_qw230810000085", "robotWxId:wwd583450d95c39236_wkwx_qw230906000192", "robotWxId:ww7809deb62d31506e_wkwx_qw240202000415", "robotWxId:wwd583450d95c39236_wkwx_qw220811000513", "robotWxId:wwd583450d95c39236_wkwx_qw230516000046", "robotWxId:ww7809deb62d31506e_wkwx_qw230808000865", "robotWxId:wwd583450d95c39236_wkwx_qw240202000638", "robotWxId:wwd583450d95c39236_wkwx_qw220602000207", "robotWxId:wwd583450d95c39236_wkwx_qw230906000669", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001508", "robotWxId:ww7809deb62d31506e_wkwx_qw230904000448", "robotWxId:ww7809deb62d31506e_wkwx_qw240117000112", "robotWxId:wwd583450d95c39236_wkwx_qw220602000200", "robotWxId:wwd583450d95c39236_wkwx_qw230906000655", "robotWxId:wwd583450d95c39236_wkwx_qw220811000390", "robotWxId:wwd583450d95c39236_wkwx_qw220602000184", "robotWxId:wwd583450d95c39236_wkwx_qw230906000701", "robotWxId:ww7809deb62d31506e_wkwx_qw240122000463", "robotWxId:wwb5f29f00ef255b16_wkwx_qw230110000191", "robotWxId:wwd583450d95c39236_wkwx_qw240122000755", "robotWxId:ww7809deb62d31506e_wkwx_qw230808001460", "robotWxId:wwd583450d95c39236_wkwx_qw240202000616", "robotWxId:ww7809deb62d31506e_wkwx_qw230802001432", "robotWxId:wwd583450d95c39236_wkwx_qw220808000471", "robotWxId:wwd583450d95c39236_wkwx_qw230906000391", "robotWxId:wwd583450d95c39236_wkwx_qw220706000050", "robotWxId:wwb5f29f00ef255b16_wkwx_qw221115000328", "robotWxId:wwd583450d95c39236_wkwx_qw220811000489"}
	// wxIds3 := []string{"robotWxId:wwd583450d95c39236_wkwx_qw220811000390", "robotWxId:wwd583450d95c39236_wkwx_qw230516000046", "robotWxId:wwd583450d95c39236_wkwx_qw220808000471", "robotWxId:ww7809deb62d31506e_wkwx_qw230810000085", "robotWxId:wwd583450d95c39236_wkwx_qw220811000340", "robotWxId:wwd583450d95c39236_wkwx_qw230605000252", "robotWxId:ww7809deb62d31506e_wkwx_qw230808000865", "robotWxId:wwd583450d95c39236_wkwx_qw230830000773", "robotWxId:wwb5f29f00ef255b16_wkwx_qw211104000388", "robotWxId:wwd583450d95c39236_wkwx_qw230906000655", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000171", "robotWxId:ww7809deb62d31506e_wkwx_qw231226000108", "robotWxId:wwd583450d95c39236_wkwx_qw220602000184", "robotWxId:wwb5f29f00ef255b16_wkwx_qw221115000328", "robotWxId:ww7809deb62d31506e_wkwx_qw230802001478", "robotWxId:wwd583450d95c39236_wkwx_qw240122000749", "robotWxId:ww7809deb62d31506e_wkwx_qw230904000448", "robotWxId:wwd583450d95c39236_wkwx_qw230906000669", "robotWxId:ww7809deb62d31506e_wkwx_qw240202000415", "robotWxId:ww7809deb62d31506e_wkwx_BJ1-202106170732", "robotWxId:wwd583450d95c39236_wkwx_qw220811000373", "robotWxId:wwd583450d95c39236_wkwx_qw220811000304", "robotWxId:ww7809deb62d31506e_wkwx_qw230828000559", "robotWxId:wwd583450d95c39236_wkwx_qw220602000207", "robotWxId:ww7809deb62d31506e_wkwx_qw230802001292", "robotWxId:wwd583450d95c39236_wkwx_qw220602000200", "robotWxId:wwd583450d95c39236_wkwx_qw230906000701", "robotWxId:wwd583450d95c39236_wkwx_qw240122000755"}

	doneMap := make(map[string]bool)
	// wxIds := append(wxIds1, wxIds2...)
	// wxIds = append(wxIds, wxIds3...)

	wxIds := []string{"robotWxId:ww7809deb62d31506e_wkwx_qw230810000085", "robotWxId:wwd583450d95c39236_wkwx_qw220602000184", "robotWxId:wwd583450d95c39236_wkwx_qw230605000252", "robotWxId:ww7809deb62d31506e_wkwx_qw240122000463", "robotWxId:ww7809deb62d31506e_wkwx_qw230802001432", "robotWxId:wwb5f29f00ef255b16_wkwx_qw211019000310", "robotWxId:wwd583450d95c39236_wkwx_qw230906000655", "robotWxId:wwd583450d95c39236_wkwx_qw230906000192", "robotWxId:wwd583450d95c39236_wkwx_qw220602000207", "robotWxId:ww7809deb62d31506e_wkwx_qw231226000108", "robotWxId:ww7809deb62d31506e_wkwx_qw230823000171", "robotWxId:ww7809deb62d31506e_wkwx_BJ1-202106170732", "robotWxId:wwf68f93e668aae3b3_wkwx_qw240102000021", "robotWxId:wwd583450d95c39236_wkwx_qw220811000489", "robotWxId:wwb5f29f00ef255b16_wkwx_qw221115000328", "robotWxId:ww7809deb62d31506e_wkwx_qw230824000200", "robotWxId:ww7809deb62d31506e_wkwx_qw230824000147", "robotWxId:wwd583450d95c39236_wkwx_qw220811000390", "robotWxId:ww7809deb62d31506e_wkwx_qw230904000448", "robotWxId:wwd583450d95c39236_wkwx_qw230906000391", "robotWxId:ww7809deb62d31506e_wkwx_qw240117000116", "robotWxId:wwd583450d95c39236_wkwx_qw220808000471", "robotWxId:wwd583450d95c39236_wkwx_qw220811000373", "robotWxId:wwd583450d95c39236_wkwx_qw230906000701", "robotWxId:wwd583450d95c39236_wkwx_qw230516000046"}
	for _, strWxId := range wxIds {
		if doneMap[strWxId] {
			fmt.Println("wxId has done:", strWxId)
			continue
		}

		temp := strings.Split(strWxId, ":")
		wxIdInfo := strings.Split(temp[1], "_wkwx_")

		if len(wxIdInfo) != 2 {
			fmt.Println("wxIdInfo error:", wxIdInfo)
			continue
		}
		requestTriggerWindow(wxIdInfo[1], wxIdInfo[0])
		doneMap[strWxId] = true
	}

}

func requestTriggerWindow(workUserId string, corpId string) {
	url := fmt.Sprintf("https://wxtools.zuoyebang.cc/wxqk-go/toolsv2/triggerwindow?workUserId=%s&corpId=%s", workUserId, corpId)

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
	req.Header.Set("cookie", "RANGERS_WEB_ID=6430599a-1243-4e7d-9d30-4de92cf0df82; RANGERS_SAMPLE=0.07367398890598964; uid=zhangxueren; ZYBIPSCAS=IPS_2026743183a82dfe845f08dbf482a2101717555104; Hm_lvt_c33960c712441eec1b994580263ccb1a=1715756176; Hm_lpvt_c33960c712441eec1b994580263ccb1a=1716189999; ZYBUSS=PwzWvo3l9NUKJIBEqGfFWPG56knpG0J71FHfn_eZZ5eDC8qqgojR-vJLf_QVeWqq; __tips__=1")
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

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 输出响应结果
	fmt.Println("requestTriggerWindow wxId:", workUserId, " response body:", string(body))
}
