package models

import (
	"cmd/ticktock/utils/logger"
	"errors"
	"net/http"
	"time"
)

type CalcTimestamps func(loc time.Location, t1, t2 time.Time) ([]string, error)

const layout = "20060102T150405Z"

var funcMap = map[string]CalcTimestamps{
	"1h":  calculateTimestampsByHour,
	"1d":  calculateTimestampsByDay,
	"1mo": calculateTimestampsByMonth,
	"1y":  calculateTimestampsByYear,
}

func GetTimestamps(reqID, period, tz, t1Str, t2Str string) ([]string, int, error) {
	log := logger.NewLogger(reqID)
	log.Info("Test dokimi")

	t1, err := time.Parse(layout, t1Str)
	if err != nil {
		log.Debug("Error: Invalid t1 format")
		return nil, http.StatusBadRequest, errors.New("Invalid t1 format")
	}
	t2, err := time.Parse(layout, t2Str)
	if err != nil {
		log.Debug("Error: Invalid t2 format")
		return nil, http.StatusBadRequest, errors.New("Invalid t2 format")
	}

	tzLoc, err := time.LoadLocation(tz)
	if err != nil {
		log.Debug("Error: Invalid timezone.  Err", err)
		return nil, http.StatusBadRequest, errors.New("Invalid time zone. Err: " + err.Error())
	}

	var tsmps []string
	t1 = t1.In(time.UTC)
	t2 = t2.In(time.UTC)
	if t1.After(t2) {
		log.Debug("Error: t1 is after t2.")
		return nil, http.StatusBadRequest, errors.New("Invalid t1 and t2. t1 is after t2")
	}

	calcTimestamps, found := funcMap[period]
	if !found {
		log.Debug("Error: Unsupported period")
		return nil, http.StatusBadRequest, errors.New("Unsupported period")
	}

	t1 = t1.Round(time.Hour)
	log.Debug("T1 round: ", t1)
	tsmps, err = calcTimestamps(*tzLoc, t1, t2)
	if err != nil {
		log.Error("Error: Could not calculate timestamps")
		return nil, http.StatusBadRequest, errors.New("Could not calculate timestamps")
	}

	return tsmps, http.StatusOK, nil
}

//calculateTimestampsByDay calculates the timestamps for a period of a duration. Duration should be less/equal than a day.
func calculateTimestampsByDuration(duration time.Duration, loc time.Location, t1, t2 time.Time) ([]string, error) {
	var tsmps []string

	//reference Summer time. (DST period)
	refSummerTime := time.Date(2023, 6, 1, 0, 0, 0, 0, &loc)
	_, refOffset := refSummerTime.Zone()

	_, tmp := t1.In(&loc).Zone()
	dif := refOffset - tmp
	//fix hour if needed
	dts := t1.Add(time.Duration(dif) * time.Second)
	tsmps = append(tsmps, dts.Format(layout))

	for {

		t1 = t1.Add(duration)
		_, tmp := t1.In(&loc).Zone()
		dif := refOffset - tmp
		//fix hour if needed
		dts = t1.Add(time.Duration(dif) * time.Second)

		if t1.After(t2) {
			break
		}
		tsmps = append(tsmps, dts.Format(layout))
	}
	return tsmps, nil
}

//calculateTimestampsByHour calculates the timestamps for a period of a hour.
func calculateTimestampsByHour(loc time.Location, t1, t2 time.Time) ([]string, error) {
	return calculateTimestampsByDuration(time.Hour, loc, t1, t2)
}

//calculateTimestampsByDay calculates the timestamps for a period of a Day.
func calculateTimestampsByDay(loc time.Location, t1, t2 time.Time) ([]string, error) {
	return calculateTimestampsByDuration(24*time.Hour, loc, t1, t2)
}

//calculateTimestampsByMonth calculates the timestamps for a period of a month.
func calculateTimestampsByMonth(loc time.Location, t1, t2 time.Time) ([]string, error) {
	var tsmps []string
	initTime := t1
	//reference Summer time. (DST period)
	refSummerTime := time.Date(2023, 6, 1, 0, 0, 0, 0, &loc)
	_, refOffset := refSummerTime.Zone()
	for {

		//find last day of the month
		firstOfMonth := time.Date(t1.Year(), t1.Month(), 1, initTime.Hour(), initTime.Minute(), initTime.Second(), 0, time.UTC)
		t1 = firstOfMonth.AddDate(0, 1, -1)

		_, tmp := t1.In(&loc).Zone()
		dif := refOffset - tmp
		//fix hour if needed

		dts := t1.Add(time.Duration(dif) * time.Second)
		////move to next month
		t1 = t1.AddDate(0, 0, 1)
		if t1.After(t2) {
			break
		}
		tsmps = append(tsmps, dts.Format(layout))
	}
	return tsmps, nil
}

//calculateTimestampsByYear calculates the timestamps for a period of a year.
func calculateTimestampsByYear(loc time.Location, t1, t2 time.Time) ([]string, error) {
	var tsmps []string

	refSummerTime := time.Date(2023, 6, 1, 0, 0, 0, 0, &loc)
	_, refOffset := refSummerTime.Zone()
	_, tmp := t1.In(&loc).Zone()
	dif := refOffset - tmp

	t1 = time.Date(t1.Year(), 12, 31, t1.Hour(), t1.Minute(), t1.Second(), 0, time.UTC)
	//fix hour if needed
	t1 = t1.Add(time.Duration(dif) * time.Second)
	tsmps = append(tsmps, t1.Format(layout))

	for {
		//increase t1 by a year
		t1 = t1.AddDate(1, 0, 0)
		if t1.After(t2) {
			break
		}
		tsmps = append(tsmps, t1.Format(layout))
	}
	return tsmps, nil
}
