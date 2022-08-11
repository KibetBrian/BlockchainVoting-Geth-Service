package api

import (
	"log"
	"net/http"

	"github.com/KibetBrian/geth/admin"
	"github.com/KibetBrian/geth/election"
	"github.com/gin-gonic/gin"
)

func (s *Server) SayHello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"Message":"Hello there!"})
}

func (s *Server) RegisterVoter(c *gin.Context){

	var req AddVoterRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("Invalid json format")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	
	registrationPhase := admin.GetRegistrationPhase();
	if !registrationPhase{
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "registration phase not active"})
		return;
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
	tx := admin.RegisterVoter(req.Address)
	c.JSON(http.StatusOK,tx)
}

func (s *Server) RegisterCandidate(c *gin.Context){
	var req AddContestantRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Invalid json format"})
		return
	}
	registrationPhase := admin.GetRegistrationPhase();
	if !registrationPhase{
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "registration phase not active"})
		return;
	}

	presidentialCandidates := admin.GetGorvernorCandidates();
	gubernitorialGandidates := admin.GetGorvernorCandidates();
	
	isRegistered  := CandidateRegistered(presidentialCandidates, gubernitorialGandidates,req.CandidatesAddress)
	if isRegistered{
		c.JSON(http.StatusConflict, gin.H{"message":"Candidate already registered"})
		return
	}

	if req.Admin!= admin.Admin{
		c.JSON(http.StatusUnauthorized, gin.H{"message":"You are not allowed to do this operation because you are not admin"})
		return
	}

	res := admin.RegisterCandidate(req.CandidatesAddress, req.CandidatesName, req.Position)	
	c.JSON(http.StatusOK,res)
}

func (s *Server) GetVoters(c *gin.Context){
	voters := admin.GetVoters()
	c.JSON(http.StatusOK, gin.H{"Voters": voters})
}

func (s *Server)GetSpecificCandidate(c *gin.Context){
	candidate := admin.GetSpecificCandidate("0xd67edd183254c4e1274b31a7f01b4de859a4db0d")
	c.JSON(http.StatusOK, gin.H{"Candidate": candidate})
}

func (s *Server) GetGorvernorCandidates(c *gin.Context){
	candidates := admin.GetGorvernorCandidates();
	c.JSON(http.StatusOK, gin.H{"Candidates": candidates})
}
func (s *Server) GetPresedentialCandidates(c *gin.Context){
	candidates := admin.GetPresidentCandidates();
	c.JSON(http.StatusOK, gin.H{"Candidates": candidates})
}

func (s *Server) ChangeRegistrationPhase(c *gin.Context){
	tx := admin.ChangeRegistrationPhase();
	c.JSON(http.StatusOK, gin.H{"message": "registration phase changed","Transaction":tx });
}

func (s *Server) ChangeVotingPhase(c *gin.Context){
	tx := admin.ChangeVotingPhase();
	c.JSON(http.StatusOK, gin.H{"message": "voting phase changed", "Transaction": tx});
}

func (s *Server) GetVotingPhase(c *gin.Context){
	state := admin.GetVotingPhase();
	c.JSON(http.StatusOK, gin.H{"votingPhase": state})
}

func (s *Server) GetRegisrationPhase(c *gin.Context){
	state := admin.GetRegistrationPhase();
	c.JSON(http.StatusOK, gin.H{"registrationPhase": state})
}

type Addresses struct{
	Addresses []string `json:"candidatesAddresses"`
}

//Process
func (s *Server)GetProcessedResults(c *gin.Context){
	var addresses Addresses;
	err := c.ShouldBindJSON(&addresses);
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}

	candidates := []election.ElectionCandidate{}

	//Get individual candidate data
	for _, v := range addresses.Addresses{
		candidate := admin.GetSpecificCandidate(v)
		candidates = append(candidates, candidate)		
	}
	c.JSON(http.StatusOK, candidates)
}


