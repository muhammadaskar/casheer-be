package user

import "github.com/muhammadaskar/casheer-be/domains"

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func FormatUser(user domains.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return formatter
}
