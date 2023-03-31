package commons

type DataTableRequest struct {
	Columns []DataTableColumn `json:"columns"`
	Filters []DataTableFilter `json:"filters"`
	Orders  []DataTableOrder  `json:"orders"`
	Page    int               `json:"page" form:"page"`
	Search  *string           `json:"search"`
	Length  int               `json:"length" form:"length"`
}

type DataTableFilter struct {
	Column   string `json:"column" form:"column"`
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
	Column    string `json:"column" form:"column"`
	Direction string `json:"direction" form:"direction"`
}

type DataTableDefaultOrder struct {
	Column string `json:"column" form:"column"`
	Dir    string `json:"dir" form:"dir"`
}
