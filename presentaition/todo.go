package presentation

import (
	"fmt"
	"himaplus-api/application"
	"himaplus-api/common/responder"
	"himaplus-api/dto/clireq"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	s *application.TodoService
}

// ファクトリー関数
func NewTodoHandler(s *application.TodoService) *TodoHandler {
	return &TodoHandler{
		s: s,
	}
}

// Todo登録
func (h *TodoHandler) RegisterTodoHandler(ctx *gin.Context) {
	// 構造体にマッピング
	var bTodo []clireq.RegisterTodo // 構造体のインスタンス
	if err := ctx.ShouldBindJSON(&bTodo); err != nil {
		fmt.Println("Binding failed:", err)
		responder.SendFailedBindJSON(ctx, err)
		return
	}

	// サービス処理
	ids, err := h.s.RegisterTodoService(bTodo)
	if err != nil {
		fmt.Println(err)
	}

	// 成功レスポンス
	responder.SendSuccess(ctx, http.StatusCreated, ids)
}

// todo取得
func (h *TodoHandler) GetAllTodoHandler(ctx *gin.Context) {

	// userid取得
	// id, _ := ctx.Get("id")
	// idAdjusted := id.(string) // アサーション
	// fmt.Println(idAdjusted)   //　アサーションの確認
	idAdjusted := "16228a6b-d768-4b30-aeaa-fc455922865c"

	// サービス処理
	todos, err := h.s.FindAllTodoService(idAdjusted)
	if err != nil {
		fmt.Println(err)
	}

	// 成功レスポンス
	responder.SendSuccess(ctx, http.StatusOK, todos)
}

// todoGroup取得
func (h *TodoHandler) GetTodoGroupHandler(ctx *gin.Context) {
	// userid取得
	// id, _ := ctx.Get("id")
	// idAdjusted := id.(string) // アサーション
	// fmt.Println(idAdjusted)   //　アサーションの確認

	idAdjusted := "16228a6b-d768-4b30-aeaa-fc455922865c"
	//notice_uuidの取得
	todoGroupUuid := ctx.Param("todo_group_uuid")

	// サービス処理
	todos, err := h.s.FindTodoGroupService(idAdjusted, todoGroupUuid)
	// TODO: todoGroupがなかったときのエラーハンドリング未実装
	if err != nil {
		fmt.Println(err)
	}

	// 成功レスポンス
	responder.SendSuccess(ctx, http.StatusOK, todos)
}
