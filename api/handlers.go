package api

import (
	"log"
	"net/http"

	"github.com/KibetBrian/geth/admin"
	"github.com/gin-gonic/gin"
)

func (s *Server) SayHello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"Message":"Hello there!"})
	return
}

func (s *Server) AddVoter(c *gin.Context){

	var req AddVoterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("Invalid json format")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	voters := admin.GetVoters()
	registered := VoterRegistered(voters,req.Address)
	if registered{
		c.JSON(http.StatusForbidden, gin.H{"message": "The address is already registered", "success": false})
		return
	}

	if req.Admin != electionAdmin{
		c.JSON(http.StatusForbidden, gin.H{"Message": "Only admins are allowed to do this operation", "success": false})
		return
	}
	tx := admin.AddVoter(req.Address)
	c.JSON(http.StatusOK,tx)
}

func (s *Server) GetVoters(c *gin.Context){
	voters := admin.GetVoters()
	c.JSON(http.StatusOK, gin.H{"Voters": voters})
	return
}
