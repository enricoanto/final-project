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
		ID         int       `json:"id"`
		Title      string    `json:"title"`
		Price      int       `json:"price"`
		Stock      int       `json:"stock"`
		CategoryID int       `json:"category_id"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}

	Category struct {
		ID                int
		Type              string
		SoldProductAmount int
		Products          []Product
		CreatedAt         time.Time
		UpdatedAt         time.Time
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
