package outgroup

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go-example/helper"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var GetGroupDetailWg sync.WaitGroup
var fileItemCnt int
var fileName = "tmp/群明细导出/外部群明细导出-1006.csv"
var IdsfileName = "tmp/群明细导出/外部群id.txt"
var GetGroupDetailGroupIds []string

func GetGroupDetailInfo() {
	InitFileInfo()
	trunks := splitSlice(GetGroupDetailGroupIds, 30000)
	fmt.Println("分片数量：", len(trunks), " 已导出数量:", fileItemCnt, " 剩余未导出数量：", len(GetGroupDetailGroupIds))

	if fileItemCnt == 0 {
		appendRecord(fileName, []string{"群ID", "群名称", "主体名称", "建群时间", "在群人数", "群活动", "四级分类", "五级分类", "六级分类", "所属战队分组", "mars", "slave", "slave个数", "mars状态", "slave状态", "群主角色"})
	}

	for _, trunk := range trunks {
		GetGroupDetailWg.Add(1)
		go process(trunk)
	}

	GetGroupDetailWg.Wait()
}

func InitFileInfo() {
	// 打开 CSV 文件
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		// 创建一个文件
		file, err = os.Create(fileName)
		if err != nil {
			// 如果创建文件失败，打印错误信息并退出程序
			panic(err)
		}
	}
	defer file.Close()

	// 解析 CSV 文件
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("无法解析文件:", err)
		return
	}

	// 遍历 CSV 记录并将 id 字段存储到 seenIDs map 中
	existIdsMap := make(map[string]bool, 1)
	for _, row := range records {
		// 假设每行的第一个元素是 id
		if len(row) > 0 {
			fileItemCnt++
			cleanedWxGroupId := strings.Replace(row[0], "_", "", -1)
			existIdsMap[cleanedWxGroupId] = true
		}
	}

	content, err := os.ReadFile(IdsfileName)
	if err != nil {
		fmt.Printf("无法读取文件 %s: %v\n", IdsfileName, err)
		return
	}

	// 将内容转换为字符串
	strContent := string(content)

	// 使用逗号作为分隔符分割字符串
	items := strings.Split(strContent, "\n")
	for _, item := range items {
		if _, ok := existIdsMap[item]; ok {
			continue
		}

		GetGroupDetailGroupIds = append(GetGroupDetailGroupIds, item)
	}
}

// splitSlice 将字符串切片按照指定大小分割成多个子切片。
// 如果无法平均分割，则最后一个子切片可能包含较少的元素。
func splitSlice(slice []string, chunkSize int) [][]string {
	var result [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		result = append(result, slice[i:end])
	}
	return result
}

func appendRecord(fileName string, record []string) error {
	// 打开文件以追加模式
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建 CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入新的记录
	if err := writer.Write(record); err != nil {
		return err
	}

	return nil
}

func process(groupIds []string) {
	defer GetGroupDetailWg.Done()
	for _, wxGroupId := range groupIds {
		wxGroupDetail, _ := getGroupMemberDetail(wxGroupId)
		strRow := fmt.Sprintf("%s,%s,%s,%s,%d,%s,%s,%s,%s,%s,%s,%s,%d,%s,%s,%s", wxGroupId+"_", wxGroupDetail.WxGroupName, wxGroupDetail.CorpName, wxGroupDetail.CreateDate, wxGroupDetail.GroupCnt, wxGroupDetail.ActivityName, wxGroupDetail.FourLevelName, wxGroupDetail.FiveLevelName, wxGroupDetail.SixLevelName, wxGroupDetail.OwnerTeamGroupName, wxGroupDetail.MasterWxIds, wxGroupDetail.SlaveWxIds, wxGroupDetail.SlaveCnt, wxGroupDetail.MasterIsOnline, wxGroupDetail.SlaveIsOnline, wxGroupDetail.OwnerRole)
		arrRow := strings.Split(strRow, ",")
		appendRecord(fileName, arrRow)
		fmt.Println(strRow)

	}
}

type MemberDetail struct {
	OwnerRole    string `json:"ownerRole"`    //群主角色
	WxGroupName  string `json:"wxGroupName"`  //群名称
	CorpName     string `json:"corpName"`     //主体名称
	ActivityName string `json:"activityName"` //活动名称

	CreateDate    string `json:"createDate"`    //创建时间
	GroupCnt      int    `json:"groupCnt"`      //群人数
	FourLevelName string `json:"fourLevelName"` //四级分类
	FiveLevelName string `json:"fiveLevelName"` //五级分类
	SixLevelName  string `json:"sixLevelName"`  //六级分类

	OwnerTeamId        interface{} `json:"ownerTeamId"`        //群主战队ID
	OwnerTeamGroupName string      `json:"ownerTeamGroupName"` //群主战队分组名称
	MasterWxIds        string      `json:"masterWxIds"`        //群主微信号
	MasterIsOnline     string      `json:"masterIsOnline"`     //群主是否在线
	SlaveWxIds         string      `json:"slaveWxIds"`         //副群主微信号
	SlaveCnt           int         `json:"slaveCnt"`           //slave个数
	SlaveIsOnline      string      `json:"slaveIsOnline"`      //副群主是否在线

}

