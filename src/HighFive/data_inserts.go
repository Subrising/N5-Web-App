package HighFive

import (
	"models"
	"gopkg.in/volatiletech/null.v6"
	"log"
	"time"
	"math/rand"
	"strconv"
)

func (five *HighFive) InitData() (error){
	log.Println("Yo")
	initMeeting := &models.Meeting{
		MeetingName: "Gold Coast",
	}
	err := initMeeting.Insert(five.DB)
	if err != nil{
		log.Println("Owner Insert error = ", err)
		panic(err)
	}

	log.Println("Yo 2")
	five.AddRace(initMeeting, "Greyhound", 1 * 10000000000)
	five.AddRace(initMeeting, "Harness", 2 * 10000000000)
	five.AddRace(initMeeting, "Thoroughbred", 3 * 10000000000)
	five.AddRace(initMeeting, "Greyhound", 4 * 10000000000)
	five.AddRace(initMeeting, "Harness", 5 * 10000000000)
	log.Println("Yo 3")

	if err != nil {
		log.Println("Insert error = ", err)
		panic(err)
	}

	return nil
}

func (five *HighFive) AddRace(meeting *models.Meeting, raceType string, addedTime int) {
	log.Println("Race yo 1")
	now := time.Now().UTC()
	log.Println("Added time = ",  time.Duration(addedTime))

	timeAdd := time.Duration(addedTime)

	now.Add(timeAdd)
	
	race := &models.Race {
		MeetingID: null.StringFrom(meeting.MeetingID),
		RaceType: raceType,
		ClosingTime: time.Now().UTC().Add(timeAdd),
	}

	race.Insert(five.DB)
	log.Println("Race yo 3")

	for i := 1; i < 5; i++ {
		name := "Racer " + strconv.Itoa(i)
		five.AddCompetitor(race, name)
	}
	log.Println("Race yo 4")
}

func (five *HighFive) AddCompetitor(race *models.Race, name string) {
	competitor := &models.Competitor {
		RaceID: null.StringFrom(race.RaceID),
		Type: race.RaceType,
		Name: null.StringFrom(name),
		Position: rand.Intn(20),
	}

	competitor.Insert(five.DB)
}


