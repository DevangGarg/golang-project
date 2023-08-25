package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/techschool/simplebank/db/sqlc"
)

// Server serves HTTP requests for our banking services
type Server struct {
	store db.Store
	router *gin.Engine
}

// New Server creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v,ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// add routes to router
	router.POST("/accounts",server.createAccount)
	router.GET("/accounts/:id",server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id",server.deleteAccount)
	router.POST("/accounts/:id",server.updateAccount)

	router.POST("/transfers",server.createTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server on specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error" : err.Error()}
}