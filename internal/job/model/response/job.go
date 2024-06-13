package response

type GetJobsResponse struct {
	StatusCode int                   `json:"status_code" example:"200"`
	Message    string                `json:"message" example:"success get job list"`
	Data       []GetJobsResponseData `json:"data"`
}

type GetJobsResponseData struct {
	CompanyName string `json:"company_name" example:"Redikru"`
	Title       string `json:"title" example:"Backend Engineer"`
	Description string `json:"description" example:"We need passionate backend engineer!"`
}
