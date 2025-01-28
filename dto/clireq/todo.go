package clireq

import (
	"time"
)

// requestに必要な値
// TODOテーブル
type RegisterTodo struct {
	GroupHost     bool          `json:"groupHost"`     // 親要素かのフラグ
	UserUUID      string        `json:"userUUID"`      // ユーザーUUID
	Titel         string        `json:"title"`         // todoタイトル
	Priority      int           `json:"priority"`      // 重要度	1:高、2:中、3:低
	RequiredTime  time.Duration `json:"requiredTime"`  // 必要時間
	Memo          string        `json:"memo"`          // memo
	TodoGroupUuid *string       `json:"todoGroupUuid"` // 親要素のUUID
}
