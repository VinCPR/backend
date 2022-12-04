package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	db "github.com/VinCPR/backend/db/sqlc"
	"github.com/VinCPR/backend/token"
	"github.com/VinCPR/backend/util"
)

type Server struct {
	config     util.Config
	tokenMaker token.ITokenMaker
	store      *db.Store
	router     *gin.Engine
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize token maker")
	}

	server := &Server{config: config, tokenMaker: tokenMaker, store: store}
	router := gin.Default()
	router.Use(CORS())
	router.POST("v1/users", server.createUser)
	router.POST("v1/users/login", server.loginUser)
	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	server.router = router
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
