package entity

import "time"

type Company struct {
	ID   string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name string `gorm:"unique_index:idx_name"`
}

type Job struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4()"`
	CompanyID   string
	Title       string
	Description string
	CreatedAt   time.Time
}

type JobToCompany struct {
	JobID       string
	Company     string
	Title       string
	Description string
	CreatedAt   time.Time
}
