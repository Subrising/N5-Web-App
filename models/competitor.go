// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"gopkg.in/volatiletech/null.v6"
)

// Competitor is an object representing the database table.
type Competitor struct {
	CompetitorID string      `boil:"competitor_id" json:"competitor_id" toml:"competitor_id" yaml:"competitor_id"`
	RaceID       null.String `boil:"race_id" json:"race_id,omitempty" toml:"race_id" yaml:"race_id,omitempty"`
	Position     int         `boil:"position" json:"position" toml:"position" yaml:"position"`
	Type         string      `boil:"type" json:"type" toml:"type" yaml:"type"`
	Name         null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`

	R *competitorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L competitorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CompetitorColumns = struct {
	CompetitorID string
	RaceID       string
	Position     string
	Type         string
	Name         string
}{
	CompetitorID: "competitor_id",
	RaceID:       "race_id",
	Position:     "position",
	Type:         "type",
	Name:         "name",
}

// competitorR is where relationships are stored.
type competitorR struct {
	Race *Race
}

// competitorL is where Load methods for each relationship are stored.
type competitorL struct{}

var (
	competitorColumns               = []string{"competitor_id", "race_id", "position", "type", "name"}
	competitorColumnsWithoutDefault = []string{"race_id", "position", "type", "name"}
	competitorColumnsWithDefault    = []string{"competitor_id"}
	competitorPrimaryKeyColumns     = []string{"competitor_id"}
)

type (
	// CompetitorSlice is an alias for a slice of pointers to Competitor.
	// This should generally be used opposed to []Competitor.
	CompetitorSlice []*Competitor
	// CompetitorHook is the signature for custom Competitor hook methods
	CompetitorHook func(boil.Executor, *Competitor) error

	competitorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	competitorType                 = reflect.TypeOf(&Competitor{})
	competitorMapping              = queries.MakeStructMapping(competitorType)
	competitorPrimaryKeyMapping, _ = queries.BindMapping(competitorType, competitorMapping, competitorPrimaryKeyColumns)
	competitorInsertCacheMut       sync.RWMutex
	competitorInsertCache          = make(map[string]insertCache)
	competitorUpdateCacheMut       sync.RWMutex
	competitorUpdateCache          = make(map[string]updateCache)
	competitorUpsertCacheMut       sync.RWMutex
	competitorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var competitorBeforeInsertHooks []CompetitorHook
var competitorBeforeUpdateHooks []CompetitorHook
var competitorBeforeDeleteHooks []CompetitorHook
var competitorBeforeUpsertHooks []CompetitorHook

var competitorAfterInsertHooks []CompetitorHook
var competitorAfterSelectHooks []CompetitorHook
var competitorAfterUpdateHooks []CompetitorHook
var competitorAfterDeleteHooks []CompetitorHook
var competitorAfterUpsertHooks []CompetitorHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Competitor) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range competitorBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Competitor) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range competitorBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Competitor) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range competitorBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Competitor) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range competitorBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Competitor) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range competitorAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Competitor) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range competitorAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Competitor) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range competitorAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Competitor) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range competitorAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Competitor) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range competitorAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCompetitorHook registers your hook function for all future operations.
func AddCompetitorHook(hookPoint boil.HookPoint, competitorHook CompetitorHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		competitorBeforeInsertHooks = append(competitorBeforeInsertHooks, competitorHook)
	case boil.BeforeUpdateHook:
		competitorBeforeUpdateHooks = append(competitorBeforeUpdateHooks, competitorHook)
	case boil.BeforeDeleteHook:
		competitorBeforeDeleteHooks = append(competitorBeforeDeleteHooks, competitorHook)
	case boil.BeforeUpsertHook:
		competitorBeforeUpsertHooks = append(competitorBeforeUpsertHooks, competitorHook)
	case boil.AfterInsertHook:
		competitorAfterInsertHooks = append(competitorAfterInsertHooks, competitorHook)
	case boil.AfterSelectHook:
		competitorAfterSelectHooks = append(competitorAfterSelectHooks, competitorHook)
	case boil.AfterUpdateHook:
		competitorAfterUpdateHooks = append(competitorAfterUpdateHooks, competitorHook)
	case boil.AfterDeleteHook:
		competitorAfterDeleteHooks = append(competitorAfterDeleteHooks, competitorHook)
	case boil.AfterUpsertHook:
		competitorAfterUpsertHooks = append(competitorAfterUpsertHooks, competitorHook)
	}
}

// OneP returns a single competitor record from the query, and panics on error.
func (q competitorQuery) OneP() *Competitor {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single competitor record from the query.
func (q competitorQuery) One() (*Competitor, error) {
	o := &Competitor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for competitor")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Competitor records from the query, and panics on error.
func (q competitorQuery) AllP() CompetitorSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Competitor records from the query.
func (q competitorQuery) All() (CompetitorSlice, error) {
	var o []*Competitor

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Competitor slice")
	}

	if len(competitorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Competitor records in the query, and panics on error.
func (q competitorQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Competitor records in the query.
func (q competitorQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count competitor rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q competitorQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q competitorQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if competitor exists")
	}

	return count > 0, nil
}

// RaceG pointed to by the foreign key.
func (o *Competitor) RaceG(mods ...qm.QueryMod) raceQuery {
	return o.Race(boil.GetDB(), mods...)
}

// Race pointed to by the foreign key.
func (o *Competitor) Race(exec boil.Executor, mods ...qm.QueryMod) raceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("race_id=?", o.RaceID),
	}

	queryMods = append(queryMods, mods...)

	query := Races(exec, queryMods...)
	queries.SetFrom(query.Query, "\"race\"")

	return query
} // LoadRace allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (competitorL) LoadRace(e boil.Executor, singular bool, maybeCompetitor interface{}) error {
	var slice []*Competitor
	var object *Competitor

	count := 1
	if singular {
		object = maybeCompetitor.(*Competitor)
	} else {
		slice = *maybeCompetitor.(*[]*Competitor)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &competitorR{}
		}
		args[0] = object.RaceID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &competitorR{}
			}
			args[i] = obj.RaceID
		}
	}

	query := fmt.Sprintf(
		"select * from \"race\" where \"race_id\" in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Race")
	}
	defer results.Close()

	var resultSlice []*Race
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Race")
	}

	if len(competitorAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Race = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RaceID.String == foreign.RaceID {
				local.R.Race = foreign
				break
			}
		}
	}

	return nil
}

// SetRaceG of the competitor to the related item.
// Sets o.R.Race to related.
// Adds o to related.R.Competitors.
// Uses the global database handle.
func (o *Competitor) SetRaceG(insert bool, related *Race) error {
	return o.SetRace(boil.GetDB(), insert, related)
}

// SetRaceP of the competitor to the related item.
// Sets o.R.Race to related.
// Adds o to related.R.Competitors.
// Panics on error.
func (o *Competitor) SetRaceP(exec boil.Executor, insert bool, related *Race) {
	if err := o.SetRace(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetRaceGP of the competitor to the related item.
// Sets o.R.Race to related.
// Adds o to related.R.Competitors.
// Uses the global database handle and panics on error.
func (o *Competitor) SetRaceGP(insert bool, related *Race) {
	if err := o.SetRace(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetRace of the competitor to the related item.
// Sets o.R.Race to related.
// Adds o to related.R.Competitors.
func (o *Competitor) SetRace(exec boil.Executor, insert bool, related *Race) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"competitor\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"race_id"}),
		strmangle.WhereClause("\"", "\"", 2, competitorPrimaryKeyColumns),
	)
	values := []interface{}{related.RaceID, o.CompetitorID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RaceID.String = related.RaceID
	o.RaceID.Valid = true

	if o.R == nil {
		o.R = &competitorR{
			Race: related,
		}
	} else {
		o.R.Race = related
	}

	if related.R == nil {
		related.R = &raceR{
			Competitors: CompetitorSlice{o},
		}
	} else {
		related.R.Competitors = append(related.R.Competitors, o)
	}

	return nil
}

// RemoveRaceG relationship.
// Sets o.R.Race to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *Competitor) RemoveRaceG(related *Race) error {
	return o.RemoveRace(boil.GetDB(), related)
}

// RemoveRaceP relationship.
// Sets o.R.Race to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Panics on error.
func (o *Competitor) RemoveRaceP(exec boil.Executor, related *Race) {
	if err := o.RemoveRace(exec, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemoveRaceGP relationship.
// Sets o.R.Race to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle and panics on error.
func (o *Competitor) RemoveRaceGP(related *Race) {
	if err := o.RemoveRace(boil.GetDB(), related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemoveRace relationship.
// Sets o.R.Race to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Competitor) RemoveRace(exec boil.Executor, related *Race) error {
	var err error

	o.RaceID.Valid = false
	if err = o.Update(exec, "race_id"); err != nil {
		o.RaceID.Valid = true
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Race = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Competitors {
		if o.RaceID.String != ri.RaceID.String {
			continue
		}

		ln := len(related.R.Competitors)
		if ln > 1 && i < ln-1 {
			related.R.Competitors[i] = related.R.Competitors[ln-1]
		}
		related.R.Competitors = related.R.Competitors[:ln-1]
		break
	}
	return nil
}

// CompetitorsG retrieves all records.
func CompetitorsG(mods ...qm.QueryMod) competitorQuery {
	return Competitors(boil.GetDB(), mods...)
}

// Competitors retrieves all the records using an executor.
func Competitors(exec boil.Executor, mods ...qm.QueryMod) competitorQuery {
	mods = append(mods, qm.From("\"competitor\""))
	return competitorQuery{NewQuery(exec, mods...)}
}

// FindCompetitorG retrieves a single record by ID.
func FindCompetitorG(competitorID string, selectCols ...string) (*Competitor, error) {
	return FindCompetitor(boil.GetDB(), competitorID, selectCols...)
}

// FindCompetitorGP retrieves a single record by ID, and panics on error.
func FindCompetitorGP(competitorID string, selectCols ...string) *Competitor {
	retobj, err := FindCompetitor(boil.GetDB(), competitorID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCompetitor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCompetitor(exec boil.Executor, competitorID string, selectCols ...string) (*Competitor, error) {
	competitorObj := &Competitor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"competitor\" where \"competitor_id\"=$1", sel,
	)

	q := queries.Raw(exec, query, competitorID)

	err := q.Bind(competitorObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from competitor")
	}

	return competitorObj, nil
}

// FindCompetitorP retrieves a single record by ID with an executor, and panics on error.
func FindCompetitorP(exec boil.Executor, competitorID string, selectCols ...string) *Competitor {
	retobj, err := FindCompetitor(exec, competitorID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Competitor) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Competitor) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Competitor) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Competitor) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no competitor provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(competitorColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	competitorInsertCacheMut.RLock()
	cache, cached := competitorInsertCache[key]
	competitorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			competitorColumns,
			competitorColumnsWithDefault,
			competitorColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(competitorType, competitorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(competitorType, competitorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"competitor\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"competitor\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into competitor")
	}

	if !cached {
		competitorInsertCacheMut.Lock()
		competitorInsertCache[key] = cache
		competitorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Competitor record. See Update for
// whitelist behavior description.
func (o *Competitor) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Competitor record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Competitor) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Competitor, and panics on error.
// See Update for whitelist behavior description.
func (o *Competitor) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Competitor.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Competitor) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	competitorUpdateCacheMut.RLock()
	cache, cached := competitorUpdateCache[key]
	competitorUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			competitorColumns,
			competitorPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update competitor, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"competitor\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, competitorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(competitorType, competitorMapping, append(wl, competitorPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update competitor row")
	}

	if !cached {
		competitorUpdateCacheMut.Lock()
		competitorUpdateCache[key] = cache
		competitorUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q competitorQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q competitorQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for competitor")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CompetitorSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CompetitorSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CompetitorSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CompetitorSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), competitorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"competitor\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, competitorPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in competitor slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Competitor) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Competitor) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Competitor) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Competitor) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no competitor provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(competitorColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()

	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	competitorUpsertCacheMut.RLock()
	cache, cached := competitorUpsertCache[key]
	competitorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			competitorColumns,
			competitorColumnsWithDefault,
			competitorColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			competitorColumns,
			competitorPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert competitor, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(competitorPrimaryKeyColumns))
			copy(conflict, competitorPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"competitor\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(competitorType, competitorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(competitorType, competitorMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert competitor")
	}

	if !cached {
		competitorUpsertCacheMut.Lock()
		competitorUpsertCache[key] = cache
		competitorUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Competitor record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Competitor) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Competitor record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Competitor) DeleteG() error {
	if o == nil {
		return errors.New("models: no Competitor provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Competitor record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Competitor) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Competitor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Competitor) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Competitor provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), competitorPrimaryKeyMapping)
	sql := "DELETE FROM \"competitor\" WHERE \"competitor_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from competitor")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q competitorQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q competitorQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no competitorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from competitor")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CompetitorSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CompetitorSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Competitor slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CompetitorSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CompetitorSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Competitor slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(competitorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), competitorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"competitor\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, competitorPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from competitor slice")
	}

	if len(competitorAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Competitor) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Competitor) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Competitor) ReloadG() error {
	if o == nil {
		return errors.New("models: no Competitor provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Competitor) Reload(exec boil.Executor) error {
	ret, err := FindCompetitor(exec, o.CompetitorID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CompetitorSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CompetitorSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompetitorSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty CompetitorSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CompetitorSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	competitors := CompetitorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), competitorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"competitor\".* FROM \"competitor\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, competitorPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&competitors)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CompetitorSlice")
	}

	*o = competitors

	return nil
}

// CompetitorExists checks if the Competitor row exists.
func CompetitorExists(exec boil.Executor, competitorID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"competitor\" where \"competitor_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, competitorID)
	}

	row := exec.QueryRow(sql, competitorID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if competitor exists")
	}

	return exists, nil
}

// CompetitorExistsG checks if the Competitor row exists.
func CompetitorExistsG(competitorID string) (bool, error) {
	return CompetitorExists(boil.GetDB(), competitorID)
}

// CompetitorExistsGP checks if the Competitor row exists. Panics on error.
func CompetitorExistsGP(competitorID string) bool {
	e, err := CompetitorExists(boil.GetDB(), competitorID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CompetitorExistsP checks if the Competitor row exists. Panics on error.
func CompetitorExistsP(exec boil.Executor, competitorID string) bool {
	e, err := CompetitorExists(exec, competitorID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
