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

func (dbA *Adapter) InitConn() {
	var err error
	conf := dbA.config.GetDB()
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Database,
	)
	fmt.Println(sc)

	dbA.DB, err = sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}
}

func (dbA *Adapter) Ping() error {
	return dbA.DB.Ping()
}
