package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのデフォルトのエンジンを作成
	r := gin.Default()

	// カスタムのhttp.Clientを作成
	client := &http.Client{
		Timeout: time.Millisecond * 6000, // 0.01秒のタイムアウト
	}

	// ロギングミドルウェアを追加
	r.Use(gin.Logger())

	// ルートパスに対するハンドラを定義
	r.GET("/", func(c *gin.Context) {
		// yahoo.comにGETリクエストを送信
		resp, err := client.Get("https://www.yahoo.com")
		if err != nil {
			// エラーログを出力
			c.Error(err)
			c.String(http.StatusInternalServerError, "Failed to fetch page from Yahoo!")
			return
		}
		defer resp.Body.Close()

		// レスポンスボディを読み取り
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			// エラーログを出力
			c.Error(err)
			c.String(http.StatusInternalServerError, "Failed to read response body")
			return
		}

		// レスポンスをクライアントに返す
		c.String(http.StatusOK, string(body))
	})

	// ポート8080でサーバーを起動
	r.Run(":8080")
}
