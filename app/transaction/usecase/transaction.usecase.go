package usecase

import (
	"errors"
	"math/rand"
	"time"

	discountMysql "github.com/muhammadaskar/casheer-be/app/discount/repository/mysql"
	memberMysql "github.com/muhammadaskar/casheer-be/app/member/repository/mysql"
	productMysql "github.com/muhammadaskar/casheer-be/app/product/repository/mysql"
	"github.com/muhammadaskar/casheer-be/app/transaction"
	transacationMysql "github.com/muhammadaskar/casheer-be/app/transaction/repository/mysql"
	"github.com/muhammadaskar/casheer-be/domains"
)

type TransacationUseCase interface {
	Create(input transaction.CreateInput) (domains.Transacation, error)
}

type usecase struct {
	transactionRepo transacationMysql.Repository
	memberRepo      memberMysql.Repository
	productRepo     productMysql.Repository
	discountRepo    discountMysql.Repository
}

func NewUseCase(transactionRepo transacationMysql.Repository,
	memberRepo memberMysql.Repository,
	productRepo productMysql.Repository,
	discountRepo discountMysql.Repository) *usecase {
	return &usecase{transactionRepo, memberRepo, productRepo, discountRepo}
}

func (u *usecase) Create(input transaction.CreateInput) (domains.Transacation, error) {
	transaction := domains.Transacation{}
	memberCode := input.MemberCode
	// kondisi jika member, maka akan mendapatkan discount
	if memberCode != "" {
		isMember, err := u.isAvailableMember(memberCode)
		if err != nil {
			return transaction, err
		}

		if !isMember {
			return transaction, errors.New("Member code is not available")
		}

		discount, err := u.getDiscount()
		if err != nil {
			return transaction, err
		}

		transaction.MemberCode = input.MemberCode
		transaction.TransactionCode = generateTransactionCode()

		product, err := u.productRepo.FindByProductID(input.ProductID)
		if err != nil {
			return transaction, err
		}

		if product.ID == 0 {
			return transaction, errors.New("product id is not available")
		}

		if input.Quantity > product.Quantity {
			return transaction, errors.New("product quantity is not enough")
		}

		transaction.ProductID = input.ProductID
		transaction.Quantity = input.Quantity

		result := calculateTotalPrice(float64(product.Price), float64(discount), input.Quantity)
		transaction.Amount = int(result)
		transaction.UserID = input.User.ID

		newTransaction, err := u.transactionRepo.Create(transaction)
		if err != nil {
			return newTransaction, err
		}

		productQuantity := product.Quantity - input.Quantity
		product.Quantity = productQuantity

		_, err = u.productRepo.Update(product)
		if err != nil {
			return newTransaction, err
		}

		return newTransaction, nil
	} else {
		// kondisi jika bukan member, maka tidak akan mendapatkan discount
		transaction.MemberCode = input.MemberCode
		transaction.TransactionCode = generateTransactionCode()

		product, err := u.productRepo.FindByProductID(input.ProductID)
		if err != nil {
			return transaction, err
		}

		if product.ID == 0 {
			return transaction, errors.New("product id is not available")
		}

		if input.Quantity > product.Quantity {
			return transaction, errors.New("product quantity is not enough")
		}

		transaction.ProductID = input.ProductID
		transaction.Quantity = input.Quantity

		result := (product.Price * input.Quantity)
		transaction.Amount = result
		transaction.UserID = input.User.ID

		newTransaction, err := u.transactionRepo.Create(transaction)
		if err != nil {
			return newTransaction, err
		}

		productQuantity := product.Quantity - input.Quantity
		product.Quantity = productQuantity

		_, err = u.productRepo.Update(product)
		if err != nil {
			return newTransaction, err
		}

		return newTransaction, nil
	}
}

func (u *usecase) isAvailableMember(memberCode string) (bool, error) {
	member, err := u.memberRepo.FindByMemberCode(memberCode)
	if err != nil {
		return false, err
	}

	if member.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (u *usecase) getDiscount() (int, error) {
	discount, err := u.discountRepo.FindById()
	if err != nil {
		return discount.Discount, err
	}

	return discount.Discount, nil
}

func (u *usecase) getProduct(id int) (int, error) {
	product, err := u.productRepo.FindByProductID(id)
	if err != nil {
		return 0, err
	}

	if product.ID != 0 {
		return product.ID, nil
	}

	return 0, nil
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
