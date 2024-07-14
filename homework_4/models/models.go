package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type Account struct {
	Name   string
	Amount int64
}

type AccountDatabase struct {
	users *sql.DB
}

func ConnectAccountDatabase(db *sql.DB) *AccountDatabase {
	return &AccountDatabase{db}
}

func (db *AccountDatabase) IsAccountExist(name string) (bool, error) {
	if len(name) == 0 {
		return false, errors.New("account name is empty")
	}

	res, err := db.users.Query("SELECT EXISTS(SELECT name FROM accounts WHERE name = $1)", name)

	defer func() {
		_ = res.Close()
	}()

	if err != nil {
		return false, err
	}
	var exist bool
	res.Next()
	err = res.Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (db *AccountDatabase) CreateAccount(name string) error {
	if len(name) == 0 {
		return errors.New("account name is empty")
	}

	exist, err := db.IsAccountExist(name)
	if err != nil {
		return err
	}
	if exist {
		return errors.New(fmt.Sprintf("account with name '%s' already exists", name))
	}

	_, err = db.users.Exec("INSERT INTO accounts(name, balance) VALUES ($1, $2)", name, 0)
	if err != nil {
		return err
	}
	return nil
}

func (db *AccountDatabase) DeleteAccount(name string) error {
	if len(name) == 0 {
		return errors.New("account name is empty")
	}

	exist, err := db.IsAccountExist(name)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New(fmt.Sprintf("account with name '%s' not found", name))
	}

	_, err = db.users.Exec("DELETE FROM accounts WHERE name = $1", name)
	if err != nil {
		return err
	}
	return nil
}

func (db *AccountDatabase) UpdateAmount(name string, amount int64) error {
	if len(name) == 0 {
		return errors.New("account name is empty")
	}

	exist, err := db.IsAccountExist(name)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New(fmt.Sprintf("account with name '%s' not found", name))
	}

	row := db.users.QueryRow("SELECT balance FROM accounts WHERE name = $1", name)
	newBalance := int64(0)
	err = row.Scan(&newBalance)
	if err != nil {
		return err
	}
	newBalance += amount

	_, err = db.users.Exec("UPDATE accounts SET balance = $1 WHERE name = $2", newBalance, name)
	return nil
}

func (db *AccountDatabase) UpdateName(name string, newName string) error {
	if len(name) == 0 {
		return errors.New("account name is empty")
	}
	if len(newName) == 0 {
		return errors.New("account new name is empty")
	}

	existName, err := db.IsAccountExist(name)
	if err != nil {
		return err
	}
	if !existName {
		return errors.New(fmt.Sprintf("account with name '%s' not found", name))
	}

	existNewName, err := db.IsAccountExist(newName)
	if err != nil {
		return err
	}
	if existNewName {
		return errors.New(fmt.Sprintf("account with name '%s' already exists", newName))
	}

	_, err = db.users.Exec("UPDATE accounts SET name = $1 WHERE name = $2", newName, name)
	if err != nil {
		return err
	}
	return nil
}

func (db *AccountDatabase) GetAccount(name string) (*Account, error) {
	if len(name) == 0 {
		return nil, errors.New("account name is empty")
	}

	exist, err := db.IsAccountExist(name)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New(fmt.Sprintf("account with name '%s' not found", name))
	}

	row := db.users.QueryRow("SELECT name, balance FROM accounts WHERE name = $1", name)
	var acc Account
	err = row.Scan(&acc.Name, &acc.Amount)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}
