package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	// サーバーを別のゴルーチンで起動
	go startServer()

	// サーバーが起動するのを少し待つ
	time.Sleep(time.Second * 1)

	// クライアントからリクエストを送信
	sendRequest()
}

// 簡易HTTPサーバー
func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Cookieの設定
		http.SetCookie(w, &http.Cookie{
			Name:     "sessionid",
			Value:    "abc123",
			Path:     "/",
			HttpOnly: true,  // HttpOnly属性を設定
			Secure:   false, // テスト環境用にfalseに設定
		})

		// リクエストヘッダの情報を表示
		fmt.Println("Content-Type:", r.Header.Get("Content-Type"))
		fmt.Println("Custom-Header:", r.Header.Get("Custom-Header"))
		fmt.Println("Accept:", r.Header.Get("Accept"))
		fmt.Println("User-Agent:", r.Header.Get("User-Agent"))
		fmt.Println("Host:", r.Header.Get("Host"))
		fmt.Println("Cache-Control:", r.Header.Get("Cache-Control"))
		fmt.Println("Referer:", r.Header.Get("Referer"))
		// ヘッダからセッションIDとユーザ名を取得
		sessionID := r.Header.Get("Cookie")
		cookie, err := r.Cookie("username")
		if err != nil {
			fmt.Println(w, "Username not found in cookie.")
			return
		}
		username := cookie.Value

		// セッションIDが特定の値に一致するか確認（ここでは例として単純な文字列比較を行う）
		if strings.Contains(sessionID, "sessionId=abc123") {
			fmt.Fprintln(w, "Connection success!")
		} else {
			fmt.Fprintln(w, "Invalid session ID")
			return
		}

		// ユーザ名をレスポンスに組み込む
		fmt.Fprintf(w, "Hello, %s!", username)
	})

	http.ListenAndServe(":8080", nil)
}

// HTTPクライアントからのリクエスト送信
func sendRequest() {
	url := "http://localhost:8080" // リクエスト先のURL

	// リクエストの作成
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// ヘッダ追加
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Custom-Header", "value")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "MyApp/1.0.0")
	req.Header.Set("Host", "example.com")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Referer", "http://localhost:8080")
	req.Header.Set("Cookie", "sessionId=abc123; username=nukko")

	// HTTPクライアントの作成
	client := &http.Client{}

	// リクエストの送信
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// レスポンスの読み込み
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
