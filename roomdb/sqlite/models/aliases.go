// SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021
//
// SPDX-License-Identifier: MIT

// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
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

// Alias is an object representing the database table.
type Alias struct {
	ID        int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string `boil:"name" json:"name" toml:"name" yaml:"name"`
	MemberID  int64  `boil:"member_id" json:"member_id" toml:"member_id" yaml:"member_id"`
	Signature []byte `boil:"signature" json:"signature" toml:"signature" yaml:"signature"`

	R *aliasR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L aliasL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AliasColumns = struct {
	ID        string
	Name      string
	MemberID  string
	Signature string
}{
	ID:        "id",
	Name:      "name",
	MemberID:  "member_id",
	Signature: "signature",
}

// Generated where

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var AliasWhere = struct {
	ID        whereHelperint64
	Name      whereHelperstring
	MemberID  whereHelperint64
	Signature whereHelper__byte
}{
	ID:        whereHelperint64{field: "\"aliases\".\"id\""},
	Name:      whereHelperstring{field: "\"aliases\".\"name\""},
	MemberID:  whereHelperint64{field: "\"aliases\".\"member_id\""},
	Signature: whereHelper__byte{field: "\"aliases\".\"signature\""},
}

// AliasRels is where relationship names are stored.
var AliasRels = struct {
	Member string
}{
	Member: "Member",
}

// aliasR is where relationships are stored.
type aliasR struct {
	Member *Member `boil:"Member" json:"Member" toml:"Member" yaml:"Member"`
}

// NewStruct creates a new relationship struct
func (*aliasR) NewStruct() *aliasR {
	return &aliasR{}
}

// aliasL is where Load methods for each relationship are stored.
type aliasL struct{}

var (
	aliasAllColumns            = []string{"id", "name", "member_id", "signature"}
	aliasColumnsWithoutDefault = []string{}
	aliasColumnsWithDefault    = []string{"id", "name", "member_id", "signature"}
	aliasPrimaryKeyColumns     = []string{"id"}
)

type (
	// AliasSlice is an alias for a slice of pointers to Alias.
	// This should generally be used opposed to []Alias.
	AliasSlice []*Alias
	// AliasHook is the signature for custom Alias hook methods
	AliasHook func(context.Context, boil.ContextExecutor, *Alias) error

	aliasQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	aliasType                 = reflect.TypeOf(&Alias{})
	aliasMapping              = queries.MakeStructMapping(aliasType)
	aliasPrimaryKeyMapping, _ = queries.BindMapping(aliasType, aliasMapping, aliasPrimaryKeyColumns)
	aliasInsertCacheMut       sync.RWMutex
	aliasInsertCache          = make(map[string]insertCache)
	aliasUpdateCacheMut       sync.RWMutex
	aliasUpdateCache          = make(map[string]updateCache)
	aliasUpsertCacheMut       sync.RWMutex
	aliasUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var aliasBeforeInsertHooks []AliasHook
var aliasBeforeUpdateHooks []AliasHook
var aliasBeforeDeleteHooks []AliasHook
var aliasBeforeUpsertHooks []AliasHook

var aliasAfterInsertHooks []AliasHook
var aliasAfterSelectHooks []AliasHook
var aliasAfterUpdateHooks []AliasHook
var aliasAfterDeleteHooks []AliasHook
var aliasAfterUpsertHooks []AliasHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Alias) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aliasBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Alias) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aliasBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Alias) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aliasBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Alias) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aliasBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Alias) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aliasAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Alias) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aliasAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Alias) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aliasAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Alias) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aliasAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Alias) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range aliasAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAliasHook registers your hook function for all future operations.
func AddAliasHook(hookPoint boil.HookPoint, aliasHook AliasHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		aliasBeforeInsertHooks = append(aliasBeforeInsertHooks, aliasHook)
	case boil.BeforeUpdateHook:
		aliasBeforeUpdateHooks = append(aliasBeforeUpdateHooks, aliasHook)
	case boil.BeforeDeleteHook:
		aliasBeforeDeleteHooks = append(aliasBeforeDeleteHooks, aliasHook)
	case boil.BeforeUpsertHook:
		aliasBeforeUpsertHooks = append(aliasBeforeUpsertHooks, aliasHook)
	case boil.AfterInsertHook:
		aliasAfterInsertHooks = append(aliasAfterInsertHooks, aliasHook)
	case boil.AfterSelectHook:
		aliasAfterSelectHooks = append(aliasAfterSelectHooks, aliasHook)
	case boil.AfterUpdateHook:
		aliasAfterUpdateHooks = append(aliasAfterUpdateHooks, aliasHook)
	case boil.AfterDeleteHook:
		aliasAfterDeleteHooks = append(aliasAfterDeleteHooks, aliasHook)
	case boil.AfterUpsertHook:
		aliasAfterUpsertHooks = append(aliasAfterUpsertHooks, aliasHook)
	}
}

