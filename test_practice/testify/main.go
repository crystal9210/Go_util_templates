package main

import (
	"fmt"
	"log"
	"net/http"
)

// HTTPプローブを実行するシンプルなwebサーバの提供

var defaultProbe Prober = new(probe)

type Prober interface {
	Probe(string) bool
}

// 定義と宣言の違いについて：specファイル内に参照先記述
// probe構造体の定義・宣言
type probe struct{}

// probe構造体に対しProbeメソッドを定義・宣言
// urlが正常に応答するかどうかを確認するメソッド
func (p *probe) Probe(url string) bool {
	res, err := http.Get(url)
	if err != nil {
		return false
	}

	return res.StatusCode == http.StatusOK
}

// 特定のURLへのプローブ(疎通確認)をしてその結果に基づいてHTTPレスポンスを送信する関数
func handleProbe(w http.ResponseWriter, req *http.Request) {
	// ret:プローブの成功(1)or失敗(0)を表現するために使用する変数
	var ret int
	if defaultProbe.Probe("http://example.com") {
		ret = 1
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("probe_success %d\n", ret)))
}

func main() {
	// /probeパスに対するリクエストをhandleProbe関数で処理するようにHTTPサーバを設定
	http.HandleFunc("/probe", handleProbe)

	// httpサーバをデフォルトのマルチプレクサ；http.DefaultServeMuxを使用して起動。
	log.Fatal(http.ListenAndServe(":8080", nil))
}
