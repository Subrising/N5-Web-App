// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testRaces(t *testing.T) {
	t.Parallel()

	query := Races(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testRacesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = race.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Races(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRacesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Races(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Races(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRacesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := RaceSlice{race}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Races(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testRacesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := RaceExists(tx, race.RaceID)
	if err != nil {
		t.Errorf("Unable to check if Race exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RaceExistsG to return true, but got false.")
	}
}
func testRacesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	raceFound, err := FindRace(tx, race.RaceID)
	if err != nil {
		t.Error(err)
	}

	if raceFound == nil {
		t.Error("want a record, got nil")
	}
}
func testRacesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Races(tx).Bind(race); err != nil {
		t.Error(err)
	}
}

func testRacesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Races(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRacesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	raceOne := &Race{}
	raceTwo := &Race{}
	if err = randomize.Struct(seed, raceOne, raceDBTypes, false, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}
	if err = randomize.Struct(seed, raceTwo, raceDBTypes, false, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = raceOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = raceTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Races(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRacesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	raceOne := &Race{}
	raceTwo := &Race{}
	if err = randomize.Struct(seed, raceOne, raceDBTypes, false, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}
	if err = randomize.Struct(seed, raceTwo, raceDBTypes, false, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = raceOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = raceTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Races(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func raceBeforeInsertHook(e boil.Executor, o *Race) error {
	*o = Race{}
	return nil
}

func raceAfterInsertHook(e boil.Executor, o *Race) error {
	*o = Race{}
	return nil
}

func raceAfterSelectHook(e boil.Executor, o *Race) error {
	*o = Race{}
	return nil
}

func raceBeforeUpdateHook(e boil.Executor, o *Race) error {
	*o = Race{}
	return nil
}

func raceAfterUpdateHook(e boil.Executor, o *Race) error {
	*o = Race{}
	return nil
}

func raceBeforeDeleteHook(e boil.Executor, o *Race) error {
	*o = Race{}
	return nil
}

func raceAfterDeleteHook(e boil.Executor, o *Race) error {
	*o = Race{}
	return nil
}

func raceBeforeUpsertHook(e boil.Executor, o *Race) error {
	*o = Race{}
	return nil
}

func raceAfterUpsertHook(e boil.Executor, o *Race) error {
	*o = Race{}
	return nil
}

func testRacesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Race{}
	o := &Race{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, raceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Race object: %s", err)
	}

	AddRaceHook(boil.BeforeInsertHook, raceBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	raceBeforeInsertHooks = []RaceHook{}

	AddRaceHook(boil.AfterInsertHook, raceAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	raceAfterInsertHooks = []RaceHook{}

	AddRaceHook(boil.AfterSelectHook, raceAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	raceAfterSelectHooks = []RaceHook{}

	AddRaceHook(boil.BeforeUpdateHook, raceBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	raceBeforeUpdateHooks = []RaceHook{}

	AddRaceHook(boil.AfterUpdateHook, raceAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	raceAfterUpdateHooks = []RaceHook{}

	AddRaceHook(boil.BeforeDeleteHook, raceBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	raceBeforeDeleteHooks = []RaceHook{}

	AddRaceHook(boil.AfterDeleteHook, raceAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	raceAfterDeleteHooks = []RaceHook{}

	AddRaceHook(boil.BeforeUpsertHook, raceBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	raceBeforeUpsertHooks = []RaceHook{}

	AddRaceHook(boil.AfterUpsertHook, raceAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	raceAfterUpsertHooks = []RaceHook{}
}
func testRacesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Races(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRacesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx, raceColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Races(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRaceToManyCompetitors(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Race
	var b, c Competitor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, competitorDBTypes, false, competitorColumnsWithDefault...)
	randomize.Struct(seed, &c, competitorDBTypes, false, competitorColumnsWithDefault...)

	b.RaceID.Valid = true
	c.RaceID.Valid = true
	b.RaceID.String = a.RaceID
	c.RaceID.String = a.RaceID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	competitor, err := a.Competitors(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range competitor {
		if v.RaceID.String == b.RaceID.String {
			bFound = true
		}
		if v.RaceID.String == c.RaceID.String {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RaceSlice{&a}
	if err = a.L.LoadCompetitors(tx, false, (*[]*Race)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Competitors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Competitors = nil
	if err = a.L.LoadCompetitors(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Competitors); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", competitor)
	}
}

func testRaceToManyAddOpCompetitors(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Race
	var b, c, d, e Competitor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, raceDBTypes, false, strmangle.SetComplement(racePrimaryKeyColumns, raceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Competitor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, competitorDBTypes, false, strmangle.SetComplement(competitorPrimaryKeyColumns, competitorColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Competitor{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCompetitors(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.RaceID != first.RaceID.String {
			t.Error("foreign key was wrong value", a.RaceID, first.RaceID.String)
		}
		if a.RaceID != second.RaceID.String {
			t.Error("foreign key was wrong value", a.RaceID, second.RaceID.String)
		}

		if first.R.Race != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Race != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Competitors[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Competitors[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Competitors(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testRaceToManySetOpCompetitors(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Race
	var b, c, d, e Competitor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, raceDBTypes, false, strmangle.SetComplement(racePrimaryKeyColumns, raceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Competitor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, competitorDBTypes, false, strmangle.SetComplement(competitorPrimaryKeyColumns, competitorColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetCompetitors(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Competitors(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetCompetitors(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Competitors(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.RaceID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.RaceID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.RaceID != d.RaceID.String {
		t.Error("foreign key was wrong value", a.RaceID, d.RaceID.String)
	}
	if a.RaceID != e.RaceID.String {
		t.Error("foreign key was wrong value", a.RaceID, e.RaceID.String)
	}

	if b.R.Race != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Race != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Race != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Race != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Competitors[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Competitors[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testRaceToManyRemoveOpCompetitors(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Race
	var b, c, d, e Competitor

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, raceDBTypes, false, strmangle.SetComplement(racePrimaryKeyColumns, raceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Competitor{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, competitorDBTypes, false, strmangle.SetComplement(competitorPrimaryKeyColumns, competitorColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddCompetitors(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Competitors(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveCompetitors(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Competitors(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.RaceID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.RaceID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Race != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Race != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Race != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Race != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Competitors) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Competitors[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Competitors[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testRaceToOneMeetingUsingMeeting(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Race
	var foreign Meeting

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, meetingDBTypes, false, meetingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meeting struct: %s", err)
	}

	local.MeetingID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.MeetingID.String = foreign.MeetingID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Meeting(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.MeetingID != foreign.MeetingID {
		t.Errorf("want: %v, got %v", foreign.MeetingID, check.MeetingID)
	}

	slice := RaceSlice{&local}
	if err = local.L.LoadMeeting(tx, false, (*[]*Race)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Meeting == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Meeting = nil
	if err = local.L.LoadMeeting(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Meeting == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRaceToOneSetOpMeetingUsingMeeting(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Race
	var b, c Meeting

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, raceDBTypes, false, strmangle.SetComplement(racePrimaryKeyColumns, raceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, meetingDBTypes, false, strmangle.SetComplement(meetingPrimaryKeyColumns, meetingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, meetingDBTypes, false, strmangle.SetComplement(meetingPrimaryKeyColumns, meetingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Meeting{&b, &c} {
		err = a.SetMeeting(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Meeting != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Races[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.MeetingID.String != x.MeetingID {
			t.Error("foreign key was wrong value", a.MeetingID.String)
		}

		zero := reflect.Zero(reflect.TypeOf(a.MeetingID.String))
		reflect.Indirect(reflect.ValueOf(&a.MeetingID.String)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.MeetingID.String != x.MeetingID {
			t.Error("foreign key was wrong value", a.MeetingID.String, x.MeetingID)
		}
	}
}

func testRaceToOneRemoveOpMeetingUsingMeeting(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Race
	var b Meeting

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, raceDBTypes, false, strmangle.SetComplement(racePrimaryKeyColumns, raceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, meetingDBTypes, false, strmangle.SetComplement(meetingPrimaryKeyColumns, meetingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetMeeting(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveMeeting(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Meeting(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Meeting != nil {
		t.Error("R struct entry should be nil")
	}

	if a.MeetingID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Races) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testRacesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = race.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testRacesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := RaceSlice{race}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testRacesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Races(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	raceDBTypes = map[string]string{`ClosingTime`: `timestamp without time zone`, `MeetingID`: `uuid`, `RaceID`: `uuid`}
	_           = bytes.MinRead
)

func testRacesUpdate(t *testing.T) {
	t.Parallel()

	if len(raceColumns) == len(racePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Races(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	if err = race.Update(tx); err != nil {
		t.Error(err)
	}
}

func testRacesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(raceColumns) == len(racePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	race := &Race{}
	if err = randomize.Struct(seed, race, raceDBTypes, true, raceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Races(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, race, raceDBTypes, true, racePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(raceColumns, racePrimaryKeyColumns) {
		fields = raceColumns
	} else {
		fields = strmangle.SetComplement(
			raceColumns,
			racePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(race))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := RaceSlice{race}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testRacesUpsert(t *testing.T) {
	t.Parallel()

	if len(raceColumns) == len(racePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	race := Race{}
	if err = randomize.Struct(seed, &race, raceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = race.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Race: %s", err)
	}

	count, err := Races(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &race, raceDBTypes, false, racePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Race struct: %s", err)
	}

	if err = race.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Race: %s", err)
	}

	count, err = Races(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
