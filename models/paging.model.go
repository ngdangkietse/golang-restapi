package models

type Paging struct {
	PageIndex     int    `json:"pageIndex"`
	PageSize      int    `json:"pageSize"`
	PageSort      string `json:"pageSort"`
	PageDirection string `json:"pageDirection"`
}
