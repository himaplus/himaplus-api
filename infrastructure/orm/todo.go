package orm

import (
	"himaplus-api/infrastructure/orm/model"

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

// Todo登録
func (i *TodoInfrastruture) CreateTodo(record model.Todo) (int64, error) {
	affected, err := i.db.Insert(record)
	return affected, err
}