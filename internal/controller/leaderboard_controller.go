package controller

import (
	"net/http"
	"strconv"

	"leaderboard/internal/model"
	"leaderboard/internal/service"

	"github.com/gin-gonic/gin"
)

type LeaderboardController struct {
	svc *service.LeaderboardService
}

func NewLeaderboardController(s *service.LeaderboardService) *LeaderboardController {
	return &LeaderboardController{svc: s}
}

func (ctl *LeaderboardController) UpdateScore(c *gin.Context) {
	var req model.UpdateScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := ctl.svc.UpdateScore(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, resp)
}

func (ctl *LeaderboardController) TopN(c *gin.Context) {
	nStr := c.Query("n")
	if nStr == "" {
		nStr = "10"
	}

	n, err := strconv.ParseInt(nStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid n parameter"})
		return
	}

	resp, err := ctl.svc.GetTopN(c, n)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp) //gin不需要return resp，直接操作*gin.Context就行
}

func (ctl *LeaderboardController) GetRankByPlayerID(c *gin.Context) {

	playerID := c.Param("player_id")
	if playerID == "" {
		c.JSON(400, gin.H{"error": "player_id required"})
		return
	}

	resp, err := ctl.svc.GetRankByPlayerID(c, playerID)
	if err != nil {
		c.JSON(500, gin.H{"ServiceError": err.Error()})
		return
	}

	c.JSON(200, resp)
}

func (ctl *LeaderboardController) Reset(c *gin.Context) {
	if err := ctl.svc.Reset(c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"reset": true})
}
