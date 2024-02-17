package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

const defaultEnableServerPush = true

func main() {
	log.SetOutput(os.Stderr)
	var enableServerPush bool
	flag.BoolVar(&enableServerPush, "enableServerPush", defaultEnableServerPush, "server push option")
	flag.Parse()

	log.Printf("enableServerPush: %t", enableServerPush)
	e := NewServerPushServer(enableServerPush)
	e.StartTLS(":18443", "server.crt", "server.key")
}

func NewServerPushServer(enableServerPush bool) *echo.Echo {
	e := echo.New()
	if enableServerPush {
		e.GET("/", serverPush)
	} else {
		e.Static("/", "static")
	}
	return e
}

func serverPush(c echo.Context) error {
	pusher, ok := c.Response().Writer.(http.Pusher)
	if !ok {

		log.Printf("error: HTTP/2 push not supported")
		return c.File("static/index.html")

	}

	// 絶対パスを生成
	// absolutePath, err := filepath.Abs("lena.png")
	// if err != nil {
	// 	log.Printf("error: failed to get absolute path - %v", err)
	// 	return c.File("static/index.html")
	// }

	// 絶対パスを使用してプッシュ
	if err := pusher.Push("/home/crystal9210/Go_util_templates/serverpush/lena.png", nil); err != nil {
		log.Printf("error: failed to push /lena.png - %v", err)
		// エラーが発生しても、処理を続行
	}

	c.Response().Header().Set(echo.HeaderContentType, "image/png")

	// ファイルの内容をレスポンスに書き込み
	file, err := os.Open("/home/crystal9210/Go_util_templates/serverpush/static/lena.png")
	if err != nil {
		log.Printf("error: failed to open /lena.png - %v", err)
		return c.File("static/index.html")
	}
	defer file.Close()

	if _, err := io.Copy(c.Response(), file); err != nil {
		log.Printf("error: failed to write /static/lena.png - %v", err)
		return c.File("static/index.html")
	}

	return c.File("static/index.html")
}
