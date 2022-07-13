package api

import "github.com/gin-gonic/gin"


func NewServer () *Server{

	server := Server{
		router: gin.Default(),
	}

	server.router.GET("/", server.SayHello)
	server.router.GET("/voter/all", server.GetVoters)
	server.router.POST("/voter/add", server.RegisterVoter)
	server.router.POST("/candidate/add", server.RegisterCandidate)
	server.router.GET("/candidates/governor", server.GetGorvernorCandidates)
	server.router.POST("/phase/registration", server.ChangeRegistrationPhase)
	server.router.POST("/phase/voting", server.ChangeVotingPhase)
	server.router.GET("/registration_phase", server.GetRegisrationPhase)
	server.router.GET("/voting_phase", server.GetVotingPhase)
	server.router.GET("/candidates/presidential", server.GetPresedentialCandidates)
	
	return &server
}