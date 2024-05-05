package main

// TCPサーバを作成し、特定のIPアドレスからの接続のみを許可するプログラム
import (
	"fmt"
	"net"
	"os"
)

func main() {
	listenPort := "8080"
	allowedClient := "127.0.0.1" // 許可するクライアントのIpアドレス→これがローカルのアドレス

	listener, err := net.Listen("tcp", listenPort)
	if err != nil {
		fmt.Println("Error listeneng:", err.Error())
		os.Exit(1)
	}
	defer listener.Close() // defer:含まれるブロック処理が終了するときに、最後に追加で後続の処理を行うためのシンボル(deferは「延期する」)→エラーハンドリング、リソースのクリーンアップに利用される(ファイル、ネットワーク接続のクローズなど)
	fmt.Println("Listeneng on ", listenPort)

	for {
		conn, err := listener.Accept() // conn:
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}

		// クライアントのIPアドレスを取得
		clientAddr := conn.RemoteAddr().(*net.TCPAddr).IP.String()
		// アドレスが許可されているかチェック(==)
		if clientAddr == allowedClient {
			go handleRequest(conn) // 許可されたクライアントからのリクエストを処理
		} else {
			conn.Close() // 許可されていないクライアントからの通信は閉じる
		}
	}

}

// ここにhandleRequest関数を実装(クライアントから受け取ったリクエストを処理するための関数)
