package member

import "github.com/muhammadaskar/casheer-be/domains"

type CreateInput struct {
	Name  string `json:"name" binding:"required,min=5"`
	Phone string `json:"phone" binding:"required,min=11,max=13"`
	User  domains.User
}

type GetMemberIDInput struct {
	ID int `uri:"id" binding:"required"`
}
