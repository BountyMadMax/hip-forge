package models

import (
	"time"
)

type DNSRecordType string

const (
	A     DNSRecordType = "A"
	AAAA  DNSRecordType = "AAAA"
	NS    DNSRecordType = "NS"
	MX    DNSRecordType = "MX"
	CNAME DNSRecordType = "CNAME"
	RP    DNSRecordType = "RP"
	TXT   DNSRecordType = "TXT"
	SOA   DNSRecordType = "SOA"
	HINFO DNSRecordType = "HINFO"
	SRV   DNSRecordType = "SRV"
	DANE  DNSRecordType = "DANE"
	TLSA  DNSRecordType = "TLSA"
	DS    DNSRecordType = "DS"
	CAA   DNSRecordType = "CAA"
)

type DNSRecord struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	TTL       int
	Type      DNSRecordType
	Value     string
	ZoneId    int
	Created   time.Time
	Updated   time.Time
	AccountID int
	Account   Account
}
