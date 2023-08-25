package notification

// "github.com/muhammadaskar/casheer-be/app/product"
// "github.com/muhammadaskar/casheer-be/app/user"

type GetInputID struct {
	ID int `uri:"id" binding:"required"`
}

type InputUpdateNotification struct {
	IsRead int `json:"is_read"`
}
