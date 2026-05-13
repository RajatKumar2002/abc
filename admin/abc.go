package admin

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func sendResponse(w http.ResponseWriter, statusCode int, data interface{}, _ string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	resp := APIResponse{
		License:    "copyright@2020 tcp",
		StatusCode: statusCode,
		Content:    data,
	}
	json.NewEncoder(w).Encode(resp)
}
func BindMatch(row MatchRow) MatchInfo {
	obj := MatchInfo{}
	obj.ID = row.MatchID
	if row.Date != nil {
		obj.Date = *row.Date
	}
	if row.Time != nil {
		obj.Time = *row.Time
	}
	if row.CounterDate != nil {
		obj.CounterDate = *row.CounterDate
	}
	if row.CounterTime != nil {
		obj.CounterTime = *row.CounterTime
	}
	if row.Status != nil {
		obj.Status = *row.Status
	}
	if row.HomeID != nil {
		obj.HomeTeam.ID = *row.HomeID
	}
	if row.HomeName != nil {
		obj.HomeTeam.Name = *row.HomeName
	}
	if row.HomeURL != nil {
		obj.HomeTeam.URL = *row.HomeURL
	}
	if row.HomeAbbr != nil {
		obj.HomeTeam.Abbr = *row.HomeAbbr
	}
	if row.HomeH1 != nil {
		obj.HomeTeam.Score.H1 = *row.HomeH1
	}
	if row.HomeH2 != nil {
		obj.HomeTeam.Score.H2 = *row.HomeH2
	}
	if row.HomeET != nil {
		obj.HomeTeam.Score.ET = *row.HomeET
	}
	if row.HomeOT != nil {
		obj.HomeTeam.Score.OT = *row.HomeOT
	}
	if row.HomeScore != nil {
		obj.HomeTeam.Score.TotalScore = *row.HomeScore
	}
	if row.AwayID != nil {
		obj.AwayTeam.ID = *row.AwayID
	}
	if row.AwayName != nil {
		obj.AwayTeam.Name = *row.AwayName
	}
	if row.AwayURL != nil {
		obj.AwayTeam.URL = *row.AwayURL
	}
	if row.AwayAbbr != nil {
		obj.AwayTeam.Abbr = *row.AwayAbbr
	}
	if row.AwayH1 != nil {
		obj.AwayTeam.Score.H1 = *row.AwayH1
	}
	if row.AwayH2 != nil {
		obj.AwayTeam.Score.H2 = *row.AwayH2
	}
	if row.AwayET != nil {
		obj.AwayTeam.Score.ET = *row.AwayET
	}
	if row.AwayOT != nil {
		obj.AwayTeam.Score.OT = *row.AwayOT
	}
	if row.AwayScore != nil {
		obj.AwayTeam.Score.TotalScore = *row.AwayScore
	}
	return obj
}

func GetMatchList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sport := strings.TrimSpace(ps.ByName("sport"))
	league := strings.TrimSpace(ps.ByName("league"))
	season := strings.TrimSpace(ps.ByName("season"))

	sportID, sportName, err := GetSportID(sport)
	if err != nil {
		sendResponse(w, http.StatusNotFound, nil, err.Error())
		return
	}
	leagueID, leagueName, err := GetLeagueID(league, sportID)
	if err != nil {
		sendResponse(w, http.StatusNotFound, nil, err.Error())
		return
	}
	seasonID, err := GetSeasonID(season, sportID, leagueID)
	if err != nil {
		sendResponse(w, http.StatusNotFound, nil, err.Error())
		return
	}
	matchRows, err := GetMatches(leagueID, seasonID)
	if err != nil {
		sendResponse(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	var matches []MatchInfo
	for _, row := range matchRows {
		matches = append(matches, BindMatch(row))
	}

	var objSport SportInfo
	objSport.ID = sportID
	objSport.Name = sportName
	objSport.URL = sport

	var objLeague LeagueInfo
	objLeague.ID = leagueID
	objLeague.Name = leagueName
	objLeague.URL = league

	var objContent MatchListContent
	objContent.Sport = objSport
	objContent.League = objLeague
	objContent.MatchInfo = matches

	sendResponse(w, http.StatusOK, objContent, "")
}
