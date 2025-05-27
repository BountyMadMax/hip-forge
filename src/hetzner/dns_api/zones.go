package dns_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Zone struct {
	// Hash string.
	Id string `json:"id"`
	// Name of the domain.
	Name          string `json:"name"`
	Ttl           uint64 `json:"ttl"`
	Registrar     string `json:"registrar"`
	LegacyDnsHost string `json:"legacy_dns_host"`
	// List of domains.
	LegacyNs []string `json:"legacy_ns"`
	// List of domains.
	Ns         []string    `json:"ns"`
	Created    HetznerTime `json:"created"`
	Verified   HetznerTime `json:"verified"`
	Modified   HetznerTime `json:"modified"`
	Project    string      `json:"project"`
	Owner      string      `json:"owner"`
	Permission string      `json:"permission"`
	ZoneType   ZoneType    `json:"zone_type"`
	// "verified", "failed" or "pending".
	Status          string          `json:"status"`
	Paused          bool            `json:"paused"`
	IsSecondaryDns  bool            `json:"is_secondary_dns"`
	TxtVerification TxtVerification `json:"txt_verification"`
	RecordsCount    int             `json:"records_count"`
}

type ZoneType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Prices      string `json:"prices"`
}

type TxtVerification struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type ZonesResponse struct {
	Zones []Zone `json:"zones"`
}

type ZoneResponse struct {
	Zone Zone `json:"zone"`
}

type ZoneCreateBody struct {
	Name string  `json:"name"`
	Ttl  *uint64 `json:"ttl"`
}

type ZoneUpdateBody struct {
	Name string  `json:"name"`
	Ttl  *uint64 `json:"ttl"`
}

func GetZones(token string) ([]Zone, error) {
	req, err := NewHetznerRequest("GET", "/zones", nil)

	if err != nil {
		return nil, err
	}

	AddAcceptHeader(req)
	AddTokenHeader(req, token)

	respBody, err := SendRequest(CreateClient(), req, []int{http.StatusOK})

	if err != nil {
		return nil, err
	}

	r := &ZonesResponse{}

	err = json.Unmarshal(respBody, r)

	if err != nil {
		return nil, err
	}

	return r.Zones, nil
}

func GetZone(token string, zone_id string) (Zone, error) {
	req, err := NewHetznerRequest("GET", fmt.Sprintf("/zones/%s", zone_id), nil)

	if err != nil {
		return Zone{}, err
	}

	AddAcceptHeader(req)
	AddTokenHeader(req, token)

	respBody, err := SendRequest(CreateClient(), req, []int{http.StatusOK})

	if err != nil {
		return Zone{}, nil
	}

	r := &ZoneResponse{}

	err = json.Unmarshal(respBody, r)

	if err != nil {
		return Zone{}, err
	}

	return r.Zone, nil
}

func CreateZone(token string, zone ZoneCreateBody) (Zone, error) {
	body, err := json.Marshal(zone)

	if err != nil {
		return Zone{}, err
	}

	req, err := NewHetznerRequest("POST", "/zones", bytes.NewBuffer(body))

	AddContentTypeHeader(req)
	AddTokenHeader(req, token)

	respBody, err := SendRequest(CreateClient(), req, []int{http.StatusCreated})

	if err != nil {
		return Zone{}, err
	}

	r := &ZoneResponse{}

	err = json.Unmarshal(respBody, r)

	if err != nil {
		return Zone{}, err
	}

	return r.Zone, nil
}

func UpdateZone(token string, zone_id string, zone ZoneUpdateBody) (Zone, error) {
	body, err := json.Marshal(zone)

	if err != nil {
		return Zone{}, err
	}

	req, err := NewHetznerRequest("PUT", fmt.Sprintf("zones/%s", zone_id), bytes.NewBuffer(body))

	AddContentTypeHeader(req)
	AddTokenHeader(req, token)

	respBody, err := SendRequest(CreateClient(), req, []int{http.StatusOK})

	if err != nil {
		return Zone{}, err
	}

	r := &ZoneResponse{}

	err = json.Unmarshal(respBody, r)

	if err != nil {
		return Zone{}, err
	}

	return r.Zone, nil
}

func DeleteZone(token string, zone_id string) error {
	req, err := NewHetznerRequest("DELETE", fmt.Sprintf("/zones/%s", zone_id), nil)

	if err != nil {
		return err
	}

	AddTokenHeader(req, token)

	_, err = SendRequest(CreateClient(), req, []int{http.StatusOK})

	return err
}
