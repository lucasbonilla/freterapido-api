package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
	freterapidoapi "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/freterapidoapi"
)

type Adapter struct {
	DB     *sql.DB
	config ports.Config
}

type Carrier struct {
	IDCarrier   int    `json:"id_carrier"`
	CarrierName string `json:"carrier_name"`
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

func (dbA *Adapter) Close() error {
	return dbA.DB.Close()
}

func (dbA *Adapter) Ping() error {
	return dbA.DB.Ping()
}

func (dbA *Adapter) AddCarrier(offers []freterapidoapi.Offers) error {
	var insertSQL string = `INSERT INTO freterapidoapi.carrier (id_carrier, carrier_name) VALUES`
	values := []interface{}{}
	for index, offer := range offers {
		insertSQL += fmt.Sprintf(" ($%d, $%d)", index*2+1, index*2+2)
		values = append(values, offer.Carrier.Reference, offer.Carrier.Name)
		if index < len(offers)-1 {
			insertSQL += ","
		}
	}
	insertSQL += " ON CONFLICT DO NOTHING;"
	_, err := dbA.DB.Exec(insertSQL, values...)
	if err != nil {
		return err
	}
	return nil
}

func (dbA *Adapter) AddQuote(offers []freterapidoapi.Offers) error {
	var insertSQL string = `INSERT INTO freterapidoapi."quote" (id_carrier,	price_quote) VALUES`
	values := []interface{}{}
	for index, offer := range offers {
		insertSQL += fmt.Sprintf(" ($%d, $%d)", index*2+1, index*2+2)
		values = append(values, offer.Carrier.Reference, offer.CostPrice)
		if index < len(offers)-1 {
			insertSQL += ","
		}
	}
	insertSQL += ";"
	_, err := dbA.DB.Exec(insertSQL, values...)
	if err != nil {
		return err
	}

	return nil
}
