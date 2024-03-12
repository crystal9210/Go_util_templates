package animals_test

// _test接尾辞：animalsパッケージの外部テストを意味し、animalsパッケージの公開APIのみを使用してテストすることを示す

import (
	"log"
	"testing"

	"test_mod/easy_subtest/animals"
	"test_mod/easy_subtest/foods"
)

// Duck型の振る舞いをテストするための関数；関数名がTestで始まるため、これがGoのテストランナーによって自動的に認識され、テストとして実行される形式
func TestDuck(t *testing.T) {
	// animals構造体のインスタンスの宣言
	duck := animals.NewDuck("tarou")

	// 第一引数：サブテストの名前
	t.Run("it says quack", func(t *testing.T) {
		actual := duck.Say()
		expected := "tarou says quack"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		} else {
			log.Printf("got: %v\nwant: %v", actual, expected)
		}
	})

	t.Run("it ate apple", func(t *testing.T) {
		apple := foods.NewApple("sunfuji")

		actual := duck.Eat(apple.String())
		expected := "tarou ate sunfuji"
		if actual != expected {
			t.Errorf("got: %v\nwant: %v", actual, expected)
		} else {
			log.Printf("got: %v\nwant: %v", actual, expected)
		}
	})
}
