package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	discountMysql "github.com/muhammadaskar/casheer-be/app/discount/repository/mysql"
	memberMysql "github.com/muhammadaskar/casheer-be/app/member/repository/mysql"
	productMysql "github.com/muhammadaskar/casheer-be/app/product/repository/mysql"
	"github.com/muhammadaskar/casheer-be/app/transaction"
	transactionMysql "github.com/muhammadaskar/casheer-be/app/transaction/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
)

type TransactionUseCase interface {
	FindAll() ([]domains.CustomTransaction, error)
	FindById(input domains.GetInputIdTransaction) (domains.CustomTransaction, error)
	FindAllMember() ([]domains.CustomTransactionMember, error)
	FindMemberById(input domains.GetInputIdTransaction) (domains.CustomTransactionMember, error)
	GetAmountOneMonthAgo() (domains.CustomTransactionAmount, error)
	GetItemOutOneMonthAgo() (domains.CustomTransactionTotalQuantity, error)
	GetCountTransactionThisYear() ([]domains.GetCountTransactionThisYear, error)
	Create(input transaction.CreateInput) (domains.Transaction, error)
}

type usecase struct {
	transactionRepo transactionMysql.Repository
	memberRepo      memberMysql.Repository
	productRepo     productMysql.Repository
	discountRepo    discountMysql.Repository
}

func NewUseCase(transactionRepo transactionMysql.Repository,
	memberRepo memberMysql.Repository,
	productRepo productMysql.Repository,
	discountRepo discountMysql.Repository) *usecase {
	return &usecase{transactionRepo, memberRepo, productRepo, discountRepo}
}