// One returns a single alias record from the query.
func (q aliasQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Alias, error) {
	o := &Alias{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for aliases")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Alias records from the query.
func (q aliasQuery) All(ctx context.Context, exec boil.ContextExecutor) (AliasSlice, error) {
	var o []*Alias

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Alias slice")
	}

	if len(aliasAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Alias records in the query.
func (q aliasQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count aliases rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q aliasQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if aliases exists")
	}

	return count > 0, nil
}

// Member pointed to by the foreign key.
func (o *Alias) Member(mods ...qm.QueryMod) memberQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.MemberID),
	}

	queryMods = append(queryMods, mods...)

	query := Members(queryMods...)
	queries.SetFrom(query.Query, "\"members\"")

	return query
}

// LoadMember allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (aliasL) LoadMember(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAlias interface{}, mods queries.Applicator) error {
	var slice []*Alias
	var object *Alias

	if singular {
		object = maybeAlias.(*Alias)
	} else {
		slice = *maybeAlias.(*[]*Alias)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &aliasR{}
		}
		args = append(args, object.MemberID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &aliasR{}
			}

			for _, a := range args {
				if a == obj.MemberID {
					continue Outer
				}
			}

			args = append(args, obj.MemberID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`members`),
		qm.WhereIn(`members.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Member")
	}

	var resultSlice []*Member
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Member")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for members")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for members")
	}

	if len(aliasAfterSelectHooks) != 0 {
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
		object.R.Member = foreign
		if foreign.R == nil {
			foreign.R = &memberR{}
		}
		foreign.R.Aliases = append(foreign.R.Aliases, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.MemberID == foreign.ID {
				local.R.Member = foreign
				if foreign.R == nil {
					foreign.R = &memberR{}
				}
				foreign.R.Aliases = append(foreign.R.Aliases, local)
				break
			}
		}
	}

	return nil
}

// SetMember of the alias to the related item.
// Sets o.R.Member to related.
// Adds o to related.R.Aliases.
func (o *Alias) SetMember(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Member) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"aliases\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"member_id"}),
		strmangle.WhereClause("\"", "\"", 0, aliasPrimaryKeyColumns),
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

	o.MemberID = related.ID
	if o.R == nil {
		o.R = &aliasR{
			Member: related,
		}
	} else {
		o.R.Member = related
	}

	if related.R == nil {
		related.R = &memberR{
			Aliases: AliasSlice{o},
		}
	} else {
		related.R.Aliases = append(related.R.Aliases, o)
	}

	return nil
}

// Aliases retrieves all the records using an executor.
func Aliases(mods ...qm.QueryMod) aliasQuery {
	mods = append(mods, qm.From("\"aliases\""))
	return aliasQuery{NewQuery(mods...)}
}

// FindAlias retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAlias(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Alias, error) {
	aliasObj := &Alias{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"aliases\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, aliasObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from aliases")
	}

	return aliasObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Alias) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no aliases provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(aliasColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	aliasInsertCacheMut.RLock()
	cache, cached := aliasInsertCache[key]
	aliasInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			aliasAllColumns,
			aliasColumnsWithDefault,
			aliasColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(aliasType, aliasMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(aliasType, aliasMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"aliases\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"aliases\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"aliases\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, aliasPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into aliases")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == aliasMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for aliases")
	}

CacheNoHooks:
	if !cached {
		aliasInsertCacheMut.Lock()
		aliasInsertCache[key] = cache
		aliasInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Alias.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Alias) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	aliasUpdateCacheMut.RLock()
	cache, cached := aliasUpdateCache[key]
	aliasUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			aliasAllColumns,
			aliasPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update aliases, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"aliases\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, aliasPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(aliasType, aliasMapping, append(wl, aliasPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update aliases row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for aliases")
	}

	if !cached {
		aliasUpdateCacheMut.Lock()
		aliasUpdateCache[key] = cache
		aliasUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q aliasQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for aliases")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for aliases")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AliasSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aliasPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"aliases\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, aliasPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in alias slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all alias")
	}
	return rowsAff, nil
}

// Delete deletes a single Alias record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Alias) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Alias provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), aliasPrimaryKeyMapping)
	sql := "DELETE FROM \"aliases\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from aliases")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for aliases")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q aliasQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no aliasQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from aliases")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for aliases")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AliasSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(aliasBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aliasPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"aliases\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, aliasPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from alias slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for aliases")
	}

	if len(aliasAfterDeleteHooks) != 0 {
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
func (o *Alias) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAlias(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AliasSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AliasSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), aliasPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"aliases\".* FROM \"aliases\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, aliasPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AliasSlice")
	}

	*o = slice

	return nil
}

// AliasExists checks if the Alias row exists.
func AliasExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"aliases\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if aliases exists")
	}

	return exists, nil
}
