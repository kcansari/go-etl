// Package models contains generated code for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// Actor represents a row from 'public.actor'.
type Actor struct {
	ActorID    int       `json:"actor_id"`    // actor_id
	FirstName  string    `json:"first_name"`  // first_name
	LastName   string    `json:"last_name"`   // last_name
	LastUpdate time.Time `json:"last_update"` // last_update
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Actor] exists in the database.
func (a *Actor) Exists() bool {
	return a._exists
}

// Deleted returns true when the [Actor] has been marked for deletion
// from the database.
func (a *Actor) Deleted() bool {
	return a._deleted
}

// Insert inserts the [Actor] to the database.
func (a *Actor) Insert(ctx context.Context, db DB) error {
	switch {
	case a._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case a._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.actor (` +
		`actor_id, first_name, last_name, last_update` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`
	// run
	logf(sqlstr, a.ActorID, a.FirstName, a.LastName, a.LastUpdate)
	if _, err := db.ExecContext(ctx, sqlstr, a.ActorID, a.FirstName, a.LastName, a.LastUpdate); err != nil {
		return logerror(err)
	}
	// set exists
	a._exists = true
	return nil
}

// Update updates a [Actor] in the database.
func (a *Actor) Update(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case a._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.actor SET ` +
		`first_name = $1, last_name = $2, last_update = $3 ` +
		`WHERE actor_id = $4`
	// run
	logf(sqlstr, a.FirstName, a.LastName, a.LastUpdate, a.ActorID)
	if _, err := db.ExecContext(ctx, sqlstr, a.FirstName, a.LastName, a.LastUpdate, a.ActorID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Actor] to the database.
func (a *Actor) Save(ctx context.Context, db DB) error {
	if a.Exists() {
		return a.Update(ctx, db)
	}
	return a.Insert(ctx, db)
}

// Upsert performs an upsert for [Actor].
func (a *Actor) Upsert(ctx context.Context, db DB) error {
	switch {
	case a._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.actor (` +
		`actor_id, first_name, last_name, last_update` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (actor_id) DO ` +
		`UPDATE SET ` +
		`first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name, last_update = EXCLUDED.last_update `
	// run
	logf(sqlstr, a.ActorID, a.FirstName, a.LastName, a.LastUpdate)
	if _, err := db.ExecContext(ctx, sqlstr, a.ActorID, a.FirstName, a.LastName, a.LastUpdate); err != nil {
		return logerror(err)
	}
	// set exists
	a._exists = true
	return nil
}

// Delete deletes the [Actor] from the database.
func (a *Actor) Delete(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return nil
	case a._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.actor ` +
		`WHERE actor_id = $1`
	// run
	logf(sqlstr, a.ActorID)
	if _, err := db.ExecContext(ctx, sqlstr, a.ActorID); err != nil {
		return logerror(err)
	}
	// set deleted
	a._deleted = true
	return nil
}

// ActorByActorID retrieves a row from 'public.actor' as a [Actor].
//
// Generated from index 'actor_pkey'.
func ActorByActorID(ctx context.Context, db DB, actorID int) (*Actor, error) {
	// query
	const sqlstr = `SELECT ` +
		`actor_id, first_name, last_name, last_update ` +
		`FROM public.actor ` +
		`WHERE actor_id = $1`
	// run
	logf(sqlstr, actorID)
	a := Actor{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, actorID).Scan(&a.ActorID, &a.FirstName, &a.LastName, &a.LastUpdate); err != nil {
		return nil, logerror(err)
	}
	return &a, nil
}

// ActorByLastName retrieves a row from 'public.actor' as a [Actor].
//
// Generated from index 'idx_actor_last_name'.
func ActorByLastName(ctx context.Context, db DB, lastName string) ([]*Actor, error) {
	// query
	const sqlstr = `SELECT ` +
		`actor_id, first_name, last_name, last_update ` +
		`FROM public.actor ` +
		`WHERE last_name = $1`
	// run
	logf(sqlstr, lastName)
	rows, err := db.QueryContext(ctx, sqlstr, lastName)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Actor
	for rows.Next() {
		a := Actor{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&a.ActorID, &a.FirstName, &a.LastName, &a.LastUpdate); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &a)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
