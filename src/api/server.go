package api

func (a *api) StartHttpServer() {

	a.hs.Listen()

}
