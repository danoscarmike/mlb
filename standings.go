package mlb

import (
	"fmt"
	"strings"
	"time"
)

// GetDivisions returns all or specific divisions based on if any divisionIds are provided or not
func (m *Mlb) GetStandings(leagueIds ...int) ([]Record, error) {

	params := map[string]string{}

	if len(leagueIds) > 0 {
		params["leagueId"] = strings.Trim(strings.Join(strings.Split(fmt.Sprint(leagueIds), " "), ","), "[]")
	}

	resp, err := m.Call("/standings", params)

	if err != nil {
		return []Record{}, err
	}

	return resp.Records, nil
}

// GetDivisions returns all or specific divisions based on if any divisionIds are provided or not
func (m *Mlb) GetStandingsByDate(date time.Time, leagueIds ...int) ([]Record, error) {

	params := map[string]string{}

	if len(leagueIds) > 0 {
		params["leagueId"] = strings.Trim(strings.Join(strings.Split(fmt.Sprint(leagueIds), " "), ","), "[]")
	}
	params["season"] = date.Format("2006")
	params["date"] = date.Format("2006-01-02")

	resp, err := m.Call("/standings", params)

	if err != nil {
		return []Record{}, err
	}

	return resp.Records, nil
}
