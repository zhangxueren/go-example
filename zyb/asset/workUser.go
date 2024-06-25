package asset

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-example/zyb"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

type workItem struct {
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
}
type response struct {
	ErrNo  int    `json:"errNo"`
	ErrStr string `json:"errStr"`
	Data   struct {
		Meta struct {
			Total    int `json:"total"`
			PageNo   int `json:"pageNo"`
			PageSize int `json:"pageSize"`
		} `json:"meta"`
		List    []workItem `json:"list"`
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

func ExportUserListToExcel(outputFilePath string) {
	workUserIds := []string{}
	// respJson := ``
	// var resp = make([]map[string]interface{}, 0)
	// _ = json.Unmarshal([]byte(respJson), &resp)
	// for _, v := range resp {
	// 	workUserIds = append(workUserIds, cast.ToString(v["workUserId"]))
	// }

	//fmt.Println("workUserId, workName, 业务线, 主体名称, 新分组, 在架状态, 双向好友数, 运营组, 人设, 年级")
	f := excelize.NewFile()

	// data := [][]string{
	// 	{"workUserId", "workName", "业务线", "主体名称", "新分组", "在架状态", "双向好友数", "运营组", "人设", "年级"},
	// }

	f.SetCellValue("Sheet1", "A1", "workUserId")
	f.SetCellValue("Sheet1", "B1", "workName")
	f.SetCellValue("Sheet1", "C1", "业务线")
	f.SetCellValue("Sheet1", "D1", "主体名称")
	f.SetCellValue("Sheet1", "E1", "新分组")
	f.SetCellValue("Sheet1", "F1", "在架状态")
	f.SetCellValue("Sheet1", "G1", "双向好友数")
	f.SetCellValue("Sheet1", "H1", "运营组")
	f.SetCellValue("Sheet1", "I1", "人设")
	f.SetCellValue("Sheet1", "J1", "年级")
	f.SetCellValue("Sheet1", "K1", "企微状态")
	f.SetCellValue("Sheet1", "L1", "接入时间")
	f.SetCellValue("Sheet1", "M1", "登出时间")

	if len(workUserIds) > 0 {
		for i, workUserId := range workUserIds {
			v := fetchWorkDetailByWorkUserId(workUserId)
			strIndex := cast.ToString(i + 2)
			robotStatus := ""
			if v.KpStatus == 0 {
				robotStatus = "正常"
			} else {
				robotStatus = "异常"
			}
			f.SetCellValue("Sheet1", "A"+strIndex, v.WorkUserID)
			f.SetCellValue("Sheet1", "B"+strIndex, v.WorkName)
			f.SetCellValue("Sheet1", "C"+strIndex, v.AppName)
			f.SetCellValue("Sheet1", "D"+strIndex, v.CorpName)
			f.SetCellValue("Sheet1", "E"+strIndex, v.GroupName)
			f.SetCellValue("Sheet1", "F"+strIndex, v.OnTheShelfStr)
			f.SetCellValue("Sheet1", "G"+strIndex, v.NoSingleFriendNum)
			f.SetCellValue("Sheet1", "H"+strIndex, v.OpGroupName)
			f.SetCellValue("Sheet1", "I"+strIndex, v.AcPersonaName)
			f.SetCellValue("Sheet1", "J"+strIndex, v.AcGradeName)
			f.SetCellValue("Sheet1", "K"+strIndex, robotStatus)
			f.SetCellValue("Sheet1", "L"+strIndex, v.CreateTime)
			f.SetCellValue("Sheet1", "M"+strIndex, v.LogoutTime)

			fmt.Println("index:"+strIndex, "appName:"+v.AppName)
		}
	} else { // fetch all
		pageNo := 1
		rowNumber := 2
		for {
			list := fetchWorkListByPage(pageNo)
			if len(list) == 0 {
				break
			}
			for _, v := range list {
				strIndex := cast.ToString(rowNumber)
				robotStatus := ""
				if v.KpStatus == 0 {
					robotStatus = "正常"
				} else {
					robotStatus = "异常"
				}
				f.SetCellValue("Sheet1", "A"+strIndex, v.WorkUserID)
				f.SetCellValue("Sheet1", "B"+strIndex, v.WorkName)
				f.SetCellValue("Sheet1", "C"+strIndex, v.AppName)
				f.SetCellValue("Sheet1", "D"+strIndex, v.CorpName)
				f.SetCellValue("Sheet1", "E"+strIndex, v.GroupName)
				f.SetCellValue("Sheet1", "F"+strIndex, v.OnTheShelfStr)
				f.SetCellValue("Sheet1", "G"+strIndex, v.NoSingleFriendNum)
				f.SetCellValue("Sheet1", "H"+strIndex, v.OpGroupName)
				f.SetCellValue("Sheet1", "I"+strIndex, v.AcPersonaName)
				f.SetCellValue("Sheet1", "J"+strIndex, v.AcGradeName)
				f.SetCellValue("Sheet1", "K"+strIndex, robotStatus)
				f.SetCellValue("Sheet1", "L"+strIndex, v.CreateTime)
				f.SetCellValue("Sheet1", "M"+strIndex, v.LogoutTime)

				fmt.Println("index:"+strIndex, "appName:"+v.AppName)
				rowNumber++
			}
			pageNo++
		}
	}

	if err := f.SaveAs(outputFilePath); err != nil {
		fmt.Println(err)
	}
}

func fetchWorkDetailByWorkUserId(workUserId string) workItem {
	url := "https://wxtools.zuoyebang.cc/wxwork/user/workeruserlist"

	// Prepare the request body as JSON
	requestBody := map[string]interface{}{
		"isOnline":            "-1",
		"onTheShelf":          "-1",
		"batchId":             "",
		"workName":            "",
		"workUserId":          workUserId,
		"corpId":              "",
		"machineCode":         "",
		"searchType":          "workName",
		"accessDate":          "",
		"appId":               "",
		"robotStatus":         "-1",
		"searchGroupId":       "",
		"corpGroupIdRelation": map[string]interface{}{},
		"classified":          0,
		"opGroupId":           []int{},
		"acRecoverType":       []string{""},
		"acPersona":           []string{""},
		"acGrade":             []string{""},
		"page_size":           100,
		"page_no":             1, // This will be incremented for each page
		"pageNo":              1,
		"pageSize":            100,
		"pn":                  1,
		"rn":                  100,
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
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "RANGERS_WEB_ID=6430599a-1243-4e7d-9d30-4de92cf0df82; RANGERS_SAMPLE=0.07367398890598964; uid=zhangxueren; ZYBUSS=PwzWvo3l9NUKJIBEqGfFWPG56knpG0J71FHfn25rH5eZGJKMrZmIiqJtAIwYJGHX; ZYBIPSCAS="+zyb.COOKIE+"; Hm_lvt_c33960c712441eec1b994580263ccb1a=1715756176;")

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

	for _, v := range r.Data.List {
		return v
	}
	return workItem{}
}

func fetchWorkListByPage(pageNo int) []workItem {
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
		"opGroupId":           []int{},
		"acRecoverType":       []string{""},
		"acPersona":           []string{""},
		"acGrade":             []string{""},
		"page_size":           100,
		"page_no":             pageNo, // This will be incremented for each page
		"pageNo":              pageNo,
		"pageSize":            100,
		"pn":                  pageNo,
		"rn":                  100,
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
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "RANGERS_WEB_ID=6430599a-1243-4e7d-9d30-4de92cf0df82; RANGERS_SAMPLE=0.07367398890598964; uid=zhangxueren; ZYBUSS=PwzWvo3l9NUKJIBEqGfFWPG56knpG0J71FHfn25rH5eZGJKMrZmIiqJtAIwYJGHX; ZYBIPSCAS="+zyb.COOKIE+"; Hm_lvt_c33960c712441eec1b994580263ccb1a=1715756176;")

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

	return r.Data.List
}
