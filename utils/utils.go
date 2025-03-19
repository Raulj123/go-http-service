package utils

import (
	"database/sql"
	"fmt"
	"time"
)

func GetCurrentDate() string {
	today := time.Now().Format("2006-01-02T15:04:05.000")
	fmt.Println(today)
	return today
}

func GetFutureDate() string {
	future := time.Now().AddDate(0,0,14).Format("2006-01-02T15:04:05.000")
	return future
}

func StringtoNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{Valid: true}
}
