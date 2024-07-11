package models

import (
	"errors"
	"fmt"
)

type Account struct {
	Name   string
	Amount int64
}

type AccountDatabase struct {
	data map[string]*Account
}

func NewAccountDatabase() *AccountDatabase {
	return &AccountDatabase{make(map[string]*Account)}
}

func (db *AccountDatabase) CreateAccount(name string) error {
	if len(name) == 0 {
		return errors.New("Account name is empty")
	}
	if _, ok := db.data[name]; ok {
		return errors.New(fmt.Sprintf("Account with name '%s' already exists", name))
	}
	db.data[name] = &Account{name, 0}
	return nil
}

func (db *AccountDatabase) DeleteAccount(name string) error {
	if len(name) == 0 {
		return errors.New("Account name is empty")
	}
	if _, ok := db.data[name]; !ok {
		return errors.New(fmt.Sprintf("Account with name '%s' not found", name))
	}
	delete(db.data, name)
	return nil
}

func (db *AccountDatabase) UpdateAmount(name string, amount int64) error {
	if len(name) == 0 {
		return errors.New("Account name is empty")
	}
	if _, ok := db.data[name]; !ok {
		return errors.New(fmt.Sprintf("Account with name '%s' not found", name))
	}
	db.data[name].Amount += amount
	return nil
}

func (db *AccountDatabase) UpdateName(name string, newName string) error {
	if len(name) == 0 {
		return errors.New("Account name is empty")
	}
	if len(newName) == 0 {
		return errors.New("Account new name is empty")
	}
	if _, ok := db.data[name]; !ok {
		return errors.New(fmt.Sprintf("Account with name '%s' not found", name))
	}
	account := db.data[name]
	account.Name = newName
	delete(db.data, name)
	db.data[newName] = account
	return nil
}

func (db *AccountDatabase) GetAccount(name string) (*Account, error) {
	if len(name) == 0 {
		return nil, errors.New("Account name is empty")
	}
	if _, ok := db.data[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Account with name '%s' not found", name))
	}
	return db.data[name], nil
}
