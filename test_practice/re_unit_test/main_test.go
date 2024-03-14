package main

import (
	"log"
	"testing"
)

// 特定範囲の同じ対象に対して複数のテストケースおよびテスト用関数群をテストスイートと呼ぶ

// 【説明】
// テストスイートは、通常、特定のパッケージ、モジュール、あるいはクラスなどの単位で複数のテスト関数をまとめるものとして使われます。そのため、1つのオブジェクトに対するテスト関数群を1つのテストスイートとしてまとめる場合もありますが、必ずしもそれが限定されるものではありません。

// 具体的には、特定の機能やコンポーネントに関連する複数のテスト関数をまとめて、その機能やコンポーネントが正しく動作するかどうかを確認するために使用されます。この場合、特定の関数に対する複数のテスト関数を含む場合もあります。

// 要するに、テストスイートはテストの組織化の手段であり、テスト対象の単位によって異なる範囲をカバーすることができます。そのため、オブジェクトに対してのみならず、特定の関数に対しての複数のテスト関数を含む場合も、それらをテストスイートと呼ぶことができます。

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	} else {
		log.Printf("You got expected 2 + 2 to equal 4")
	}
}

// 内部でmainパッケージ空間のCalculate関数を呼び出してテストしている；入力と期待する結果の組み合わせを用いたデータ群(テーブル)を用いているので"テーブル駆動テスト"と呼ばれる。
func TestTableCalculate(t *testing.T) {
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
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
