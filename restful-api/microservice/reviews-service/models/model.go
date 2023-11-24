package models

// Reviews

type Reviews struct {
	ID        int    `json:"id" uri:"id" gorm:"primaryKey"`
	Comment   string `json:Comments`
	ProductID int    `json:product_id`
}
