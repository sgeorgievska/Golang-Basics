package main

import "fmt"

type racer struct {
	racerName     string
	raceNo        int
	currentTeam   string
	previousTeams []string
}

func main() {
	aRacer := racer{
		racerName:   "Sebastian Vettel",
		raceNo:      5,
		currentTeam: "Ferrari",
		previousTeams: []string{
			"RedBull",
			"Toro Rosso",
		},
	}
	fmt.Println(aRacer.racerName)
}
