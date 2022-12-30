// ルーティング用パッケージ
package router

import (
	"context"

	"github.com/app-todos/library/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	Port string = ":3000" // ポート番号
)

var (
	Origin []string = []string{"http://localhost:8080"}            // ベースURL
	Method []string = []string{"GET", "POST", "DELETE", "OPTIONS"} // HTTP通信メソッド
	Header []string = []string{"*"}                                // ヘッダー
)

// ルーティング情報構造体
type Route struct {
	Gin    *gin.Engine // Ginインスタンス
	Origin []string    // ベースURL
	Method []string    // HTTP通信メソッド
	Header []string    // ヘッダー
	Port   string      // ポート番号
}

// NewRoute
// Route構造体を生成
// return *Route
func NewRoute() *Route {
	return &Route{
		Gin:    gin.Default(),
		Origin: Origin,
		Method: Method,
		Header: Header,
		Port:   Port,
	}
}

// SetRoute
// Route構造体をもとにルーティングを設定
func (r *Route) SetRoute() {
	r.Gin.Use(cors.New(cors.Config{
		AllowOrigins: r.Origin,
		AllowMethods: r.Method,
		AllowHeaders: r.Header,
	}))
}

// Run
// サーバ起動処理
// param ctx : コンテキスト
func Run(ctx context.Context) {
	logger.Log.Info("START - router")
	defer logger.Log.Info("END - router")

	// ルーティング設定
	r := NewRoute()
	r.SetRoute()
	// ルーティング定義
	r.Gin.Run(r.Port)
}