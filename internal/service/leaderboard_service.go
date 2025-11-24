package service

import (
	"context"
	"fmt"
	"leaderboard/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type LeaderboardService struct {
	rdb *redis.Client
}

func NewLeaderboardService(r *redis.Client) *LeaderboardService {
	return &LeaderboardService{rdb: r}
}

func (s *LeaderboardService) UpdateScore(ctx context.Context, req model.UpdateScoreRequest) (*model.PlayerRankResponse, error) {
	newScore, err := s.rdb.ZIncrBy(ctx, "leaderboard", req.Score, req.PlayerID).Result()
	if err != nil {
		return nil, err
	}

	rank, err := s.rdb.ZRevRank(ctx, "leaderboard", req.PlayerID).Result()
	if err != nil {
		return nil, err
	}

	return &model.PlayerRankResponse{
		PlayerID: req.PlayerID,
		Score:    newScore,
		Rank:     rank + 1,
	}, nil
}

// func (s *LeaderboardService) GetPlayerRank()
func (s *LeaderboardService) GetTopN(c *gin.Context, n int64) ([]model.TopPlayer, error) {
	result, err := s.rdb.ZRevRangeWithScores(c, "leaderboard", 0, n-1).Result()
	if err != nil {
		return nil, err
	}

	top := make([]model.TopPlayer, 0, len(result))
	for i, v := range result {
		top = append(top, model.TopPlayer{
			PlayerID: v.Member.(string),
			Score:    v.Score,
			Rank:     int64(i + 1),
		})
	}
	return top, nil

}

func (s *LeaderboardService) GetRankByPlayerID(c *gin.Context, id string) (*model.PlayerRankResponse, error) {
	rank, err := s.rdb.ZRevRank(c, "leaderboard", id).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("player not found")
	} else if err != nil {
		return nil, err
	}

	score, err := s.rdb.ZScore(c, "leaderboard", id).Result()
	if err != nil {
		return nil, err
	}

	return &model.PlayerRankResponse{
		PlayerID: id,
		Score:    score,
		Rank:     rank + 1,
	}, nil

}

func (s *LeaderboardService) Reset(c *gin.Context) error {
	_, err := s.rdb.Del(c, "leaderboard").Result()
	return err
}
