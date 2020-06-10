package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchRank(steamID int64) (int, error) {
	id := steamID - 76561197960265728
	url := fmt.Sprintf("https://api.opendota.com/api/players/%v", id)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var result struct {
		RankTier int `json:"rank_tier"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result.RankTier, nil
}
