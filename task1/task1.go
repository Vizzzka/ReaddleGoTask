package task1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	layoutISO      = "2006-01-02"
	layoutUS       = "January 2, 2006"
	nextHolidayURL = "https://date.nager.at/Api/v2/NextPublicHolidays/UA"
)

type holiday struct {
	Date      string
	LocalName string
}

func isTodayHoliday(date time.Time) bool {
	currentDate := time.Now().Format(layoutISO)
	holidayDate := date.Format(layoutISO)
	return currentDate == holidayDate
}

func adjacentWeekend(date time.Time) time.Time {
	if date.Weekday() == time.Friday {
		return date.AddDate(0, 0, 2)
	}
	if date.Weekday() == time.Saturday {
		return date.AddDate(0, 0, 1)
	}
	return date
}

func stringToDate(date string) time.Time {
	t, _ := time.Parse(layoutISO, date)
	return t
}

func getNextHoliday() holiday {
	// get next holidays in UA
	resp, err := http.Get(nextHolidayURL)
	if err != nil {
		log.Fatalln(err)
	}
	// close body to avoid leaks
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	// get
	keys := make([]holiday, 0)
	json.Unmarshal(body, &keys)
	return keys[0] // return closest holiday
}

func Work() {
	var result string
	nextHoliday := getNextHoliday()
	nextHolidayDate, nextHolidayName := stringToDate(nextHoliday.Date), nextHoliday.LocalName

	if isTodayHoliday(nextHolidayDate) {
		result = fmt.Sprintf("Today holiday is %s, %s.",
			nextHolidayName, nextHolidayDate.Format(layoutUS))
	} else {
		result = fmt.Sprintf("The next holiday is %s, %s.",
			nextHolidayName, nextHolidayDate.Format(layoutUS))
	}

	endOfWeekend := adjacentWeekend(nextHolidayDate)
	if nextHolidayDate != endOfWeekend {
		result += fmt.Sprintf(" Weekend will last from %s till %s.",
			nextHolidayDate.Format(layoutUS), endOfWeekend.Format(layoutUS))
	}
	fmt.Print(result)
}
