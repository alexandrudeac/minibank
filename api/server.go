package api

import (
	"github.com/gin-gonic/gin"
	db "gitlab.com/alexandrudeac/minibank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store}
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	return &Server{
		store:  store,
		router: router,
	}
}
func (server *Server) Run(addr string) error {
	return server.router.Run(addr)
}
