package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	// var ListBarang database.Database
	ListBarang, _ := s.database.GetCartItems()
	Barang, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	}
	if quantity >= 1 {
		BarangTambah := entity.CartItem{ProductName: Barang.Name, Price: Barang.Price, Quantity: quantity}
		ListBarang = append(ListBarang, BarangTambah)
		s.database.SaveCartItems(ListBarang)
		return nil
	}
	return errors.New("invalid quantity") // TODO: replace this
}

func (s *Service) RemoveCart(productName string) error {
	ListBarang, _ := s.database.GetCartItems()
	var newListBarang []entity.CartItem
	check := 0

	for i, data := range ListBarang {
		if data.ProductName == productName {
			check = i
		}
	}
	if check != 0 {
		for i, data := range ListBarang {
			if i == check {
				continue
			}
			newListBarang = append(newListBarang, data)
		}
		s.database.SaveCartItems(newListBarang)
		return nil // TODO: replace this
	} else {
		return errors.New("product not found")
	}
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (s *Service) ResetCart() error {
	s.database.SaveCartItems([]entity.CartItem{})
	return nil // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	return s.database.GetProductData(), nil // TODO: replace this
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	ListBarang, _ := s.database.GetCartItems()
	TotalPrice := 0
	for _, data := range ListBarang {
		TotalPrice += data.Price * data.Quantity

	}
	if money-TotalPrice<=0{
		return entity.PaymentInformation{},errors.New("money is not enough")
	}
	s.ResetCart()
	return entity.PaymentInformation{
		ProductList: ListBarang,
		TotalPrice:  TotalPrice,
		MoneyPaid:   money,
		Change:      money - TotalPrice,
	}, nil // TODO: replace this
}
