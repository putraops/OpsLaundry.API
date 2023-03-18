package commons

type DataTableRequest struct {
	Columns []DataTableColumn `json:"columns"`
	Filters []DataTableFilter `json:"filters"`
	Orders  []DataTableOrder  `json:"orders"`
	Page    int               `json:"page" form:"page"`
	Search  *string           `json:"search"`
	Size    int               `json:"size" form:"size"`
}

type DataTableFilter struct {
	Field    string `json:"field" form:"field"`
	Operator string `json:"operator" form:"operator"`
	Value    string `json:"value" form:"value"`
}

type DataTableColumn struct {
	Name        string `json:"name" form:"name"`
	Orderable   *bool  `json:"orderable" form:"orderable"`
	Searchable  *bool  `json:"searchable" form:"searchable"`
	SearchValue string `json:"searchValue" form:"searchValue"`
}

type DataTableSearch struct {
	Value string `json:"value" form:"value"`
	Regex bool   `json:"regex" form:"regex"`
}

type DataTableOrder struct {
	Field     string `json:"field" form:"field"`
	Direction string `json:"direction" form:"direction"`
}

type DataTableDefaultOrder struct {
	Column string `json:"column" form:"column"`
	Dir    string `json:"dir" form:"dir"`
}
