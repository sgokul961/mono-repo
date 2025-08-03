package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sgokul961/GO-PROJECT/pkg/handler"
)

func AdminRoutes(
	r *gin.RouterGroup,
	adminHandler *handler.AdminHandler,

) {
	// Unprotected route
	r.POST("/adminlogin", adminHandler.LoginHandler)

	// Create a subrouter for protected routes
	//adminSubRouter := r.PathPrefix("").Subrouter()
	//adminSubRouter.Use(middleware.AdminAuthMiddleware)

	// // /users group
	// adminSubRouter.HandleFunc("/users/block", adminHandler.BlockUser).Methods("POST")
	// adminSubRouter.HandleFunc("/users/unblock", adminHandler.UnblockUser).Methods("POST")
	// adminSubRouter.HandleFunc("/users/getusers", adminHandler.GetUsers).Methods("GET")

}
