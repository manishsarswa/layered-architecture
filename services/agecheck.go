package services

import (
	"strings"
	"time"
)
//date in seconds takes date as a parameter in (DD/MM/YYYY) format and will return time in seconds
func dateInSeconds(d1 string) int {
	Slice := strings.Split(d1, "/")

	newDate := Slice[2] + "/" + Slice[1] + "/" + Slice[0]
	myDate, ok := time.Parse("2006/01/02", newDate)

	if ok != nil {
		panic(ok)
	}

	return int(time.Now().Unix() - myDate.Unix())
}
