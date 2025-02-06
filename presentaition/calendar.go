package presentation

import (
	"errors"
	"fmt"
	"himaplus-api/application"
	"himaplus-api/common/responder"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
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
	accessToken, exists := ctx.Get("accessToken") // token
	if !exists {
		responder.SendJSON(ctx, http.StatusUnauthorized, "Access token not available.", errors.New("could not get value from context"), gin.H{})
	}
	accessTokenAjusted := accessToken.(string) // アサーション
	fmt.Printf("accessTokenAjusted: %v\n", accessTokenAjusted)

	// // 1. トークンからクライアントを作成

	// // トークンを作成
	// o2Token := &oauth2.Token{
	// 	AccessToken: accessTokenAjusted,
	// }

	// // Google Calendar APIのクライアントを作成
	// client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(o2Token))

	// // Google Calendar APIのサービスを作成
	// srv, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"mssage": "Failed to create API service."})
	// }

	// 2. 生のトークンを直接サービスに渡す

	// Google Calendar APIのクライアントを作成
	srv, err := calendar.NewService(ctx, option.WithTokenSource(oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: accessTokenAjusted,
		},
	)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"mssage": "Failed to create API service."})
	}

	// カレンダーAPIを使う

	// 今月の分が欲しいので、
	now := time.Now() // 今の日付 // .Format("2006-01-02 00:00:00")
	year, month, _ := now.Date()
	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)                   // 月初
	endOfMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Second).Format(time.RFC3339) // 月終わり // 今日基準の翌月から引く

	fmt.Printf("startOfMonth: %v\n", startOfMonth)
	fmt.Printf("endOfMonth: %v\n", endOfMonth)

	// events
	events, err := srv.Events.List("primary"). // カレンダーリスト内のプライマリーカレンダーを取得
							ShowDeleted(false).    // 削除済みイベントを取得しない
							TimeMin(startOfMonth). // 月初
							TimeMax(endOfMonth).   // 月終わり
							OrderBy("startTime").  // 開始時間順
							SingleEvents(true).    // 繰り返しイベントを一つずつのイベントして分ける
							Do()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve event.",
		})
	}

	// t := time.Now().Format(time.RFC3339)
	// events, err := srv.Events.List("primary").ShowDeleted(false).
	// 	SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	// }
	// fmt.Println("Upcoming events:")
	// if len(events.Items) == 0 {
	// 	fmt.Println("No upcoming events found.")
	// } else {
	// 	for _, item := range events.Items {
	// 		date := item.Start.DateTime
	// 		if date == "" {
	// 			date = item.Start.Date
	// 		}
	// 		fmt.Printf("%v (%v)\n", item.Summary, date)
	// 	}
	// }

	ctx.JSON(200, gin.H{
		"message":   "Successfully retrieved events.",
		"calendars": events.Items, // 認証失敗で取れてないとここでにるぽる
	})
}
