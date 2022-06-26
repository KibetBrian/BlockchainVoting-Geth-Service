package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

const electionAdmin = "brian"

type Server struct{
	router *gin.Engine
}

type AddVoterRequest struct{
	Admin string `json:"admin" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func (s *Server) Start(port string){
	err := s.router.Run(port)
	if err != nil {
		log.Fatalf("Server Start Failed. Error: %v", err)
	}
}
