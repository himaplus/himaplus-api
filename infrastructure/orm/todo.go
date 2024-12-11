package orm

import (
	"xorm.io/xorm"
)

// サービスの構造体
type TodoInfrastruture struct {
	db *xorm.Engine
}

// ファクトリー関数
func NewTodoInfrastruture(db *xorm.Engine) *TodoInfrastruture {
	return &TodoInfrastruture{db: db}
}

