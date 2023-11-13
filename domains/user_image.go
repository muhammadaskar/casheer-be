package domains

import "time"

type UserImage struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserImageInput struct {
	UserID int
	Image  string `json:"image" binding:"required"`
}
