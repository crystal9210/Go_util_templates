package main

import (
	"integration/handler"
	"integration/repository"
	"log"
	"net/http"
)

// ☆main.goファイル内では直接modelパッケージの機能を呼び出さないためインポートしなくてもok.
// 参考(GPT):main.goはhandlerとrepositoryの実装にのみ依存し、これらのモジュールが内部でどのようにmodelパッケージとやり取りしているかについては知る必要がありません。そのため、modelパッケージの構造や振る舞いが変更された場合でも、main.goのコードを修正する必要がなく、変更の影響を局所化できます。

func main() {
	// モックリポジトリのインスタンスを作成します。
	mockRepo := &repository.MockUserRepository{}

	// UserHandlerのインスタンスを作成し、モックリポジトリを注入します。
	userHandler := &handler.UserHandler{
		Repo: mockRepo,
	}

	// HTTPリクエストハンドラを設定します。
	http.HandleFunc("/user", userHandler.GetUserByID)

	// HTTPサーバーを起動します。
	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
