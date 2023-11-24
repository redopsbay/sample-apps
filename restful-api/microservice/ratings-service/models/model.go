package models

// Ratings

type Ratings struct {
	ID        int    `json:"id" uri:"id" gorm:"primaryKey"`
	Rates     int    `json:rates`
	Category  string `json:category`
	ProductID int    `json:product_id`
}
