package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/category"
	"github.com/muhammadaskar/casheer-be/app/helper"
)

type CategoryHandler struct {
	categoryService category.Service
}

func NewCategoryHandler(categoryService category.Service) *CategoryHandler {
	return &CategoryHandler{categoryService}
}

func (h *CategoryHandler) FindAll(c *gin.Context) {
	categories, err := h.categoryService.FindAll()
	if err != nil {
		response := helper.APIResponse("Error to get categories", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of categories", http.StatusOK, "success", category.FormatCategories(categories))
	c.JSON(http.StatusOK, response)
}
