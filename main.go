package main

import (
	"fmt"
	"go-example/zyb/asset"
	"go-example/zyb/wxgroup"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go resetRecoverTask")
		return
	}

	action := args[1]
	if action == "resetRecoverTask" {
		wxgroup.ResetRecoverTask()
	}

	if action == "updateGroupClassify" {
		// baseDir, _ := os.Getwd()
		// file1 := "tmp/0605群分组调整/待分组0611.xlsx"
		// absFilePath := fmt.Sprintf("%s/%s", baseDir, file1)
		// wxgroup.UpdateGroupClassifyByFile(absFilePath, "Sheet1", 20, 16, 18, 19)
		// return

		// mapRel := map[int]map[string]int{
		// 	5556: {
		// 		"classify1Id": 9168,
		// 		"classify2Id": 9276,
		// 	},
		// }

		// for sourceGroupId, descGroupInfo := range mapRel {
		// 	fmt.Printf("sourceGroupId: %d, classify1Id: %d, classify2Id: %d \n", sourceGroupId, descGroupInfo["classify1Id"], descGroupInfo["classify2Id"])
		// 	wxgroup.UpdateGroupClassifyByClassify2Id(sourceGroupId, descGroupInfo["classify1Id"], descGroupInfo["classify2Id"], 0)
		// }

		wxGroupIds := []string{"10756098601840071", "10968288448278894", "10832418706942157", "10850168083104519", "10731039140716344", "10714143174597433", "10897044612963516", "10709626505137375", "10708815563872304", "10708023478749816", "10876890678705062", "10702617742627386", "10713279763821409", "10836489784424480", "10810911586155304", "10863092771596132", "10838994963566381", "10750426033727086", "10753819130593305", "10795117583498109", "10915976549202282", "10836669909973730"}
		wxgroup.UpdateGroupClassifyForBatch(wxGroupIds, 9242, 9243, 0)
	}

	if action == "exportAsset" {
		baseDir, _ := os.Getwd()
		fileName := "tmp/0618-个人号全量资产信息.xlsx"
		absFilePath := fmt.Sprintf("%s/%s", baseDir, fileName)
		asset.ExportUserListToExcel(absFilePath)
	}

	// 触发企微弹窗
	if action == "triggerWindow" {
		wxgroup.TriggerWindow()
	}

}
