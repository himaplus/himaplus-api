package model

import "time"

// todoグループテーブル
type TodoGroup struct {
	UserUuid      string    `xorm:"varchar(36) pk" json:"userUUID"`                       // ユーザーUUID
	TodoGroupUuid string    `xorm:"varchar(36) pk" json:"parentUUID"`                     // 親要素のUUID
	Title         string    `xorm:"varchar(36) not null" json:"Title" binding:"required"` // タイトル
	Priority      int       `xorm:"int not null" json:"priority" binding:"required"`      // 重要度
	Date          time.Time `xorm:"DATETIME not null" json:"date"`                        // 登録した時間
}
