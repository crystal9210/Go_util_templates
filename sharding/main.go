package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// SecureRandomInt：0からmaxまでの暗号学的に安全なランダム整数を生成
func SecureRandomInt(max int) (int, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()), nil
}

// 実際に動かしてみるとわかるが、アクセスにバラツキが生じる→このようなアクセスの偏りや性質を考慮したうえでシャードキーを選択することが必要、他に一貫性ハッシュの利用など

func main() {

	// シャードの数
	const numOfShards = 10

	// シャードの初期化
	shards := make([]*Shard, numOfShards)
	for i := 0; i < numOfShards; i++ {
		shards[i] = NewShard(i)
	}

	// データの生成＋シャードに挿入
	data := GenerateData(100)
	for _, item := range data {
		shardID := ShardFunction(item.Key, numOfShards)
		shards[shardID].InsertData(item.Key, item.Value)
	}

	// データをランダムに10個取得して表示
	for i := 0; i < 10; i++ {
		randomIndex, err := SecureRandomInt(len(data))
		if err != nil {
			fmt.Println("Error generating random index:", err)
			continue
		}
		randomKey := data[randomIndex].Key
		shardID := ShardFunction(randomKey, numOfShards)
		if value, exists := shards[shardID].GetData(randomKey); exists {
			fmt.Printf("Found: %s in Shard %d\n", value, shardID)
		}
	}

	// 存在しないデータを5回検索して、見つからない場合の処理を表示

	// キーに対しハッシュ値を計算した結果としてシャードIDを生成し、そのシャードIDのシャード内にnonexistenkeyがあるかどうかを判定している
	for i := 0; i < 5; i++ {
		// 存在しないキーを生成
		nonExistentKey := fmt.Sprintf("nonexistent%d", i)
		shardID := ShardFunction(nonExistentKey, numOfShards)
		// 上記で生成したキーのハッシュ値から導出したシャードIDのシャードに対し、キーが存在するかどうかを判定
		if _, exists := shards[shardID].GetData(nonExistentKey); !exists {
			fmt.Printf("Data for key '%s' not found in Shard %d\n", nonExistentKey, shardID)
		}
	}
}
