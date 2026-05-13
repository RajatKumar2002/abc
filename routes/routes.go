package routes

import (
	"abc/admin"
	"github.com/julienschmidt/httprouter"
)

func RegisterRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/match-list/:sport/:league/:season", admin.GetMatchList)

	return router
}

