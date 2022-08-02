package main

import (
	"fmt"
	"time"

	//_ "demon/common"
	//_ "demon/module/socket"
	log "github.com/sirupsen/logrus"
	"strings"
)

const TimeFormat = "2006-01-02 15:04:05.000000000"

func main() {
	log.Info("demon main")
	temp := "AntiProxyLog_20220727"
	dateStr := strings.Replace(temp, "AntiProxyLog_", "", -1)
	fmt.Println(dateStr)
	date, err := ParseTime(dateStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(date)
}

func ParseTime(str string) (time.Time, error) {
	var (
		date time.Time
		err  error
	)
	// 2006/01/02 15:10:19
	if strings.Contains(str, "/") {
		date, err = time.ParseInLocation("2006/1/2 15:04:05", str, time.Local)
	} else if strings.Contains(str, "-") {
		date, err = time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	} else if len(str) == 8 {
		date, err = time.ParseInLocation("20060102", str, time.Local)
	} else {
		date, err = time.ParseInLocation("20060102 15:04", str, time.Local)
	}

	if err != nil {
		return date, err
	}
	return date, nil
}

func convertTime(t time.Time, loc *time.Location) time.Time {
	y, mon, d := t.Date()
	h, m, s := t.Clock()
	return time.Date(y, mon, d, h, m, s, t.Nanosecond(), loc)
}
