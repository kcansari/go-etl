package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// RewardsReport calls the stored function 'public.rewards_report(integer, numeric) customer' on db.
func RewardsReport(ctx context.Context, db DB, minMonthlyPurchases int, minDollarAmountPurchased float64) (Customer, error) {
	// call public.rewards_report
	const sqlstr = `SELECT * FROM public.rewards_report($1, $2)`
	// run
	var r0 Customer
	logf(sqlstr, minMonthlyPurchases, minDollarAmountPurchased)
	if err := db.QueryRowContext(ctx, sqlstr, minMonthlyPurchases, minDollarAmountPurchased).Scan(&r0); err != nil {
		return Customer{}, logerror(err)
	}
	return r0, nil
}
