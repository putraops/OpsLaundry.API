package commons

type DataTableResponse struct {
	RecordsTotal    int         `json:"recordsTotal" form:"recordsTotal"`
	RecordsFiltered int         `json:"recordsFiltered" form:"recordsFiltered"`
	Data            interface{} `json:"data" form:"data"`
	Draw            int         `json:"draw" form:"draw"`
	Error           interface{} `json:"error" form:"error"`
}
