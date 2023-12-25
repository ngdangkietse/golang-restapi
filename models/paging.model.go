package models

type Paging struct {
	PageSize  int `json:"pageSize"`
	PageLimit int `json:"pageLimit"`
}
