package main

import "log"

type item struct {
	raceID   string
	racename string
	name     string
	votes    int32
}

type race struct {
	raceID   string
	racename string
	racers
}

type racers []racer

type racer struct {
	name  string
	votes int32
}

func main() {
	items := []item{
		item{"1", "racename1", "scooby", 1000},
		item{"1", "racename1", "scrappy", 9000},
		item{"1", "racename1", "shaggy", 1000},
		item{"2", "racename2", "shaggy", 450},
		item{"2", "racename2", "fred", 550},
	}
	races := []race{}
	for _, i := range items {
		log.Println("i:", i)
		races = append(races, race{i.raceID, i.racename, []racer{}})

	}

	log.Println("races:", races)
}
