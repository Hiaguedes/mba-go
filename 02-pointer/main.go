package main

import "fmt"

type Client struct {
	id   uint
	name string
	Account
}

type Account struct {
	balance float32
	active  bool
}

func NewAccount() *Account {
	return &Account{balance: 0.0, active: true}
}

func (account *Account) Transfer(value float32) {
	account.balance += value
}

func (account *Account) DeactivateAccount() {
	account.active = false
}

func main() {
	client := Client{
		id:   1,
		name: "Hiago",
	}

	account := NewAccount()

	client.Account = *account

	client.Account.Transfer(540.0)

	fmt.Printf("Client %+v", client)
}