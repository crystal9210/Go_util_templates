package handler_test

import (
	"encoding/json"
	"integration/handler"
	"integration/model"
	"integration/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_GetUserByID(t *testing.T) {
	// repository.MockUserRepositoryインスタンスの作成＋handler.UserHandlerに注入
	repo := &repository.MockUserRepository{}
	h := handler.UserHandler{Repo: repo}

	//	httptestパッケージでテスト用サーバの構築(http.HandlerFuncを使うことで、h.GetUserByIDをHTTPリクエストを処理するために満たすことが必要なGoのhttp.Handlerインターフェースを満たすように変換)
	server := httptest.NewServer(http.HandlerFunc(h.GetUserByID))
	defer server.Close()

	// テストケースの作成
	tests := []struct {
		name       string
		userID     string
		wantStatus int
		wantName   string
	}{
		{"Valid ID", "1", http.StatusOK, "John Doe"},
		{"Invalid ID", "abc", http.StatusBadRequest, ""},
		{"Non-existing ID", "999", http.StatusNotFound, ""},
	}

	// 各テストケースを実行
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(server.URL + "?id=" + tc.userID)
			if err != nil {
				t.Fatalf("Could not make GET request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.wantStatus {
				t.Errorf("Expected status %v; got %v", tc.wantStatus, resp.Status)
			} else {
				// 正常な処理の場合の確認用ログ
				t.Logf("Successfully matched the expected status code: %v for userID %s", tc.wantStatus, tc.userID)
			}

			if tc.wantStatus == http.StatusOK {
				var user model.User
				if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
					t.Fatalf("Could not decode response body: %v", err)
				}
				if user.Name != tc.wantName {
					t.Errorf("Expected user name %s; got %s", tc.wantName, user.Name)
				} else {
					t.Logf("Successfully matched the expected user name: %s for userID %s", tc.wantName, tc.userID)
				}
			}
		})
	}

}
