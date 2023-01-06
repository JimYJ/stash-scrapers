package sqlite

import (
	"fmt"
	"stash-scrapers/common/config"
	"stash-scrapers/services/log"
	"time"

	"github.com/jmoiron/sqlx"
	// https://github.com/mattn/go-sqlite3
)

func Conn() *sqlx.DB {
	url := fmt.Sprintf("file:%s?_journal=WAL&_sync=NORMAL", config.SqliteInfo.Path)
	conn, err := sqlx.Open("sqlite3ex", url)
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(4)
	conn.SetConnMaxLifetime(30 * time.Second)
	if err != nil {
		log.Fatalf("db.Open(): %w", err)
		return nil
	}
	return conn
}
