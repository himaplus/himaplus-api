package responses

// responseで返す値たち
// TODOテーブル
type RegisterTodo struct {
	UserUuid     string        `xorm:"varchar(36) pk" json:"userUUID"`                           // ユーザーUUID
	TodoUuid     string        `xorm:"varchar(36) pk" json:"todoUUID"`                           // todoUUID
	TodoTitle   string        `xorm:"varchar(36) not null" json:"todotitel" binding:"required"` // todoタイトル
}
