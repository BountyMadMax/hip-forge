package models

import (
	"time"
)

type RecordType string

const (
	A     RecordType = "A"
	AAAA  RecordType = "AAAA"
	NS    RecordType = "NS"
	MX    RecordType = "MX"
	CNAME RecordType = "CNAME"
	RP    RecordType = "RP"
	TXT   RecordType = "TXT"
	SOA   RecordType = "SOA"
	HINFO RecordType = "HINFO"
	SRV   RecordType = "SRV"
	DANE  RecordType = "DANE"
	TLSA  RecordType = "TLSA"
	DS    RecordType = "DS"
	CAA   RecordType = "CAA"
)

type Record struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	TTL       int
	Type      RecordType
	Value     string
	Created   time.Time
	Updated   time.Time
	ZoneID    uint
	Zone      Zone
	AccountID uint
	Account   Account
	Hidden    bool `gorm:"default:true"`
}
