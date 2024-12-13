package application

import (
	"fmt"
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

// 自己満で登録されたTODOを確認したいのでそれ用の構造体
type TodoInfo struct { // typeで型の定義, structは構造体
	TodoUuid string `json:"todoUUID"` //UUID
	Title    string `json:"title"`    //タイトル
}

// Todo新規登録
func (s *TodoService) RegisterTodoService(req []requests.RegisterTodo) ([]TodoInfo, error) {

	// 他のtodo登録で必要なgroupHostの情報を保持する用変数
	var groupUuid string
	hostPriority := 0 // TODO:置き換わってるかを判断できるように0入れてるけど、別の方法ありそうだったから変える

	// 結果格納用
	var todoInfos []TodoInfo
	
	// 複数登録するときのために
	for _, todo := range req {

		fmt.Println(todo) // 情報確認

		if todo.GroupHost {
			// uuidを生成
			hostId, err := uuid.NewRandom() //新しいuuidの作成
			if err != nil {
				return []TodoInfo{}, err
			}

			// 小要素のtodo登録に必要なので保持しておく
			groupUuid = hostId.String()
			hostPriority = todo.Priority

			// 構造体をレコード登録処理に投げる
			_, err = s.i.CreateTodoGroup(model.TodoGroup{
				UserUuid:      todo.UserUUID,
				TodoGroupUuid: groupUuid,
				Title:         todo.Titel,
				Priority:      todo.Priority,
				Date:          time.Now(),
			})
			if err != nil {
				return []TodoInfo{}, err
			}

			// 情報格納
			todoInfo := TodoInfo{
				TodoUuid: groupUuid,
				Title:    todo.Titel,
			}
			// 情報格納
			todoInfos = append(todoInfos, todoInfo)

		} else {

			// uuidを生成
			uuid, err := uuid.NewRandom() //新しいuuidの作成
			if err != nil {
				return []TodoInfo{}, err
			}

			// hostがいるならhostPriorityに値があるハズなかったら自分の使う
			if hostPriority == 0 {
				hostPriority = todo.Priority
			}

			// 構造体をレコード登録処理に投げる
			_, err = s.i.CreateTodo(model.Todo{
				UserUuid:      todo.UserUUID,
				TodoUuid:      uuid.String(),
				Title:         todo.Titel,
				Priority:      hostPriority,
				RequiredTime:  todo.RequiredTime,
				Memo:          todo.Memo,
				Date:          time.Now(),
				TodoGroupUuid: &groupUuid,
			})
			if err != nil {
				return []TodoInfo{}, err
			}

			// 情報格納
			todoInfo := TodoInfo{
				TodoUuid: uuid.String(),
				Title:    todo.Titel,
			}
			// 情報格納
			todoInfos = append(todoInfos, todoInfo)
		}


	}

	return todoInfos, nil
}
