package model

// User ...
type User struct {
	ID      int `json:"id"`
	Balance int `json:"balance"`
}

type ReservedFund struct {
	ID        int `json:"id"`
	IDService int `json:"id_service"`
	IDOrder   int `json:"id_order"`
	Price     int `json:"price"`
}

type AccountingReport struct {
	IDService int `json:"id_service"`
	Month     int `json:"month"`
	Year      int `json:"year"`
	Funds     int `json:"funds"`
}