type GroupMemberResp struct {
	Data GroupMemberRespData `json:"data"`
}

// 定义一个结构体来匹配JSON数据
type GroupMemberRespData struct {
	List []struct {
		IsAdmin  int    `json:"isAdmin"`
		IsOwner  int    `json:"isOwner"`
		Role     int    `json:"role"`
		WxId     string `json:"wxId"`
		RoleName string `json:"roleName"`
		IsDelete int    `json:"isDelete"`
	} `json:"list"`
}

func getGroupMemberDetail(wxGroupId string) (detail MemberDetail, err error) {

	url := "https://wxtools.zuoyebang.cc/wxqk/groupmessage/memberlist"
	params := "?page_size=10&page_no=1&pageSize=10&pageNo=1&pn=1&curPage=1&rn=10&total=4&activityId=17676&groupId=" + wxGroupId + "&label=%5Bnull%5D"

	req, err := http.NewRequest("GET", url+params, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return detail, fmt.Errorf(err.Error())
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", "HMACCOUNT=19A5FA5D6642C585; uid=zhangxueren; ZYBIPSUN=7a68616e6778756572656e; Hm_lvt_c33960c712441eec1b994580263ccb1a=1721738711; Hm_lpvt_c33960c712441eec1b994580263ccb1a=1723450262; fp=b1185116131254e0f9d7cca202f46803; ZYBIPSCAS="+helper.ZYBIPSCAS+"; RANGERS_WEB_ID=5d80359c-cb85-4c27-82d0-45da799bfe41; RANGERS_SAMPLE=0.4995041547977497;")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Google Chrome\";v=\"128\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return detail, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return detail, fmt.Errorf(err.Error())
	}

	var respData GroupMemberResp
	if err = json.Unmarshal(body, &respData); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return detail, fmt.Errorf(err.Error())
	}

	for _, item := range respData.Data.List {
		if item.IsDelete == 1 {
			continue
		}
		if item.IsOwner == 1 {
			detail.OwnerRole = item.RoleName
		}

		wxGroupDetail, _ := GetWxGroupDetail(wxGroupId)
		if wxGroupDetail.WxGroupId != "" {
			// 假设有一个时间戳
			timestamp := int64(wxGroupDetail.CreateTime) // 例如：2021-10-01 00:00:00 UTC
			// 将时间戳转换为time.Time类型
			t := time.Unix(timestamp, 0)
			// 将time.Time类型格式化为字符串
			timeString := t.Format("2006-01-02 15:04:05")
			detail.WxGroupName = wxGroupDetail.WxGroupName
			detail.CorpName = wxGroupDetail.CorpName
			detail.CreateDate = timeString
			detail.GroupCnt = wxGroupDetail.MemberCnt
			detail.ActivityName = wxGroupDetail.ActTypeDesc
			detail.FourLevelName = wxGroupDetail.RecoverTypeName
			detail.FiveLevelName = wxGroupDetail.GroupAttributeName
			detail.SixLevelName = wxGroupDetail.GradeLabel
		}

		if item.RoleName == "Mars" {
			robotDetail, err := getRobotDetail(item.WxId)
			if err != nil {
				detail.MasterIsOnline = err.Error()
			} else {
				detail.MasterIsOnline = robotDetail.IsOnlineAlias
				if item.IsOwner == 1 {
					detail.OwnerTeamGroupName = robotDetail.TeamName
					detail.OwnerTeamId = robotDetail.TeamId
				}
			}

			detail.MasterWxIds += item.WxId + "|"
		}

		if item.RoleName == "SLAVE" {
			robotDetail, err := getRobotDetail(item.WxId)
			if err != nil {
				detail.SlaveIsOnline = err.Error()
			} else {
				detail.SlaveIsOnline = robotDetail.IsOnlineAlias
				if item.IsOwner == 1 {
					detail.OwnerTeamGroupName = robotDetail.TeamName
					detail.OwnerTeamId = robotDetail.TeamId
				}
			}

			detail.SlaveWxIds += item.WxId + "|"
			detail.SlaveCnt += 1
		}
	}

	return detail, nil
}

type RobotDetail struct {
	TeamName      string      `json:"teamName"`
	IsOnlineAlias string      `json:"isOnlineAlias"`
	TeamId        interface{} `json:"teamId"`
	QkTeamId      int         `json:"qkTeamId"`
}

