package application

import (
	"himaplus-api/dto/requests"
	"himaplus-api/infrastructure/orm"
	// "github.com/google/uuid"
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
	// // 複数登録するときのために
	// for _, todo := range req {

	// 	// uuidを生成
	// 	todoId, err := uuid.NewRandom() //新しいuuidの作成
	// 	if err != nil {
	// 		return "", err
	// 	}

	// }

	return "", nil
}
