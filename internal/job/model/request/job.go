package request

import "github.com/handikacatur/jobs-api/internal/model"

type GetJobsRequest struct {
	Keyword    string
	Pagination model.PaginationMetaMessage
}
