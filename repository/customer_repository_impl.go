package repository

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/vnnyx/golang-api/entity"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repo *CustomerRepositoryImpl) CreateCustomer(ctx context.Context, tx *sqlx.Tx, customer entity.Customer) (entity.Customer, error) {
	SQL := "INSERT into customers(id, username, email, password, gender) VALUES (?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, SQL, customer.Id, customer.Username, customer.Email, customer.Password, customer.Gender)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (repo *CustomerRepositoryImpl) GetAllCustomer(ctx context.Context, tx *sqlx.Tx) ([]entity.Customer, error) {
	SQL := "SELECT * FROM customers"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return []entity.Customer{}, err
	}
	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		customer := entity.Customer{}
		err = rows.Scan(&customer.Id, &customer.Username, &customer.Email, &customer.Password, &customer.Gender, &customer.CreatedAt)
		if err != nil {
			return []entity.Customer{}, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (repo *CustomerRepositoryImpl) GetCustomerById(ctx context.Context, tx *sqlx.Tx, customerId uint32) (entity.Customer, error) {
	SQL := "SELECT * FROM customers WHERE id=?"
	rows, err := tx.QueryContext(ctx, SQL, customerId)
	if err != nil {
		return entity.Customer{}, err
	}
	defer rows.Close()

	customer := entity.Customer{}
	if rows.Next() {
		err = rows.Scan(&customer.Id, &customer.Username, &customer.Email, &customer.Password, &customer.Gender, &customer.CreatedAt)
		if err != nil {
			return customer, err
		}
		return customer, nil
	} else {
		return customer, errors.New("customer not found")
	}
}

func (repo *CustomerRepositoryImpl) UpdateCustomer(ctx context.Context, tx *sqlx.Tx, customer entity.Customer) (entity.Customer, error) {
	SQL := "UPDATE customers SET username=?, email=?, password=?, gender=? WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, customer.Username, customer.Email, customer.Password, customer.Gender, customer.Id)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (repo *CustomerRepositoryImpl) DeleteCustomer(ctx context.Context, tx *sqlx.Tx, customerId uint32) error {
	SQL := "DELETE FROM customers WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, customerId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CustomerRepositoryImpl) GetUserByUsername(ctx context.Context, tx *sqlx.Tx, username string) (entity.Customer, error) {
	SQL := "SELECT * FROM customers WHERE username=?"
	rows, err := tx.QueryContext(ctx, SQL, username)
	if err != nil {
		return entity.Customer{}, err
	}
	defer rows.Close()

	customer := entity.Customer{}
	if rows.Next() {
		err = rows.Scan(&customer.Id, &customer.Username, &customer.Email, &customer.Password, &customer.Gender, &customer.CreatedAt)
		if err != nil {
			return customer, err
		}
		return customer, nil
	}
	return customer, errors.New("customer not found")

}

func (repo *CustomerRepositoryImpl) DeleteAllCustomer(ctx context.Context, tx *sqlx.Tx) error {
	SQL := "DELETE FROM customers"
	_, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return err
	}
	return nil
}
