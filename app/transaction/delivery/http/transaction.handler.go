package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadaskar/casheer-be/app/transaction"
	"github.com/muhammadaskar/casheer-be/app/transaction/usecase"
	"github.com/muhammadaskar/casheer-be/domains"
	customresponse "github.com/muhammadaskar/casheer-be/utils/custom_response"
)

type TransactionHandler struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewTransactionHandler(transactionUseCase usecase.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{transactionUseCase}
}

func (h *TransactionHandler) FindAll(c *gin.Context) {
	transactions, err := h.transactionUseCase.FindAll()
	if err != nil {
		response := customresponse.APIResponse("Failed to get transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get transactions", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) FindAllMember(c *gin.Context) {
	transactions, err := h.transactionUseCase.FindAllMember()
	if err != nil {
		response := customresponse.APIResponse("Failed to get transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get transactions", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) FindById(c *gin.Context) {
	var input domains.GetInputIdTransaction

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to get transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transaction, err := h.transactionUseCase.FindById(input)
	if err != nil {
		response := customresponse.APIResponse("Failed to get transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get transactions", http.StatusOK, "success", transaction)
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) FindMemberById(c *gin.Context) {
	var input domains.GetInputIdTransaction

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to get transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	transaction, err := h.transactionUseCase.FindMemberById(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := customresponse.APIResponse("Failed to get transactions", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get transactions", http.StatusOK, "success", transaction)
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetAmountOneMonthAgo(c *gin.Context) {
	transactions, err := h.transactionUseCase.GetAmountOneMonthAgo()
	if err != nil {
		response := customresponse.APIResponse("Failed to get amount transaction for one month ago", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get amount transaction for one month ago", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetItemOutOneMonthAgo(c *gin.Context) {
	transactions, err := h.transactionUseCase.GetItemOutOneMonthAgo()
	if err != nil {
		response := customresponse.APIResponse("Failed to get items transaction for one month ago", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get items transaction for one month ago", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetCountTransactionThisYear(c *gin.Context) {
	transactions, err := h.transactionUseCase.GetCountTransactionThisYear()
	if err != nil {
		response := customresponse.APIResponse("Failed to get count transaction for this year", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get count transaction for this year", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetAmountTransactionThisYear(c *gin.Context) {
	transactions, err := h.transactionUseCase.GetAmountTransactionThisYear()
	if err != nil {
		response := customresponse.APIResponse("Failed to get amount transaction for this year", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to get amount transaction for this year", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := customresponse.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := customresponse.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(domains.User)
	input.User = currentUser

	newTransaction, err := h.transactionUseCase.Create(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := customresponse.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := customresponse.APIResponse("Success to create transaction", http.StatusCreated, "success", newTransaction)
	c.JSON(http.StatusCreated, response)
}
