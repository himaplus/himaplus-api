package route

import (
	"himaplus-api/presentaition"
	"himaplus-api/view"

	"github.com/gin-gonic/gin"
)

// エンドポイントのルーティング
func routing(engine *gin.Engine, handlers Handlers) {

	// checkグループ
	check := engine.Group("/check")
	{
		// confirmation and response json test
		check.GET("/echo", presentation.ConfirmationReq) // /check/echo

		// sandbox
		check.GET("/sandbox", presentation.Try) // /check/sandbox
	}

	// ver1グループ
	v1 := engine.Group("/v1")
	{

		// todosグループ
		todos := v1.Group("/todos")
		{
			// todo新規登録
			todos.POST("/register", handlers.TodoHandler.RegisterTodoHandler) //  v1/todos/register

			// todo取得
			todos.GET("/todos", handlers.TodoHandler.GetAllTodoHandler)	// v1/todos

			// todo詳細取得
			todos.GET("/:todo_uuid", handlers.TodoHandler.GetTodoDetailHandler)	// va/todos/{todo_uuid}

			// todoGroup取得 TODO:
			todos.GET("/todo_groups/:todo_group_uuid", handlers.TodoHandler.GetTodoGroupHandler) // v1/todos/todo_groups/{todo_group_uuid}
		}
	}
}

// エンジンを作成して返す
func SetupRouter(handlers Handlers) (*gin.Engine, error) {
	// エンジンを作成
	engine := gin.Default()

	// 静的ファイル設定
	err := view.LoadingStaticFile(engine)
	if err != nil {
		return nil, err
	}

	// マルチパートフォームのメモリ使用制限を設定
	engine.MaxMultipartMemory = 8 << 20 // 20bit左シフトで8MiB

	// ルーティング
	routing(engine, handlers)

	// router設定されたengineを返す。
	return engine, nil
}
