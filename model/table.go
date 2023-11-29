package model

type Table struct {
	Table_id     int    `json:"table_id"`
	Table_number int    `json:"table_number"`
	Capacity     int    `json:"capacity"`
	Location     string `json:"location"`
	Description  string `json:"description"`
}

type Table_num struct {
	Table_number int
}
