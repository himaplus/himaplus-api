package presentation

import (
	"errors"
	"fmt"
	"himaplus-api/application"
	"himaplus-api/common/custom"
	"himaplus-api/common/logging"
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
	//todo_group_uuidの取得
	todoGroupUuid := ctx.Param("todo_group_uuid")

	// サービス処理
	todos, err := h.s.FindTodoGroupService(idAdjusted, todoGroupUuid)
	// TODO: todoGroupがなかったときのエラーハンドリング未実装
	if err != nil {
		// 処理で発生したエラーのうちカスタムエラーのみ
		var serviceErr *custom.CustomErr
		if errors.As(err, &serviceErr) {
			switch serviceErr.Type {
			case custom.ErrTypePermissionDenied:	// 見れるグループない
				// エラーログ(権限無し)
				logging.ErrorLog("Do not have the necessary permissions", err)
				// レスポンス
				resStatusCode := http.StatusForbidden
				ctx.JSON(resStatusCode, gin.H{
					"srvResMsg":  http.StatusText(resStatusCode),
					"srvResData": gin.H{},
				})
				return

			default:
				// エラーログ
				logging.ErrorLog("aiueos", err)
				// レスポンス
				resStatusCode := http.StatusBadRequest
				ctx.JSON(resStatusCode, gin.H{
					"srvResMsg":  http.StatusText(resStatusCode),
					"srvResData": gin.H{},
				})
			}
		}
		// エラーログ
		logging.ErrorLog("todoGroup find error", err)
		// レスポンス(StatusInternalServerError サーバーエラー500番)
		resStatusCode := http.StatusInternalServerError
		ctx.JSON(resStatusCode, gin.H{
			"srvResMsg":  http.StatusText(resStatusCode),
			"srvResData": gin.H{},
		})
		return //　<-返すよって型指定してないから切り上げるだけ
	}

	// 成功レスポンス
	responder.SendSuccess(ctx, http.StatusOK, todos)
}

// todo詳細取得
func (h *TodoHandler) GetTodoDetailHandler(ctx *gin.Context) {
	// userid取得
	// id, _ := ctx.Get("id")
	// idAdjusted := id.(string) // アサーション
	// fmt.Println(idAdjusted)   //　アサーションの確認

	idAdjusted := "16228a6b-d768-4b30-aeaa-fc455922865c"
	//todo_uuidの取得
	todoUuid := ctx.Param("todo_uuid")

	// サービス処理
	todo, err := h.s.GetTodoDetailaService(idAdjusted, todoUuid)
	// TODO: todoGroupがなかったときのエラーハンドリング未実装
	if err != nil {
		// 処理で発生したエラーのうちカスタムエラーのみ
		var serviceErr *custom.CustomErr
		if errors.As(err, &serviceErr) {
			switch serviceErr.Type {
			case custom.ErrTypePermissionDenied:	// 見れるグループない
				// エラーログ(権限無し)
				logging.ErrorLog("Do not have the necessary permissions", err)
				// レスポンス
				resStatusCode := http.StatusForbidden
				ctx.JSON(resStatusCode, gin.H{
					"srvResMsg":  http.StatusText(resStatusCode),
					"srvResData": gin.H{},
				})
				return

			default:
				// エラーログ
				logging.ErrorLog("aiueos", err)
				// レスポンス
				resStatusCode := http.StatusBadRequest
				ctx.JSON(resStatusCode, gin.H{
					"srvResMsg":  http.StatusText(resStatusCode),
					"srvResData": gin.H{},
				})
			}
		}
		// エラーログ
		logging.ErrorLog("todo find error", err)
		// レスポンス(StatusInternalServerError サーバーエラー500番)
		resStatusCode := http.StatusInternalServerError
		ctx.JSON(resStatusCode, gin.H{
			"srvResMsg":  http.StatusText(resStatusCode),
			"srvResData": gin.H{},
		})
		return //　<-返すよって型指定してないから切り上げるだけ
	}

	// 成功レスポンス
	responder.SendSuccess(ctx, http.StatusOK, todo)
}
