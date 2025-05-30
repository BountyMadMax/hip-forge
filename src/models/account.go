package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name    string
	Token   string
	ZoneID  int
	Zone    Zone
	Records []Record
}

func (account *Account) UnhiddenRecords(db *gorm.DB) []Record {
	records := []Record{}
	result := db.Where(&Record{AccountID: account.ID, Hidden: false}).Find(&records)

	if result.Error != nil {
		return nil
	}

	return records
}

func (account *Account) HiddenRecords(db *gorm.DB) []Record {
	records := []Record{}
	result := db.Where(&Record{AccountID: account.ID, Hidden: true}).Find(&records)

	if result.Error != nil {
		return nil
	}

	return records
}
