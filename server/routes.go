package server

import (
	"github.com/gin-gonic/gin"
	"myGo/adapter/route"
	"myGo/handlers"
	"net/http"
)

func routes(engine *gin.Engine) {
	route.Route(engine, http.MethodGet, "/ping", handlers.PingHandler)
	route.Route(engine, http.MethodGet, "/user/info", handlers.UserInfoHandler)
}
