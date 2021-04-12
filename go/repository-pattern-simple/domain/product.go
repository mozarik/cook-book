package domain

type ID uint

type Product struct {
	ID         ID      `json:"id"`
	Name       string  `json:"name"`
	Price      uint    `json:"price"`
	Rate       float32 `json:"rate"`
	CategoryID uint    `json:"categoryId"`
	Review     uint    `json:"review"`
	ImageURL   string  `json:"imageUrl"`
}
