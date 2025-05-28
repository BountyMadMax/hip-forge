package dns_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Record struct {
	Created  HetznerTime `json:"created"`
	Id       string      `json:"id"`
	Modified HetznerTime `json:"modified"`
	Name     string      `json:"name"`
	Ttl      uint64      `json:"ttl"`
	Type     string      `json:"type"`
	Value    string      `json:"value"`
	ZoneId   string      `json:"zone_id"`
}

type RecordsResponse struct {
	Records []Record `json:"records"`
}

type RecordResponse struct {
	Record Record `json:"record"`
}

type BulkCreateResponse struct {
	Records        []Record           `json:"records"`
	ValidRecords   []RecordCreateBody `json:"valid_records"`
	InvalidRecords []RecordCreateBody `json:"invalid_records"`
}

type BulkUpdateResponse struct {
	Records       []Record           `json:"records"`
	FailedRecords []RecordUpdateBody `json:"failed_records"`
}

type RecordCreateBody struct {
	Name   string  `json:"name"`
	Ttl    *uint64 `json:"ttl"`
	Type   string  `json:"type"`
	Value  string  `json:"value"`
	ZoneId string  `json:"zone_id"`
}

type RecordUpdateBody struct {
	Name   string  `json:"name"`
	Ttl    *uint64 `json:"ttl"`
	Type   string  `json:"type"`
	Value  string  `json:"value"`
	ZoneId string  `json:"zone_id"`
}

type BulkCreateRecordsBody struct {
	Records []RecordCreateBody `json:"records"`
}

type BulkUpdateRecordsBody struct {
	Records []RecordUpdateBody `json:"records"`
}

func GetRecords(token string, zone_id *string) ([]Record, error) {
	url := "/records"

	if zone_id != nil {
		url = fmt.Sprintf("%s?zone_id=%s", url, *zone_id)
	}

	req, err := NewHetznerRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	AddTokenHeader(req, token)

	err = req.ParseForm()

	if err != nil {
		return nil, err
	}

	respBody, err := SendRequest(CreateClient(), req, []int{http.StatusOK})

	if err != nil {
		return nil, err
	}

	r := &RecordsResponse{}

	err = json.Unmarshal(respBody, r)

	if err != nil {
		return nil, err
	}

	return r.Records, nil
}

func GetRecord(token string, record_id string) (Record, error) {
	req, err := NewHetznerRequest("GET", fmt.Sprintf("/records/%s", record_id), nil)

	if err != nil {
		return Record{}, err
	}

	AddTokenHeader(req, token)

	respBody, err := SendRequest(CreateClient(), req, []int{http.StatusOK})

	if err != nil {
		return Record{}, err
	}

	r := &RecordResponse{}

	err = json.Unmarshal(respBody, r)

	if err != nil {
		return Record{}, err
	}

	return r.Record, nil
}

func CreateRecord(token string, record RecordCreateBody) (Record, error) {
	body, err := json.Marshal(record)

	if err != nil {
		return Record{}, nil
	}

	req, err := NewHetznerRequest("POST", "/records", bytes.NewBuffer(body))

	AddContentTypeHeader(req)
	AddTokenHeader(req, token)

	respBody, err := SendRequest(CreateClient(), req, []int{http.StatusOK})

	if err != nil {
		return Record{}, nil
	}

	r := &RecordResponse{}

	err = json.Unmarshal(respBody, r)

	if err != nil {
		return Record{}, nil
	}

	return r.Record, nil
}

func UpdateRecord(token string, record_id string, record RecordUpdateBody) (Record, error) {
	body, err := json.Marshal(record)

	if err != nil {
		return Record{}, err
	}

	req, err := NewHetznerRequest("PUT", fmt.Sprintf("/records/%s", record_id), bytes.NewBuffer(body))

	AddContentTypeHeader(req)
	AddTokenHeader(req, token)

	respBody, err := SendRequest(CreateClient(), req, []int{http.StatusOK})

	if err != nil {
		return Record{}, err
	}

	r := &RecordResponse{}

	err = json.Unmarshal(respBody, r)

	if err != nil {
		return Record{}, err
	}

	return r.Record, nil
}

func DeleteRecord(token string, record_id string) error {
	req, err := NewHetznerRequest("DELETE", fmt.Sprintf("/records/%s", record_id), nil)

	if err != nil {
		return err
	}

	AddTokenHeader(req, token)

	_, err = SendRequest(CreateClient(), req, []int{http.StatusOK})

	return err
}

func BulkCreateRecords(token string, records BulkCreateRecordsBody) (BulkCreateResponse, error) {
	body, err := json.Marshal(records)

	if err != nil {
		return BulkCreateResponse{}, err
	}

	req, err := NewHetznerRequest("POST", "/records/bulk", bytes.NewBuffer(body))

	AddContentTypeHeader(req)
	AddTokenHeader(req, token)

	respBody, err := SendRequest(CreateClient(), req, []int{http.StatusOK})

	if err != nil {
		return BulkCreateResponse{}, err
	}

	r := &BulkCreateResponse{}

	err = json.Unmarshal(respBody, r)

	if err != nil {
		return BulkCreateResponse{}, err
	}

	return *r, nil
}

func BulkUpdateRecords(token string, records BulkUpdateRecordsBody) (BulkUpdateResponse, error) {
	body, err := json.Marshal(records)

	if err != nil {
		return BulkUpdateResponse{}, err
	}

	req, err := NewHetznerRequest("PUT", "/records/bulk", bytes.NewBuffer(body))

	AddContentTypeHeader(req)
	AddTokenHeader(req, token)

	respBody, err := SendRequest(CreateClient(), req, []int{http.StatusOK})

	if err != nil {
		return BulkUpdateResponse{}, err
	}

	r := &BulkUpdateResponse{}

	err = json.Unmarshal(respBody, r)

	if err != nil {
		return BulkUpdateResponse{}, err
	}

	return *r, nil
}
