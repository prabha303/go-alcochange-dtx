package googleAnalyticsService

import (
	"context"
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/internals/adapter/googleAnalyticsAdapter"
	"cyberliver/go-alcochange-dtx/internals/daos"

	"github.com/FenixAra/go-util/log"
)

type GoogleAnalytics struct {
	l               *log.Logger
	dbConnMSSQL     *mssqlcon.DBConn
	googleAnalytics daos.GoogleAnalyticsDao
}

func New(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *GoogleAnalytics {
	return &GoogleAnalytics{
		l:               l,
		dbConnMSSQL:     dbConnMSSQL,
		googleAnalytics: daos.NewGoogleAnalytics(l, dbConnMSSQL),
	}
}

func (ga *GoogleAnalytics) InsertGoogleAnalytics(date string, ctx context.Context) (*GoogleAnalyticsResponse, error) {

	response := GoogleAnalyticsResponse{}
	googleRes := googleAnalyticsAdapter.NewGoogleAnalyticsDtx(ga.l)
	resp, err := googleRes.GetGoogleAnalytics(date, ctx)
	if err != nil {
		ga.l.Error("InsertGoogleAnalytics Error - ", err)
		return nil, err
	}

	errG := ga.googleAnalytics.InsertGoogleAnalyticsData(resp)
	if errG != nil {
		ga.l.Error("InsertGoogleAnalytics Error - ", errG)
		return nil, errG
	}

	response.Message = "Google Analytics data Inserted successfully"

	return &response, nil
}

type GoogleAnalyticsResponse struct {
	Message string `json:"message"`
}
