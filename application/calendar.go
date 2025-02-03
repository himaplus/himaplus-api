package application

import (
	apicli "himaplus-api/infrastructure/extapi"
)

type CalendarService struct {
	apiCli *apicli.CalendarApiCli
}

// ファクトリー関数
func NewCalendarService(api *apicli.CalendarApiCli) *CalendarService {
	return &CalendarService{
		apiCli: api,
	}
}

// サービス処理
