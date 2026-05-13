package admin

import (
	"abc/data"
	"database/sql"
	"fmt"
)

func GetSportID(sport string) (int, string, error) {
	var sportID int
	var sportName string
	query := `
		SELECT sport_id, sport_name
		FROM isg_sports
		WHERE sport_url = ?
		AND status = 1
		LIMIT 1`
	err := data.DB.QueryRow(query, sport).Scan(&sportID, &sportName)
	if err != nil {
	if err == sql.ErrNoRows {
		return 0, "", fmt.Errorf("sport '%s' not found", sport)
	}
	return 0, "", fmt.Errorf("query error for sport: %v", err)
	}
	return sportID, sportName, nil
}

func GetLeagueID(league string, sportID int) (int, string, error) {
	var leagueID int
	var leagueName string
	if sportID == 7 {
	query := `
		SELECT league_id, league_name
		FROM isg_rugby_league
		WHERE league_url = ?
		AND status = 1
		LIMIT 1`
	err := data.DB.QueryRow(query, league).Scan(&leagueID, &leagueName)
	if err != nil {
	    if err == sql.ErrNoRows {
			return 0, "", fmt.Errorf("league '%s' not found", league)
		}
			return 0, "", fmt.Errorf("query error: %v", err)
	    }
	}
	return leagueID, leagueName, nil
}

func GetSeasonID(season string, sportID int, leagueID int) (int, error) {
	var seasonID int
	query := `
	SELECT season_id
	FROM isg_sports_league_seasons
	WHERE season_url = ?
	AND sport_id = ? AND league_id = ? AND status = 1
	LIMIT 1`
	err := data.DB.QueryRow(query, season, sportID, leagueID).Scan(&seasonID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("season '%s' not found for this sport and league", season)
		}
		return 0, fmt.Errorf("query error for season: %v", err)
	}
	return seasonID, nil
}

func GetMatches(leagueID int, seasonID int) ([]MatchRow, error) {
	query := `
	SELECT
	m.match_id, m.match_date, m.match_time, m.counter_date, m.counter_time, m.status,
	ht.team_id, ht.team_name, ht.url, ht.abbreviation,
	awt.team_id, awt.team_name, awt.url, awt.abbreviation,
	s.home_score, s.away_score, s.home_h1, s.home_h2, s.home_et, s.home_ot,
	s.away_h1, s.away_h2, s.away_et, s.away_ot
	FROM isg_rugby_league_matches m
	JOIN isg_team ht ON m.home_team_id = ht.team_id
	JOIN isg_team awt ON m.away_team_id = awt.team_id
	LEFT JOIN isg_rugby_league_matches_scores s ON m.match_id = s.match_id
	WHERE m.league_id = ?
	AND m.season_id = ?
	ORDER BY m.status DESC, m.match_date DESC, m.match_time DESC`

	rows, err := data.DB.Query(query, leagueID, seasonID)
	if err != nil {
		return nil, fmt.Errorf("query error for matches: %v", err)
	}
	defer rows.Close()
	var matchRows []MatchRow
	for rows.Next() {
		var row MatchRow
		err := rows.Scan(
			&row.MatchID, &row.Date, &row.Time, &row.CounterDate, &row.CounterTime, &row.Status,
			&row.HomeID, &row.HomeName, &row.HomeURL, &row.HomeAbbr,
			&row.AwayID, &row.AwayName, &row.AwayURL, &row.AwayAbbr,
			&row.HomeScore, &row.AwayScore,
			&row.HomeH1, &row.HomeH2, &row.HomeET, &row.HomeOT,
			&row.AwayH1, &row.AwayH2, &row.AwayET, &row.AwayOT,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		matchRows = append(matchRows, row)
	}
	return matchRows, nil
}
