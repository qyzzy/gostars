package models

import "time"

type Authority struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `gorm:"index"`
	AuthorityID   string     `json:"authorityid" gorm:"not null;unique;primary_key;size:90"`
	AuthorityName string     `json:"authorityname" gorm:""`
}

func AuthorityTableName() string {
	return "authorities"
}
