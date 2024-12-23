package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// FilmCategory represents a row from 'public.film_category'.
type FilmCategory struct {
	FilmID     int16     `json:"film_id"`     // film_id
	CategoryID int16     `json:"category_id"` // category_id
	LastUpdate time.Time `json:"last_update"` // last_update
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [FilmCategory] exists in the database.
func (fc *FilmCategory) Exists() bool {
	return fc._exists
}

// Deleted returns true when the [FilmCategory] has been marked for deletion
// from the database.
func (fc *FilmCategory) Deleted() bool {
	return fc._deleted
}

// Insert inserts the [FilmCategory] to the database.
func (fc *FilmCategory) Insert(ctx context.Context, db DB) error {
	switch {
	case fc._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case fc._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.film_category (` +
		`film_id, category_id, last_update` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)`
	// run
	logf(sqlstr, fc.FilmID, fc.CategoryID, fc.LastUpdate)
	if _, err := db.ExecContext(ctx, sqlstr, fc.FilmID, fc.CategoryID, fc.LastUpdate); err != nil {
		return logerror(err)
	}
	// set exists
	fc._exists = true
	return nil
}

// Update updates a [FilmCategory] in the database.
func (fc *FilmCategory) Update(ctx context.Context, db DB) error {
	switch {
	case !fc._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case fc._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.film_category SET ` +
		`last_update = $1 ` +
		`WHERE film_id = $2 AND category_id = $3`
	// run
	logf(sqlstr, fc.LastUpdate, fc.FilmID, fc.CategoryID)
	if _, err := db.ExecContext(ctx, sqlstr, fc.LastUpdate, fc.FilmID, fc.CategoryID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [FilmCategory] to the database.
func (fc *FilmCategory) Save(ctx context.Context, db DB) error {
	if fc.Exists() {
		return fc.Update(ctx, db)
	}
	return fc.Insert(ctx, db)
}

// Upsert performs an upsert for [FilmCategory].
func (fc *FilmCategory) Upsert(ctx context.Context, db DB) error {
	switch {
	case fc._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.film_category (` +
		`film_id, category_id, last_update` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (film_id, category_id) DO ` +
		`UPDATE SET ` +
		`last_update = EXCLUDED.last_update `
	// run
	logf(sqlstr, fc.FilmID, fc.CategoryID, fc.LastUpdate)
	if _, err := db.ExecContext(ctx, sqlstr, fc.FilmID, fc.CategoryID, fc.LastUpdate); err != nil {
		return logerror(err)
	}
	// set exists
	fc._exists = true
	return nil
}

// Delete deletes the [FilmCategory] from the database.
func (fc *FilmCategory) Delete(ctx context.Context, db DB) error {
	switch {
	case !fc._exists: // doesn't exist
		return nil
	case fc._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM public.film_category ` +
		`WHERE film_id = $1 AND category_id = $2`
	// run
	logf(sqlstr, fc.FilmID, fc.CategoryID)
	if _, err := db.ExecContext(ctx, sqlstr, fc.FilmID, fc.CategoryID); err != nil {
		return logerror(err)
	}
	// set deleted
	fc._deleted = true
	return nil
}

// FilmCategoryByFilmIDCategoryID retrieves a row from 'public.film_category' as a [FilmCategory].
//
// Generated from index 'film_category_pkey'.
func FilmCategoryByFilmIDCategoryID(ctx context.Context, db DB, filmID, categoryID int16) (*FilmCategory, error) {
	// query
	const sqlstr = `SELECT ` +
		`film_id, category_id, last_update ` +
		`FROM public.film_category ` +
		`WHERE film_id = $1 AND category_id = $2`
	// run
	logf(sqlstr, filmID, categoryID)
	fc := FilmCategory{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, filmID, categoryID).Scan(&fc.FilmID, &fc.CategoryID, &fc.LastUpdate); err != nil {
		return nil, logerror(err)
	}
	return &fc, nil
}

// Category returns the Category associated with the [FilmCategory]'s (CategoryID).
//
// Generated from foreign key 'film_category_category_id_fkey'.
func (fc *FilmCategory) Category(ctx context.Context, db DB) (*Category, error) {
	return CategoryByCategoryID(ctx, db, int(fc.CategoryID))
}

// Film returns the Film associated with the [FilmCategory]'s (FilmID).
//
// Generated from foreign key 'film_category_film_id_fkey'.
func (fc *FilmCategory) Film(ctx context.Context, db DB) (*Film, error) {
	return FilmByFilmID(ctx, db, int(fc.FilmID))
}
