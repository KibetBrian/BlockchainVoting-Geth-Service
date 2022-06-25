package api

import (
	"log"
	"net/http"
	"github.com/KibetBrian/geth/admin"

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
func (s *Server) SayHello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"Message":"Hello there!"})
	return
}

func (s *Server) GetVoters(c *gin.Context){
	voters := admin.GetVoters()
	c.JSON(http.StatusOK, gin.H{"Voters": voters})
	return
}

func (s *Server) AddVoter(c *gin.Context){

	var req AddVoterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if req.Admin != electionAdmin{
		log.Println(req.Admin)
		log.Println(electionAdmin)
		c.JSON(http.StatusForbidden, "Only admins are allowed to do this operation")
		return
	}
	tx := admin.AddVoter(req.Address)
	c.JSON(http.StatusOK,tx)
}