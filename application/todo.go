package application

import (
	"himaplus-api/dto/requests"
	"himaplus-api/infrastructure/orm"
	"himaplus-api/infrastructure/orm/model"
	"time"

	"github.com/google/uuid"
)

// サービスの構造体
type TodoService struct {
	i *orm.TodoInfrastruture
}

func NewTodoService(i *orm.TodoInfrastruture) *TodoService {
	return &TodoService{
		i: i,
	}
}

// Todo新規登録
func (s *TodoService) RegisterTodoService(req []requests.RegisterTodo) (string, error) {
	// 複数登録するときのために
	for _, Todo := range req {

		// uuidを生成
		todoId, err := uuid.NewRandom() //新しいuuidの作成
		if err != nil {
			return "", err
		}


		// 構造体をレコード登録処理に投げる
		_, err = s.i.CreateTodo(model.Todo{
			UserUuid:     Todo.UserUUID,      // userUuid
			TodoUuid:     todoId.String(),   // TodoUuid
			Title:        Todo.Titel,        // タイトル
			Importance:   Todo.Importance,   // 重要度
			RequiredTime: Todo.RequiredTime, // 必要時間
			Memo:         Todo.Memo,         // メモ
			Date:         time.Time{},       // 登録された時間
			ParentUuid:   nil,               // 親要素のUUID
		})
		if err != nil {
			return "", err
		}
	}

	return "", nil
}
