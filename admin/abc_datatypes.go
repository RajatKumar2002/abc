package admin

type APIResponse struct {
	License   string      `json:"license"`
	StatusCode  int         `json:"status-code"`
	Content    interface{} `json:"content"`
}

type SportInfo struct {
	ID  int    `json:"id"`
	Name  string `json:"name"`
	URL  string `json:"url"`
}

type LeagueInfo struct {
	ID  int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}
type ScoreInfo struct {
	H1 int `json:"h1"`
	H2 int `json:"h2"`
	ET int `json:"et"`
    OT int `json:"ot"`
	TotalScore int `json:"total_score"`
}

type HomeTeam struct {
	ID   int    `json:"id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Abbr   string `json:"abbr"`
	Score   ScoreInfo   `json:"home-score"`
}

type AwayTeam struct {
	ID    int    `json:"id"`
	Name   string `json:"name"`
	URL   string `json:"url"`
	Abbr   string `json:"abbr"`
	Score  ScoreInfo   `json:"away-score"`
}

type MatchInfo struct {
	ID      int      `json:"id"`
	Date    string   `json:"date"`
	Time    string   `json:"time"`
	CounterDate  string   `json:"counter_date"`
	CounterTime  string   `json:"counter_time"`
	Status      string   `json:"status"`
	HomeTeam    HomeTeam `json:"home_team"`
	AwayTeam    AwayTeam `json:"away_team"`
}

type MatchRow struct {
	MatchID     int
	Date       *string
	Time       *string
	CounterDate   *string
	CounterTime   *string
	Status      *string
	HomeID      *int
	HomeName     *string
	HomeURL     *string
	HomeAbbr     *string
	HomeScore    *int
	HomeH1      *int
	HomeH2      *int
	HomeET      *int
	HomeOT      *int
	AwayID      *int
	AwayName     *string
	AwayURL     *string
	AwayAbbr    *string
	AwayScore   *int
	AwayH1      *int
	AwayH2      *int
	AwayET      *int
	AwayOT      *int
}
type MatchListContent struct {
	Sport     SportInfo   `json:"sport"`
	League    LeagueInfo  `json:"league"`
	MatchInfo []MatchInfo `json:"match_info"`
}
