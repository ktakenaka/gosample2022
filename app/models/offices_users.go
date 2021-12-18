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

// OfficesUser is an object representing the database table.
type OfficesUser struct {
	ID       []byte `boil:"id" json:"id" toml:"id" yaml:"id"`
	OfficeID []byte `boil:"office_id" json:"office_id" toml:"office_id" yaml:"office_id"`
	UserID   []byte `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Name     string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *officesUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L officesUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OfficesUserColumns = struct {
	ID       string
	OfficeID string
	UserID   string
	Name     string
}{
	ID:       "id",
	OfficeID: "office_id",
	UserID:   "user_id",
	Name:     "name",
}

var OfficesUserTableColumns = struct {
	ID       string
	OfficeID string
	UserID   string
	Name     string
}{
	ID:       "offices_users.id",
	OfficeID: "offices_users.office_id",
	UserID:   "offices_users.user_id",
	Name:     "offices_users.name",
}

// Generated where

var OfficesUserWhere = struct {
	ID       whereHelper__byte
	OfficeID whereHelper__byte
	UserID   whereHelper__byte
	Name     whereHelperstring
}{
	ID:       whereHelper__byte{field: "`offices_users`.`id`"},
	OfficeID: whereHelper__byte{field: "`offices_users`.`office_id`"},
	UserID:   whereHelper__byte{field: "`offices_users`.`user_id`"},
	Name:     whereHelperstring{field: "`offices_users`.`name`"},
}

// OfficesUserRels is where relationship names are stored.
var OfficesUserRels = struct {
	Office string
	User   string
}{
	Office: "Office",
	User:   "User",
}

// officesUserR is where relationships are stored.
type officesUserR struct {
	Office *Office `boil:"Office" json:"Office" toml:"Office" yaml:"Office"`
	User   *User   `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*officesUserR) NewStruct() *officesUserR {
	return &officesUserR{}
}

// officesUserL is where Load methods for each relationship are stored.
type officesUserL struct{}

var (
	officesUserAllColumns            = []string{"id", "office_id", "user_id", "name"}
	officesUserColumnsWithoutDefault = []string{"id", "office_id", "user_id", "name"}
	officesUserColumnsWithDefault    = []string{}
	officesUserPrimaryKeyColumns     = []string{"id"}
)

type (
	// OfficesUserSlice is an alias for a slice of pointers to OfficesUser.
	// This should almost always be used instead of []OfficesUser.
	OfficesUserSlice []*OfficesUser
	// OfficesUserHook is the signature for custom OfficesUser hook methods
	OfficesUserHook func(context.Context, boil.ContextExecutor, *OfficesUser) error

	officesUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	officesUserType                 = reflect.TypeOf(&OfficesUser{})
	officesUserMapping              = queries.MakeStructMapping(officesUserType)
	officesUserPrimaryKeyMapping, _ = queries.BindMapping(officesUserType, officesUserMapping, officesUserPrimaryKeyColumns)
	officesUserInsertCacheMut       sync.RWMutex
	officesUserInsertCache          = make(map[string]insertCache)
	officesUserUpdateCacheMut       sync.RWMutex
	officesUserUpdateCache          = make(map[string]updateCache)
	officesUserUpsertCacheMut       sync.RWMutex
	officesUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var officesUserBeforeInsertHooks []OfficesUserHook
var officesUserBeforeUpdateHooks []OfficesUserHook
var officesUserBeforeDeleteHooks []OfficesUserHook
var officesUserBeforeUpsertHooks []OfficesUserHook

var officesUserAfterInsertHooks []OfficesUserHook
var officesUserAfterSelectHooks []OfficesUserHook
var officesUserAfterUpdateHooks []OfficesUserHook
var officesUserAfterDeleteHooks []OfficesUserHook
var officesUserAfterUpsertHooks []OfficesUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OfficesUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officesUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OfficesUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officesUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OfficesUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officesUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OfficesUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officesUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OfficesUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officesUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OfficesUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officesUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OfficesUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officesUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OfficesUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officesUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OfficesUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range officesUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOfficesUserHook registers your hook function for all future operations.
func AddOfficesUserHook(hookPoint boil.HookPoint, officesUserHook OfficesUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		officesUserBeforeInsertHooks = append(officesUserBeforeInsertHooks, officesUserHook)
	case boil.BeforeUpdateHook:
		officesUserBeforeUpdateHooks = append(officesUserBeforeUpdateHooks, officesUserHook)
	case boil.BeforeDeleteHook:
		officesUserBeforeDeleteHooks = append(officesUserBeforeDeleteHooks, officesUserHook)
	case boil.BeforeUpsertHook:
		officesUserBeforeUpsertHooks = append(officesUserBeforeUpsertHooks, officesUserHook)
	case boil.AfterInsertHook:
		officesUserAfterInsertHooks = append(officesUserAfterInsertHooks, officesUserHook)
	case boil.AfterSelectHook:
		officesUserAfterSelectHooks = append(officesUserAfterSelectHooks, officesUserHook)
	case boil.AfterUpdateHook:
		officesUserAfterUpdateHooks = append(officesUserAfterUpdateHooks, officesUserHook)
	case boil.AfterDeleteHook:
		officesUserAfterDeleteHooks = append(officesUserAfterDeleteHooks, officesUserHook)
	case boil.AfterUpsertHook:
		officesUserAfterUpsertHooks = append(officesUserAfterUpsertHooks, officesUserHook)
	}
}

