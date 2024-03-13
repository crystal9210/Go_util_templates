package handler

import (
	"encoding/json"
	"integration/repository"
	"net/http"
	"strconv"
)

type UserHandler struct {
	Repo repository.UserRepository
}

// HTTPリクエストのクエリパラメータのIDに基づいてユーザー情報をレスポンスとして返す関数
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// URLのクエリパラメータから"id"を取得
	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	// クエリパラメータの"id"を整数に変換；urlに組み込まれている時点；デフォルトでは文字列であるため
	// strconv.Atoi関数：文字列を整数に変換しようと試みるが、提供された文字列が整数として解釈できない場合（例えば、文字列に数字以外の文字が含まれている場合など）、変換エラーを返す
	id, err := strconv.Atoi(ids[0])
	if err != nil {
		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
		return
	}

	// リポジトリを使用してユーザー情報を取得します。
	user, err := h.Repo.FindByID(id)
	// userがnilまたはエラーがある場合、ユーザーが見つからなかったとして扱う；この処理がないとテストのステータスコードが期待するものと一致せず、404のところが200になりFAILする
	if err != nil || user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// ユーザー情報をJSON形式でレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// 注意：下のコードではテストが正常に通らなかった。ステータスコードが違うものが返された。
// 【説明】MockUserRepositoryのFindByIDメソッドが、存在しないユーザーIDに対してnil, nilを返している点にあります。この実装では、ユーザーが見つからない場合にnilユーザーオブジェクトとnilエラーを返していますが、handler.UserHandlerのGetUserByIDメソッド内でこのケースを適切にハンドリングしていないため、ユーザーが見つからない場合でも200 OKレスポンスが返されてしまいます。
// 特定の問題点
// handler.UserHandlerのGetUserByIDメソッドでは、FindByIDからnilユーザーとnilエラーが返された場合に対する特別な処理が実装されていません。これにより、ユーザーが存在しない場合でも、空のユーザーオブジェクトがレスポンスとして返され、ステータスコード200 OKがクライアントに返されることになります。
// func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
// 	// URLのクエリパラメータにidが含まれるか確認
// 	ids, ok := r.URL.Query()["id"]
// 	if !ok || len(ids[0]) < 1 {
// 		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
// 		return
// 	}

// 	// URLに含まれるクエリパラメータの値は　文字列として最初扱われるので文字列を数値に変換
// 	id, err := strconv.Atoi(ids[0])
// 	if err != nil {
// 		http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
// 		return
// 	}

// 	// idを元にuserを取得
// 	user, err := h.Repo.FindByID(id)
// 	if err != nil {
// 		http.Error(w, "User not found", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// }
