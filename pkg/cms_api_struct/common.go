package cms_api_struct

type RequestPagination struct {
	PageNumber int `json:"pageNumber" binding:"required"`
	ShowNumber int `json:"showNumber" binding:"required"`
}

type ResponsePagination struct {
	CurrentPage int `json:"currentPage"`
	ShowNumber  int `json:"showNumber"`
}
