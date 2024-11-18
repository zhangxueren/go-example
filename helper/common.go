package helper

import (
	"bufio"
	"os"
)

// reverseSlice 是一个泛型函数，用于翻转任何类型的切片。
func ReverseSlice[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// 快速将文本写入文件
func WriteFile(filePath string, text string) error {
	// 创建或打开一个文件用于写入
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建一个bufio.Writer对象，用于缓冲写操作
	writer := bufio.NewWriter(file)

	// 使用writer.Write([]byte)方法写入文本
	_, err = writer.WriteString(text)
	if err != nil {
		return err
	}

	// 别忘了调用Flush()方法，确保所有数据都被写入到文件中
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

// 实现一个简单的泛型函数，用于对slice中的元素进行去重
func RemoveDuplicate[T comparable](s []T) []T {
	m := make(map[T]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}
	result := make([]T, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}