// One returns a single officesUser record from the query.
func (q officesUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OfficesUser, error) {
	o := &OfficesUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for offices_users")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OfficesUser records from the query.
func (q officesUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (OfficesUserSlice, error) {
	var o []*OfficesUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OfficesUser slice")
	}

	if len(officesUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OfficesUser records in the query.
func (q officesUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count offices_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q officesUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if offices_users exists")
	}

	return count > 0, nil
}

// Office pointed to by the foreign key.
func (o *OfficesUser) Office(mods ...qm.QueryMod) officeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.OfficeID),
	}

	queryMods = append(queryMods, mods...)

	query := Offices(queryMods...)
	queries.SetFrom(query.Query, "`offices`")

	return query
}

// User pointed to by the foreign key.
func (o *OfficesUser) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// LoadOffice allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (officesUserL) LoadOffice(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOfficesUser interface{}, mods queries.Applicator) error {
	var slice []*OfficesUser
	var object *OfficesUser

	if singular {
		object = maybeOfficesUser.(*OfficesUser)
	} else {
		slice = *maybeOfficesUser.(*[]*OfficesUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &officesUserR{}
		}
		if !queries.IsNil(object.OfficeID) {
			args = append(args, object.OfficeID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &officesUserR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.OfficeID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.OfficeID) {
				args = append(args, obj.OfficeID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`offices`),
		qm.WhereIn(`offices.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Office")
	}

	var resultSlice []*Office
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Office")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for offices")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for offices")
	}

	if len(officesUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Office = foreign
		if foreign.R == nil {
			foreign.R = &officeR{}
		}
		foreign.R.OfficesUsers = append(foreign.R.OfficesUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.OfficeID, foreign.ID) {
				local.R.Office = foreign
				if foreign.R == nil {
					foreign.R = &officeR{}
				}
				foreign.R.OfficesUsers = append(foreign.R.OfficesUsers, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (officesUserL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOfficesUser interface{}, mods queries.Applicator) error {
	var slice []*OfficesUser
	var object *OfficesUser

	if singular {
		object = maybeOfficesUser.(*OfficesUser)
	} else {
		slice = *maybeOfficesUser.(*[]*OfficesUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &officesUserR{}
		}
		if !queries.IsNil(object.UserID) {
			args = append(args, object.UserID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &officesUserR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.UserID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.UserID) {
				args = append(args, obj.UserID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(officesUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.OfficesUsers = append(foreign.R.OfficesUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.UserID, foreign.ID) {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.OfficesUsers = append(foreign.R.OfficesUsers, local)
				break
			}
		}
	}

	return nil
}

// SetOffice of the officesUser to the related item.
// Sets o.R.Office to related.
// Adds o to related.R.OfficesUsers.
func (o *OfficesUser) SetOffice(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Office) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `offices_users` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"office_id"}),
		strmangle.WhereClause("`", "`", 0, officesUserPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.OfficeID, related.ID)
	if o.R == nil {
		o.R = &officesUserR{
			Office: related,
		}
	} else {
		o.R.Office = related
	}

	if related.R == nil {
		related.R = &officeR{
			OfficesUsers: OfficesUserSlice{o},
		}
	} else {
		related.R.OfficesUsers = append(related.R.OfficesUsers, o)
	}

	return nil
}

// SetUser of the officesUser to the related item.
// Sets o.R.User to related.
// Adds o to related.R.OfficesUsers.
func (o *OfficesUser) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `offices_users` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, officesUserPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.UserID, related.ID)
	if o.R == nil {
		o.R = &officesUserR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			OfficesUsers: OfficesUserSlice{o},
		}
	} else {
		related.R.OfficesUsers = append(related.R.OfficesUsers, o)
	}

	return nil
}

// OfficesUsers retrieves all the records using an executor.
func OfficesUsers(mods ...qm.QueryMod) officesUserQuery {
	mods = append(mods, qm.From("`offices_users`"))
	return officesUserQuery{NewQuery(mods...)}
}

// FindOfficesUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOfficesUser(ctx context.Context, exec boil.ContextExecutor, iD []byte, selectCols ...string) (*OfficesUser, error) {
	officesUserObj := &OfficesUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `offices_users` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, officesUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from offices_users")
	}

	if err = officesUserObj.doAfterSelectHooks(ctx, exec); err != nil {
		return officesUserObj, err
	}

	return officesUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OfficesUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no offices_users provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(officesUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	officesUserInsertCacheMut.RLock()
	cache, cached := officesUserInsertCache[key]
	officesUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			officesUserAllColumns,
			officesUserColumnsWithDefault,
			officesUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(officesUserType, officesUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(officesUserType, officesUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `offices_users` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `offices_users` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `offices_users` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, officesUserPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into offices_users")
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
		return errors.Wrap(err, "models: unable to populate default values for offices_users")
	}

CacheNoHooks:
	if !cached {
		officesUserInsertCacheMut.Lock()
		officesUserInsertCache[key] = cache
		officesUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OfficesUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OfficesUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	officesUserUpdateCacheMut.RLock()
	cache, cached := officesUserUpdateCache[key]
	officesUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			officesUserAllColumns,
			officesUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update offices_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `offices_users` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, officesUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(officesUserType, officesUserMapping, append(wl, officesUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update offices_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for offices_users")
	}

	if !cached {
		officesUserUpdateCacheMut.Lock()
		officesUserUpdateCache[key] = cache
		officesUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q officesUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for offices_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for offices_users")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OfficesUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), officesUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `offices_users` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, officesUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in officesUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all officesUser")
	}
	return rowsAff, nil
}

var mySQLOfficesUserUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OfficesUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no offices_users provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(officesUserColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOfficesUserUniqueColumns, o)

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

	officesUserUpsertCacheMut.RLock()
	cache, cached := officesUserUpsertCache[key]
	officesUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			officesUserAllColumns,
			officesUserColumnsWithDefault,
			officesUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			officesUserAllColumns,
			officesUserPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert offices_users, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`offices_users`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `offices_users` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(officesUserType, officesUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(officesUserType, officesUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for offices_users")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(officesUserType, officesUserMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for offices_users")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for offices_users")
	}

CacheNoHooks:
	if !cached {
		officesUserUpsertCacheMut.Lock()
		officesUserUpsertCache[key] = cache
		officesUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OfficesUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OfficesUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OfficesUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), officesUserPrimaryKeyMapping)
	sql := "DELETE FROM `offices_users` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from offices_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for offices_users")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q officesUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no officesUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from offices_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for offices_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OfficesUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(officesUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), officesUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `offices_users` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, officesUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from officesUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for offices_users")
	}

	if len(officesUserAfterDeleteHooks) != 0 {
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
func (o *OfficesUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOfficesUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OfficesUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OfficesUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), officesUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `offices_users`.* FROM `offices_users` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, officesUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OfficesUserSlice")
	}

	*o = slice

	return nil
}

// OfficesUserExists checks if the OfficesUser row exists.
func OfficesUserExists(ctx context.Context, exec boil.ContextExecutor, iD []byte) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `offices_users` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if offices_users exists")
	}

	return exists, nil
}
