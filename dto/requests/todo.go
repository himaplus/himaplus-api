package requests

import (
	"time"
)

// requestに必要な値
// TODOテーブル
type RegisterTodo struct {
	UserUUID     string        `json:"userUUID" binding:"required"`     // ユーザーUUID
	Titel        string        `json:"Titel" binding:"required"`        // todoタイトル
	Importance   int           `json:"importance" binding:"required"`   // 重要度	1:高、2:中、3:低
	RequiredTime time.Duration `json:"requiredTime" binding:"required"` // 必要時間
	Memo         string        `json:"Memo"`                            // memo
	ParentUuid   *string       `xorm:"varchar(36)" json:"parentUUID"`   // 親要素のUUID
}
