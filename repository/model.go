package model

import "time"

type (
	User struct {
		ID        int       `json:"id"`
		FullName  string    `json:"full_name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Role      string    `json:"role"`
		Balance   int       `json:"balance"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Product struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Price     int       `json:"price"`
		Stock     int       `json:"stock"`
		CategoryID int `json:"category_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Category struct {
		ID                int `json:"id"`
		Type              string `json:"type"`
		SoldProductAmount int `json:"sold_product_amount"`
		Products          []Product `json:"Products,omitempty"`
		CreatedAt         time.Time `json:"created_at,omitempty"`
		UpdatedAt         time.Time `json:"updated_at,omitempty"`
	}

	TransactionHistory struct {
		ID         int
		ProductID  int
		UserID     int
		Quantity   int
		TotalPrice int
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
)
