package request

type GetJobsRequest struct {
	Keyword     string `query:"keyword"`
	CompanyName string `query:"companyName"`
}
