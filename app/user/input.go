package user

type RegisterInput struct {
	Name     string `json:"name" binding:"required,min=5"`
	Username string `json:"username" binding:"required,min=6,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type GetUserIDInput struct {
	ID int `uri:"id" binding:"required"`
}
