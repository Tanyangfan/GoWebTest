package lib

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DbConnect = new(sqlx.DB)

func InitDbConnect() {
	connString := fmt.Sprintf("host=192.168.2.235 user=postgres password=mima123 dbname=postgres sslmode=disable")
	conn, err := sqlx.Open("postgres", connString)
	if err != nil {
		LogFatal("InitDbConnect", fmt.Sprintf("Open connection  failed:%s", err.Error()))
	}
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(5)
	DbConnect = conn
}
