package mlb

import (
	"errors"
	"strconv"
	"time"
)

// Get games by specific date.  Date is assumed to be relevant local to game timezone.
func (m *Mlb) GetGamesByDate(date time.Time, hydrate bool) ([]Game, error) {
	return m.GetGamesByDateRange(date, date, hydrate)
}

// Get games by specific date range.  Date is assumed to be relevant local to game timezone.
func (m *Mlb) GetGamesByDateRange(start, end time.Time, hydrate bool) ([]Game, error) {
	return m.GetGames(start, end, 0, hydrate)
}

// Get games by date for a specific team.  Date is assumed to be relevant local to game timezone.
func (m *Mlb) GetGamesByDateForTeam(date time.Time, teamId int, hydrate bool) ([]Game, error) {
	return m.GetGames(date, date, teamId, hydrate)
}

// Get games in date range for specific team.  Date is assumed to be relevant local to game timezone.
func (m *Mlb) GetGamesByDateRangeForTeam(start, end time.Time, teamId int, hydrate bool) ([]Game, error) {
	return m.GetGames(start, end, teamId, hydrate)
}

func (m *Mlb) GetGames(start, end time.Time, teamId int, hydrate bool) ([]Game, error) {

	// &season=2018&startDate=2018-08-01&endDate=2018-08-31&teamId=119&eventTypes=primary&scheduleTypes=games

	if start.IsZero() {
		return nil, errors.New("invalid start date")
	}

	if end.IsZero() {
		return nil, errors.New("invalid end date")
	}

	params := map[string]string{
		"season":        start.Format("2006"),
		"eventTypes":    "primary",
		"scheduleTypes": "games",
		"startDate":     start.Format("2006-01-02"),
		"endDate":       end.Format("2006-01-02"),
	}

	if params["sportId"] == "" {
		params["sportId"] = "1"
	}

	if teamId > 0 {
		params["teamId"] = strconv.Itoa(teamId)
	}

	if hydrate {
		params["hydrate"] = "seriesStatus"
	}

	resp, err := m.Call("/schedule", params)
	if err != nil {
		ifError(err)
		return []Game{}, err
	}

	if len(resp.Dates) == 0 {
		return []Game{}, err
	}

	games := make([]Game, len(resp.Dates))

	for _, date := range resp.Dates {
		for _, game := range date.Games {
			if hydrate {
				game.WinningTeam = game.SeriesStatus.WinningTeam.Name
				game.LosingTeam = game.SeriesStatus.LosingTeam.Name
			}
			games = append(games, game)
		}
	}

	return games, err
}
