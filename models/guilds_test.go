// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testGuilds(t *testing.T) {
	t.Parallel()

	query := Guilds()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGuildsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Guilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuildsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Guilds().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Guilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuildsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GuildSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Guilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGuildsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GuildExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Guild exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GuildExists to return true, but got false.")
	}
}

func testGuildsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	guildFound, err := FindGuild(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if guildFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGuildsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Guilds().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGuildsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Guilds().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGuildsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	guildOne := &Guild{}
	guildTwo := &Guild{}
	if err = randomize.Struct(seed, guildOne, guildDBTypes, false, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}
	if err = randomize.Struct(seed, guildTwo, guildDBTypes, false, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = guildOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = guildTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Guilds().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGuildsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	guildOne := &Guild{}
	guildTwo := &Guild{}
	if err = randomize.Struct(seed, guildOne, guildDBTypes, false, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}
	if err = randomize.Struct(seed, guildTwo, guildDBTypes, false, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = guildOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = guildTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Guilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func guildBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Guild) error {
	*o = Guild{}
	return nil
}

func guildAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Guild) error {
	*o = Guild{}
	return nil
}

func guildAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Guild) error {
	*o = Guild{}
	return nil
}

func guildBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Guild) error {
	*o = Guild{}
	return nil
}

func guildAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Guild) error {
	*o = Guild{}
	return nil
}

func guildBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Guild) error {
	*o = Guild{}
	return nil
}

func guildAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Guild) error {
	*o = Guild{}
	return nil
}

func guildBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Guild) error {
	*o = Guild{}
	return nil
}

func guildAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Guild) error {
	*o = Guild{}
	return nil
}

func testGuildsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Guild{}
	o := &Guild{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, guildDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Guild object: %s", err)
	}

	AddGuildHook(boil.BeforeInsertHook, guildBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	guildBeforeInsertHooks = []GuildHook{}

	AddGuildHook(boil.AfterInsertHook, guildAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	guildAfterInsertHooks = []GuildHook{}

	AddGuildHook(boil.AfterSelectHook, guildAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	guildAfterSelectHooks = []GuildHook{}

	AddGuildHook(boil.BeforeUpdateHook, guildBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	guildBeforeUpdateHooks = []GuildHook{}

	AddGuildHook(boil.AfterUpdateHook, guildAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	guildAfterUpdateHooks = []GuildHook{}

	AddGuildHook(boil.BeforeDeleteHook, guildBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	guildBeforeDeleteHooks = []GuildHook{}

	AddGuildHook(boil.AfterDeleteHook, guildAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	guildAfterDeleteHooks = []GuildHook{}

	AddGuildHook(boil.BeforeUpsertHook, guildBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	guildBeforeUpsertHooks = []GuildHook{}

	AddGuildHook(boil.AfterUpsertHook, guildAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	guildAfterUpsertHooks = []GuildHook{}
}

func testGuildsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Guilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGuildsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(guildColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Guilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGuildToManyRoles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Guild
	var b, c Role

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.GuildID = a.ID
	c.GuildID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Roles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.GuildID == b.GuildID {
			bFound = true
		}
		if v.GuildID == c.GuildID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := GuildSlice{&a}
	if err = a.L.LoadRoles(ctx, tx, false, (*[]*Guild)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Roles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Roles = nil
	if err = a.L.LoadRoles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Roles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testGuildToManyAddOpRoles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Guild
	var b, c, d, e Role

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, guildDBTypes, false, strmangle.SetComplement(guildPrimaryKeyColumns, guildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Role{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Role{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRoles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.GuildID {
			t.Error("foreign key was wrong value", a.ID, first.GuildID)
		}
		if a.ID != second.GuildID {
			t.Error("foreign key was wrong value", a.ID, second.GuildID)
		}

		if first.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Guild != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Roles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Roles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Roles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testGuildsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGuildsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GuildSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGuildsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Guilds().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	guildDBTypes = map[string]string{`ID`: `bigint`}
	_            = bytes.MinRead
)

func testGuildsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(guildPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(guildAllColumns) == len(guildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Guilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, guildDBTypes, true, guildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGuildsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(guildAllColumns) == len(guildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Guild{}
	if err = randomize.Struct(seed, o, guildDBTypes, true, guildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Guilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, guildDBTypes, true, guildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(guildAllColumns, guildPrimaryKeyColumns) {
		fields = guildAllColumns
	} else {
		fields = strmangle.SetComplement(
			guildAllColumns,
			guildPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := GuildSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGuildsUpsert(t *testing.T) {
	t.Parallel()

	if len(guildAllColumns) == len(guildPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Guild{}
	if err = randomize.Struct(seed, &o, guildDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Guild: %s", err)
	}

	count, err := Guilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, guildDBTypes, false, guildPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Guild struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Guild: %s", err)
	}

	count, err = Guilds().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
