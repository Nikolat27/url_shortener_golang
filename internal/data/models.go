package data

import "database/sql"

type Models struct {
	Urls UrlModel
}

func NewModels(DB *sql.DB) Models {
	return Models{
		Urls: UrlModel{DB},
	}
}
