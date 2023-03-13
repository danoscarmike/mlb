package mlb

type Response struct {
	Copyright     string `json:"copyright"`
	Message       string `json:"message,omitempty"`
	MessageNumber int    `json:"messageNumber,omitempty"`

	TotalItems           int        `json:"totalItems,omitempty"`
	TotalEvents          int        `json:"totalEvents,omitempty"`
	TotalGames           int        `json:"totalGames,omitempty"`
	TotalGamesInProgress int        `json:"totalGamesInProgress,omitempty"`
	Dates                []Date     `json:"dates,omitempty"`
	Teams                []Team     `json:"teams,omitempty"`
	Venues               []Venue    `json:"venues,omitempty"`
	Divisions            []Division `json:"divisions,omitempty"`
	Records              []Record   `json:"records,omitempty"`
}

type Date struct {
	Date                 string `json:"date"`
	TotalItems           int    `json:"totalItems"`
	TotalEvents          int    `json:"totalEvents"`
	TotalGames           int    `json:"totalGames"`
	TotalGamesInProgress int    `json:"totalGamesInProgress"`
	Games                []Game `json:"games"`
}

type Game struct {
	Id      int                 `json:"gamePk"`
	Link    string              `json:"link"`
	Type    string              `json:"gameType"`
	Season  string              `json:"season"`
	Date    string              `json:"gameDate"`
	Status  GameStatus          `json:"status"`
	Teams   map[string]GameTeam `json:"teams"`
	Venue   Venue               `json:"venue"`
	Content struct {
		Link string `json:"link"`
	} `json:"content"`
	SeriesStatus           SeriesStatus `json:"seriesStatus,omitempty"`
	DoubleHeader           string       `json:"doubleHeader"`
	GamedayType            string       `json:"gamedayType"`
	Tiebreaker             string       `json:"tiebreaker"`
	GameNumber             int          `json:"gameNumber"`
	CalendarEventID        string       `json:"calendarEventID"`
	SeasonDisplay          string       `json:"seasonDisplay"`
	DayNight               string       `json:"dayNight"`
	ScheduledInnings       int          `json:"scheduledInnings"`
	GamesInSeries          int          `json:"gamesInSeries"`
	SeriesGameNumber       int          `json:"seriesGameNumber"`
	SeriesDescription      string       `json:"seriesDescription"`
	RecordSource           string       `json:"recordSource"`
	IfNecessary            string       `json:"ifNecessary"`
	IfNecessaryDescription string       `json:"ifNecessaryDescription"`
	Events                 []Event      `json:""`
}

type GameStatus struct {
	AbstractGameState string `json:"abstractGameState"`
	AbstractGameCode  string `json:"abstractGameCode"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     string `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
	StartTimeTbd      bool   `json:"startTimeTBD,omitempty"`
}

type SeriesStatus struct {
	GameNumber       int        `json:"gameNumber"`
	TotalGames       int        `json:"totalGames"`
	IsTied           bool       `json:"isTied"`
	IsOver           bool       `json:"isOver"`
	Wins             int        `json:"wins"`
	Losses           int        `json:"losses"`
	WinningTeam      SeriesTeam `json:"winningTeam"`
	LosingTeam       SeriesTeam `json:"losingTeam"`
	Description      string     `json:"description"`
	ShortDescription string     `json:"shortDescription"`
	Result           string     `json:"result"`
	ShortName        string     `json:"shortName"`
}

type SeriesTeam struct {
	SpringLeague  League `json:"springLeague"`
	AllStarStatus string `json:"allStarStatus"`
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Link          string `json:"link"`
}

type Venue struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type Event struct{}

type GameTeam struct {
	LeagueRecord LeagueRecord `json:"leagueRecord"`
	Team         Team         `json:"team"`
	IsWinner     bool         `json:"isWinner"`
	Score        int          `json:"score"`
	SplitSquad   bool         `json:"splitSquad"`
	SeriesNumber int          `json:"seriesNumber"`
}

type Team struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Link            string   `json:"link"`
	Venue           Venue    `json:"venue"`
	TeamCode        string   `json:"teamCode"`
	FileCode        string   `json:"fileCode"`
	Abbreviation    string   `json:"abbreviation"`
	TeamName        string   `json:"teamName"`
	LocationName    string   `json:"locationName"`
	FirstYearOfPlay string   `json:"firstYearOfPlay"`
	League          League   `json:"league"`
	Division        Division `json:"division"`
	Sport           Sport    `json:"sport"`
	ShortName       string   `json:"shortName"`
	SpringLeague    League   `json:"springLeague"`
	Active          bool     `json:"active"`
}

