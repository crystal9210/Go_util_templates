package main

import "fmt"

// DataItem:各シャードに格納するデータオブジェクト
type DataItem struct {
	Key   string
	Value string
}

// テストデータを生成するための関数
func GenerateData(n int) []DataItem {
	var data []DataItem
	for i := 0; i < n; i++ {
		data = append(data, DataItem{Key: fmt.Sprintf("key%d", i), Value: fmt.Sprintf("value%d", i)})
	}
	return data
}
