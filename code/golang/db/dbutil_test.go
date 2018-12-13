package dbutil

import (
	"log"
	"testing"
	"time"

	Config "../config"
	_ "github.com/lib/pq"
)

func dbOperation(dbHost, dbUser, dbPasswd, dbName string, t *testing.T) {
	conn := New(dbHost, dbUser, dbPasswd, dbName)
	defer conn.Close()

	users := conn.Rows("users", "id, openid, nickname, created")
	defer users.Close()
	for users.Next() {
		var id, openid, nickname string
		var created time.Time
		if err := users.Scan(&id, &openid, &nickname, &created); err != nil {
			t.Error(err)
		}
		log.Println(id, openid, nickname, created.Format(time.RFC3339))
	}
}

func TestDBWithYaml(t *testing.T) {
	log.Println("[test yaml]")
	_, err := Config.GetConfig("../config/database.yml")
	if err != nil {
		t.Error(err)
	}

	// dbHost, dbUser, dbPasswd, dbName := cfg.Get("development.host"),
	// 	cfg.Get("development.username"),
	// 	cfg.Get("development.password"),
	// 	cfg.Get("development.database")
	//
	// dbOperation(dbHost, dbUser, dbPasswd, dbName, t)
}

func TestDBWithToml(t *testing.T) {
	log.Println("[test toml]")
	cfg, err := Config.GetTomlConfig("../config/database.toml")
	if err != nil {
		t.Error(err)
	}

	dev := cfg.Env["development"]
	dbOperation(dev.Host, dev.Username, dev.Password, dev.Database, t)
}
