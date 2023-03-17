package mlb_test

import (
	"testing"
	"time"

	"github.com/stevepartridge/mlb"
)

const (
	validDate1   = "2017/05/17"
	invalidDate1 = "00/01/01"
	validTeam1   = mlb.Dodgers
)

func Test_GetGamesByDate_Success(t *testing.T) {

	date, _ := time.Parse("2006/01/02", validDate1)

	games, err := mlbApi.GetGamesByDate(date, false)
	if err != nil {
		t.Error("Should not error: " + err.Error())
	}

	if len(games) == 0 {
		t.Error("Games should be more than zero.")
	}
}

func Test_GetGamesByDate_Failure(t *testing.T) {

	date, _ := time.Parse("2006/01/02", invalidDate1)

	games, err := mlbApi.GetGamesByDate(date, false)
	if err == nil {
		t.Error("Should be an error")
	}

	if len(games) > 0 {
		t.Error("Games should not be more than zero.")
	}
}

func Test_GetGamesByDateRange_Failure_InvalidEndDate(t *testing.T) {

	start, _ := time.Parse("2006/01/02", validDate1)
	end, _ := time.Parse("2006/01/02", invalidDate1)

	games, err := mlbApi.GetGamesByDateRange(start, end, false)
	if err == nil {
		t.Error("Should be an error")
	}

	if len(games) > 0 {
		t.Error("Games should not be more than zero.")
	}
}

func Test_GetGamesByDateForTeam_Success(t *testing.T) {

	date, _ := time.Parse("2006/01/02", validDate1)

	games, err := mlbApi.GetGamesByDateForTeam(date, validTeam1, false)
	if err != nil {
		t.Error("Should not error: " + err.Error())
	}

	if len(games) == 0 {
		t.Error("Games should be more than zero.")
	}
}

func Test_GetGamesByDateRangeForTeam_Success(t *testing.T) {

	start, _ := time.Parse("2006/01/02", validDate1)
	end, _ := time.Parse("2006/01/02", validDate1)

	games, err := mlbApi.GetGamesByDateRangeForTeam(start, end, validTeam1, false)
	if err != nil {
		t.Error("Should not error: " + err.Error())
	}

	if len(games) == 0 {
		t.Error("Games should be more than zero.")
	}
}

func Test_GetGamesByDateWithHydrate_Success(t *testing.T) {

	date, _ := time.Parse("2006/01/02", validDate1)

	games, err := mlbApi.GetGamesByDate(date, true)
	if err != nil {
		t.Error("Should not error: " + err.Error())
	}

	if len(games) == 0 {
		t.Error("Games should be more than zero.")
	}

	if games[7].WinningTeam == games[7].LosingTeam {
		t.Error("WinningTeam should not equal LosingTeam.")
	}
}
