package discount

import "github.com/muhammadaskar/casheer-be/domains"

type CreateInput struct {
	Discount int `json:"discount" bind:"required"`
	User     domains.User
}
