package HighFive

import (
	"models"
	"gopkg.in/volatiletech/null.v6"
	"log"
	"time"
	"math/rand"
	"strconv"
)

// Creates initial database entries for the Next 5 App
func (five *HighFive) InitData() (error){
	initMeeting := &models.Meeting{
		MeetingName: "Gold Coast",
	}
	err := initMeeting.Insert(five.DB)
	if err != nil{
		log.Println("Owner Insert error = ", err)
		panic(err)
	}

	// Inserts 5 races with different times and race types
	five.AddRace(initMeeting, "Greyhound", 1 * 10000000000)
	five.AddRace(initMeeting, "Harness", 2 * 10000000000)
	five.AddRace(initMeeting, "Thoroughbred", 3 * 10000000000)
	five.AddRace(initMeeting, "Greyhound", 4 * 10000000000)
	five.AddRace(initMeeting, "Harness", 5 * 10000000000)

	if err != nil {
		log.Println("Insert error = ", err)
		panic(err)
	}

	return nil
}

// Function to add races to the database
func (five *HighFive) AddRace(meeting *models.Meeting, raceType string, addedTime int) {
	// Gets and sets time for race
	now := time.Now().UTC()
	timeAdd := time.Duration(addedTime)
	now.Add(timeAdd)

	race := &models.Race {
		MeetingID: null.StringFrom(meeting.MeetingID),
		RaceType: raceType,
		ClosingTime: time.Now().UTC().Add(timeAdd),
	}

	race.Insert(five.DB)

	// Adds competitors to each of the races
	for i := 1; i < 5; i++ {
		name := "Racer " + strconv.Itoa(i)
		five.AddCompetitor(race, name)
	}
}

// Function creating competitors for each of the races
func (five *HighFive) AddCompetitor(race *models.Race, name string) {
	competitor := &models.Competitor {
		RaceID: null.StringFrom(race.RaceID),
		Type: race.RaceType,
		Name: null.StringFrom(name),
		Position: rand.Intn(20),
	}

	competitor.Insert(five.DB)
}


