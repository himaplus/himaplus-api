package presentation

import (
	"fmt"
	"himaplus-api/application"
	"himaplus-api/common/responder"
	"himaplus-api/dto/requests"
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

	fmt.Println("Todoはんどらーです")

	// 構造体にマッピング
	var bTodo []requests.RegisterTodo // 構造体のインスタンス
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
