package models

import "time"

// Page ...
type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

// PageUtil ...
func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}

// MockData ...
type MockData struct {
	Id         int
	Path       string
	Data       string
	Tip        string
	Check      bool
	CreateTime time.Time
	UpdateTime time.Time
}
