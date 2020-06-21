// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// RankRole is an object representing the database table.
type RankRole struct {
	ID      int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID int64 `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	Rank    int   `boil:"rank" json:"rank" toml:"rank" yaml:"rank"`

	R *rankRoleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L rankRoleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RankRoleColumns = struct {
	ID      string
	GuildID string
	Rank    string
}{
	ID:      "id",
	GuildID: "guild_id",
	Rank:    "rank",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

var RankRoleWhere = struct {
	ID      whereHelperint64
	GuildID whereHelperint64
	Rank    whereHelperint
}{
	ID:      whereHelperint64{field: "\"rank_roles\".\"id\""},
	GuildID: whereHelperint64{field: "\"rank_roles\".\"guild_id\""},
	Rank:    whereHelperint{field: "\"rank_roles\".\"rank\""},
}

// RankRoleRels is where relationship names are stored.
var RankRoleRels = struct {
	Guild string
}{
	Guild: "Guild",
}

// rankRoleR is where relationships are stored.
type rankRoleR struct {
	Guild *Guild
}

// NewStruct creates a new relationship struct
func (*rankRoleR) NewStruct() *rankRoleR {
	return &rankRoleR{}
}

// rankRoleL is where Load methods for each relationship are stored.
type rankRoleL struct{}

var (
	rankRoleAllColumns            = []string{"id", "guild_id", "rank"}
	rankRoleColumnsWithoutDefault = []string{"id", "guild_id", "rank"}
	rankRoleColumnsWithDefault    = []string{}
	rankRolePrimaryKeyColumns     = []string{"id"}
)

type (
	// RankRoleSlice is an alias for a slice of pointers to RankRole.
	// This should generally be used opposed to []RankRole.
	RankRoleSlice []*RankRole
	// RankRoleHook is the signature for custom RankRole hook methods
	RankRoleHook func(context.Context, boil.ContextExecutor, *RankRole) error

	rankRoleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	rankRoleType                 = reflect.TypeOf(&RankRole{})
	rankRoleMapping              = queries.MakeStructMapping(rankRoleType)
	rankRolePrimaryKeyMapping, _ = queries.BindMapping(rankRoleType, rankRoleMapping, rankRolePrimaryKeyColumns)
	rankRoleInsertCacheMut       sync.RWMutex
	rankRoleInsertCache          = make(map[string]insertCache)
	rankRoleUpdateCacheMut       sync.RWMutex
	rankRoleUpdateCache          = make(map[string]updateCache)
	rankRoleUpsertCacheMut       sync.RWMutex
	rankRoleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var rankRoleBeforeInsertHooks []RankRoleHook
var rankRoleBeforeUpdateHooks []RankRoleHook
var rankRoleBeforeDeleteHooks []RankRoleHook
var rankRoleBeforeUpsertHooks []RankRoleHook

var rankRoleAfterInsertHooks []RankRoleHook
var rankRoleAfterSelectHooks []RankRoleHook
var rankRoleAfterUpdateHooks []RankRoleHook
var rankRoleAfterDeleteHooks []RankRoleHook
var rankRoleAfterUpsertHooks []RankRoleHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RankRole) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankRoleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RankRole) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankRoleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RankRole) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankRoleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RankRole) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankRoleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RankRole) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankRoleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RankRole) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankRoleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RankRole) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankRoleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RankRole) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankRoleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RankRole) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range rankRoleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRankRoleHook registers your hook function for all future operations.
func AddRankRoleHook(hookPoint boil.HookPoint, rankRoleHook RankRoleHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		rankRoleBeforeInsertHooks = append(rankRoleBeforeInsertHooks, rankRoleHook)
	case boil.BeforeUpdateHook:
		rankRoleBeforeUpdateHooks = append(rankRoleBeforeUpdateHooks, rankRoleHook)
	case boil.BeforeDeleteHook:
		rankRoleBeforeDeleteHooks = append(rankRoleBeforeDeleteHooks, rankRoleHook)
	case boil.BeforeUpsertHook:
		rankRoleBeforeUpsertHooks = append(rankRoleBeforeUpsertHooks, rankRoleHook)
	case boil.AfterInsertHook:
		rankRoleAfterInsertHooks = append(rankRoleAfterInsertHooks, rankRoleHook)
	case boil.AfterSelectHook:
		rankRoleAfterSelectHooks = append(rankRoleAfterSelectHooks, rankRoleHook)
	case boil.AfterUpdateHook:
		rankRoleAfterUpdateHooks = append(rankRoleAfterUpdateHooks, rankRoleHook)
	case boil.AfterDeleteHook:
		rankRoleAfterDeleteHooks = append(rankRoleAfterDeleteHooks, rankRoleHook)
	case boil.AfterUpsertHook:
		rankRoleAfterUpsertHooks = append(rankRoleAfterUpsertHooks, rankRoleHook)
	}
}

// One returns a single rankRole record from the query.
func (q rankRoleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RankRole, error) {
	o := &RankRole{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for rank_roles")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RankRole records from the query.
func (q rankRoleQuery) All(ctx context.Context, exec boil.ContextExecutor) (RankRoleSlice, error) {
	var o []*RankRole

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RankRole slice")
	}

	if len(rankRoleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RankRole records in the query.
func (q rankRoleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count rank_roles rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q rankRoleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if rank_roles exists")
	}

	return count > 0, nil
}

// Guild pointed to by the foreign key.
func (o *RankRole) Guild(mods ...qm.QueryMod) guildQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.GuildID),
	}

	queryMods = append(queryMods, mods...)

	query := Guilds(queryMods...)
	queries.SetFrom(query.Query, "\"guilds\"")

	return query
}

// LoadGuild allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (rankRoleL) LoadGuild(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRankRole interface{}, mods queries.Applicator) error {
	var slice []*RankRole
	var object *RankRole

	if singular {
		object = maybeRankRole.(*RankRole)
	} else {
		slice = *maybeRankRole.(*[]*RankRole)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &rankRoleR{}
		}
		args = append(args, object.GuildID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &rankRoleR{}
			}

			for _, a := range args {
				if a == obj.GuildID {
					continue Outer
				}
			}

			args = append(args, obj.GuildID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`guilds`), qm.WhereIn(`guilds.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Guild")
	}

	var resultSlice []*Guild
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Guild")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for guilds")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for guilds")
	}

	if len(rankRoleAfterSelectHooks) != 0 {
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
		object.R.Guild = foreign
		if foreign.R == nil {
			foreign.R = &guildR{}
		}
		foreign.R.RankRoles = append(foreign.R.RankRoles, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GuildID == foreign.ID {
				local.R.Guild = foreign
				if foreign.R == nil {
					foreign.R = &guildR{}
				}
				foreign.R.RankRoles = append(foreign.R.RankRoles, local)
				break
			}
		}
	}

	return nil
}

