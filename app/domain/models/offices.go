// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Office is an object representing the database table.
type Office struct {
	ID   []byte `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *officeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L officeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OfficeColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

var OfficeTableColumns = struct {
	ID   string
	Name string
}{
	ID:   "offices.id",
	Name: "offices.name",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var OfficeWhere = struct {
	ID   whereHelper__byte
	Name whereHelperstring
}{
	ID:   whereHelper__byte{field: "`offices`.`id`"},
	Name: whereHelperstring{field: "`offices`.`name`"},
}

// OfficeRels is where relationship names are stored.
var OfficeRels = struct {
	OfficeUsers string
	Samples     string
}{
	OfficeUsers: "OfficeUsers",
	Samples:     "Samples",
}

// officeR is where relationships are stored.
type officeR struct {
	OfficeUsers OfficeUserSlice `boil:"OfficeUsers" json:"OfficeUsers" toml:"OfficeUsers" yaml:"OfficeUsers"`
	Samples     SampleSlice     `boil:"Samples" json:"Samples" toml:"Samples" yaml:"Samples"`
}

// NewStruct creates a new relationship struct
func (*officeR) NewStruct() *officeR {
	return &officeR{}
}

// officeL is where Load methods for each relationship are stored.
type officeL struct{}

var (
	officeAllColumns            = []string{"id", "name"}
	officeColumnsWithoutDefault = []string{"id", "name"}
	officeColumnsWithDefault    = []string{}
	officePrimaryKeyColumns     = []string{"id"}
)

type (
	// OfficeSlice is an alias for a slice of pointers to Office.
	// This should almost always be used instead of []Office.
	OfficeSlice []*Office
	// OfficeHook is the signature for custom Office hook methods
	OfficeHook func(context.Context, boil.ContextExecutor, *Office) error

	officeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	officeType                 = reflect.TypeOf(&Office{})
	officeMapping              = queries.MakeStructMapping(officeType)
	officePrimaryKeyMapping, _ = queries.BindMapping(officeType, officeMapping, officePrimaryKeyColumns)
	officeInsertCacheMut       sync.RWMutex
	officeInsertCache          = make(map[string]insertCache)
	officeUpdateCacheMut       sync.RWMutex
	officeUpdateCache          = make(map[string]updateCache)
	officeUpsertCacheMut       sync.RWMutex
	officeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var officeBeforeInsertHooks []OfficeHook
var officeBeforeUpdateHooks []OfficeHook
var officeBeforeDeleteHooks []OfficeHook
var officeBeforeUpsertHooks []OfficeHook

var officeAfterInsertHooks []OfficeHook
var officeAfterSelectHooks []OfficeHook
var officeAfterUpdateHooks []OfficeHook
var officeAfterDeleteHooks []OfficeHook
var officeAfterUpsertHooks []OfficeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Office) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Office) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Office) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Office) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Office) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Office) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Office) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Office) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Office) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOfficeHook registers your hook function for all future operations.
func AddOfficeHook(hookPoint boil.HookPoint, officeHook OfficeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		officeBeforeInsertHooks = append(officeBeforeInsertHooks, officeHook)
	case boil.BeforeUpdateHook:
		officeBeforeUpdateHooks = append(officeBeforeUpdateHooks, officeHook)
	case boil.BeforeDeleteHook:
		officeBeforeDeleteHooks = append(officeBeforeDeleteHooks, officeHook)
	case boil.BeforeUpsertHook:
		officeBeforeUpsertHooks = append(officeBeforeUpsertHooks, officeHook)
	case boil.AfterInsertHook:
		officeAfterInsertHooks = append(officeAfterInsertHooks, officeHook)
	case boil.AfterSelectHook:
		officeAfterSelectHooks = append(officeAfterSelectHooks, officeHook)
	case boil.AfterUpdateHook:
		officeAfterUpdateHooks = append(officeAfterUpdateHooks, officeHook)
	case boil.AfterDeleteHook:
		officeAfterDeleteHooks = append(officeAfterDeleteHooks, officeHook)
	case boil.AfterUpsertHook:
		officeAfterUpsertHooks = append(officeAfterUpsertHooks, officeHook)
	}
}

// One returns a single office record from the query.
func (q officeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Office, error) {
	o := &Office{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for offices")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Office records from the query.
func (q officeQuery) All(ctx context.Context, exec boil.ContextExecutor) (OfficeSlice, error) {
	var o []*Office

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Office slice")
	}

	if len(officeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Office records in the query.
func (q officeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count offices rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q officeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if offices exists")
	}

	return count > 0, nil
}

// OfficeUsers retrieves all the office_user's OfficeUsers with an executor.
func (o *Office) OfficeUsers(mods ...qm.QueryMod) officeUserQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`office_users`.`office_id`=?", o.ID),
	)

	query := OfficeUsers(queryMods...)
	queries.SetFrom(query.Query, "`office_users`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`office_users`.*"})
	}

	return query
}

// Samples retrieves all the sample's Samples with an executor.
func (o *Office) Samples(mods ...qm.QueryMod) sampleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`samples`.`office_id`=?", o.ID),
	)

	query := Samples(queryMods...)
	queries.SetFrom(query.Query, "`samples`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`samples`.*"})
	}

	return query
}

// LoadOfficeUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (officeL) LoadOfficeUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOffice interface{}, mods queries.Applicator) error {
	var slice []*Office
	var object *Office

	if singular {
		object = maybeOffice.(*Office)
	} else {
		slice = *maybeOffice.(*[]*Office)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &officeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &officeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`office_users`),
		qm.WhereIn(`office_users.office_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load office_users")
	}

	var resultSlice []*OfficeUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice office_users")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on office_users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for office_users")
	}

	if len(officeUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.OfficeUsers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &officeUserR{}
			}
			foreign.R.Office = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.OfficeID) {
				local.R.OfficeUsers = append(local.R.OfficeUsers, foreign)
				if foreign.R == nil {
					foreign.R = &officeUserR{}
				}
				foreign.R.Office = local
				break
			}
		}
	}

	return nil
}

