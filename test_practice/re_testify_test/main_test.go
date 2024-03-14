package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testifyパッケージを使ってるので簡潔に記述可能
func TestCalculate(t *testing.T) {
	assert.Equal(t, Calculate(2), 4)
	fmt.Println("TestCalculate: PASS")
}

// 参考：Testifyを使わない場合：
// func TestCalculate(t *testing.T) {
//     if Calculate(2) != 4 {
//         t.Error("Expected 2 + 2 to equal 4")
//     }
// }

// アプリケーションの状態は一般的に「up（稼働中）」または「down（停止中）」のいずれか
// GetStatus()はTestStatusNotEmpty()で呼び出される
func GetStatus() string {
	// この関数は仮にステータスを返すもの
	return "up"
}

// ステータスが "down" でないことをテストする
func TestStatusNotDown(t *testing.T) {
	status := GetStatus()
	if status != "down" {
		fmt.Println("TestStatusNotDown: PASS")
	} else {
		fmt.Println("TestStatusNotDown: FAIL")
		t.Fail()
	}
}

// ステータスが空でないことをテストする
func TestStatusNotEmpty(t *testing.T) {
	status := GetStatus()
	if status != "" {
		fmt.Println("TestStatusNotEmpty: PASS")
	} else {
		fmt.Println("TestStatusNotEmpty: FAIL")
		t.Fail()
	}
}

// テーブル駆動テスト
func TestCalculateWithTable(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if assert.Equal(Calculate(test.input), test.expected) {
			fmt.Println("TestCalculateWithTable: PASS")
		} else {
			fmt.Println("TestCalculateWithTable: FAIL")
			t.Fail()
		}
	}
}
