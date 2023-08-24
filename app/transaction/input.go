package transaction

import "github.com/muhammadaskar/casheer-be/domains"

type CreateInput struct {
	// MemberCode   string `json:"member_code"`
	// Transacation string `json:"transactions" binding:"required"`
	MemberCode   string `json:"member_code"`
	Transactions string `json:"transactions" binding:"required"`
	// ProductID  string `json:"product_id" binding:"required"`
	// Quantity   string `json:"quantity" binding:"required"`
	// Amount     int    `json:"amount" binding:"required"`
	User domains.User
}
