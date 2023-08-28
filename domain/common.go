package domain

import "time"

type GenericEntity struct {
	Reference   string
	IsActive    bool
	CreatedDate time.Time
	UpdatedDate time.Time
}
