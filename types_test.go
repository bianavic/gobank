package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("a", "b", "a@email.com", "1")
	assert.Nil(t, err)

	fmt.Printf("%+v\n", acc)
}

//type LoginRequest struct {
//	Email    string `json:"email"`
//	Password string `json:"password"`
//}
//
//type TransferRequest struct {
//	ToAccount int `json:"toAccount"`
//	Amount    int `json:"amount"`
//}
//
//type CreateAccountRequest struct {
//	FirstName string `json:"firstName"`
//	LastName  string `json:"lastName"`
//	Email     string `json:"email"`
//	Password  string `json:"password"`
//}
//
//type Account struct {
//	ID                int       `json:"ID"`
//	FirstName         string    `json:"firstName"`
//	LastName          string    `json:"lastName"`
//	Email             string    `json:"email"`
//	EncryptedPassword string    `json:"-"`
//	Number            int64     `json:"number"`
//	Balance           int64     `json:"balance"`
//	CreatedAt         time.Time `json:"createdAt"`
//}
//
//func NewAccount(firstName, lastName, email, password string) (*Account, error) {
//	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		return nil, err
//	}
//
//	return &Account{
//		FirstName:         firstName,
//		LastName:          lastName,
//		Email:             email,
//		EncryptedPassword: string(encpw),
//		Number:            int64(rand.Intn(1000000)),
//		CreatedAt:         time.Now().UTC(),
//	}, nil
//
//}
