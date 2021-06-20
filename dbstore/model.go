package dbstore

import "database/sql"

type Profile struct {
	ID int
	Name sql.NullString
	Filter sql.NullString
	Sorter sql.NullString
}