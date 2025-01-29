package extres

import "time"

// 認証サーバへの検証レスポンス
type AuthUser struct {
	Data struct {
		Authenticated bool     `json:"authenticated"`
		UserInfo      UserInfo `json:"userInfo"`
	} `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// 再利用性のため切り出す
type UserInfo struct { // 認証サーバへの検証レスポンスが返すユーザーデータ
	Id          string    `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	AvatarPath  string    `json:"avatorUrl"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	AccessToken string    `json:"accessToken"`
}
