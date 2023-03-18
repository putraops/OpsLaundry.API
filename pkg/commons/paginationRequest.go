package commons

type PaginationRequest struct {
	Page   int      `json:"page" form:"page"`
	Size   int      `json:"size" form:"size"`
	Search *string  `json:"search"`
	Filter []Filter `json:"filter"`
	Order  *[]Order `json:"order"`
	// Column *[]Column `json:"column"`
}

type PaginationResponse struct {
	RecordsTotal    int         `json:"recordsTotal" form:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered" form:"recordsFiltered"`
	Data            interface{} `json:"data" form:"data"`
	Draw            int         `json:"draw" form:"draw"`
	Error           interface{} `json:"error" form:"error"`
}

type Filter struct {
	Field    string `json:"field" form:"field"`
	Operator string `json:"operator" form:"operator"`
	Value    string `json:"value" form:"value"`
}

type Order struct {
	Field     string `json:"field" form:"field"`
	Direction string `json:"direction" form:"direction"`
}

type Column struct {
	Name        string `json:"name" form:"name"`
	Orderable   *bool  `json:"orderable" form:"orderable"`
	Searchable  *bool  `json:"searchable" form:"searchable"`
	SearchValue string `json:"searchValue" form:"searchValue"`
}
