package request

type GetJobsRequest struct {
	Keyword     string `query:"keyword"`
	CompanyName string `query:"companyName"`
}

type CreateJobRequest struct {
	CompanyName string `json:"company_name" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}
