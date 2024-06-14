package response

type ResponseStatusOnly struct {
	StatusCode int    `json:"status_code" example:"400"`
	Message    string `json:"message" example:"strconv.ParseInt: parsing \"a\": invalid syntax"`
}

type Error struct {
	StatusCode int    `json:"status_code" example:"400"`
	Message    string `json:"message" example:"strconv.ParseInt: parsing \"a\": invalid syntax"`
	ErrorCode  int    `json:"error_code" example:"1001"`
}
