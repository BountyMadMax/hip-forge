package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name       string
	Token      string
	DNSRecords []DNSRecord
}
