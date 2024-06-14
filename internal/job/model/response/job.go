package response

import "time"

type GetJobsResponse struct {
	StatusCode int                   `json:"status_code" example:"200"`
	Message    string                `json:"message" example:"success get job list"`
	Data       []GetJobsResponseData `json:"data"`
}

type GetJobsResponseData struct {
	JobID       string    `json:"job_id" example:"a5024c8c-2f7c-4e38-8324-9fac9f84acbf"`
	CompanyName string    `json:"company_name" example:"Redikru"`
	Title       string    `json:"title" example:"Backend Engineer"`
	Description string    `json:"description" example:"We need passionate backend engineer!"`
	CreatedAt   time.Time `json:"created_at" example:""`
}
