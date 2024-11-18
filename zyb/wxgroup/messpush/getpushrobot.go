package messpush

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
)

var wgLock sync.WaitGroup
var fileItemCnt int
var groupIdsFileName = "tmp/wxgroup/messpush/内容群-groupIds.txt"
var outputFileName = "tmp/wxgroup/messpush/群消息推送机器人明细-内容群.csv"
var groupIds []string

func ExportPushRobots() {
	InitFileInfo()
	trunks := splitSlice(groupIds, 30000)
	fmt.Println("分片数量：", len(trunks), " 已导出数量:", fileItemCnt, " 剩余未导出数量：", len(groupIds))

	if fileItemCnt == 0 {
		appendRecord(outputFileName, []string{"群ID", "推送机器人", "角色", "所属战队", "isAdmin", "isOwner"})
	}

	for _, trunk := range trunks {
		wgLock.Add(1)
		go process(trunk)
	}

	wgLock.Wait()
}

func InitFileInfo() {
	// 打开 CSV 文件
	file, err := os.Open(outputFileName)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		// 创建一个文件
		file, err = os.Create(outputFileName)
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

	content, err := os.ReadFile(groupIdsFileName)
	if err != nil {
		fmt.Printf("无法读取文件 %s: %v\n", groupIdsFileName, err)
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

		groupIds = append(groupIds, item)
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
	defer wgLock.Done()
	for _, wxGroupId := range groupIds {
		wxGroupDetail, _ := getPushRobotDetail(wxGroupId)
		// {"群ID", "推送机器人", "角色", "所属战队", "isAdmin", "isOwner"})
		strRow := fmt.Sprintf("%s,%s,%d,%d,%d,%d", wxGroupId+"_", wxGroupDetail.WxId, wxGroupDetail.Role, wxGroupDetail.QkTeamId, wxGroupDetail.IsAdmin, wxGroupDetail.IsOwner)
		arrRow := strings.Split(strRow, ",")
		appendRecord(outputFileName, arrRow)
		fmt.Println(strRow)

	}
}

type SearchGroupRobotResp struct {
	Data SearchGroupRobotRespData `json:"data"`
}

// 定义一个结构体来匹配JSON数据
type SearchGroupRobotRespData struct {
	RobotInfo []SearchGroupRobotRespDataItem `json:"robotInfo"`
}

type SearchGroupRobotRespDataItem struct {
	IsAdmin  int    `json:"isAdmin"`
	IsOwner  int    `json:"isOwner"`
	Role     int    `json:"role"`
	WxId     string `json:"wxId"`
	QkTeamId int    `json:"qkTeamId"`
	IsOnline int    `json:"isOnline"`
}

func getPushRobotDetail(wxGroupId string) (detail SearchGroupRobotRespDataItem, err error) {

	url := "https://wxtools.zuoyebang.cc/wxgc/groupmember/searchgrouprobot?from=messagepush&wxGroupId=" + wxGroupId

	req, err := http.NewRequest("GET", url, nil)
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

	var respData SearchGroupRobotResp
	if err = json.Unmarshal(body, &respData); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return detail, fmt.Errorf(err.Error())
	}

	for _, item := range respData.Data.RobotInfo {
		return item, nil
	}
	return
}
