package main

import (
	"net/http" //HTTPクライアントとサーバーの実装のためにインポート(Goの標準パッケージ)

	"github.com/gin-contrib/cors" // CORS設定用のパッケージ
	"github.com/gin-gonic/gin"    // GoのWebフレームワーク
)

func main() {
	// Ginを使って新しいルーターを作成
	// これでHTTPリクエストの処理とルーティングできるようになる
	r := gin.Default()

	// CORS設定
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // すべてのドメインからのリクエストを許可
	r.Use(cors.New(config))       // CORS設定をここで適用

	// ルーティングの設定
	r.POST("/process_message", func(c *gin.Context) { // POSTリクエストを受け取るエンドポイントを設定。
		var data struct { // リクエストデータを格納するための構造体を定義
			Message string `json:"message" form:"message"` // Messageというstring型のフィールドを持つ構造体を定義。jsonとformのタグを付与しているので、JSON形式とフォーム形式の両方で受け取れる。
		}

		// HTTPリクエストの中身のデータを構造体(data)に割り当てる。
		// JSONとフォームデータの両方に対応
		if err := c.ShouldBind(&data); err != nil { // もし割り当てに失敗した場合(データの形式が合わないなど)、エラーレスポンスを返す。
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported Media Type"})
			return
		}

		// 処理書く場所
		// 一旦テスト用にPOSTされたメッセージの末尾に"!"を付け足すだけの機能を実装
		processedMessage := data.Message + "!"

		// 処理されたメッセージをJSON形式でレスポンスする
		c.JSON(http.StatusOK, gin.H{"response": processedMessage})
	})

	// このコードでサーバーを起動(ポートは8080)
	r.Run(":8080")
}
