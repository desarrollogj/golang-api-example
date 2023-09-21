package domain

import "time"

type GenericEntity struct {
	Reference   string
	IsActive    bool
	CreatedDate time.Time
	UpdatedDate time.Time
}

type SearchInput struct {
	Page     int
	PageSize int
}

type SearchOutput struct {
	Total    int64
	Page     int
	PageSize int
}
