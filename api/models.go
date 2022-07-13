package api

type AddContestantRequest struct {
	Admin string `json:"admin" binding:"required"`
	CandidatesAddress string `json:"candidatesAddress" binding:"required"`
	CandidatesName string `json:"candidatesName" binding:"required"`
	Position string `json:"position" binding:"required"`
}