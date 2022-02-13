package models

type PaginatedResult struct {
	Pagination Pagination        `json:"pagination"`
	Result     map[string]string `json:"result"`
}

type Pagination struct {
	Count int32 `json:"count"`
	Limit int32 `json:"limit"`
	Skip  int32 `json:"skip"`
}
