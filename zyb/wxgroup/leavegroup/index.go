package leavegroup

import (
	"encoding/csv"
	"fmt"
	"go-example/zyb/base"
	"os"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
)

var leaveGroupFile = "tmp/wxgroup/leavegroup/retry批次退群slave群明细.csv"
var resultFile = "tmp/wxgroup/leavegroup/result.csv"
var wxGroupIdsMap map[string][]string
var existsWxGroupIdsMap map[string][]string
var wxChannel = make(chan string, 100)
var wg sync.WaitGroup
var workerNum = 30

func initExistsWxGroupIdsMap() {
	existsWxGroupIdsMap = make(map[string][]string)
	file, err := os.Open(resultFile)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}

	defer file.Close()
	// 解析 CSV 文件
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("无法解析文件:", err)
		return
	}

	for _, row := range records {
		// if index == 0 {
		// 	continue
		// }

		wxId := row[0]
		wxGroupId := row[1]
		if _, ok := existsWxGroupIdsMap[wxId]; !ok {
			existsWxGroupIdsMap[wxId] = []string{}
		}
		existsWxGroupIdsMap[wxId] = append(existsWxGroupIdsMap[wxId], wxGroupId)
	}

	fmt.Println("existsWxGroupIdsMap:", len(existsWxGroupIdsMap))
}

// 退群
func LeaveGroup() {
	// 初始化已退群的群ID
	initExistsWxGroupIdsMap()

	// 初始化待退群机器人信息
	// 打开 CSV 文件
	file, err := os.Open(leaveGroupFile)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
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
	wxGroupIdsMap = make(map[string][]string)
	for index, row := range records {
		if index == 0 {
			continue
		}

		wxId := row[0]
		wxGroupId := row[1]

		// 判断是否已经处理
		if wxGroupIds, ok := existsWxGroupIdsMap[wxId]; ok {
			isExists := false
			for _, existsWxGroupId := range wxGroupIds {
				if existsWxGroupId == wxGroupId {
					isExists = true
					break
				}
			}

			if isExists {
				fmt.Printf("wxId[%s] wxGroupId[%s] 已退群，跳过\n", wxId, wxGroupId)
				continue
			}
		}

		if _, ok := wxGroupIdsMap[wxId]; !ok {
			wxGroupIdsMap[wxId] = []string{}
		}
		wxGroupIdsMap[wxId] = append(wxGroupIdsMap[wxId], wxGroupId)
	}

	// 启动协程处理
	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go leaveGroup(i)
	}

	fmt.Println("wxGroupIdsMap:", len(wxGroupIdsMap))
	//创建协程，按机器人进行并发退群处理
	for wxId := range wxGroupIdsMap {
		fmt.Printf("push wxId[%s] to chan \n", wxId)
		wxChannel <- wxId
		time.Sleep(1 * time.Second)
	}

	close(wxChannel)
	wg.Wait()
}

func leaveGroup(i int) {
	fmt.Println("start leaveGroup goroutine ", i)

	defer wg.Done()
	for wxId := range wxChannel {
		wxGroupIds := wxGroupIdsMap[wxId]
		for _, wxGroupId := range wxGroupIds {
			time.Sleep(15 * time.Second)
			ret, err := base.RunAction(wxId, fmt.Sprintf("leavegroup-system-%s", uuid.NewV4().String()), 150010, map[string]interface{}{
				"chatId": wxGroupId,
				// "corpId": "wx230802001384",
				// "userId": "userId",
			})
			if err != nil {
				fmt.Println("退群请求发送失败，err:", err)
				continue
			}

			// 获取 kpErr 字段的值
			kpErrJson, ok := ret["kpErr"].(string)
			if !ok {
				fmt.Println("kpErr is not a string")
				return
			}

			// 使用 gjson 提取 errNo 和 errStr
			errNo := gjson.Get(kpErrJson, "errNo").Int()
			errStr := gjson.Get(kpErrJson, "errStr").String()

			appendRecord(resultFile, []string{wxId, wxGroupId, cast.ToString(errNo), errStr})

			// 检查 errNo 是否为 0
			if errNo != 0 {
				fmt.Printf("退群失败，wxId[%s] wxGroupId[%s] errNo[%d] errStr[%s] \n", wxId, wxGroupId, errNo, errStr)
				continue
			}

			fmt.Printf("退群成功, wxId[%s] wxGroupId[%s] \n", wxId, wxGroupId)
		}
	}
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
