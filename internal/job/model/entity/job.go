package entity

import "time"

type Job struct {
	ID          string `gorm:"primaryKey"`
	CompanyID   string
	Title       string
	Description string
	CreatedAt   time.Time
	Tsv         string `gorm:"type:tsvector"`
}

type JobToCompany struct {
	JobID       string
	Company     string
	Title       string
	Description string
	CreatedAt   time.Time
}
