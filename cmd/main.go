package main

import (
	"leaderboard/internal/controller"
	"leaderboard/internal/service"
	myredis "leaderboard/pkg/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	rdb := myredis.NewClient()
	svc := service.NewLeaderboardService(rdb)
	ctl := controller.NewLeaderboardController(svc)

	r := gin.Default()

	r.POST("/api/v1/score/update", ctl.UpdateScore)
	// 后面的接口我们继续补
	r.GET("api/v1/leaderboard/top", ctl.TopN)
	r.GET("/api/v1/leaderboard/rank/:player_id", ctl.GetRankByPlayerID)
	r.DELETE("/api/v1/leaderboard/reset", ctl.Reset)

	r.Run(":8080")
}
