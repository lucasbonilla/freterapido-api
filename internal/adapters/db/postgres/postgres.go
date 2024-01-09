package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
)

type Adapter struct {
	DB     *sql.DB
	config ports.Config
}

func NewAdapter(config ports.Config) *Adapter {
	return &Adapter{config: config}
}

func (dbA *Adapter) InitConn() error {
	var err error
	conf := dbA.config.GetDB()
	stringConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Database,
	)
	fmt.Println(stringConn)

	dbA.DB, err = sql.Open("postgres", stringConn)
	if err != nil {
		return err
	}

	return nil
}

func (dbA *Adapter) Ping() error {
	return dbA.DB.Ping()
}