func getRobotDetail(wxId string) (detail RobotDetail, err error) {
	url := "https://wxtools.zuoyebang.cc/wxqk/robot/robotqueryes"
	params := "?page_size=10&page_no=1&pageSize=10&pageNo=1&pn=1&curPage=1&rn=10&total=10000&wxName=" + wxId + "&isWorkWx=1"

	req, err := http.NewRequest("GET", url+params, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return detail, fmt.Errorf(err.Error())
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", "HMACCOUNT=19A5FA5D6642C585; uid=zhangxueren; ZYBIPSUN=7a68616e6778756572656e; Hm_lvt_c33960c712441eec1b994580263ccb1a=1721738711; Hm_lpvt_c33960c712441eec1b994580263ccb1a=1723450262; fp=b1185116131254e0f9d7cca202f46803; ZYBIPSCAS="+helper.ZYBIPSCAS+"; RANGERS_WEB_ID=5d80359c-cb85-4c27-82d0-45da799bfe41; RANGERS_SAMPLE=0.4995041547977497;")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Google Chrome\";v=\"128\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return detail, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return detail, fmt.Errorf(err.Error())
	}

	var respData struct {
		Data struct {
			List []RobotDetail
		} `json:"data"`
	}
	if err = json.Unmarshal(body, &respData); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return detail, fmt.Errorf(err.Error())
	}

	for _, item := range respData.Data.List {
		return item, nil
	}

	return detail, fmt.Errorf("not found robotinfo")
}

type wxGroupDetail struct {
	WxGroupId          string `json:"wxGroupId"`
	WxGroupName        string `json:"wxGroupName"`
	CreateTime         int    `json:"createTime"`
	CorpName           string `json:"corpName"`
	MemberCnt          int    `json:"memberCnt"`
	ActTypeDesc        string `json:"actTypeDesc"`
	RecoverTypeName    string `json:"recoverTypeName"`
	GroupAttributeName string `json:"groupAttributeName"`
	GradeLabel         string `json:"gradeLabel"`
}

func GetWxGroupDetail(wxGroupId string) (detail wxGroupDetail, err error) {
	// url := "https://wxtools.zuoyebang.cc/wxqk/group/searchwxgrouplist"
	// params := "?page_size=20&page_no=1&pageSize=20&pageNo=1&pn=1&curPage=1&rn=20&total=0&wxGroupId=" + wxGroupId + "&isClassify=-1&label=&memberCntMin=0&memberCntMax=500&ownerWxId=&isJoin=-1&wxGroupName=&activityIsUsable=-1&productType=2&groupAttribute=&riskType=&actType=-1&startTime=0&endTime=0&corpId=&innerSpeakTimeStart=0&innerSpeakTimeEnd=0&outerSpeakTimeStart=0&outerSpeakTimeEnd=0&isHitDismissRule=-1&dismissRuleId=&dismissDateStart=&dismissDateEnd=&isDataV2=1&notAuth=1"
	url := "https://wxtools.zuoyebang.cc/wxqk/group/groupbackes?page_size=20&page_no=1&pageSize=20&pageNo=1&pn=1&curPage=1&rn=20&total=0&wxGroupName=&wxGroupId=" + wxGroupId + "&memberCntMin=0&memberCntMax=500&acceptTimeStart=&acceptTimeEnd=&lastLessonTimeStart=&lastLessonTimeEnd=&isClassify=-1&status[]=60&opIsIn=-1&corpId=&innerSpeakTimeStart=0&innerSpeakTimeEnd=0&outerSpeakTimeStart=0&outerSpeakTimeEnd=0&isHitDismissRule=-1&dismissRuleId=&dismissDateStart=0&dismissDateEnd=0&isDataV2=1&notAuth=1"
	params := ""
	req, err := http.NewRequest("GET", url+params, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return detail, fmt.Errorf(err.Error())
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", "HMACCOUNT=19A5FA5D6642C585; uid=zhangxueren; ZYBIPSUN=7a68616e6778756572656e; Hm_lvt_c33960c712441eec1b994580263ccb1a=1721738711; Hm_lpvt_c33960c712441eec1b994580263ccb1a=1723450262; fp=b1185116131254e0f9d7cca202f46803; ZYBIPSCAS="+helper.ZYBIPSCAS+"; RANGERS_WEB_ID=5d80359c-cb85-4c27-82d0-45da799bfe41; RANGERS_SAMPLE=0.4995041547977497;")
	req.Header.Set("Pragma", "no-cache")
	// req.Header.Set("Priority", "u=1, i")
	// req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Google Chrome\";v=\"128\"")
	// req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	// req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	// req.Header.Set("Sec-Fetch-Dest", "empty")
	// req.Header.Set("Sec-Fetch-Mode", "cors")
	// req.Header.Set("Sec-Fetch-Site", "same-origin")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return detail, fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return detail, fmt.Errorf(err.Error())
	}

	var respData struct {
		Data struct {
			List []wxGroupDetail
		} `json:"data"`
	}
	if err = json.Unmarshal(body, &respData); err != nil {
		fmt.Println(err.Error())
		return detail, fmt.Errorf(err.Error())
	}

	for _, item := range respData.Data.List {
		return item, nil
	}

	return detail, fmt.Errorf("not found wxgroupdetail")
}
