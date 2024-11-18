package main

import (
	"context"
	"fmt"
	"go-example/helper"
	"go-example/helper/labs"
	"go-example/zyb/asset"
	"go-example/zyb/grh"
	"go-example/zyb/grh/transuser"
	"go-example/zyb/wxgroup"
	"go-example/zyb/wxgroup/leavegroup"
	"go-example/zyb/wxgroup/messpush"
	"go-example/zyb/wxgroup/outgroup"
	"go-example/zyb/wxgroup/qrcode"
	"go-example/zyb/yxcontent"
	"go-example/zyb/yxtaskengine"
	"os"
)

func main() {
	var ctx = context.Background()
	helper.Init(ctx)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go resetRecoverTask")
		return
	}

	action := args[1]

	if action == "EmbedStruct" {
		animal := labs.NewAnimal("dog", 1)
		fmt.Println(animal.Name, animal.Age)

	}

	if action == "ResetRecoverTask" {
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
		fileName := "tmp/1021-个人号全量资产信息.xlsx"
		absFilePath := fmt.Sprintf("%s/%s", baseDir, fileName)
		asset.ExportUserListToExcel(absFilePath)
	}

	// 触发企微弹窗
	if action == "triggerWindow" {
		wxgroup.TriggerWindow()
	}

	if action == "GenDocCapture" {
		yxcontent.GenDocCapture()
	}

	if action == "ChangeRobotGroup" {
		asset.ChangeRobotGroup()
	}

	// 迁移企微分组
	if action == "RecoverAssetGroups" {
		asset.RecoverAssetGroups()
	}

	// 群继承数据分析
	if action == "inherit" {
		baseDir, _ := os.Getwd()
		fileName := "tmp/6月份-督学群继承动作日志明细.xlsx"
		absFilePath := fmt.Sprintf("%s/%s", baseDir, fileName)
		wxgroup.AnalyseInherit(absFilePath)
	}

	// 退号任务
	if action == "ReturnRobot" {
		asset.ReturnRobot()
	}

	// 退号任务
	if action == "CheckReturnRobot" {
		asset.CheckReturnRobot()
	}

	if action == "ExportCourseList" {
		asset.ExportCourseList()
	}

	//刷新高一因升年级错配督学群接回失败任务
	if action == "ResetInvalidRecoverGroupRule" {
		wxgroup.ResetInvalidRecoverGroupRule()
	}

	// 刷新升年级错配课程信息
	if action == "RepairCourse" {
		wxgroup.RepairCourse()
	}

	if action == "ResetWxidTeamid" {
		asset.ResetWxidTeamid()
	}

	if action == "InheritTask" {
		asset.InheritTask()
	}

	if action == "GetCourseDetailList" {
		wxgroup.GetCourseDetailList()
	}

	if action == "CheckFinish" {
		asset.CheckFinish()
	}

	if action == "GetKpGroupList" {
		wxgroup.GetKpGroupList()
	}

	//yxtaskengine
	if action == "CleanAddMemberCtrl" {
		yxtaskengine.CleanAddMemberCtrl()
	}

	//解散群
	if action == "DismissGroup" {
		wxgroup.DismissGroup()
	}

	//获取群解散明细
	if action == "GetGroupDismissInfo" {
		wxgroup.GetGroupDismissInfo()
	}

	if action == "GetGroupDetailInfo" {
		outgroup.GetGroupDetailInfo()
	}

	if action == "ExportGroupList" {
		grh.ExportGroupList()
	}

	if action == "ExportPushRobots" {
		messpush.ExportPushRobots()
	}

	if action == "GetGroupQrCode" {
		qrcode.GetGroupQrCode()
	}

	if action == "ResetTransUserTask" {
		transuser.ResetTransUserTask()
	}

	//退群
	if action == "LeaveGroup" {
		leavegroup.LeaveGroup()
	}

}
