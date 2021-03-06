// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Competitors", testCompetitors)
	t.Run("Meetings", testMeetings)
	t.Run("Races", testRaces)
}

func TestDelete(t *testing.T) {
	t.Run("Competitors", testCompetitorsDelete)
	t.Run("Meetings", testMeetingsDelete)
	t.Run("Races", testRacesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Competitors", testCompetitorsQueryDeleteAll)
	t.Run("Meetings", testMeetingsQueryDeleteAll)
	t.Run("Races", testRacesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Competitors", testCompetitorsSliceDeleteAll)
	t.Run("Meetings", testMeetingsSliceDeleteAll)
	t.Run("Races", testRacesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Competitors", testCompetitorsExists)
	t.Run("Meetings", testMeetingsExists)
	t.Run("Races", testRacesExists)
}

func TestFind(t *testing.T) {
	t.Run("Competitors", testCompetitorsFind)
	t.Run("Meetings", testMeetingsFind)
	t.Run("Races", testRacesFind)
}

func TestBind(t *testing.T) {
	t.Run("Competitors", testCompetitorsBind)
	t.Run("Meetings", testMeetingsBind)
	t.Run("Races", testRacesBind)
}

func TestOne(t *testing.T) {
	t.Run("Competitors", testCompetitorsOne)
	t.Run("Meetings", testMeetingsOne)
	t.Run("Races", testRacesOne)
}

func TestAll(t *testing.T) {
	t.Run("Competitors", testCompetitorsAll)
	t.Run("Meetings", testMeetingsAll)
	t.Run("Races", testRacesAll)
}

func TestCount(t *testing.T) {
	t.Run("Competitors", testCompetitorsCount)
	t.Run("Meetings", testMeetingsCount)
	t.Run("Races", testRacesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Competitors", testCompetitorsHooks)
	t.Run("Meetings", testMeetingsHooks)
	t.Run("Races", testRacesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Competitors", testCompetitorsInsert)
	t.Run("Competitors", testCompetitorsInsertWhitelist)
	t.Run("Meetings", testMeetingsInsert)
	t.Run("Meetings", testMeetingsInsertWhitelist)
	t.Run("Races", testRacesInsert)
	t.Run("Races", testRacesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CompetitorToRaceUsingRace", testCompetitorToOneRaceUsingRace)
	t.Run("RaceToMeetingUsingMeeting", testRaceToOneMeetingUsingMeeting)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("MeetingToRaces", testMeetingToManyRaces)
	t.Run("RaceToCompetitors", testRaceToManyCompetitors)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CompetitorToRaceUsingRace", testCompetitorToOneSetOpRaceUsingRace)
	t.Run("RaceToMeetingUsingMeeting", testRaceToOneSetOpMeetingUsingMeeting)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("CompetitorToRaceUsingRace", testCompetitorToOneRemoveOpRaceUsingRace)
	t.Run("RaceToMeetingUsingMeeting", testRaceToOneRemoveOpMeetingUsingMeeting)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("MeetingToRaces", testMeetingToManyAddOpRaces)
	t.Run("RaceToCompetitors", testRaceToManyAddOpCompetitors)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("MeetingToRaces", testMeetingToManySetOpRaces)
	t.Run("RaceToCompetitors", testRaceToManySetOpCompetitors)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("MeetingToRaces", testMeetingToManyRemoveOpRaces)
	t.Run("RaceToCompetitors", testRaceToManyRemoveOpCompetitors)
}

func TestReload(t *testing.T) {
	t.Run("Competitors", testCompetitorsReload)
	t.Run("Meetings", testMeetingsReload)
	t.Run("Races", testRacesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Competitors", testCompetitorsReloadAll)
	t.Run("Meetings", testMeetingsReloadAll)
	t.Run("Races", testRacesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Competitors", testCompetitorsSelect)
	t.Run("Meetings", testMeetingsSelect)
	t.Run("Races", testRacesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Competitors", testCompetitorsUpdate)
	t.Run("Meetings", testMeetingsUpdate)
	t.Run("Races", testRacesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Competitors", testCompetitorsSliceUpdateAll)
	t.Run("Meetings", testMeetingsSliceUpdateAll)
	t.Run("Races", testRacesSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Competitors", testCompetitorsUpsert)
	t.Run("Meetings", testMeetingsUpsert)
	t.Run("Races", testRacesUpsert)
}
