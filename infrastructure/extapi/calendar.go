package apicli

import (
	"net/http"
	"time"
)

// // インフラ処理

// サービスの構造体
type CalendarApiCli struct {
	client *http.Client
}

// ファクトリー関数
func NewCalendarApiCli(client *http.Client) *CalendarApiCli {
	return &CalendarApiCli{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

// インフラ処理
