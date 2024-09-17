package main

import (
	"fmt"
	"time"
)

func main() {

	parsedTime, _ := time.Parse("15:04", "23:09")
	fmt.Println(parsedTime.Hour()*60 + parsedTime.Minute())
}
