package main

import (
	"flag"
	"fmt"
	"github.com/foize/go.sgr"
	"github.com/imroc/req"
	"github.com/joho/godotenv"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Stores the contents of file as a string
type dataProcessing struct {
	Path    string
	Name    string
	HTML    string
	Content string
}

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

// Read a file given its path/name
func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func processFile(filePath string) {
	fileName := strings.Split(filePath, ".txt")[0]
	namedHTML := fileName + ".html"
	fileContents := readFile(filePath)

	info := dataProcessing{
		Path:    filePath,
		Name:    fileName,
		HTML:    namedHTML,
		Content: fileContents,
	}

	applyTemplate("template.tmpl", info)
}

// Create an Html File Based off the provided template and processed data
func applyTemplate(path string, data dataProcessing) {
	t := template.Must(template.New(path).ParseFiles(path))
	newFile, _ := os.Create(data.HTML)
	// err := t.Execute(os.Stdout, data)
	err := t.Execute(newFile, data)
	if err != nil {
		panic(err)
	}
}

func searchDirectory(dirPath string) {
	libRegEx, e := regexp.Compile("^.+\\.(txt)$")
	if e != nil {
		log.Fatal(e)
	}
	count, size := 0, 0.0

	e = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) {
			count = count + 1
			fmt.Println(info.Name(), float64(info.Size()), float64(os.Getpagesize()))
			// size = size + float64(info.Size())
			size = size + float64(os.Getpagesize())/1000.0
			processFile(info.Name())
		}
		return nil
	})
	if e != nil {
		log.Fatal(e)
	}
	sgr.Printf("[fg-green bold] Success! [reset] Generated [bold] %d [reset] pages (%6.1fkB total). \n", count, size)
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
	t, e := time.Parse(`2006-01-02T15:04:05.000-07:00`, date) //
	if e != nil {
		panic(e)
	}

	return int(t.UTC().UnixNano() / 1000000)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var dirPath, filePath, newFlyer string
	flag.StringVar(&dirPath, "dir", "", "Directory Path")
	flag.StringVar(&filePath, "file", "", "Name or Path to a text file")
	flag.StringVar(&newFlyer, "flyer", "", "Create A New Chess Tournament and Flyer")
	flag.Parse()

	switch {
	case dirPath != "":
		searchDirectory(dirPath)
	case filePath != "":
		processFile(filePath)
	case newFlyer != "":
		timeStamp := getTimeStamp(`2021-02-21T14:00:00.000-08:00`)
		createTournamentAndFlyer(timeStamp)
	default:
		fmt.Print("No Option Selected")
	}
}
