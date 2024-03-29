package daos

import (
	"cyberliver/go-alcochange-dtx/dbcon/mssqlcon"
	"cyberliver/go-alcochange-dtx/dtos"
	"fmt"

	"github.com/FenixAra/go-util/log"
	"github.com/Masterminds/structable"
)

type GoogleAnalytics struct {
	l           *log.Logger
	dbConnMSSQL *mssqlcon.DBConn
}

func NewGoogleAnalytics(l *log.Logger, dbConnMSSQL *mssqlcon.DBConn) *GoogleAnalytics {
	return &GoogleAnalytics{
		l:           l,
		dbConnMSSQL: dbConnMSSQL,
	}
}

type GoogleAnalyticsDao interface {
	InsertGoogleAnalyticsData(*[]dtos.GoogleAnalytics) error
}

func (ga *GoogleAnalytics) InsertGoogleAnalyticsData(googleAnalyticsIns *[]dtos.GoogleAnalytics) error {

	db := mssqlcon.NewDBConn()

	// db := getDb() // Get a sql.Db. You're on  the hook to do this part.

	// Create a new structable.Recorder and tell it to
	// bind the given struct as a row in the given table.
	r := structable.New(db, "mysql").Bind("test_table", googleAnalyticsIns)

	// This will insert the stool into the test_table.
	err := r.Insert()

	if err != nil {
		fmt.Println("------------------------------ : ", err)
	}

	fmt.Println("--------- : ", googleAnalyticsIns)

	return nil

}
