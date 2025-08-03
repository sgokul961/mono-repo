package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sgokul961/GO-PROJECT/pkg/handler"
	"github.com/sgokul961/GO-PROJECT/pkg/routes"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	adminHandler *handler.AdminHandler,

) *ServerHTTP {
	gin.SetMode(gin.ReleaseMode) //for production mode
	// Create a new Gin engine
	engine := gin.New()
	engine.Use(gin.Recovery()) // Recover from panics and log them
	engine.Use(gin.Logger())

	// Swagger UI (serve static files)
	//engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Attach routes
	routes.AdminRoutes(engine.Group("/admin"), adminHandler)
	//routes.UserRoutes(r.PathPrefix("/user").Subrouter(), userHandler, otpHandler, inventoryHandler, cartHandler, orderHandler, paymentHandler)

	return &ServerHTTP{
		engine: engine,
	}
}

func (s *ServerHTTP) Start() {
	log.Println("Server starting on port 3001...")
	s.engine.Run(":3001") // Start the server on port 3001
}