// LoadSamples allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (officeL) LoadSamples(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOffice interface{}, mods queries.Applicator) error {
	var slice []*Office
	var object *Office

	if singular {
		object = maybeOffice.(*Office)
	} else {
		slice = *maybeOffice.(*[]*Office)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &officeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &officeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`samples`),
		qm.WhereIn(`samples.office_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load samples")
	}

	var resultSlice []*Sample
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice samples")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on samples")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for samples")
	}

	if len(sampleAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Samples = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &sampleR{}
			}
			foreign.R.Office = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.OfficeID) {
				local.R.Samples = append(local.R.Samples, foreign)
				if foreign.R == nil {
					foreign.R = &sampleR{}
				}
				foreign.R.Office = local
				break
			}
		}
	}

	return nil
}

// AddOfficeUsers adds the given related objects to the existing relationships
// of the office, optionally inserting them as new records.
// Appends related to o.R.OfficeUsers.
// Sets related.R.Office appropriately.
func (o *Office) AddOfficeUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*OfficeUser) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.OfficeID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `office_users` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"office_id"}),
				strmangle.WhereClause("`", "`", 0, officeUserPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.OfficeID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &officeR{
			OfficeUsers: related,
		}
	} else {
		o.R.OfficeUsers = append(o.R.OfficeUsers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &officeUserR{
				Office: o,
			}
		} else {
			rel.R.Office = o
		}
	}
	return nil
}

// AddSamples adds the given related objects to the existing relationships
// of the office, optionally inserting them as new records.
// Appends related to o.R.Samples.
// Sets related.R.Office appropriately.
func (o *Office) AddSamples(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Sample) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.OfficeID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `samples` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"office_id"}),
				strmangle.WhereClause("`", "`", 0, samplePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.OfficeID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &officeR{
			Samples: related,
		}
	} else {
		o.R.Samples = append(o.R.Samples, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &sampleR{
				Office: o,
			}
		} else {
			rel.R.Office = o
		}
	}
	return nil
}

// Offices retrieves all the records using an executor.
func Offices(mods ...qm.QueryMod) officeQuery {
	mods = append(mods, qm.From("`offices`"))
	return officeQuery{NewQuery(mods...)}
}

// FindOffice retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOffice(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*Office, error) {
	officeObj := &Office{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `offices` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, officeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from offices")
	}

	if err = officeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return officeObj, err
	}

	return officeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Office) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no offices provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(officeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	officeInsertCacheMut.RLock()
	cache, cached := officeInsertCache[key]
	officeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			officeAllColumns,
			officeColumnsWithDefault,
			officeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(officeType, officeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(officeType, officeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `offices` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `offices` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `offices` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, officePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into offices")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for offices")
	}

CacheNoHooks:
	if !cached {
		officeInsertCacheMut.Lock()
		officeInsertCache[key] = cache
		officeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Office.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Office) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	officeUpdateCacheMut.RLock()
	cache, cached := officeUpdateCache[key]
	officeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			officeAllColumns,
			officePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update offices, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `offices` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, officePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(officeType, officeMapping, append(wl, officePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update offices row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for offices")
	}

	if !cached {
		officeUpdateCacheMut.Lock()
		officeUpdateCache[key] = cache
		officeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q officeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for offices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for offices")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OfficeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), officePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `offices` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, officePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in office slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all office")
	}
	return rowsAff, nil
}

var mySQLOfficeUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Office) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no offices provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(officeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOfficeUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	officeUpsertCacheMut.RLock()
	cache, cached := officeUpsertCache[key]
	officeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			officeAllColumns,
			officeColumnsWithDefault,
			officeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			officeAllColumns,
			officePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert offices, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`offices`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `offices` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(officeType, officeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(officeType, officeMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for offices")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(officeType, officeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for offices")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for offices")
	}

CacheNoHooks:
	if !cached {
		officeUpsertCacheMut.Lock()
		officeUpsertCache[key] = cache
		officeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Office record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Office) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Office provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), officePrimaryKeyMapping)
	sql := "DELETE FROM `offices` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from offices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for offices")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q officeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no officeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from offices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for offices")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OfficeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(officeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), officePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `offices` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, officePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from office slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for offices")
	}

	if len(officeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Office) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOffice(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OfficeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OfficeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), officePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `offices`.* FROM `offices` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, officePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OfficeSlice")
	}

	*o = slice

	return nil
}

// OfficeExists checks if the Office row exists.
func OfficeExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `offices` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if offices exists")
	}

	return exists, nil
}
