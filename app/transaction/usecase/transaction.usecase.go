package usecase

import (
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

func (u *usecase) Create(input transaction.CreateInput) (domains.Transaction, error) {
	transaction := domains.Transaction{}
	memberCode := input.MemberCode
	var totalAmount int

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

		for _, proudctQuantity := range productQuantities {
			product, err := u.productRepo.FindByProductID(proudctQuantity.ProductID)
			if err != nil {
				fmt.Printf("Error finding product with ID %d: %s\n", proudctQuantity.ProductID, err)
				continue
			}

			if product.ID == 0 {
				return transaction, errors.New("product id is not available")
			}

			if proudctQuantity.Quantity > product.Quantity {
				return transaction, errors.New("Product quantity " + product.Name + " is not enough")
			}

			result := calculateTotalPrice(float64(product.Price), float64(discount.Discount), proudctQuantity.Quantity)
			totalAmount += int(result)

			productQuantity := product.Quantity - proudctQuantity.Quantity
			product.Quantity = productQuantity

			_, err = u.productRepo.Update(product)
			if err != nil {
				return transaction, err
			}
		}

		transaction.Amount = totalAmount
		transaction.UserID = input.User.ID
		transaction.Transactions = input.Transactions
		newTransaction, err := u.transactionRepo.Create(transaction)
		if err != nil {
			return newTransaction, err
		}
		return newTransaction, nil

	} else {
		// kondisi jika bukan member, maka tidak akan mendapatkan discount
		transaction.MemberCode = input.MemberCode
		transaction.TransactionCode = generateTransactionCode()

		// productQuantities := parseInput(input.Transactions)
		productQuantities, err := u.productParse(input)
		if err != nil {
			return transaction, err
		}

		for _, proudctQuantity := range productQuantities {
			product, err := u.productRepo.FindByProductID(proudctQuantity.ProductID)
			if err != nil {
				fmt.Printf("Error finding product with ID %d: %s\n", proudctQuantity.ProductID, err)
				continue
			}

			if err != nil {
				return transaction, err
			}
			if product.ID == 0 {
				return transaction, errors.New("product id is not available")
			}

			if proudctQuantity.Quantity > product.Quantity {
				return transaction, errors.New("Product quantity " + product.Name + " is not enough")
			}

			result := (product.Price * proudctQuantity.Quantity)
			totalAmount += result

			productQuantity := product.Quantity - proudctQuantity.Quantity
			product.Quantity = productQuantity

			_, err = u.productRepo.Update(product)
			if err != nil {
				return transaction, err
			}
		}

		transaction.Amount = totalAmount
		transaction.UserID = input.User.ID
		transaction.Transactions = input.Transactions

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
	input = strings.Trim(input, "{}")     // Remove outer curly braces
	pairs := strings.Split(input, "}, {") // Split into individual pairs

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
		if product.ID == 0 {
			return productQuantities, errors.New("product id is not available")
		}

		if proudctQuantity.Quantity > product.Quantity {
			return productQuantities, errors.New("product quantity " + product.Name + " is not enough")
		}
	}
	return productQuantities, nil
}
