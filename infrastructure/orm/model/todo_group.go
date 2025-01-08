package model

import "time"

// todoグループテーブル
type TodoGroup struct {
	UserUuid      string    `xorm:"varchar(36)" json:"userUUID"`                       // ユーザーUUID
	TodoGroupUuid string    `xorm:"varchar(36) pk" json:"todoGroupUUID"`                     // 親要素のUUID
	Title         string    `xorm:"varchar(36) not null" json:"title" binding:"required"` // タイトル
	Priority      int       `xorm:"int not null" json:"priority" binding:"required"`      // 重要度
	Date          time.Time `xorm:"DATETIME not null" json:"date"`                        // 登録した時間
}

func CreateTodoGroupTestData() []TodoGroup {
	return []TodoGroup{
		{
			UserUuid:      "16228a6b-d768-4b30-aeaa-fc455922865c",
			TodoGroupUuid: "7ec51405-03f4-47f6-a69e-8e52395d796b",
			Title:         "どっかー勉強",
			Priority:      1,
			Date:          time.Now().Add(72 * time.Hour),
		},
	}
}

