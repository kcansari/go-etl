package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// SELECT film.film_id AS fid,     film.title,     film.description,     category.name AS category,     film.rental_rate AS price,     film.length,     film.rating,     group_concat((((upper("substring"((actor.first_name)::text, 1, 1)) || lower("substring"((actor.first_name)::text, 2))) || upper("substring"((actor.last_name)::text, 1, 1))) || lower("substring"((actor.last_name)::text, 2)))) AS actors    FROM ((((category      LEFT JOIN film_category ON ((category.category_id = film_category.category_id)))      LEFT JOIN film ON ((film_category.film_id = film.film_id)))      JOIN film_actor ON ((film.film_id = film_actor.film_id)))      JOIN actor ON ((film_actor.actor_id = actor.actor_id)))   GROUP BY film.film_id, film.title, film.description, category.name, film.rental_rate, film.length, film.rating;
type NicerButSlowerFilmList struct {
	Fid         sql.NullInt64   `json:"fid"`         // fid
	Title       sql.NullString  `json:"title"`       // title
	Description sql.NullString  `json:"description"` // description
	Category    sql.NullString  `json:"category"`    // category
	Price       sql.NullFloat64 `json:"price"`       // price
	Length      sql.NullInt64   `json:"length"`      // length
	Rating      NullMpaaRating  `json:"rating"`      // rating
	Actors      sql.NullString  `json:"actors"`      // actors
}