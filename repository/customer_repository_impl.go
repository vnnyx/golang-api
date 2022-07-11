package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-simple-api/entity"
	"golang-simple-api/exception"
	"golang-simple-api/helper"
)

type CustomerRepositoryImpl struct {
	DB *sql.DB
}

func NewCustomerRepository(DB *sql.DB) CustomerRepository {
	return &CustomerRepositoryImpl{DB: DB}
}

func (repo *CustomerRepositoryImpl) CreateCustomer(ctx context.Context, customer entity.Customer) entity.Customer {
	tx, err := repo.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)
	SQL := "INSERT into customers(id, username, email, password, gender) VALUES (?,?,?,?,?)"
	_, err = tx.ExecContext(ctx, SQL, customer.Id, customer.Username, customer.Email, customer.Password, customer.Gender)
	exception.PanicIfNeeded(err)

	return customer
}

func (repo *CustomerRepositoryImpl) GetAllCustomer(ctx context.Context) []entity.Customer {
	tx, err := repo.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)
	SQL := "SELECT * FROM customers"
	rows, err := tx.QueryContext(ctx, SQL)
	exception.PanicIfNeeded(err)

	var customers []entity.Customer
	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Id, &customer.Username, &customer.Email, &customer.Password, &customer.Gender, &customer.CreatedAt)
		exception.PanicIfNeeded(err)
		customers = append(customers, customer)
	}
	return customers
}

func (repo *CustomerRepositoryImpl) GetCustomerById(ctx context.Context, customerId uint32) (entity.Customer, error) {
	tx, err := repo.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)
	SQL := "SELECT * FROM customers WHERE id=?"
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	exception.PanicIfNeeded(err)
	defer rows.Close()

	customer := entity.Customer{}
	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Username, &customer.Email, &customer.Password, &customer.Gender, &customer.CreatedAt)
		exception.PanicIfNeeded(err)
		return customer, nil
	} else {
		return customer, errors.New("customer not found")
	}
}

func (repo *CustomerRepositoryImpl) UpdateCustomer(ctx context.Context, customer entity.Customer) entity.Customer {
	tx, err := repo.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)
	SQL := "UPDATE customers SET username=?, email=?, password=?, gender=? WHERE id=?"
	_, err = tx.ExecContext(ctx, SQL, customer.Username, customer.Email, customer.Password, customer.Gender, customer.Id)
	exception.PanicIfNeeded(err)

	return customer
}

func (repo *CustomerRepositoryImpl) DeleteCustomer(ctx context.Context, customerId uint32) {
	tx, err := repo.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)
	SQL := "DELETE FROM customers WHERE id=?"
	_, err = tx.ExecContext(ctx, SQL, customerId)
	exception.PanicIfNeeded(err)
}

func (repo *CustomerRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (entity.Customer, error) {
	tx, err := repo.DB.Begin()
	exception.PanicIfNeeded(err)
	defer helper.CommitOrRollback(tx)
	SQL := "SELECT * FROM customers WHERE username=?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	exception.PanicIfNeeded(err)
	defer rows.Close()

	customer := entity.Customer{}
	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Username, &customer.Email, &customer.Password, &customer.Gender, &customer.CreatedAt)
		exception.PanicIfNeeded(err)
		return customer, nil
	} else {
		return customer, errors.New("customer not found")
	}
}
