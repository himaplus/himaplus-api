package middleware

import (
	"himaplus-api/common/logging"
	"time"

	"github.com/gin-gonic/gin"
)

// ロギング
func LoggingMid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ラップした関数を利用（特に変わらない）

		// リクエストを受け取った時のログ
		logging.SimpleLog("Received request.\n")                          // リクエストの受理ログ
		logging.SimpleLog("Time: ", time.Now(), "\n")                     // 時刻
		logging.SimpleLog("Request method: ", ctx.Request.Method, "\n")   // メソッドの種類
		logging.SimpleLog("Request path: ", ctx.Request.URL.Path, "\n\n") // リクエストパラメータ

		// リクエストを次のハンドラに渡す
		ctx.Next()

		// レスポンスを返した後のログ
		logging.SimpleLog("Sent response.\n")                               // レスポンスの送信ログ
		logging.SimpleLog("Time: ", time.Now(), "\n")                       // 時刻
		logging.SimpleLog("Response Status: ", ctx.Writer.Status(), "\n\n") // ステータスコード

		// 組み込みのlog pkgの関数を利用

		// // リクエストを受け取った時のログ
		// log.Printf("Received request.\n")                        // リクエストの受理ログ
		// log.Printf("Time: %v\n", time.Now())                     // 時刻
		// log.Printf("Request method: %s\n", ctx.Request.Method)   // メソッドの種類
		// log.Printf("Request path: %s\n\n", ctx.Request.URL.Path) // リクエストパラメータ

		// // リクエストを次のハンドラに渡す
		// ctx.Next()

		// // レスポンスを返した後のログ
		// log.Printf("Sent response.\n")                             // レスポンスの送信ログ
		// log.Printf("Time: %v\n", time.Now())                       // 時刻
		// log.Printf("Response Status: %d\n\n", ctx.Writer.Status()) // ステータスコード
	}
}
