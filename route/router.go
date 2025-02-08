package route

import (
	"himaplus-api/middleware"
	presentation "himaplus-api/presentaition"
	"himaplus-api/view"

	"github.com/gin-gonic/gin"
)

// エンドポイントのルーティング
func routing(engine *gin.Engine, handlers Handlers) {
	// MidLog all

	// logging
	engine.Use(middleware.LoggingMid())

	// endpoints

	// root page
	engine.GET("/", presentation.ShowRootPage) // /

	// checkグループ
	check := engine.Group("/check")
	{
		// confirmation and response json test
		check.GET("/echo", presentation.ConfirmationReq) // /check/echo

		// sandbox
		check.GET("/sandbox", presentation.Try) // /check/sandbox

		// 認証ミドルウェア
		check.GET("/auth", middleware.AuthToken(), presentation.Auth) // /check/auth
	}

	// ver1グループ
	v1 := engine.Group("/v1")
	{
		// todo新規登録 TODO: /v1/todosグループから/v1/auth/todosグループに移行して下行のコメントも外す
		// v1.POST("/register", handlers.TodoHandler.RegisterTodoHandler) //  /v1/todos/register

		// authグループ
		auth := v1.Group("/auth", middleware.AuthToken())
		{
			// todosグループ
			todos := auth.Group("/todos")
			{
				// todo新規登録
				todos.POST("/register", handlers.TodoHandler.RegisterTodoHandler) //  v1/auth/todos/register

				// todo取得
				todos.GET("/todos", handlers.TodoHandler.GetAllTodoHandler) // v1/auth/todos

				// todoGroup取得 TODO:
				todos.GET("/todo_groups/:todo_group_uuid", handlers.TodoHandler.GetTodoGroupHandler) // v1/todos/todo_groups/{todo_group_uuid}
			}

			// usersグループ
			users := auth.Group("/users")
			{

				// カレンダー取得
				users.GET("/calendars", handlers.CalendarHandler.GetCalender) // /v1/auth/users/calendars

				// カレンダーに予定追加
				users.POST("/calendars/events")

				// // calendarsグループ // TODO: あとで追加
				// calendars := users.Group("calendars")
				// {

				// }
			}

		}

				// todo詳細取得
				todos.GET("/:todo_uuid", handlers.TodoHandler.GetTodoDetailHandler) // v1/auth/todos/{todo_uuid}

				// // todo更新
				// todos.PUT("/update/:todo_uuid", handlers.TodoHandler.UpdateTodoHandler) // v1/auth/todos/update/{todo_uuid}

			}
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
