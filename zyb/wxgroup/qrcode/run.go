package qrcode

import (
	"fmt"
	"go-example/zyb/base"
)

func GetGroupQrCode() {
	//获取群码
	// {
	// 	"wxId": "ww7809deb62d31506e_wkwx_qw230824000156",
	// 	"bizRequestId": "qkstrategy_f684276aworkGetQRf074workGetQRcbedworkGetQR020bworkGetQRbba0514d1c4b",
	// 	"bizLine": "laxinqk",
	// 	"actionType": 150014,
	// 	"actionContent": {
	// 		"chatId": "10735270342700561",
	// 		"pt": "{\"wxGroupId\":\"10735270342700561\",\"wxId\":\"ww7809deb62d31506e_wkwx_qw230824000156\"}"
	// 	}
	// }

	// 定义请求参数
	actionType := 150014
	wxId := "wwf68f93e668aae3b3_wkwx_qw240202000700"
	chatId := "10855763463089606"
	bizRequestId := "qkstrategy_testGetQR_20241107_qw230824000156_03"
	actionContent := map[string]interface{}{
		"chatId": chatId,
		"pt":     fmt.Sprintf("{\"wxGroupId\":\"%s\",\"wxId\":\"%s\"}", chatId, wxId),
	}

	ret, err := base.RunAction(wxId, bizRequestId, actionType, actionContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ret)
}
