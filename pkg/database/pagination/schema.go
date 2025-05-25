package pagination

type PageRequest struct {
	Page  int64  `json:"page" validate:"omitempty,numeric"`
	Limit int64  `json:"limit" validate:"omitempty,numeric"`
	Sort  string `json:"sort" validate:"omitempty"`
}

type PageResponse struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	HasNext    bool  `json:"hasNext"`
	HasPrev    bool  `json:"hasPrev"`
	NextPage   int64 `json:"nextPage"`
	PrevPage   int64 `json:"prevPage"`
	From       int64 `json:"from"`
	To         int64 `json:"to"`
}
