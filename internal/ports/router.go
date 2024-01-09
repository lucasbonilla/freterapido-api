package ports

type Router interface {
	InitializeRoutes()
	Serve(listenAdr string) error
}
