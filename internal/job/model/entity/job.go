package entity

import "time"

type Job struct {
	ID          string    `db:"id"`
	CompanyID   string    `db:"company_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}
