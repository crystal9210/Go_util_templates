package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func Test_Context_Data(t *testing.T) {
	// テスト用のHTTPサーバーをセットアップ
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		// コンテキストのタイムアウトを設定
		ctxWithTimeout, cancel := context.WithTimeout(ctx.Request.Context(), time.Millisecond*30) // 30→タイムアウト、3000→【200 OK.】
		defer cancel()

		// HTTPリクエストを生成
		req, err := http.NewRequestWithContext(ctxWithTimeout, http.MethodGet, "https://yahoo.com", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		// HTTPリクエストを実行
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to perform request: %v", err)
		}
		defer res.Body.Close()

		// HTTPレスポンスのボディを読み取り
		imageData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		// レスポンスをクライアントに返す
		ctx.Data(http.StatusOK, "text/html", imageData)
	})

	// テスト用のHTTPリクエストを作成
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	// HTTPハンドラーを呼び出し、リクエストを処理
	r.ServeHTTP(w, req)

	// レスポンスのステータスコードを検証
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code to be %d, got %d", http.StatusOK, w.Code)
	}

	// レスポンスのコンテンツタイプを検証
	contentType := w.Header().Get("Content-Type")
	if contentType != "text/html; charset=utf-8" {
		t.Errorf("Expected content type to be text/html; charset=utf-8, got %s", contentType)
	}

	// レスポンスのボディを検証
	expectedBody := []byte("<!DOCTYPE html>")
	actualBody := w.Body.Bytes()
	if !bytes.Contains(actualBody, expectedBody) {
		t.Errorf("Response body does not contain expected content: %s", expectedBody)
	}
}
