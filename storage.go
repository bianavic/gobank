package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "gobank"
)

type Storage interface {
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	GetAccountByEmail(email string) (*Account, error)
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
}

type PostgreSQLStore struct {
	db *sql.DB
}

func NewPostgreSQLStore() (*PostgreSQLStore, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	//defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("The database is connected")

	return &PostgreSQLStore{
		db: db,
	}, nil

}

func (s *PostgreSQLStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgreSQLStore) createAccountTable() error {
	query := `create table if not exists account (
	id serial primary key,
	first_name varchar(50),
	last_name varchar(50),
    number serial,
	balance serial,
	created_at timestamp,
	encrypted_password varchar(100),
    email varchar(255))`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgreSQLStore) CreateAccount(acc *Account) error {
	query := `insert into account
	(first_name, last_name, number, balance, created_at, encrypted_password, email)
	values ($1, $2, $3, $4, $5, $6, $7)`

	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
		acc.EncryptedPassword,
		acc.Email)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgreSQLStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, fmt.Errorf("error fetching accounts: %w", err)
	}

	// create an empty slice to store actual accounts
	accounts := []*Account{}

	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		// append the scanned account to the slice
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (s *PostgreSQLStore) GetAccountByID(id int) (*Account, error) {

	query := "SELECT * FROM account WHERE id = $1"
	rows, err := s.db.Query(query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("account %d not found", id)
		}
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgreSQLStore) GetAccountByEmail(email string) (*Account, error) {

	query := "SELECT * FROM account WHERE email = $1"
	rows, err := s.db.Query(query, email)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		return account, nil
	}

	return nil, fmt.Errorf("email account %s not found", email)
}

func (s *PostgreSQLStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgreSQLStore) DeleteAccount(id int) error {
	query := "DELETE FROM account WHERE id = $1"
	_, err := s.db.Query(query, id)

	fmt.Printf("query %s and id %d", query, id)

	return err
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := &Account{}
	var createdAt time.Time
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
		&account.EncryptedPassword,
		&account.Email,
	)
	account.CreatedAt = createdAt

	return account, err
}