type LeagueRecord struct {
	Wins    int    `json:"wins"`
	Losses  int    `json:"losses"`
	Ties    int    `json:"ties,omitempty"`
	Percent string `json:"pct"`
	League  League `json:"league,omitempty"`
}

type Sport struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type League struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
}

type Division struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	NameShort    string `json:"nameShort,omitempty"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation,omitempty"`
	League       League `json:"league,omitempty"`
	Sport        Sport  `json:"sport,omitempty"`
	HasWildcard  bool   `json:"hasWildcard,omitempty"`
}

type Record struct {
	StandingsType string       `json:"standingsType"`
	League        League       `json:"league,omitempty"`
	Division      Division     `json:"division,omitempty"`
	Sport         Sport        `json:"sport,omitempty"`
	LastUpdated   string       `json:"lastUpdated,omitempty"`
	TeamRecords   []TeamRecord `json:"teamRecords,omitempty"`
}

type Streak struct {
	StreakType   string `json:"streakType,omitempty"`
	StreakNumber int    `json:"streakNumber,omitempty"`
	StreakCode   string `json:"streakCode,omitempty"`
}

type TeamRecordDetails struct {
	SplitRecords    []SplitRecord    `json:"splitRecords,omitempty"`
	DivisionRecords []DivisionRecord `json:"divisionRecords,omitempty"`
	OverallRecords  []SplitRecord    `json:"overallRecords,omitempty"`
	LeagueRecords   []LeagueRecord   `json:"leagueRecords,omitempty"`
	ExpectedRecords []SplitRecord    `json:"expectedRecords,omitempty"`
}

type SplitRecord struct {
	Wins   int    `json:"wins,omitempty"`
	Losses int    `json:"losses,omitempty"`
	Type   string `json:"type,omitempty"`
	Pct    string `json:"pct,omitempty"`
}

type DivisionRecord struct {
	Wins     int      `json:"wins,omitempty"`
	Losses   int      `json:"losses,omitempty"`
	Pct      string   `json:"pct,omitempty"`
	Division Division `json:"division,omitempty"`
}

type TeamRecord struct {
	Team                      Team              `json:"team"`
	Season                    string            `json:"season"`
	Streak                    Streak            `json:"streak,omitempty"`
	ClinchIndicator           string            `json:"clinchIndicator,omitempty"`
	DivisionRank              string            `json:"divisionRank,omitempty"`
	LeagueRank                string            `json:"leagueRank,omitempty"`
	SportRank                 string            `json:"sportRank,omitempty"`
	GamesPlayed               int               `json:"gamesPlayed,omitempty"`
	GamesBack                 string            `json:"gamesBack,omitempty"`
	WildCardGamesBack         string            `json:"wildCardGamesBack,omitempty"`
	LeagueGamesBack           string            `json:"leagueGamesBack,omitempty"`
	SpringLeagueGamesBack     string            `json:"springLeagueGamesBack,omitempty"`
	SportGamesBack            string            `json:"sportGamesBack,omitempty"`
	DivisionGamesBack         string            `json:"divisionGamesBack,omitempty"`
	ConferenceGamesBack       string            `json:"conferenceGamesBack,omitempty"`
	LeagueRecord              LeagueRecord      `json:"leagueRecord,omitempty"`
	LastUpdated               string            `json:"lastUpdated,omitempty"`
	Records                   TeamRecordDetails `json:"records,omitempty"`
	RunsAllowed               int               `json:"runsAllowed,omitempty"`
	RunsScored                int               `json:"runsScored,omitempty"`
	DivisionChamp             bool              `json:"divisionChamp,omitempty"`
	DivisionLeader            bool              `json:"divisionLeader,omitempty"`
	HasWildcard               bool              `json:"hasWildCard,omitempty"`
	Clinched                  bool              `json:"clinched,omitempty"`
	EliminationNumber         string            `json:"eliminationNumber,omitempty"`
	WildCardEliminationNumber string            `json:"wildCardEliminationNumber,omitempty"`
	MagicNumber               string            `json:"magicNumber,omitempty"`
	Wins                      int               `json:"wins,omitempty"`
	Losses                    int               `json:"losses,omitempty"`
	RunDifferential           int               `json:"runDifferential,omitempty"`
	WinningPercentage         string            `json:"winningPercentage,omitempty"`
}
