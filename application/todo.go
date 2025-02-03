package application

import (
	"fmt"
	"himaplus-api/common/custom"
	"himaplus-api/common/logging"
	"himaplus-api/dto/clireq"
	"himaplus-api/infrastructure/orm"
	"himaplus-api/infrastructure/orm/model"
	"sort"
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
func (s *TodoService) RegisterTodoService(req []clireq.RegisterTodo) ([]TodoInfo, error) {

	// 他のtodo登録で必要なgroupHostの情報を保持する用変数
	var groupUuid string
	hostPriority := 0 // TODO:置き換わってるかを判断できるように0入れてるけど、別の方法ありそうだったら変える

	// 結果格納用
	var todoInfos []TodoInfo

	// 複数登録するときのために　// TODO:コード書き換えます。きしょいので
	for _, todo := range req {
		fmt.Printf("todo: %+v\n", todo) // フィールド内容をすべて確認
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

			// 登録する構造体の宣言
			bTodo := model.Todo{
				UserUuid:     todo.UserUUID,
				TodoUuid:     uuid.String(),
				Title:        todo.Titel,
				Priority:     hostPriority,
				RequiredTime: todo.RequiredTime,
				Memo:         todo.Memo,
				Date:         time.Now(),
			}

			// hostがあった場合、todoGroupUuidが保持されるため単体登録の時、登録されないので
			// groupUuidに値が入っているときにその項目を増やすようにしました
			if groupUuid != "" {
				bTodo.TodoGroupUuid = &groupUuid
			}

			// 構造体をレコード登録処理に投げる
			_, err = s.i.CreateTodo(bTodo)
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

// todo一覧取得の際のテーブル
type FindAllTodo struct {
	TodoHostUuid string `json:"todoUUID"`  // 親要素か単体todoのUUID
	Title        string `json:"title"`     // タイトル
	Priority     int    `json:"priority"`  // 重要度
	GroupHost    bool   `json:"groupHost"` // GroupHostなのかを保持
}

// todo一覧取得
func (s *TodoService) FindAllTodoService(userUuid string) ([]FindAllTodo, error) {

	// 取得してきた値を格納する用のスライス
	var FindAllTodos []FindAllTodo

	// TODO:grouphostと単体todoをまとめて撮ってくるには私の精神が持たないので分けます
	// TodoGroupに登録されているのを取得してくる
	todoGroups, err := s.i.FindAllTodoGroup(userUuid)
	if err != nil {
		return []FindAllTodo{}, err
	}

	// 整形して必要な値を取り出していくために取ってきた分だけ繰り返す
	for _, todoGroup := range todoGroups {

		// 整形して必要な値を取り出す
		FindAllTodo := FindAllTodo{
			TodoHostUuid: todoGroup.TodoGroupUuid,
			Title:        todoGroup.Title,
			Priority:     todoGroup.Priority,
			GroupHost:    true,
		}

		// スライスに追加していく
		FindAllTodos = append(FindAllTodos, FindAllTodo)
	}

	// 単体todoを取得してくる
	singleTodos, err := s.i.FindAllSingleTodo(userUuid)
	// 整形して必要な値を取り出していくために取ってきた分だけ繰り返す
	for _, singleTodo := range singleTodos {

		// 整形して必要な値を取り出す
		FindAllTodo := FindAllTodo{
			TodoHostUuid: singleTodo.TodoUuid,
			Title:        singleTodo.Title,
			Priority:     singleTodo.Priority,
			GroupHost:    false,
		}

		// スライスに追加していく
		FindAllTodos = append(FindAllTodos, FindAllTodo)
	}

	// 重要度の高い順に並び替える
	sort.Slice(FindAllTodos, func(i, j int) bool {
		return FindAllTodos[i].Priority < FindAllTodos[j].Priority
	})

	return FindAllTodos, err
}

// todoGroup取得
func (s *TodoService) FindTodoGroupService(userUuid string, todoGroupUuid string) ([]model.Todo, error) {

	// todoGroupが存在しているのか判断する
	isTodoGroup, err := s.i.IsTodoGroup(userUuid, todoGroupUuid)
	if err != nil {
		return []model.Todo{}, err
	}
	if !isTodoGroup { // なかったらエラー TODO:グループなかったときのエラーって権限なし？
		logging.ErrorLog("Do not have the necessary permissions", nil)
		return nil, custom.NewErr(custom.ErrTypePermissionDenied)
	}

	// 取得してきたgroupに関するtodoを全取得
	todoGroup, err := s.i.GetTodoByTodoGroup(userUuid, todoGroupUuid)
	if err != nil {
		return []model.Todo{}, err
	}

	fmt.Println(todoGroup)

	// 重要度の高い順に並び替える
	sort.Slice(todoGroup, func(i, j int) bool {
		return todoGroup[i].Priority < todoGroup[j].Priority
	})

	return todoGroup, err
}

// todo詳細取得
func (s *TodoService) GetTodoDetailaService(userUuid string, todoUuid string) (model.Todo, error) {

	// todo取得
	todoDetail, err := s.i.GetTodoDetail(userUuid, todoUuid)
	if err != nil {
		return model.Todo{}, err
	}

	// 取得できてるか確認　なかったらエラー
	if todoDetail == nil { // 取得できなかった
		fmt.Println("todoDetail is nil")
		return model.Todo{}, custom.NewErr(custom.ErrTypeNoResourceExist)
	}

	fmt.Println(todoDetail)

	return *todoDetail, err
}

// todo更新
func (s *TodoService) UpdateTodoService(req clireq.RegisterTodo, todoUuid string) (TodoInfo, error) {

	// todoが存在するか reqの中にあるuserUuidを使う
	todoDetail, err := s.i.GetTodoDetail(req.UserUUID, todoUuid)
	if err != nil {
		return TodoInfo{}, err
	}

	// 取得できてるか確認　なかったらエラー
	if todoDetail == nil { // 取得できなかった
		fmt.Println("todoDetail is nil")
		return TodoInfo{}, custom.NewErr(custom.ErrTypeNoResourceExist)
	}

	// 更新の際に使用
	bTodo := model.Todo{
		UserUuid:     req.UserUUID,
		TodoUuid:     todoUuid,
		Title:        req.Titel,
		Priority:     req.Priority,
		RequiredTime: req.RequiredTime,
		Memo:         req.Memo,
		Date:         time.Now(),
	}

	_, err = s.i.UpdateTodo(req.UserUUID, todoDetail.TodoUuid, bTodo)
	if err != nil {
		return TodoInfo{}, err
	}

	// 情報格納
	todoInfo := TodoInfo{
		TodoUuid: todoUuid,
		Title:    bTodo.Title,
	}

	return todoInfo, nil
}
