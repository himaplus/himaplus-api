package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"himaplus-api/common/responder"
	"himaplus-api/dto/extres"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 認証サーバに
func AuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// クライアントからのリクエストのヘッダ情報取得
		headerAuthorization := ctx.Request.Header.Get("Authorization")
		if headerAuthorization == "" { // ヘッダーが存在しない場合
			responder.SendJSON(ctx, http.StatusUnauthorized, "Authentication unsuccessful. Header information does not contain authentication information.", nil, gin.H{})
			ctx.Abort() // 次のルーティングに進まないよう処理を止める。
			return      // 早期リターンで終了
		}

		// トークンの検証
		userInfo, err := VerityTokenWithAuthSrv(headerAuthorization)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			responder.SendJSON(ctx, http.StatusUnauthorized, "Authentication unsuccessful.", nil, gin.H{})
			ctx.Abort() // 次のルーティングに進まないよう処理を止める。
			return      // 早期リターンで終了
		}

		// コンテキストに設定
		ctx.Set("token", headerAuthorization)        // クライアントのトークン
		ctx.Set("id", userInfo.Id)                   // クライアントのID
		ctx.Set("accessToken", userInfo.AccessToken) // クライアントのアクセストークン
		ctx.Set("user", userInfo)                    // クライアントの全情報

		ctx.Next() // エンドポイントの処理に以降
	}
}

func VerityTokenWithAuthSrv(token string) (extres.UserInfo, error) {
	// 認証サーバへリクエストを送る

	// リクエストの作成
	method := "GET"                                        // メソッド
	endopoint := "http://pb-authn-srv:8090" + "/auth/user" // URL
	req, err := http.NewRequest(method, endopoint, nil)    // リクエストの作成
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return extres.UserInfo{}, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// リクエストを送る
	client := &http.Client{ // クライアントを作成
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req) // リクエストを送信しレスポンスを受け取る
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return extres.UserInfo{}, err
	}
	defer res.Body.Close() // リソースの解放 必須

	// リクエスト結果から成功しているかどうか
	if res.StatusCode != 200 {
		return extres.UserInfo{}, errors.New(strconv.Itoa(res.StatusCode))
	}

	// 構造体にマッピング
	var authUserRes extres.AuthUser
	err = json.NewDecoder(res.Body).Decode(&authUserRes)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return extres.UserInfo{}, err
	}

	return authUserRes.Data.UserInfo, nil
}
