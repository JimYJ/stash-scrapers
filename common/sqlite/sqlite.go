package sqlite

import (
	"fmt"
	"stash-scrapers/common/config"
	"stash-scrapers/services/log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	// https://github.com/mattn/go-sqlite3
)

var (
	conn *sqlx.DB
	once sync.Once
)

func Conn() *sqlx.DB {
	once.Do(func() {
		var err error
		url := fmt.Sprintf("file:%s?_journal=WAL&_sync=NORMAL", config.SqliteInfo.Path)
		conn, err = sqlx.Open("sqlite3", url)
		if err != nil {
			log.Fatalf("db.Open(): %w", err)
		}
		conn.SetMaxOpenConns(25)
		conn.SetMaxIdleConns(4)
		conn.SetConnMaxLifetime(30 * time.Second)
		conn.Ping()
	})
	return conn
}