func (u *usecase) FindAll() ([]domains.CustomTransaction, error) {
	transactions, err := u.transactionRepo.FindAll()
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (u *usecase) FindAllMember() ([]domains.CustomTransactionMember, error) {
	transactions, err := u.transactionRepo.FindAllMember()
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (u *usecase) FindById(input domains.GetInputIdTransaction) (domains.CustomTransaction, error) {
	transaction, err := u.transactionRepo.FindById(input.ID)
	if err != nil {
		return transaction, err
	}
	if transaction.ID == 0 {
		return transaction, errors.New("Transaction not found")
	}

	return transaction, nil
}

func (u *usecase) FindMemberById(input domains.GetInputIdTransaction) (domains.CustomTransactionMember, error) {
	transaction, err := u.transactionRepo.FindMemberById(input.ID)
	if err != nil {
		return transaction, err
	}

	if transaction.ID == 0 {
		return transaction, errors.New("Transaction not found")
	}

	if transaction.MemberCode == "" {
		return transaction, errors.New("Member not found")
	}

	return transaction, nil
}

func (u *usecase) GetAmountOneMonthAgo() (domains.CustomTransactionAmount, error) {
	currentTime := time.Now()
	oneMonthAgo := currentTime.AddDate(0, -1, 0)

	layout := "2006-01-02 15:04:05.999" // Format layout

	currentFormatted := currentTime.Format(layout)
	oneMonthAgoFormatted := oneMonthAgo.Format(layout)

	transaction, err := u.transactionRepo.GetAmountOneMonthAgo(currentFormatted, oneMonthAgoFormatted)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (u *usecase) GetItemOutOneMonthAgo() (domains.CustomTransactionTotalQuantity, error) {
	currentTime := time.Now()
	oneMonthAgo := currentTime.AddDate(0, -1, 0)

	layout := "2006-01-02 15:04:05.999" // Format layout

	currentFormatted := currentTime.Format(layout)
	oneMonthAgoFormatted := oneMonthAgo.Format(layout)

	transaction, err := u.transactionRepo.GetItemOneOutMonthAgo(currentFormatted, oneMonthAgoFormatted)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (u *usecase) GetCountTransactionThisYear() ([]domains.GetCountTransactionThisYear, error) {
	now := time.Now()
	currentYear := now.Year()
	currentMonth := now.Month()
	start := strconv.Itoa(currentYear) + "-01-01 00:00:00.00"
	end := strconv.Itoa(currentYear) + "-12-31 23:59:00.00"

	transactions, err := u.transactionRepo.GetCountTransactionThisYear(start, end)
	if err != nil {
		return transactions, err
	}

	currentMonthInt := int(currentMonth)

	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	var filteredTransactions []domains.GetCountTransactionThisYear
	for _, transaction := range transactions {
		month, err := strconv.Atoi(transaction.Month)
		if err != nil {
			// Handle kesalahan konversi jika diperlukan
			continue
		}

		if month <= currentMonthInt {
			filteredTransactions = append(filteredTransactions, transaction)
		}
	}

	transactions = filteredTransactions

	for i := range transactions {
		transactions[i].Month = months[i]
	}

	return transactions, nil
}

func (u *usecase) Create(input transaction.CreateInput) (domains.Transaction, error) {
	transaction := domains.Transaction{}
	memberCode := input.MemberCode
	var totalAmount int
	var totalQuantity int

	if memberCode != "" {
		member, err := u.memberRepo.FindByMemberCode(memberCode)
		if err != nil {
			return transaction, err
		}

		if member.ID == 0 {
			return transaction, errors.New("Member code is not available")
		}

		if member.IsActive != 0 {
			return transaction, errors.New("Member is not active")
		}

		discount, err := u.discountRepo.FindById()
		if err != nil {
			return transaction, err
		}

		transaction.MemberCode = input.MemberCode
		transaction.TransactionCode = generateTransactionCode()

		productQuantities, err := u.productParse(input)
		if err != nil {
			return transaction, err
		}
		var productData []map[string]string

		for _, pq := range productQuantities {
			product, err := u.productRepo.FindByProductID(pq.ProductID)
			if err != nil {
				fmt.Printf("Error finding product with ID %d: %s\n", pq.ProductID, err)
				continue
			}

			result := calculateTotalPrice(float64(product.Price), float64(discount.Discount), pq.Quantity)
			totalAmount += int(result)

			productQuantity := product.Quantity - pq.Quantity
			product.Quantity = productQuantity

			productData = append(productData, map[string]string{"product_id": strconv.Itoa(product.ID), "product_name": product.Name, "quantity": strconv.Itoa(pq.Quantity)})

			_, err = u.productRepo.Update(product)
			if err != nil {
				return transaction, err
			}
		}

		quantities, err := u.quantityParse(input)
		if err != nil {
			return transaction, err
		}

		for _, quantity := range quantities {
			totalQuantity += quantity.Quantity
		}

		jsonBytes, err := json.Marshal(productData)
		jsonString := string(jsonBytes)

		transaction.Amount = totalAmount
		transaction.UserID = input.User.ID
		transaction.Transactions = jsonString
		transaction.TotalQuantity = totalQuantity

		newTransaction, err := u.transactionRepo.Create(transaction)
		if err != nil {
			return newTransaction, err
		}
		return newTransaction, nil
	} else {
		// kondisi jika bukan member, maka tidak akan mendapatkan discount
		transaction.MemberCode = input.MemberCode
		transaction.TransactionCode = generateTransactionCode()

		productQuantities, err := u.productParse(input)
		if err != nil {
			return transaction, err
		}

		var productData []map[string]string

		for _, pq := range productQuantities {
			product, err := u.productRepo.FindByProductID(pq.ProductID)
			if err != nil {
				fmt.Printf("Error finding product with ID %d: %s\n", pq.ProductID, err)
				continue
			}

			result := (product.Price * pq.Quantity)
			totalAmount += result

			productQuantity := product.Quantity - pq.Quantity
			product.Quantity = productQuantity

			productData = append(productData, map[string]string{"product_id": strconv.Itoa(product.ID), "product_name": product.Name, "quantity": strconv.Itoa(pq.Quantity)})

			_, err = u.productRepo.Update(product)
			if err != nil {
				return transaction, err
			}
		}

		quantities, err := u.quantityParse(input)
		if err != nil {
			return transaction, err
		}

		for _, quantity := range quantities {
			totalQuantity += quantity.Quantity
		}
		jsonBytes, err := json.Marshal(productData)
		jsonString := string(jsonBytes)

		transaction.Amount = totalAmount
		transaction.UserID = input.User.ID
		transaction.Transactions = jsonString
		transaction.TotalQuantity = totalQuantity

		newTransaction, err := u.transactionRepo.Create(transaction)
		if err != nil {
			return newTransaction, err
		}
		return newTransaction, nil
	}
}

func generateTransactionCode() string {
	rand.Seed(time.Now().UnixNano())
	length := 16

	const availableChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(availableChars))
		code[i] = availableChars[randomIndex]
	}

	return string(code)
}

func calculateTotalPrice(originalPrice float64, discountPercentage float64, quantity int) float64 {
	totalOriginalPrice := originalPrice * float64(quantity)
	discountAmount := (discountPercentage / 100) * totalOriginalPrice
	totalPriceAfterDiscount := totalOriginalPrice - discountAmount
	return totalPriceAfterDiscount
}

func parseInput(input string) []domains.TransactionProductQuantity {
	input = strings.Trim(input, "{}")    // Remove outer curly braces
	pairs := strings.Split(input, "},{") // Split into individual pairs

	var transactions []domains.TransactionProductQuantity

	for _, pair := range pairs {
		parts := strings.Split(pair, ",")
		if len(parts) != 2 {
			continue
		}

		productID, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		quantity, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

		if err1 == nil && err2 == nil {
			transactions = append(transactions, domains.TransactionProductQuantity{ProductID: productID, Quantity: quantity})
		}
	}

	return transactions
}

func (u *usecase) productParse(input transaction.CreateInput) ([]domains.TransactionProductQuantity, error) {
	productQuantities := parseInput(input.Transactions)
	for _, proudctQuantity := range productQuantities {
		product, err := u.productRepo.FindByProductID(proudctQuantity.ProductID)
		if err != nil {
			fmt.Printf("Error finding product with ID %d: %s\n", proudctQuantity.ProductID, err)
			continue
		}

		if err != nil {
			return productQuantities, err
		}

		if product.ID == 0 || product.IsDeleted == 0 {
			return productQuantities, errors.New("product id is not available")
		}

		if proudctQuantity.Quantity > product.Quantity {
			return productQuantities, errors.New("product quantity " + product.Name + " is not enough")
		}
	}
	return productQuantities, nil
}

func (u *usecase) quantityParse(input transaction.CreateInput) ([]domains.TransactionProductQuantity, error) {
	productQuantities := parseInput(input.Transactions)
	return productQuantities, nil
}
