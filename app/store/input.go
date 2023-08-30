package store

type CreateInput struct {
	Name  string `json:"name" binding:"required"`
	Image string `json:"image" binding:"required"`
}
