package category

type GetCategoryInputID struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCategoryInput struct {
	Name string `json:"name" binding:"required"`
}
