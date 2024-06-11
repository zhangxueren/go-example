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
		// file1 := "tmp/0605群分组调整/群信息-2024-05-09 15_29_07.xlsx"
		// absFilePath := fmt.Sprintf("%s/%s", baseDir, file1)
		// wxgroup.UpdateGroupClassifyByFile(absFilePath, 16, 0, 14, 15)
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
	}

	if action == "exportAsset" {
		baseDir, _ := os.Getwd()
		fileName := "tmp/5月份未触达用户资产信息.xlsx"
		absFilePath := fmt.Sprintf("%s/%s", baseDir, fileName)
		asset.ExportUserListToExcel(absFilePath)
	}

}