// SetGuild of the rankRole to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.RankRoles.
func (o *RankRole) SetGuild(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Guild) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"rank_roles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
		strmangle.WhereClause("\"", "\"", 2, rankRolePrimaryKeyColumns),
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

	o.GuildID = related.ID
	if o.R == nil {
		o.R = &rankRoleR{
			Guild: related,
		}
	} else {
		o.R.Guild = related
	}

	if related.R == nil {
		related.R = &guildR{
			RankRoles: RankRoleSlice{o},
		}
	} else {
		related.R.RankRoles = append(related.R.RankRoles, o)
	}

	return nil
}

// RankRoles retrieves all the records using an executor.
func RankRoles(mods ...qm.QueryMod) rankRoleQuery {
	mods = append(mods, qm.From("\"rank_roles\""))
	return rankRoleQuery{NewQuery(mods...)}
}

// FindRankRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRankRole(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*RankRole, error) {
	rankRoleObj := &RankRole{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"rank_roles\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, rankRoleObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from rank_roles")
	}

	return rankRoleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RankRole) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rank_roles provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rankRoleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	rankRoleInsertCacheMut.RLock()
	cache, cached := rankRoleInsertCache[key]
	rankRoleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			rankRoleAllColumns,
			rankRoleColumnsWithDefault,
			rankRoleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(rankRoleType, rankRoleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(rankRoleType, rankRoleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"rank_roles\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"rank_roles\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into rank_roles")
	}

	if !cached {
		rankRoleInsertCacheMut.Lock()
		rankRoleInsertCache[key] = cache
		rankRoleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RankRole.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RankRole) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	rankRoleUpdateCacheMut.RLock()
	cache, cached := rankRoleUpdateCache[key]
	rankRoleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			rankRoleAllColumns,
			rankRolePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update rank_roles, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"rank_roles\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, rankRolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(rankRoleType, rankRoleMapping, append(wl, rankRolePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update rank_roles row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for rank_roles")
	}

	if !cached {
		rankRoleUpdateCacheMut.Lock()
		rankRoleUpdateCache[key] = cache
		rankRoleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q rankRoleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for rank_roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for rank_roles")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RankRoleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rankRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"rank_roles\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, rankRolePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in rankRole slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all rankRole")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RankRole) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no rank_roles provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(rankRoleColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	rankRoleUpsertCacheMut.RLock()
	cache, cached := rankRoleUpsertCache[key]
	rankRoleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			rankRoleAllColumns,
			rankRoleColumnsWithDefault,
			rankRoleColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			rankRoleAllColumns,
			rankRolePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert rank_roles, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(rankRolePrimaryKeyColumns))
			copy(conflict, rankRolePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"rank_roles\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(rankRoleType, rankRoleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(rankRoleType, rankRoleMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert rank_roles")
	}

	if !cached {
		rankRoleUpsertCacheMut.Lock()
		rankRoleUpsertCache[key] = cache
		rankRoleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RankRole record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RankRole) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RankRole provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rankRolePrimaryKeyMapping)
	sql := "DELETE FROM \"rank_roles\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from rank_roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for rank_roles")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q rankRoleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no rankRoleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rank_roles")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for rank_roles")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RankRoleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(rankRoleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rankRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"rank_roles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rankRolePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from rankRole slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for rank_roles")
	}

	if len(rankRoleAfterDeleteHooks) != 0 {
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
func (o *RankRole) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRankRole(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RankRoleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RankRoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rankRolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"rank_roles\".* FROM \"rank_roles\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, rankRolePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RankRoleSlice")
	}

	*o = slice

	return nil
}

// RankRoleExists checks if the RankRole row exists.
func RankRoleExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"rank_roles\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if rank_roles exists")
	}

	return exists, nil
}