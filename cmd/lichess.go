package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/imroc/req"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lichessCommand)
}

var lichessCommand = &cobra.Command{
	Use:   "lichess",
	Short: "Make API Calls to lichess.org",
	Long:  "Interact with lichess.org's API to manage makeschool chess club events",
	Run: func(cmd *cobra.Command, args []string) {
		timeStamp := getTimeStamp(`2021-02-21T14:00:00.000-08:00`)
		createTournamentAndFlyer(timeStamp)
	},
}

func createTournamentAndFlyer(date int) {
	apiKey := os.Getenv("LICHESS_TOKEN")

	// Create a new session
	session := req.New()

	header := req.Header{
		"Accept":        "application/json",
		"Authorization": "Bearer " + apiKey,
	}

	params := req.Param{
		"name":           "MakeSchool",
		"clockTime":      10,
		"clockIncrement": 0,
		"minutes":        90,
		"waitMinutes":    60,
		"startDate":      date,
		"variant":        "standard", //"standard" "chess960" "crazyhouse" "antichess" "atomic" "horde" "kingOfTheHill" "racingKings" "threeCheck"
		"rated":          true,
		// "position": ""
		"berserkable": false,
		"streakable":  false,
		"hasChat":     true,
		// "description": "",
		// "password": "",
		// "teambBattleByTeam": "",
		// "conditions.teamMember.teamId": "",
		// "conditions.minRating.rating": 500,
		// "conditions.maxRating.rating": 1200,
		// "conditions.nbRatedGame.nb": 0,
	}

	creationResponse, err := session.Post("https://lichess.org/api/tournament", header, params)
	// Make sure we get a 200 status code from our request
	if creationResponse.Response().StatusCode != 200 || err != nil {
		log.Fatalf("Error Code: %d This request is invalid because we got a non-ok status code...\n %s",
			creationResponse.Response().StatusCode, creationResponse)
	}

	fmt.Print(creationResponse)
}

func getTimeStamp(date string) int {
	// https://play.golang.org/p/ouiDtIVjQI
	t, e := time.Parse(`2006-01-02T15:04:05.000-07:00`, date)
	if e != nil {
		panic(e)
	}

	return int(t.UTC().UnixNano() / 1000000)
}
