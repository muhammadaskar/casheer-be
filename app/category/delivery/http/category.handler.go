package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/category"
	"github.com/muhammadaskar/casheer-be/app/category/usecase"
	"github.com/muhammadaskar/casheer-be/app/helper"
)

type CategoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(categoryUseCase usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUseCase}
}

func (h *CategoryHandler) FindAll(c *gin.Context) {
	categories, err := h.categoryUseCase.FindAll()
	if err != nil {
		response := helper.APIResponse("Error to get categories", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of categories", http.StatusOK, "success", category.FormatCategories(categories))
	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) FindById(c *gin.Context) {
	var input category.GetCategoryInputID

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed to get detail of category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getCategory, err := h.categoryUseCase.FindById(input)
	if err != nil {
		response := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail of category", http.StatusOK, "success", category.FormatCategory(getCategory))
	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var input category.CreateCategoryInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Failed to create category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newCategory, err := h.categoryUseCase.Create(input)
	if err != nil {
		response := helper.APIResponse("Failed to create category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("category created successfully", http.StatusCreated, "success", category.FormatCategory(newCategory))
	c.JSON(http.StatusCreated, response)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	var inputID category.GetCategoryInputID

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update cateogry", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData category.CreateCategoryInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update category", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updateCategory, err := h.categoryUseCase.UpdateCategory(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update category", http.StatusOK, "success", category.FormatCategory(updateCategory))
	c.JSON(http.StatusOK, response)
}
