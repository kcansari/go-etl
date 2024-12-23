package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// SELECT cu.customer_id AS id,     (((cu.first_name)::text || ' '::text) || (cu.last_name)::text) AS name,     a.address,     a.postal_code AS "zip code",     a.phone,     city.city,     country.country,         CASE             WHEN cu.activebool THEN 'active'::text             ELSE ”::text         END AS notes,     cu.store_id AS sid    FROM (((customer cu      JOIN address a ON ((cu.address_id = a.address_id)))      JOIN city ON ((a.city_id = city.city_id)))      JOIN country ON ((city.country_id = country.country_id)));
type CustomerList struct {
	ID      sql.NullInt64  `json:"id"`       // id
	Name    sql.NullString `json:"name"`     // name
	Address sql.NullString `json:"address"`  // address
	ZipCode sql.NullString `json:"zip code"` // zip code
	Phone   sql.NullString `json:"phone"`    // phone
	City    sql.NullString `json:"city"`     // city
	Country sql.NullString `json:"country"`  // country
	Notes   sql.NullString `json:"notes"`    // notes
	Sid     sql.NullInt64  `json:"sid"`      // sid
}
