package utils

import (
	"io/ioutil"
	"net/http"
)

func DownloadFile(url string) ([]byte, error) {
	// 发起HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应的内容
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
