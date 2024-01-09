package ports

type Postgres interface {
	InitConn() error
	Ping() error
}
