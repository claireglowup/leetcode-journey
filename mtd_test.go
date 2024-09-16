package main

import (
	"fmt"
	"os"
	"sort"
	"testing"
	"time"
)

//minimum time difference - medium

func TestMtd(t *testing.T) {

	timePoints := []string{"23:57", "00:00"}

	if len(timePoints) < 2 {
		os.Exit(8)
	}

	minutesList := make([]int, len(timePoints))
	for i, t := range timePoints {
		minutesList[i] = timeToMinutes(t)
	}

	sort.Ints(minutesList)

	minDifference := 1440

	for i := 1; i < len(minutesList); i++ {
		diff := minutesList[i] - minutesList[i-1]
		if diff < minDifference {
			minDifference = diff
		}
	}

	circularDiff := 1440 - (minutesList[len(minutesList)-1] - minutesList[0])
	if circularDiff < minDifference {
		minDifference = circularDiff
	}

	fmt.Println(minDifference)
}

func timeToMinutes(t string) int {
	parsedTime, _ := time.Parse("15:04", t)
	return parsedTime.Hour()*60 + parsedTime.Minute()
}
