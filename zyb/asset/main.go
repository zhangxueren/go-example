package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type response struct {
	ErrNo  int    `json:"errNo"`
	ErrStr string `json:"errStr"`
	Data   struct {
		Meta struct {
			Total    int `json:"total"`
			PageNo   int `json:"pageNo"`
			PageSize int `json:"pageSize"`
		} `json:"meta"`
		List []struct {
			FlowStatus          string `json:"flowStatus"`
			WorkID              int    `json:"workId"`
			WorkUserID          string `json:"workUserId"`
			WorkName            string `json:"workName"`
			Avatar              string `json:"avatar"`
			CreateTime          string `json:"createTime"`
			IsDeleted           int    `json:"isDeleted"`
			IsOnline            int    `json:"isOnline"`
			KpStatus            int    `json:"kpStatus"`
			LogoutTime          string `json:"logoutTime"`
			FriendNum           int    `json:"friendNum"`
			BatchID             int    `json:"batchId"`
			GreetWordID         int    `json:"greetWordId"`
			GreetWordName       string `json:"greetWordName"`
			BatchName           string `json:"batchName"`
			ReplyNum            int    `json:"replyNum"`
			CorpName            string `json:"corpName"`
			CorpID              string `json:"corpId"`
			QrcodeCnt           int    `json:"qrcodeCnt"`
			WaitingQrcodeStatus int    `json:"waitingQrcodeStatus"`
			AppName             string `json:"appName"`
			RobotStatus         string `json:"robotStatus"`
			GroupName           string `json:"groupName"`
			OnTheShelfStr       string `json:"onTheShelfStr"`
			OnTheShelf          int    `json:"onTheShelf"`
			Gender              string `json:"gender"`
			NoSingleFriendNum   int    `json:"noSingleFriendNum"`
			OpGroupID           int    `json:"opGroupId"`
			OpGroupName         string `json:"opGroupIdName"`
			AcField             int    `json:"acField"`
			AcScene             int    `json:"acScene"`
			AcSceneSegmented    int    `json:"acSceneSegmented"`
			AcRecoverType       int    `json:"acRecoverType"`
			AcPersona           int    `json:"acPersona"`
			AcPersonaName       string `json:"acPersonaName"`
			AcGrade             int    `json:"acGrade"`
			AcGradeName         string `json:"acGradeName"`
		} `json:"list"`
		WorkNum struct {
			Offline       int `json:"offline"`
			Online        int `json:"online"`
			Total         int `json:"total"`
			UnBind        int `json:"unBind"`
			LimitRobotCnt int `json:"limitRobotCnt"`
		} `json:"workNum"`
	} `json:"data"`
	LogID string `json:"logId"`
}

func main() {

	file, err := os.Create("/Users/shawn/studyspace/golang/go-example/output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the header to the file
	header := "workId,workUserId,workName,friendNum,noSingleFriendNum,corpName,corpId\n"
	_, err = file.WriteString(header)
	if err != nil {
		panic(err)
	}

	page := 1
	maxPages := 56
	fmt.Println("workId,workUserId,workName,friendNum,noSingleFriendNum,corpName,corpId")
	for {
		fetch(page, file)
		page++
		if page > maxPages {
			break
		}

		// time.Sleep(5 * time.Second)
	}
}

func fetch(page int, file *os.File) {
	url := "https://wxtools.zuoyebang.cc/wxwork/user/workeruserlist"

	// Prepare the request body as JSON
	requestBody := map[string]interface{}{
		"isOnline":            "-1",
		"onTheShelf":          "-1",
		"batchId":             "",
		"workName":            "",
		"workUserId":          "",
		"corpId":              "",
		"machineCode":         "",
		"searchType":          "workName",
		"accessDate":          "",
		"appId":               "",
		"robotStatus":         "-1",
		"searchGroupId":       "",
		"corpGroupIdRelation": map[string]interface{}{},
		"classified":          0,
		"opGroupId":           []int{10015, 10016, 10017, 10018, 10019, 10020, 10062, 10114},
		"acRecoverType":       []string{""},
		"acPersona":           []string{""},
		"acGrade":             []string{""},
		"page_size":           50,
		"page_no":             page, // This will be incremented for each page
		"pageNo":              page,
		"pageSize":            50,
		"pn":                  page,
		"rn":                  50,
		"beginTime":           "",
		"stopTime":            "",
	}

	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	// Create the HTTP client
	client := &http.Client{}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		panic(err)
	}

	// Set headers
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "RANGERS_WEB_ID=6430599a-1243-4e7d-9d30-4de92cf0df82; RANGERS_SAMPLE=0.07367398890598964; uid=zhangxueren; ZYBUSS=PwzWvo3l9NUKJIBEqGfFWPG56knpG0J71FHfn25rH5eZGJKMrZmIiqJtAIwYJGHX; ZYBIPSCAS=IPS_1d59147c193f2fc2a3e3909dae0dca171715597013; Hm_lvt_c33960c712441eec1b994580263ccb1a=1715756176; __tips__=1; Hm_lpvt_c33960c712441eec1b994580263ccb1a=1716189999")
	req.Header.Set("origin", "https://wxtools.zuoyebang.cc")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	//fmt.Printf("Fetching page %d...\n", page)

	// Update the page_no in the request body
	requestBody["page_no"] = page
	requestBody["pn"] = page

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshal the response body into a struct
	var r response
	err = json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	// 打印list
	// 输出打印头

	for _, v := range r.Data.List {
		fmt.Printf("%d, %s, %s, %d, %d, %s, %s\n", v.WorkID, v.WorkUserID, v.WorkName, v.FriendNum, v.NoSingleFriendNum, v.CorpName, v.CorpID)
		row := fmt.Sprintf("%d,%s,%s,%d,%d,%s,%s\n", v.WorkID, v.WorkUserID, v.WorkName, v.FriendNum, v.NoSingleFriendNum, v.CorpName, v.CorpID)
		_, _ = file.WriteString(row)
	}
}
