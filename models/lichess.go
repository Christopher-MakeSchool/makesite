package models

import "time"

// Tournament ... Holds the Respons Data From Creating a new tournament on lichess.org
// https://mholt.github.io/json-to-go/
type Tournament struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	Clock    struct {
		Increment int `json:"increment"`
		Limit     int `json:"limit"`
	} `json:"clock"`
	Minutes            int       `json:"minutes"`
	CreatedBy          string    `json:"createdBy"`
	System             string    `json:"system"`
	SecondsToStart     int       `json:"secondsToStart"`
	SecondsToFinish    int       `json:"secondsToFinish"`
	IsFinished         bool      `json:"isFinished"`
	IsRecentlyFinished bool      `json:"isRecentlyFinished"`
	PairingsClosed     bool      `json:"pairingsClosed"`
	StartsAt           time.Time `json:"startsAt"`
	NbPlayers          int       `json:"nbPlayers"`
	Perf               struct {
		Icon     string `json:"icon"`
		Key      string `json:"key"`
		Name     string `json:"name"`
		Position int    `json:"position"`
	} `json:"perf"`
	Schedule struct {
		Freq  string `json:"freq"`
		Speed string `json:"speed"`
	} `json:"schedule"`
	Variant struct {
		Key   string `json:"key"`
		Name  string `json:"name"`
		Short string `json:"short"`
	} `json:"variant"`
	Duels []struct {
		ID string `json:"id"`
		P  []struct {
			N string `json:"n"`
			R int    `json:"r"`
			K int    `json:"k"`
		} `json:"p"`
	} `json:"duels"`
	Standings struct {
		Page    int `json:"page"`
		Players []struct {
			Name   string `json:"name"`
			Rank   int    `json:"rank"`
			Rating int    `json:"rating"`
			Score  int    `json:"score"`
			Sheet  struct {
				Scores []interface{} `json:"scores"`
				Total  int           `json:"total"`
				Fire   bool          `json:"fire"`
			} `json:"sheet"`
		} `json:"players"`
	} `json:"standings"`
	Featured struct {
		ID       string `json:"id"`
		Fen      string `json:"fen"`
		Color    string `json:"color"`
		LastMove string `json:"lastMove"`
		White    struct {
			Rank   int    `json:"rank"`
			Name   string `json:"name"`
			Rating int    `json:"rating"`
		} `json:"white"`
		Black struct {
			Rank   int    `json:"rank"`
			Name   string `json:"name"`
			Rating int    `json:"rating"`
		} `json:"black"`
	} `json:"featured"`
	Podium []struct {
		Name   string `json:"name"`
		Rank   int    `json:"rank"`
		Rating int    `json:"rating"`
		Score  int    `json:"score"`
		Sheet  struct {
			Scores []interface{} `json:"scores"`
			Total  int           `json:"total"`
			Fire   bool          `json:"fire"`
		} `json:"sheet"`
		Nb struct {
			Game   int `json:"game"`
			Beserk int `json:"beserk"`
			Win    int `json:"win"`
		} `json:"nb"`
		Performance int `json:"performance"`
	} `json:"podium"`
	Stats struct {
		Games         int `json:"games"`
		Moves         int `json:"moves"`
		WhiteWins     int `json:"whiteWins"`
		BlackWins     int `json:"blackWins"`
		Draws         int `json:"draws"`
		Berserks      int `json:"berserks"`
		AverageRating int `json:"averageRating"`
	} `json:"stats"`
}
