package somefunc

import (
	"testing"
)

// Client構造体のRunメソッドが正常に動作するかどうかをテスト
func TestRun(t *testing.T) {

	patterns := []struct {
		val      int
		expected int
	}{
		{2, 2},
		{8, 8},
		{-10, -10},
	}

	for idx, pattern := range patterns {
		// Clientのnewの際に、モックオブジェクトを引数にする
		c := Client{&mockCaller{}}
		actual := c.Run(pattern.val)
		if pattern.expected != actual {
			t.Errorf("pattern %d: want %d, actual %d", idx, pattern.expected, actual)
		}
	}
}

// callメソッドのレシーバをmockCallerとして宣言する。
type mockCaller struct{}

// 通常のコードではcallメソッドは引数の値をそのまま返却するが、
// モックでは、引数 + 10した値を返却するようにする。→テストがFAILする。
func (s *mockCaller) Call(val int) int {
	return val + 10
}
