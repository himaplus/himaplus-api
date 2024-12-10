package model

import (
	"time"
)

// Todoテーブル
type Todo struct {
	UserUuid     string        `xorm:"varchar(36) pk" json:"userUUID"` // ユーザーUUID
	TodoUuid     string        `xorm:"varchar(36) pk" json:"TodoUUID"` // TodoUUID
	Title        string        `xorm:"varchar(36) not null" json:"Titel" binding:"required"`
	Importance   int           `xorm:"int not null" json:"importance" binding:"required"`
	RequiredTime time.Duration `xorm:"bigint" json:"requiredTime"` // 必要時間
	Memo         string        `xorm:"text" json:"memo"`                    // memo
	Date         time.Time     `xorm:"DATETIME not null" json:"date"`       // 登録した時間
	ParentUuid   *string       `xorm:"varchar(36)" json:"parentUUID"`    // 親要素のUUID
}