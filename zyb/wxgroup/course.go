package wxgroup

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-example/helper"
	"io/ioutil"
	"net/http"
)

type RequestBody struct {
	CourseCreateTimeArr []string `json:"courseCreateTimeArr"`
	YearArr             []string `json:"yearArr"`
	SeasonArr           []string `json:"seasonArr"`
	LearnSeasonArr      []string `json:"learnSeasonArr"`
	GradeIdArr          []string `json:"gradeIdArr"`
	SubjectIdArr        []string `json:"subjectIdArr"`
	StatusArr           []string `json:"statusArr"`
	CollStatusArr       []string `json:"collStatusArr"`
	CourseStr           string   `json:"courseStr"`
	CreatorNameStr      string   `json:"creatorNameStr"`
	PageSize            int      `json:"pageSize"`
	PageNo              int      `json:"pageNo"`
}

type CourseResponse struct {
	ErrNo  int    `json:"errNo"`
	ErrStr string `json:"errStr"`
	Data   Data   `json:"data"`
	LogId  string `json:"logId"`
}

type Data struct {
	DataList []DataItem `json:"dataList"`
	Meta     Meta       `json:"meta"`
}

type DataItem struct {
	CourseId         int    `json:"courseId"`
	CourseName       string `json:"courseName"`
	Status           string `json:"status"`
	Season           string `json:"season"`
	LearnSeason      string `json:"learnSeason"`
	GradeId          string `json:"gradeId"`
	SubjectId        string `json:"subjectId"`
	Price            int    `json:"price"`
	LessonTimeRange  string `json:"lessonTimeRange"`
	CreatorName      string `json:"creatorName"`
	CourseCreateTime string `json:"courseCreateTime"`
	CollStartTime    string `json:"collStartTime"`
	CollStatus       string `json:"collStatus"`
	CollStatusChina  int    `json:"collStatusChina"`
	GroupCnt         int    `json:"groupCnt"`
	CollSucGroupCnt  string `json:"collSucGroupCnt"`
	Year             int    `json:"year"`
	FirstLessonTime  int64  `json:"firstLessonTime"`
	LastLessonTime   int64  `json:"lastLessonTime"`
	LabelId          int    `json:"labelId"`
}

type Meta struct {
	Total    int `json:"total"`
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

func fetchCourses(courseStr string) (*CourseResponse, error) {
	reqBody := RequestBody{
		CourseCreateTimeArr: nil,
		YearArr:             nil,
		SeasonArr:           nil,
		LearnSeasonArr:      nil,
		GradeIdArr:          nil,
		SubjectIdArr:        nil,
		StatusArr:           nil,
		CollStatusArr:       nil,
		CourseStr:           courseStr,
		CreatorNameStr:      "",
		PageSize:            50,
		PageNo:              1,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"https://wxtools.zuoyebang.cc/wxqk/collocation/getcourselist",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://wxtools.zuoyebang.cc")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Sec-Ch-Ua", `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`)
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")
	req.Header.Set("Cookie", "RANGERS_WEB_ID=991ac774-eed9-4805-9fbb-fe9ea61621c8; RANGERS_SAMPLE=0.04972195991847683; Hm_lvt_c33960c712441eec1b994580263ccb1a=1719804603,1721738711; Hm_lpvt_c33960c712441eec1b994580263ccb1a=1721738711; HMACCOUNT=19A5FA5D6642C585; uid=zhangxueren; ZYBIPSCAS=IPS_0a469733992542889066b3d542b05af41722491442; ZYBIPSUN=7a68616e6778756572656e")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	//fmt.Println(string(body))

	var response CourseResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func GetCourseDetailList() {
	courseStrs := []string{"2957233", "2928187", "2918688", "2928187", "2928255", "2928187", "2928319", "2928319", "2918688", "2917693", "2928187", "2928187", "2928250", "2928187", "2928187", "2928187", "2928255", "2928316", "2928187", "2928255", "2928187", "2928163", "2928187", "2928250", "2928187", "2957234", "2928255", "2928319", "2928187", "2928250", "2928250", "2918688", "2928316", "2928255", "2928187", "2918688", "2928187", "2928319", "2928187", "2928163", "2928316", "2957234", "2928255", "2928250", "2928187", "2928163", "2928316", "2928187", "2918688", "2928187", "2928187", "2928187", "2918688", "2928250", "2928187", "2928255", "2928187", "2928163", "2957234", "2918688", "2928187", "2957234", "2928316", "2928187", "2935848", "2918688", "2928187", "2918688", "2928316", "2918688", "2928163", "2928187", "2928187", "2928163", "2928316", "2928187", "2928187", "2928187", "2928187", "2928316", "2928316", "2928163", "2867681", "2928187", "2928187", "2928187", "2928187", "2928187", "2928187", "2918688", "2928187", "2928187", "2928316", "2928163", "2928187", "2928187", "2928187", "2928187", "2867676", "2928187", "2928187", "2928187", "2928163", "2928187", "2867675", "2918688", "2928316", "2928187", "2928255", "2928250"}
	courseStrs = helper.RemoveDuplicate(courseStrs)

	fmt.Println("CourseId, CourseName, Status, Season, LearnSeason, GradeId, SubjectId, Price, LessonTimeRange, CreatorName, CourseCreateTime, CollStartTime, CollStatus, CollStatusChina, GroupCnt, CollSucGroupCnt, Year, FirstLessonTime, LastLessonTime, LabelId")
	for _, courseStr := range courseStrs {
		response, err := fetchCourses(courseStr)
		if err != nil {
			fmt.Printf("Error fetching courses for %s: %v\n", courseStr, err)
			continue
		}
		//fmt.Printf("Response for %s: %+v\n", courseStr, response)

		for _, item := range response.Data.DataList {
			fmt.Printf("%d,%s,%s,%s,%s,%s,%s,%d,%s,%s,%s,%s,%s,%d,%d,%s,%d,%d,%d,%d\n",
				item.CourseId, item.CourseName, item.Status, item.Season, item.LearnSeason, item.GradeId, item.SubjectId, item.Price, item.LessonTimeRange, item.CreatorName, item.CourseCreateTime, item.CollStartTime, item.CollStatus, item.CollStatusChina, item.GroupCnt, item.CollSucGroupCnt, item.Year, item.FirstLessonTime, item.LastLessonTime, item.LabelId)
		}
	}
}
