package model

type Bill struct {
	ID            int64      `json:"id,omitempty"`
	UserID        int64      `json:"userId"`
	ShopID        int64      `json:"shopId"`
	Positions     []Position `json:"positions,omitempty"`
	TotalDiscount float64    `json:"totalDiscount"`
	Timestamp     string     `json:"timestamp,omitempty"`
}

type Position struct {
	ID          int64   `json:"id"`
	Description string  `json:"description"`
	Amount      int     `json:"amount"`
	SinglePrice float64 `json:"singlePrice"`
	Discount    float64 `json:"discount"`
	BillID      int64   `json:"billId"`
}
