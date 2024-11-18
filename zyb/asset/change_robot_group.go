package asset

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-example/zyb"
	"io/ioutil"
	"net/http"
)

func ChangeRobotGroup() {
	fmt.Println("开始迁移机器人")
	isFinish := false

	config := []struct {
		CorpId        string
		SourceGroupId int
		TargetGroupId int
		Limit         int
	}{
		// {"ww7809deb62d31506e", 516, 299, 10},
		// {"ww7809deb62d31506e", 299, 240, 10},
		// {"ww7809deb62d31506e", 240, 241, 10},
		// {"ww7809deb62d31506e", 241, 242, 10},
		// {"ww7809deb62d31506e", 495, 344, 10},
		// {"ww7809deb62d31506e", 344, 345, 10},
		// {"ww7809deb62d31506e", 345, 346, 10},
		// {"ww7809deb62d31506e", 497, 498, 10},
		// {"ww7809deb62d31506e", 498, 499, 10},
		// {"ww7809deb62d31506e", 518, 519, 10},
		// {"ww7809deb62d31506e", 519, 520, 10},
		// {"ww7809deb62d31506e", 554, 555, 10},
		// {"ww7809deb62d31506e", 555, 556, 10},
		// {"ww82a47fd1d43f9bbb", 703, 598, 10},
		// {"ww82a47fd1d43f9bbb", 598, 508, 10},
		// {"ww82a47fd1d43f9bbb", 508, 186, 10},
		// {"ww82a47fd1d43f9bbb", 186, 187, 10},
		// {"ww82a47fd1d43f9bbb", 302, 303, 10},
		// {"ww82a47fd1d43f9bbb", 303, 304, 10},
		// {"ww82a47fd1d43f9bbb", 373, 374, 10},
		// {"ww82a47fd1d43f9bbb", 402, 403, 10},
		// {"ww82a47fd1d43f9bbb", 403, 404, 10},
		// {"ww82a47fd1d43f9bbb", 466, 467, 10},
		// {"ww82a47fd1d43f9bbb", 467, 468, 10},
		// {"ww82a47fd1d43f9bbb", 463, 464, 10},
		// {"ww82a47fd1d43f9bbb", 464, 465, 10},
		// {"wwb5f29f00ef255b16", 187, 188, 10},
		// {"wwb5f29f00ef255b16", 191, 192, 10},
		// {"wwb5f29f00ef255b16", 195, 196, 10},
		// {"wwb5f29f00ef255b16", 196, 197, 10},
		// {"wwb5f29f00ef255b16", 1239, 386, 10},
		// {"wwb5f29f00ef255b16", 386, 387, 10},
		// {"wwb5f29f00ef255b16", 387, 388, 10},
		// {"wwb5f29f00ef255b16", 1089, 410, 10},
		// {"wwb5f29f00ef255b16", 410, 1088, 10},
		// {"wwb5f29f00ef255b16", 650, 651, 10},
		// {"wwb5f29f00ef255b16", 651, 652, 10},
		// {"wwb5f29f00ef255b16", 653, 654, 10},
		// {"wwb5f29f00ef255b16", 654, 655, 10},
		// {"wwb5f29f00ef255b16", 658, 659, 10},
		// {"wwb5f29f00ef255b16", 659, 660, 10},
		// {"wwb5f29f00ef255b16", 722, 723, 10},
		// {"wwb5f29f00ef255b16", 723, 724, 10},
		// {"wwb5f29f00ef255b16", 1040, 775, 10},
		// {"wwb5f29f00ef255b16", 775, 776, 10},
		// {"wwb5f29f00ef255b16", 776, 777, 10},
		// {"wwb5f29f00ef255b16", 779, 780, 10},
		// {"wwb5f29f00ef255b16", 780, 781, 10},
		// {"wwb5f29f00ef255b16", 818, 819, 10},
		// {"wwb5f29f00ef255b16", 819, 820, 10},
		// {"wwb5f29f00ef255b16", 991, 992, 10},
		// {"wwb5f29f00ef255b16", 992, 993, 10},
		// {"wwb5f29f00ef255b16", 814, 815, 10},
		// {"wwb5f29f00ef255b16", 815, 816, 10},
		// {"wwb5f29f00ef255b16", 1078, 1079, 10},
		// {"wwb5f29f00ef255b16", 1079, 1080, 10},
		// {"wwb5f29f00ef255b16", 1149, 1150, 10},
		// {"wwb5f29f00ef255b16", 1150, 1151, 10},
		// {"wwb5f29f00ef255b16", 1186, 1187, 10},
		// {"wwb5f29f00ef255b16", 1187, 1188, 10},
		// {"ww1ed06a6af05bc9a8", 426, 427, 10},
		// {"ww1ed06a6af05bc9a8", 427, 428, 10},
		// {"ww1ed06a6af05bc9a8", 326, 327, 10},
		// {"ww1ed06a6af05bc9a8", 329, 330, 10},
		// {"ww1ed06a6af05bc9a8", 512, 509, 10},
		// {"ww1ed06a6af05bc9a8", 321, 322, 10},
		// {"ww1ed06a6af05bc9a8", 513, 322, 10},
		// {"ww1ed06a6af05bc9a8", 310, 311, 10},
		// {"ww1ed06a6af05bc9a8", 311, 312, 10},
		// {"ww1ed06a6af05bc9a8", 296, 297, 10},
		// {"ww1ed06a6af05bc9a8", 431, 432, 10},
		// {"ww1ed06a6af05bc9a8", 432, 433, 10},
		// {"ww1ed06a6af05bc9a8", 433, 434, 10},
		// {"ww7809deb62d31506e", 142, 143, 10},
		// {"ww7809deb62d31506e", 143, 144, 10},
		// {"ww7809deb62d31506e", 144, 145, 10},
		// {"ww7809deb62d31506e", 145, 146, 10},
		// {"ww7809deb62d31506e", 146, 147, 10},
		// {"ww7809deb62d31506e", 151, 152, 10},
		// {"ww7809deb62d31506e", 152, 153, 10},
		// {"ww7809deb62d31506e", 153, 154, 10},
		// {"ww7809deb62d31506e", 374, 375, 10},
		// {"ww7809deb62d31506e", 375, 376, 10},
		// {"ww7809deb62d31506e", 376, 377, 10},
		// {"ww7809deb62d31506e", 377, 378, 10},
		// {"ww7809deb62d31506e", 378, 379, 10},
		// {"wwb5f29f00ef255b16", 179, 180, 10},
		// {"wwb5f29f00ef255b16", 180, 181, 10},
		// {"wwb5f29f00ef255b16", 181, 182, 10},
		// {"wwb5f29f00ef255b16", 182, 183, 10},
		// {"wwb5f29f00ef255b16", 183, 202, 10},
		// {"wwb5f29f00ef255b16", 636, 637, 10},
		// {"wwb5f29f00ef255b16", 637, 638, 10},
		// {"wwb5f29f00ef255b16", 638, 639, 10},
		// {"wwb5f29f00ef255b16", 639, 640, 10},
		// {"wwb5f29f00ef255b16", 640, 641, 10},
		// {"wwb5f29f00ef255b16", 207, 208, 10},
		// {"wwb5f29f00ef255b16", 208, 209, 10},
		// {"wwb5f29f00ef255b16", 209, 210, 10},
		// {"wwb5f29f00ef255b16", 210, 211, 10},
		// {"wwb5f29f00ef255b16", 211, 212, 10},
		// {"wwb5f29f00ef255b16", 888, 889, 10},
		// {"wwb5f29f00ef255b16", 889, 890, 10},
		// {"wwb5f29f00ef255b16", 890, 891, 10},
		// {"wwb5f29f00ef255b16", 891, 892, 10},
		// {"wwb5f29f00ef255b16", 892, 893, 10},
		// {"wwb5f29f00ef255b16", 304, 305, 10},
		// {"wwb5f29f00ef255b16", 305, 306, 10},
		// {"wwb5f29f00ef255b16", 306, 307, 10},
		// {"wwb5f29f00ef255b16", 307, 308, 10},
		// {"wwb5f29f00ef255b16", 308, 309, 10},
		// {"wwb5f29f00ef255b16", 669, 670, 10},
		// {"wwb5f29f00ef255b16", 670, 671, 10},
		// {"wwb5f29f00ef255b16", 671, 672, 10},
		// {"wwb5f29f00ef255b16", 672, 673, 10},
		// {"wwb5f29f00ef255b16", 673, 674, 10},
		// {"wwb5f29f00ef255b16", 844, 845, 10},
		// {"wwb5f29f00ef255b16", 845, 846, 10},
		// {"wwb5f29f00ef255b16", 846, 847, 10},
		// {"wwb5f29f00ef255b16", 847, 848, 10},
		// {"wwb5f29f00ef255b16", 848, 849, 10},
		// {"wwb5f29f00ef255b16", 851, 852, 10},
		// {"wwb5f29f00ef255b16", 852, 853, 10},
		// {"wwb5f29f00ef255b16", 853, 854, 10},
		// {"wwb5f29f00ef255b16", 854, 855, 10},
		// {"wwb5f29f00ef255b16", 855, 856, 10},
		// {"wwb5f29f00ef255b16", 858, 859, 10},
		// {"wwb5f29f00ef255b16", 859, 860, 10},
		// {"wwb5f29f00ef255b16", 860, 861, 10},
		// {"wwb5f29f00ef255b16", 861, 862, 10},
		// {"wwb5f29f00ef255b16", 862, 863, 10},
		// {"ww82a47fd1d43f9bbb", 609, 191, 10},
		// {"ww82a47fd1d43f9bbb", 191, 192, 10},
		// {"ww82a47fd1d43f9bbb", 192, 193, 10},
		// {"ww82a47fd1d43f9bbb", 193, 194, 10},
		// {"ww82a47fd1d43f9bbb", 194, 195, 10},
		// {"ww82a47fd1d43f9bbb", 348, 349, 10},
		// {"ww82a47fd1d43f9bbb", 349, 350, 10},
		// {"ww82a47fd1d43f9bbb", 350, 351, 10},
		// {"ww82a47fd1d43f9bbb", 351, 352, 10},
		// {"ww82a47fd1d43f9bbb", 352, 353, 10},
		// {"ww82a47fd1d43f9bbb", 407, 408, 10},
		// {"ww82a47fd1d43f9bbb", 408, 409, 10},
		// {"ww82a47fd1d43f9bbb", 409, 410, 10},
		// {"ww82a47fd1d43f9bbb", 410, 411, 10},
		// {"ww82a47fd1d43f9bbb", 411, 412, 10},
		// {"ww82a47fd1d43f9bbb", 435, 436, 10},
		// {"ww82a47fd1d43f9bbb", 436, 437, 10},
		// {"ww82a47fd1d43f9bbb", 437, 438, 10},
		// {"ww82a47fd1d43f9bbb", 438, 439, 10},
		// {"ww82a47fd1d43f9bbb", 439, 440, 10},

		//好课
		{"wwd583450d95c39236", 161, 162, 10},
		{"wwd583450d95c39236", 160, 161, 10},
		{"wwd583450d95c39236", 146, 147, 10},
		{"wwd583450d95c39236", 145, 146, 10},
		{"wwd583450d95c39236", 114, 115, 10},
		{"wwd583450d95c39236", 103, 104, 10},
		{"wwd583450d95c39236", 102, 103, 10},
		{"wwd583450d95c39236", 220, 221, 10},
		{"wwd583450d95c39236", 91, 92, 10},
		{"wwd583450d95c39236", 90, 91, 10},
		{"wwd583450d95c39236", 75, 76, 10},
		{"wwd583450d95c39236", 74, 75, 10},
		{"wwd583450d95c39236", 72, 73, 10},
		{"wwd583450d95c39236", 71, 72, 10},
		{"wwd583450d95c39236", 69, 70, 10},
		{"wwd583450d95c39236", 68, 69, 10},
		{"wwd583450d95c39236", 32, 33, 10},
		{"wwd583450d95c39236", 31, 32, 10},
		{"wwd583450d95c39236", 194, 31, 10},
		{"wwd583450d95c39236", 53, 54, 10},
		{"wwd583450d95c39236", 52, 53, 10},
		{"wwd583450d95c39236", 191, 52, 10},

		//济南
		{"ww1ed06a6af05bc9a8", 433, 434, 10},
		{"ww1ed06a6af05bc9a8", 432, 433, 10},

		//西安
		{"wwb5f29f00ef255b16", 1329, 812, 10},
		{"wwb5f29f00ef255b16", 811, 812, 10},
		{"wwb5f29f00ef255b16", 810, 811, 10},
		{"wwb5f29f00ef255b16", 804, 872, 10},
		{"wwb5f29f00ef255b16", 711, 712, 10},
		{"wwb5f29f00ef255b16", 686, 687, 10},
		{"wwb5f29f00ef255b16", 300, 301, 10},
		{"wwb5f29f00ef255b16", 830, 831, 10},
		{"wwb5f29f00ef255b16", 256, 257, 10},
		{"wwb5f29f00ef255b16", 255, 256, 10},
		{"wwb5f29f00ef255b16", 1233, 255, 10},
		{"wwb5f29f00ef255b16", 247, 248, 10},
		{"wwb5f29f00ef255b16", 246, 247, 10},
		{"wwb5f29f00ef255b16", 243, 244, 10},
		{"wwb5f29f00ef255b16", 242, 243, 10},
		{"wwb5f29f00ef255b16", 239, 240, 10},
		{"wwb5f29f00ef255b16", 238, 239, 10},
		{"wwb5f29f00ef255b16", 235, 236, 10},
		{"wwb5f29f00ef255b16", 234, 235, 10},
		{"wwb5f29f00ef255b16", 230, 231, 10},

		//合肥
		{"ww82a47fd1d43f9bbb", 423, 424, 10},
		{"ww82a47fd1d43f9bbb", 608, 203, 10},
		{"ww82a47fd1d43f9bbb", 600, 203, 10},
		{"ww82a47fd1d43f9bbb", 507, 203, 10},
		{"ww82a47fd1d43f9bbb", 201, 203, 10},

		//北京
		{"ww7809deb62d31506e", 213, 215, 10},
		{"ww7809deb62d31506e", 214, 215, 10},
	}

	//revConfig := helper.ReverseSlice(config)

	for _, c := range config {
		isFinish = false
		for !isFinish {
			isFinish = doChangeRobotGroup(c.CorpId, c.SourceGroupId, c.TargetGroupId, c.Limit)
		}

		fmt.Println("迁移机器人完成, corpId: ", c.CorpId, "sourceGroupId: ", c.SourceGroupId, "targetGroupId: ", c.TargetGroupId)
	}

}

func doChangeRobotGroup(corpId string, sourceGroupId int, targetGroupId int, limit int) (isFinish bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			isFinish = true
		}
	}()

	url := "https://wxtools.zuoyebang.cc/wxwork/tool/main"

	// Prepare the request body as JSON
	requestBody := map[string]interface{}{
		"corpId":        corpId,
		"sourceGroupId": sourceGroupId,
		"targetGroupId": targetGroupId,
		"limit":         limit,
		"renderType":    "json",
		"op":            "changeRobotGroup",
		"page":          "form",
		"menuId":        2,
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

	fmt.Println(string(body))

	var r map[string]interface{}

	err = json.Unmarshal(body, &r)
	if err != nil {
		panic(err)
	}

	data, ok := r["data"].(map[string]interface{})
	if !ok {
		fmt.Println("Error: 'data' field not found or incorrect type")
		return false
	}

	ret, ok := data["ret"].(string)
	if !ok {
		return false
	}

	// Check if the operation is finished
	if ret == "没有需要迁移的机器人" {
		return true
	}

	return false
}
