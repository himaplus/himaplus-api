package model

import (
	"time"
)

// Todoテーブル
type Todo struct {
	UserUuid      string        `xorm:"varchar(36) pk" json:"userUUID"`                       // ユーザーUUID
	TodoUuid      string        `xorm:"varchar(36) pk" json:"TodoUUID"`                       // TodoUUID
	Title         string        `xorm:"varchar(36) not null" json:"Title" binding:"required"` // タイトル
	Priority      int           `xorm:"int not null" json:"priority" binding:"required"`      // 重要度
	RequiredTime  time.Duration `xorm:"bigint not null" json:"requiredTime" binding:"required"`        // 必要時間
	Memo          string        `xorm:"text" json:"memo"`                                     // memo
	Date          time.Time     `xorm:"DATETIME not null" json:"date"`                        // 登録した時間
	TodoGroupUuid *string       `xorm:"varchar(36)" json:"parentUUID"`                        // 親要素のUUID
}

func CreateTodoTestData() []Todo {
	return []Todo{
		{
			UserUuid:      "16228a6b-d768-4b30-aeaa-fc455922865c",
			TodoUuid:      "db2d30de-127e-47cf-aa26-772398e004f4",
			Title:         "買い物",
			Priority:      3,
			RequiredTime:  3600000,
			Memo:          "じゃがいもかいます",
			Date:          time.Now().Add(24 * time.Hour),
		},
		{
			UserUuid:      "16228a6b-d768-4b30-aeaa-fc455922865c",
			TodoUuid:      "97c2e621-4067-480b-90ed-2ad69af04b8b",
			Title:         "資料作成",
			Priority:      2,
			RequiredTime:  3600000,
			Memo:          "地球祭資料の作成",
			Date:          time.Now().Add(48 * time.Hour),
		},
	}
}
