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

// TodoGroup登録
func (i *TodoInfrastruture) CreateTodoGroup(record model.TodoGroup) (int64, error) {
	affected, err := i.db.Insert(record)
	return affected, err
}

// Todo登録
func (i *TodoInfrastruture) CreateTodo(record model.Todo) (int64, error) {
	affected, err := i.db.Insert(record)
	return affected, err
}

// todoGroupを取得
func (i *TodoInfrastruture) FindAllTodoGroup(userUuid string) ([]model.TodoGroup, error) {
	// 結果格納用
	var todoGroups []model.TodoGroup

	// userUuidで絞り込んだデータを全権取得
	err := i.db.In("user_uuid", userUuid).Find(&todoGroups)
	// データが取得できなかったらerrを返す
	if err != nil {
		return nil, err
	}

	// エラーが出なければ取得結果を返す
	return todoGroups, nil
}

// 単体todoの取得
func (i *TodoInfrastruture) FindAllSingleTodo(userUuid string) ([]model.Todo, error) {
	// 結果格納用
	var singleTodos []model.Todo

	// userUuidとGroupUuidがnullで絞り込んだデータを全権取得
	err := i.db.In("user_uuid", userUuid).Where("todo_group_uuid IS NULL").Find(&singleTodos)
	// データが取得できなかったらerrを返す
	if err != nil {
		return nil, err
	}

	// エラーが出なければ取得結果を返す
	return singleTodos, nil

}

// todoGroupあるかの確認
func (i *TodoInfrastruture) IsTodoGroup(userUuid string, todoGroupUuid string) (bool, error) {
	var todoGroup model.TodoGroup

	isGroup, err := i.db.Where("user_uuid = ?", userUuid).And("todo_group_uuid=?", todoGroupUuid).Exist(&todoGroup)
	if err != nil {
		return false, err
	}

	return isGroup, nil
}

// todoGroupに関するtodo取得
func (i *TodoInfrastruture) GetTodoByTodoGroup(userUuid string, todoGroupUuid string) ([]model.Todo, error) {
	var todos []model.Todo

	// TODO:ここってfind?Get?
	err := i.db.Where("user_uuid	= ?", userUuid).And("todo_group_uuid=?", todoGroupUuid).Find(&todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// todo詳細取得
func (i *TodoInfrastruture) GetTodoDetail(userUuid string, todoUuid string) (*model.Todo, error) {
	// 格納用変数
	var todoDateil model.Todo

	// 詳細取得
	found, err := i.db.Where("user_uuid	= ?", userUuid).And("todo_uuid=?", todoUuid).Get(&todoDateil)
	if err != nil {
		return nil, err
	}

	// データが取得できなかったら
	if !found {
		return nil, nil
	}

	return &todoDateil, nil
}

// todo更新
func (i *TodoInfrastruture) UpdateTodo(userUuid string, todoUuid string, record model.Todo) (int64, error) {
	affected, err := i.db.
			Where("user_uuid = ? AND todo_uuid = ?", userUuid, todoUuid).
			Update(&record)
	return affected, err
}
