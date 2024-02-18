package main

import "hash/crc32"

// シャードを表現する簡単な構造体、ここで各シャードは一意のID(キー)を持ち、特定のデータセットを格納する
type Shard struct {
	ID   int
	Data map[string]string // キーと値のペアを格納する簡単な例
}

// 新しいシャードを作成する関数
func NewShard(id int) *Shard {
	return &Shard{ID: id, Data: make(map[string]string)}
}

// シャーディング関数；データをどのシャードに格納するかを決定
func ShardFunction(key string, numberOfShards int) int {
	// 簡単な例として、キーのハッシュ値を計算し、シャード数で割った余りをシャードIDとして使用
	hash := crc32.ChecksumIEEE([]byte(key))
	return int(hash) % numberOfShards
}

// データ挿入関数 for shard
func (s *Shard) InsertData(key, value string) {
	s.Data[key] = value
}

// データ取得関数 for shard
func (s *Shard) GetData(key string) (string, bool) {
	value, exists := s.Data[key]
	return value, exists
}
