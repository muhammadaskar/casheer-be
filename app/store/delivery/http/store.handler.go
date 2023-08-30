package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/store"
	"github.com/muhammadaskar/casheer-be/app/store/usecase"
	customresponse "github.com/muhammadaskar/casheer-be/utils/custom_response"
)

type StoreHandler struct {
	usecase usecase.StoreUseCase
}

func NewStoreHandler(usecase usecase.StoreUseCase) *StoreHandler {
	return &StoreHandler{usecase}
}

func (h *StoreHandler) FindOne(c *gin.Context) {
	store, err := h.usecase.FindOne()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse("Failed to get store information", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get store information", http.StatusOK, "success", store)
	c.JSON(http.StatusOK, response)
}

func (h *StoreHandler) Create(c *gin.Context) {
	var input store.CreateInput
	err := c.BindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to create store information", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newStore, err := h.usecase.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to create store information", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := customresponse.APIResponse("Success to create store information", http.StatusCreated, "error", newStore)
	c.JSON(http.StatusCreated, response)
}

func (h *StoreHandler) Update(c *gin.Context) {
	var input store.CreateInput
	err := c.BindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to update store information", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newStore, err := h.usecase.Update(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to update store information", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := customresponse.APIResponse("Success to update store information", http.StatusOK, "success", newStore)
	c.JSON(http.StatusOK, response)
}
