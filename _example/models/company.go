package models

import "database/sql"

// Company ...
type Company struct {
	ID   uint           `gorm:"primary_key;AUTO_INCREMENT" json:"id" form:"id"`
	Name string         `json:"name" form:"name"`
	URL  sql.NullString `json:"url" form:"url"`
	Jobs []*Job         `json:"jobs" form:"jobs"`
}
