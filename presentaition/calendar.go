package presentation

import (
	"himaplus-api/application"

	"github.com/gin-gonic/gin"
)

type CalendarHandler struct {
	s *application.CalendarService
}

// ファクトリー関数
func NewCalendarHandler(s *application.CalendarService) *CalendarHandler {
	return &CalendarHandler{
		s: s,
	}
}

// ハンドラー

// カレンダー取得
func (h *CalendarHandler) GetCalender(ctx *gin.Context) {

}
