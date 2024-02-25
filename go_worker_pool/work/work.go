// 完了するまでに時間がかかる処理をシミュレートするためにランダムな文字列を使って沢山のジョブを作成する

package job

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"math/big"
	"os"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunesは指定された長さのランダムな文字列を生成します。
func RandStringRunes(n int) (string, error) {
	b := make([]rune, n)
	for i := range b {
		// crypto/randを使用して安全な乱数を生成
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		if err != nil {
			return "", err // エラーを返す
		}
		b[i] = letterRunes[num.Int64()]
	}
	return string(b), nil
}

// CreateJobsは指定された数のジョブを生成します。
func CreateJobs(amount int) ([]string, error) {
	var jobs []string
	for i := 0; i < amount; i++ {
		job, err := RandStringRunes(8)
		if err != nil {
			return nil, err // エラーを返す
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

// mimics any type of job that can be run concurrently
func DoWork(word string, id int) {
	h := fnv.New32a()
	h.Write([]byte(word))
	time.Sleep(time.Second)
	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("worker [%d] - created hash [%d] from word [%s]\n", id, h.Sum32(), word)
	}
}
