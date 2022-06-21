package api

func (a *api) StartHttpServer() {

	a.registerRoutes()

	a.hs.Listen()

}
