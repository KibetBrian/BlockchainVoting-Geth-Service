package api

import "github.com/gin-gonic/gin"


func NewServer () *Server{

	server := Server{
		router: gin.Default(),
	}

	server.router.GET("/", server.SayHello)
	server.router.GET("/voter/all", server.GetVoters)
	server.router.POST("/voter/add", server.AddVoter)
	

	return &server
}