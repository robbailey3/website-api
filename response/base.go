package response

type Pagination struct {
	CurrentPage  int `json:"currentPage"`
	ItemsPerPage int `json:"itemsPerPage"`
	CurrentItems int `json:"currentItems"`
	TotalItems   int `json:"totalItems"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BaseResponse struct {
	Result    interface{}    `json:"result,omitempty"`
	Timestamp int64          `json:"timestamp"`
	Success   bool           `json:"success"`
	Error     *ErrorResponse `json:"error,omitempty"`
	*Pagination
}
