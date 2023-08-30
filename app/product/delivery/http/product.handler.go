package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/product"
	"github.com/muhammadaskar/casheer-be/app/product/usecase"
	"github.com/muhammadaskar/casheer-be/domains"
	customresponse "github.com/muhammadaskar/casheer-be/utils/custom_response"
)

type ProductHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductHandler(productUseCase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{productUseCase}
}

func (h *ProductHandler) FindAll(c *gin.Context) {
	var query product.GetProductsQueryInput
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response := customresponse.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, isLastPage, err := h.productUseCase.FindAll(query)
	if err != nil {
		response := customresponse.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_last_page": isLastPage,
		"products":     products,
	}

	response := customresponse.APIResponse("List of products", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	var query product.GetProductsQueryInput
	err := c.ShouldBindQuery(&query)
	if err != nil {
		response := customresponse.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	products, err := h.productUseCase.GetAll()
	if err != nil {
		response := customresponse.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("List of products", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) CountProducts(c *gin.Context) {
	count, err := h.productUseCase.CountAll()
	if err != nil {
		response := customresponse.APIResponse("Error to get count products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	response := customresponse.APIResponse("Get count data products", http.StatusOK, "success", count)
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) FindById(c *gin.Context) {
	var input product.GetProductDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := customresponse.APIResponse("Error to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	productDetail, err := h.productUseCase.FindById(input)
	if err != nil {
		response := customresponse.APIResponse("Error to get product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Product detail", http.StatusOK, "success", productDetail)
	c.JSON(http.StatusOK, response)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input product.CreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)
	input.User = currentUser

	newProduct, err := h.productUseCase.Create(input)
	if err != nil {
		response := customresponse.APIResponse("Failed to create product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to create product", http.StatusCreated, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusCreated, response)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var inputID product.GetProductDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to update product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input product.CreateInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to update product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)
	input.User = currentUser

	updateProduct, err := h.productUseCase.Update(inputID, input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse("Failed to update product", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to update product", http.StatusCreated, "success", product.FormatProduct(updateProduct))
	c.JSON(http.StatusCreated, response)
}

func (h *ProductHandler) UpdateProductQuantity(c *gin.Context) {
	var inputID product.GetProductDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to update product quantity", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var input product.UpdateQuantity

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to update product quantity", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.productUseCase.UpdateQuantity(inputID, input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse("Failed to update product quantity", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to update product quantity", http.StatusCreated, "success", nil)
	c.JSON(http.StatusCreated, response)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	var inputID product.GetProductDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to delete product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	_, err = h.productUseCase.Delete(inputID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to delete product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := customresponse.APIResponse("Success to delete product", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
