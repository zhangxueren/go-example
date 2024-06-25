package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 读取资料名称.csv文件并创建CONTENTID到CONTENTNAME的映射
	contentMap, err := readContentMap("资料名称.csv")
	if err != nil {
		fmt.Println("读取资料名称.csv失败:", err)
		return
	}

	// 读取趣学习星球资料下载TOP100.csv文件
	csvFile, err := os.Open("趣学习星球资料下载TOP100.csv")
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer csvFile.Close()

	// 创建CSV读取器
	reader := csv.NewReader(csvFile)

	// 读取所有记录
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("读取CSV失败:", err)
		return
	}

	// 创建新的CSV文件用于写入合并后的数据
	outputFile, err := os.Create("合并后的资料下载排行.csv")
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer outputFile.Close()

	// 创建CSV写入器
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// 写入标题头
	writer.Write([]string{"ID", "CONTENTID", "真实播放量", "外显播放量", "真实下载量", "外显下载量", "真实分享量", "创建时间", "更新时间", "CONTENTNAME"})

	// 遍历记录并进行合并
	for _, record := range records {
		if len(record) == 0 {
			continue
		}

		// 根据CONTENTID查找CONTENTNAME
		contentName, exists := contentMap[record[1]]

		if exists {
			// 写入合并后的数据
			writer.Write(append(record, contentName))
		} else {
			// 如果没有找到CONTENTNAME，写入原始记录
			writer.Write(record)
		}
	}

	// 检查并处理CSV写入错误
	if err := writer.Error(); err != nil {
		fmt.Println("写入CSV失败:", err)
	}
}

// readContentMap 从CSV文件中读取CONTENTID和CONTENTNAME，并返回映射
func readContentMap(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ',' // 假设CSV文件使用':'作为字段分隔符

	// 读取所有记录
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	contentMap := make(map[string]string)
	for _, record := range records {
		if len(record) == 2 {
			contentMap[record[0]] = record[1]
		}
	}

	return contentMap, nil
}
