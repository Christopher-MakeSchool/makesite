package cmd

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/chrisbarnes2000/makesite/models"
	"github.com/imroc/req"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lichessCommand)
	lichessCommand.AddCommand(linksCommand)
}

var lichessCommand = &cobra.Command{
	Use:   "lichess",
	Short: "Make API Calls to lichess.org",
	Long:  "Interact with lichess.org's API to manage makeschool chess club events",
	Run: func(cmd *cobra.Command, args []string) {
		createTournamentAndFlyer(getTimeStamp(args[0])) //  `2021-03-07T14:00:00.000-08:00`
	},
}

var linksCommand = &cobra.Command{
	Use:   "links",
	Short: "List links related to chess",
	Long:  "List links related to Make School's Chess Club",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("https://lichess.org/ads")
		fmt.Println("https://lichess.org/lag")
		fmt.Println("https://lichess.org/faq")
		fmt.Println("https://lichess.org/api")
		fmt.Println("----------------------------------")
		fmt.Println("https://lichess.org/@/cbarnes2000/tournaments/created")
		fmt.Println("----------------------------------")
		fmt.Println("https://lichess.org/blog/V0KrLSkAAMo3hsi4/study-chess-the-lichess-way")
		fmt.Println("https://lichess.org/analysis")
		fmt.Println("https://lichess.org/learn")
		fmt.Println("https://lichess.org/practice")
		fmt.Println("----------------------------------")
		fmt.Println("https://discord.gg/MAR5NJYpfv")
		fmt.Println("#chess-club")
		fmt.Println("https://lichess.org/team/makeschool")
		fmt.Println("https://chess.com/club/make-school/join")
		fmt.Println("https://twitch.tv/make_school")
		fmt.Println("https://instagram.com/make_school")
		fmt.Println("----------------------------------")
		fmt.Println("https://ecf.octoknight.com/")
		fmt.Println("----------------------------------")
		fmt.Println("https://www.englishchess.org.uk/ecf-membership-partners-and-benefits/")
		fmt.Println("https://www.englishchess.org.uk/ecf-membership-rates-and-joining-details/")
		fmt.Println("----------------------------------")
		fmt.Println("https://rotherhamonlinechess.azurewebsites.net/tournaments")
		fmt.Println("----------------------------------")
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

	tournamentInfo := &models.Tournament{}
	creationResponse.ToJSON(tournamentInfo)
	// fmt.Print(tournamentInfo)
	applyFlyerTemplate("flyer-template.tmpl", tournamentInfo)
}

func getTimeStamp(date string) int {
	// https://play.golang.org/p/ouiDtIVjQI
	t, e := time.Parse(`2006-01-02T15:04:05.000-07:00`, date)
	if e != nil {
		panic(e)
	}

	return int(t.UTC().UnixNano() / 1000000)
}

// Create an Html File Based off the provided template and processed data
func applyFlyerTemplate(path string, data *models.Tournament) {
	t := template.Must(template.New(path).ParseFiles(path))
	newFile, _ := os.Create("./flyers/"+string(data.StartsAt.Format(`2006-01-02T15:04`))+".html")
	// err := t.Execute(os.Stdout, data)
	err := t.Execute(newFile, data)
	if err != nil {
		panic(err)
	}
}
