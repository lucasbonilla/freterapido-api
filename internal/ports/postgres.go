package ports

type Postgres interface {
	InitConn()
	Ping() error
}
