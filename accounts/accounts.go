package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates new account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit adds amount to balance
func (a *Account) Deposit(amount int) bool {
	a.balance += amount
	return true
}

// Withdraw from account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of their account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner prints owner
func (a Account) Owner() string {
	return a.owner
}

// Balance prints balance
func (a Account) Balance() int {
	return a.balance
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
