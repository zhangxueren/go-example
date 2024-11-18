package wxgroup

import (
	"fmt"
	"go-example/helper"

	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

func AnalyseInherit(filePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	retrySuccWxGroupIdMap := make(map[string]map[int]int)
	for i, row := range rows {
		if len(row) < 4 {
			fmt.Printf("数据不全跳过，index[%d] columnNum[%d] rowInfo[%v] \n", i+1, len(row), row)
			continue
		}
		_, robotWxId, status, _ := row[0], row[1], row[2], row[3]
		if retrySuccWxGroupIdMap[robotWxId] == nil {
			retrySuccWxGroupIdMap[robotWxId] = make(map[int]int)
		}
		retrySuccWxGroupIdMap[robotWxId][cast.ToInt(status)] = 1
		//fmt.Println("wxGroupId:", wxGroupId, "robotWxId:", robotWxId, "status:", status, "createTime:", createTime)
	}

	content := ""
	for robotWxId, statusMap := range retrySuccWxGroupIdMap {
		// if len(statusMap) == 2 {
		content += fmt.Sprintf("%s\n", robotWxId)
		fmt.Println("robotWxId:", robotWxId, "statusMap:", statusMap)
		// }
	}
	helper.WriteFile("tmp/successInheritRobotWxId.txt", content)
}
