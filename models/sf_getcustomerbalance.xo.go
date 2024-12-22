package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// GetCustomerBalance calls the stored function 'public.get_customer_balance(integer, timestamp without time zone) numeric' on db.
func GetCustomerBalance(ctx context.Context, db DB, pCustomerID int, pEffectiveDate time.Time) (float64, error) {
	// call public.get_customer_balance
	const sqlstr = `SELECT * FROM public.get_customer_balance($1, $2)`
	// run
	var r0 float64
	logf(sqlstr, pCustomerID, pEffectiveDate)
	if err := db.QueryRowContext(ctx, sqlstr, pCustomerID, pEffectiveDate).Scan(&r0); err != nil {
		return 0.0, logerror(err)
	}
	return r0, nil
}