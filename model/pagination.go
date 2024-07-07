package model

type Pagination struct {
	TotalCount  int  `json:"totalCount"`
	TotalPage   int  `json:"totalPage"`
	CurrentPage int  `json:"currentPage"`
	PageSize    int  `json:"pageSize"`
	HasNext     bool `json:"hasNext"`
}
