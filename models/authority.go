package models

import "time"

type Authority struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `gorm:"index"`
	AuthorityID   int        `json:"authorityid" gorm:"not null;unique;primary_key"`
	AuthorityName string     `json:"authorityname" gorm:""`
}

func AuthorityTableName() string {
	return "authorities"
}
