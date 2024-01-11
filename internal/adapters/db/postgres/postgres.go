package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/lucasbonilla/freterapido-api/internal/ports"
	APIResp "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/api"
	freterapidoapi "github.com/lucasbonilla/freterapido-api/internal/schemas/message/response/freterapidoapi"
)

const (
	step = 2
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
		insertSQL += fmt.Sprintf(" ($%d, $%d)", index*2+1, index*2+step)
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
		insertSQL += fmt.Sprintf(" ($%d, $%d)", index*2+1, index*2+step)
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

func (dbA *Adapter) GetNumberOfQuotes(limit *int, offset *int) (APIResp.QuotesQuantity, error) {
	var quotesQuantity APIResp.QuotesQuantity
	var selectSQL string = `
	SELECT
		c.carrier_name,
		q.id_carrier,
		COUNT(*) AS number_of_quotes
	FROM
		freterapidoapi."quote" q
	JOIN freterapidoapi.carrier c ON
		q.id_carrier = c.id_carrier
	GROUP BY
		c.carrier_name,
		q.id_carrier
	ORDER BY
		number_of_quotes DESC `
	selectSQL = addLimitOffset(selectSQL, limit, offset)

	rows, err := dbA.DB.Query(selectSQL)
	if err != nil {
		return quotesQuantity, err
	}
	defer rows.Close()

	for rows.Next() {
		var quote APIResp.QuotesByQuantity
		if err := rows.Scan(&quote.CarrierName, &quote.IDCarrier, &quote.NumberOfQuotes); err != nil {
			return quotesQuantity, err
		}
		quotesQuantity.QuotesByQuantity = append(quotesQuantity.QuotesByQuantity, quote)
	}
	if err = rows.Err(); err != nil {
		return quotesQuantity, err
	}

	return quotesQuantity, nil
}

func (dbA *Adapter) GetTotalQuotes(limit *int, offset *int) (APIResp.TotalQuotesPrice, error) {
	var totalQuotes APIResp.TotalQuotesPrice
	var selectSQL string = `
	SELECT
		c.carrier_name,
		q.id_carrier,
		ROUND(SUM(price_quote)::NUMERIC, 2) AS total_price_quote
	FROM
		freterapidoapi."quote" q
	JOIN freterapidoapi.carrier c ON
		q.id_carrier = c.id_carrier
	GROUP BY
		c.carrier_name,
		q.id_carrier
	ORDER BY
		total_price_quote DESC `
	selectSQL = addLimitOffset(selectSQL, limit, offset)

	rows, err := dbA.DB.Query(selectSQL)
	if err != nil {
		return totalQuotes, err
	}
	defer rows.Close()

	for rows.Next() {
		var quote APIResp.TotalQuotesByPrice
		if err := rows.Scan(&quote.CarrierName, &quote.IDCarrier, &quote.TotalPriceQuote); err != nil {
			return totalQuotes, err
		}
		totalQuotes.TotalQuotesByPrice = append(totalQuotes.TotalQuotesByPrice, quote)
	}
	if err = rows.Err(); err != nil {
		return totalQuotes, err
	}

	return totalQuotes, nil
}

func (dbA *Adapter) GetAverageQuotes(limit *int, offset *int) (APIResp.TotalQuotesAveragePrice, error) {
	var totalQuotes APIResp.TotalQuotesAveragePrice
	var selectSQL string = `
	SELECT
			c.carrier_name,
			q.id_carrier,
			ROUND(AVG(price_quote)::NUMERIC, 2) AS average_price_quote
	FROM
			freterapidoapi."quote" q
	JOIN freterapidoapi.carrier c ON
			q.id_carrier = c.id_carrier
	GROUP BY
			c.carrier_name,
			q.id_carrier
	ORDER BY
			average_price_quote DESC `
	selectSQL = addLimitOffset(selectSQL, limit, offset)

	rows, err := dbA.DB.Query(selectSQL)
	if err != nil {
		return totalQuotes, err
	}
	defer rows.Close()

	for rows.Next() {
		var quote APIResp.TotalQuotesByAveragePrice
		if err := rows.Scan(&quote.CarrierName, &quote.IDCarrier, &quote.AveragePriceQuote); err != nil {
			return totalQuotes, err
		}
		totalQuotes.TotalQuotesByAveragePrice = append(totalQuotes.TotalQuotesByAveragePrice, quote)
	}
	if err = rows.Err(); err != nil {
		return totalQuotes, err
	}

	return totalQuotes, nil
}

func (dbA *Adapter) GetCheapestQuotes(limit *int, offset *int) (APIResp.TotalQuotesCheapestPrice, error) {
	var totalQuotes APIResp.TotalQuotesCheapestPrice
	var selectSQL string = `
	SELECT
		c.carrier_name,
		q.id_carrier,
		MIN(price_quote) AS price_quote_cheapest
	FROM
		freterapidoapi."quote" q
	JOIN freterapidoapi.carrier c ON
		q.id_carrier = c.id_carrier
	GROUP BY
		c.carrier_name,
		q.id_carrier
	ORDER BY
		price_quote_cheapest ASC `
	selectSQL = addLimitOffset(selectSQL, limit, offset)

	rows, err := dbA.DB.Query(selectSQL)
	if err != nil {
		return totalQuotes, err
	}
	defer rows.Close()

	for rows.Next() {
		var quote APIResp.TotalQuotesForCheapestPrice
		if err := rows.Scan(&quote.CarrierName, &quote.IDCarrier, &quote.PriceQuoteCheapest); err != nil {
			return totalQuotes, err
		}
		totalQuotes.TotalQuotesForCheapestPrice = append(totalQuotes.TotalQuotesForCheapestPrice, quote)
	}
	if err = rows.Err(); err != nil {
		return totalQuotes, err
	}

	return totalQuotes, nil
}

func (dbA *Adapter) GetMostExpensiveQuotes(limit *int, offset *int) (APIResp.TotalQuoteMostExpensivePrice, error) {
	var totalQuotes APIResp.TotalQuoteMostExpensivePrice
	var selectSQL string = `
	SELECT
		c.carrier_name,
		q.id_carrier,
		MAX(price_quote) AS price_quote_most_expensive
	FROM
		freterapidoapi."quote" q
	JOIN freterapidoapi.carrier c ON
		q.id_carrier = c.id_carrier
	GROUP BY
		c.carrier_name,
		q.id_carrier
	ORDER BY
		price_quote_most_expensive DESC `
	selectSQL = addLimitOffset(selectSQL, limit, offset)

	rows, err := dbA.DB.Query(selectSQL)
	if err != nil {
		return totalQuotes, err
	}
	defer rows.Close()

	for rows.Next() {
		var quote APIResp.TotalQuoteByMostExpensivePrice
		if err := rows.Scan(&quote.CarrierName, &quote.IDCarrier, &quote.PriceQuoteMostExpensive); err != nil {
			return totalQuotes, err
		}
		totalQuotes.TotalQuoteByMostExpensivePrice = append(totalQuotes.TotalQuoteByMostExpensivePrice, quote)
	}
	if err = rows.Err(); err != nil {
		return totalQuotes, err
	}

	return totalQuotes, nil
}

func addLimitOffset(selectSQL string, limit *int, offset *int) string {
	if limit != nil {
		selectSQL += fmt.Sprintf("LIMIT %v ", *limit)
		if offset != nil {
			selectSQL += fmt.Sprintf("OFFSET %v", *offset)
		}
	}
	selectSQL += ";"

	return selectSQL
}
