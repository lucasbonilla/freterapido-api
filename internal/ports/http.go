package ports

type Http interface {
	InitializeRoutes()
	Serve(listenAdr string) error
}
