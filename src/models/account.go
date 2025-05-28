package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name    string
	Token   string
	ZoneID  int
	Zone    Zone
	Records []Record
}
