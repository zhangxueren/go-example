package wxgroup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const COOKIE = "IPS_d702cc584b8a2bbd0d8f309436dabb851717554753"
const UpdateGroupClassifyUrl = "https://wxtools-wxwork1-cc.suanshubang.cc/wxqk/classify/updategroupclassify"
const GetGroupListUrlByClassify2Id = "https://wxtools-wxwork1-cc.suanshubang.cc/wxqk/classify/getgrouplist"

func UpdateGroupClassifyByClassify2Id(classify2Id int, classify1Id int, classity2Id int, autoGroupClassifyByGrade int) {
	// 定义请求地址和初始页码
	url := GetGroupListUrlByClassify2Id
	pageSize := 10
	pageNo := 1

	// 设置Cookie信息
	cookie := "RANGERS_WEB_ID=5aa89a84-50f9-42a2-a610-ae3532318a0e; RANGERS_SAMPLE=0.712393799821835; ZYBIPSCAS=" + COOKIE

	type Response struct {
		Data struct {
			List []struct {
				WxGroupId string `json:"wxGroupId"`
			} `json:"list"`
		} `json:"data"`
	}

	groupIds := make([]string, 0)
	// 发起HTTP请求，逐页获取数据直到获取完所有数据
	for {
		// 构建请求URL
		reqURL := fmt.Sprintf("%s?pageSize=%d&pageNo=%d&classify2Id=%d", url, pageSize, pageNo, classify2Id)

		// 创建HTTP请求
		req, _ := http.NewRequest("GET", reqURL, nil)

		// 设置请求Header中的Cookie信息
		req.Header.Set("Cookie", cookie)

		// 发起请求
		client := &http.Client{}
		resp, _ := client.Do(req)

		// 读取响应信息
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		var response Response
		err := json.Unmarshal(body, &response)
		if err != nil {
			log.Default().Println("Error unmarshal response:", err)
			continue
		}

		for _, group := range response.Data.List {
			groupIds = append(groupIds, group.WxGroupId)
		}

		// 假设一页数据不足pageSize条，表示已经获取全部信息
		if len(response.Data.List) < pageSize {
			break
		}

		// 更新页码
		pageNo++
	}

	fmt.Println("All groups data:")
	fmt.Println(groupIds)

	// 更新群分类
	UpdateGroupClassifyForBatch(groupIds, classify1Id, classity2Id, autoGroupClassifyByGrade)
}

func UpdateGroupClassifyForBatch(wxgroupIds []string, classify1Id int, classity2Id int, autoGroupClassifyByGrade int) {
	batchSize := 50
	batchGroupIds := make([]string, 0, batchSize)
	for i, wxgroupId := range wxgroupIds {
		batchGroupIds = append(batchGroupIds, wxgroupId)
		if (i+1)%batchSize == 0 {
			updateGroupClassify(batchGroupIds, classify1Id, classity2Id, autoGroupClassifyByGrade)
			batchGroupIds = make([]string, 0, batchSize)
		}
	}
	if len(batchGroupIds) > 0 {
		updateGroupClassify(batchGroupIds, classify1Id, classity2Id, autoGroupClassifyByGrade)
	}
}

func updateGroupClassify(wxgroupIds []string, classify1Id int, classity2Id int, autoGroupClassifyByGrade int) error {
	url := UpdateGroupClassifyUrl
	// 准备请求数据
	requestBody, _ := json.Marshal(map[string]interface{}{
		"autoGroupClassifyByGrade": autoGroupClassifyByGrade,
		"wxGroupIds":               wxgroupIds,
		"classify": []map[string]interface{}{
			{
				"classify1Id": classify1Id,
				"classify2Id": classity2Id,
			},
		},
	})

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Default().Println("Error creating request:", err)
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "RANGERS_WEB_ID=5aa89a84-50f9-42a2-a610-ae3532318a0e; RANGERS_SAMPLE=0.712393799821835; ZYBIPSCAS="+COOKIE)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Default().Println("Error creating request:", err)
		return err
	}
	defer resp.Body.Close()

	// 读取响应内容
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.String()

	fmt.Println("Response:", responseBody)
	return nil
}
