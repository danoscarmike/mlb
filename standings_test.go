package mlb_test

import (
	"testing"
	"time"

	"github.com/stevepartridge/mlb"
)

const (
	invalidLeagueId1 = 999
)

func Test_GetStandings_Success(t *testing.T) {

	standings, err := mlbApi.GetStandings()
	if err != nil {
		t.Error("Should not error when retrieving standings")
	}

	if len(standings) > 0 {
		t.Error("Should be zero when no param (leagueId) is passed")
	}

}

func Test_GetStanding_Success(t *testing.T) {
	leagueIds := []int{103, 104}
	standings, err := mlbApi.GetStandings(leagueIds...)
	if err != nil {
		t.Error("Should not error when retrieving standings")
	}

	for i := range standings {
		t.Logf("%s", standings[i].Division.League.Name)
	}

	if len(standings) != 6 {
		t.Logf("Found: %d; Expected: 6\n", len(standings))
		t.Error("Should be exactly 6 division standings found")
	}

}

func Test_GetStandings_SpecificLeague_Success(t *testing.T) {

	standings, err := mlbApi.GetStandings(mlb.NationalLeague)
	if err != nil {
		t.Error("Should not error when retrieving league standings")
	}

	if len(standings) != 3 {
		t.Error("Should be exactly 1 division found.")
	}

}

func Test_GetStandings_Failure_InvalidLeagueId(t *testing.T) {

	divisions, err := mlbApi.GetStandings(invalidLeagueId1)
	if err != nil {
		t.Error("Should not error when retrieving invalid league")
	}

	if len(divisions) != 0 {
		t.Error("Zero divisions should be found.")
	}

}

func Test_GetStandingsByDate_Success(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2022-10-05")
	divisions, err := mlbApi.GetStandingsByDate(date, mlb.NationalLeague)
	if err != nil {
		t.Error("Should not error when retrieving league standings by date")
	}

	if len(divisions) != 3 {
		t.Error("Should be exactly 3 divisions found.")
	}

	if divisions[0].TeamRecords[0].Team.Name != "Atlanta Braves" {
		t.Error("First team record should be Atlanta Braves.")
	}

	if divisions[2].TeamRecords[2].LeagueRecord.Percent != ".500" {
		t.Error("San Franscisco Giants record should be .500")
	}
}
