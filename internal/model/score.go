package model

type UpdateScoreRequest struct {
	PlayerID string  `json:"player_id" binding:"required"`
	Score    float64 `json:"score" binding:"required"`
}

type PlayerRankResponse struct {
	PlayerID string  `json:"player_id"`
	Score    float64 `json:"score"`
	Rank     int64   `json:"rank"`
}

type TopPlayer struct {
	PlayerID string  `json:"player_id"`
	Score    float64 `json:"score"`
	Rank     int64   `json:"rank"`
}
