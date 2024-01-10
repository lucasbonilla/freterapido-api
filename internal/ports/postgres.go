package ports

type Postgres interface {
	InitConn() error
	Close() error
	Ping() error
}
