package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-simple-api/entity"
	"golang-simple-api/exception"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repo *CustomerRepositoryImpl) CreateCustomer(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer {
	SQL := "INSERT into customers(username, email, password, gender) VALUES (?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, customer.Username, customer.Email, customer.Password, customer.Gender)
	exception.PanicIfNeeded(err)

	id, _ := result.LastInsertId()
	customer.Id = int(id)

	return customer
}

func (repo *CustomerRepositoryImpl) GetAllCustomer(ctx context.Context, tx *sql.Tx) []entity.Customer {
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

func (repo *CustomerRepositoryImpl) GetCustomerById(ctx context.Context, tx *sql.Tx, customerId int) (entity.Customer, error) {
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

func (repo *CustomerRepositoryImpl) UpdateCustomer(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer {
	SQL := "UPDATE customers SET username=?, email=?, password=?, gender=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, customer.Username, customer.Email, customer.Password, customer.Gender, customer.Id)
	exception.PanicIfNeeded(err)

	return customer
}

func (repo *CustomerRepositoryImpl) DeleteCustomer(ctx context.Context, tx *sql.Tx, customerId int) {
	SQL := "DELETE FROM customers WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, customerId)
	exception.PanicIfNeeded(err)
}
