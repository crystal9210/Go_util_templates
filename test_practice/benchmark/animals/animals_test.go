package animals_test

import (
	"testing"

	"test_mod/benchmark/animals"
	"test_mod/benchmark/foods"
)

func BenchmarkDuck_Say(b *testing.B) {
	duck := animals.NewDuck("tarou")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		duck.Say()
	}
	b.StopTimer()
}

// Say関数の実行結果：
// Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkDuck_Say$ test_mod/benchmark/animals

// goos: linux
// goarch: amd64
// pkg: test_mod/benchmark/animals
// cpu: 11th Gen Intel(R) Core(TM) i5-1145G7 @ 2.60GHz
// === RUN   BenchmarkDuck_Say
// BenchmarkDuck_Say
// BenchmarkDuck_Say-8     85785481                14.18 ns/op            0 B/op        0 allocs/op
// PASS
// ok      test_mod/benchmark/animals      1.700s

func BenchmarkDuck_Eat(b *testing.B) {
	duck := animals.NewDuck("tarou")
	food := foods.NewApple("sunfuji")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		duck.Eat(food.String())
	}
	b.StopTimer()
}

// Eat関数の実行結果：
// Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^BenchmarkDuck_Eat$ test_mod/benchmark/animals

// goos: linux
// goarch: amd64
// pkg: test_mod/benchmark/animals
// cpu: 11th Gen Intel(R) Core(TM) i5-1145G7 @ 2.60GHz
// === RUN   BenchmarkDuck_Eat
// BenchmarkDuck_Eat
// BenchmarkDuck_Eat-8     65990749                17.95 ns/op            0 B/op        0 allocs/op
// PASS
// ok      test_mod/benchmark/animals      2.191s
