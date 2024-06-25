package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// 读取音视频集名称.csv文件并创建PACKAGEID到PACKAGENAME的映射
	packageMap, err := readPackageMap("音视频集名称.csv")
	if err != nil {
		fmt.Println("读取音视频集名称.csv失败:", err)
		return
	}

	// 读取火龙果音视频播放排行TOP100.csv文件
	csvFile, err := os.Open("火龙果音视频播放排行TOP100.csv")
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
	outputFile, err := os.Create("合并后的音视频排行.csv")
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer outputFile.Close()

	// 创建CSV写入器
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// 写入标题头
	writer.Write([]string{"ID", "FILEID", "PLAYCNT", "PLAYCNTINIT", "CREATETIME", "UPDATETIME", "PACKAGENAME"})

	// 遍历记录并进行合并
	for _, record := range records {
		if len(record) == 0 {
			continue
		}

		// 根据FILEID查找PACKAGENAME
		fileID := record[1]
		packageName, exists := packageMap[fileID]

		if exists {
			// 写入合并后的数据
			writer.Write(append(record, packageName))
		} else {
			// 如果没有找到PACKAGENAME，写入原始记录
			writer.Write(record)
		}
	}

	// 检查并处理CSV写入错误
	if err := writer.Error(); err != nil {
		fmt.Println("写入CSV失败:", err)
	}
}

// readPackageMap 从CSV文件中读取PACKAGEID和PACKAGENAME，并返回映射
func readPackageMap(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	packageMap := make(map[string]string)
	for _, record := range records {
		if len(record) == 2 {
			packageMap[record[0]] = record[1]
		}
	}

	return packageMap, nil
}
