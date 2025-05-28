package models

import "time"

type Zone struct {
	ID               string `gorm:"primaryKey"`
	Name             string
	Ttl              uint64
	Registrar        string
	LegacyDnsHost    string
	Created          time.Time
	Verified         time.Time
	Modified         time.Time
	Project          string
	Owner            string
	Permission       string
	Status           string
	Paused           bool
	IsSecondaryDns   bool
	Records          []Record
	AccountID        int
	ZoneTypeID       int
	ZoneType         ZoneType
	Nameserver       []*Nameserver `gorm:"many2many:zones_nameservers"`
	LegacyNameserver []*Nameserver `gorm:"many2many:zones_legacy_nameservers"`
}

type ZoneType struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	Prices      string
	Zones       []Zone
}

type Nameserver struct {
	Name        string  `gorm:"primaryKey"`
	Zones       []*Zone `gorm:"many2many:zones_nameservers"`
	LegacyZones []*Zone `gorm:"many2many:zones_legacy_nameservers"`
}
