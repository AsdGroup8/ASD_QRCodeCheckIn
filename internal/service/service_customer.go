package service

import (
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/db"
	"github.com/AsdGroup8/ASD_QRCodeCheckIn/internal/model"
)

// CreateCustomer ...
func CreateCustomer(customer *model.Customer) error {
	return dbmgr.Create(customer).Error
}

// DeleteCustomer ...
func DeleteCustomer(customerID uint) error {
	return dbmgr.Delete(&model.Customer{}, customerID).Error
}

// UpdateCustomer ...
func UpdateCustomer(customer *model.Customer) error {
	return dbmgr.Model(&customer).Updates(customer).Error
}

// FindCustomerByID find one customer record according to id
func FindCustomerByID(customerID uint) (*model.Customer, error) {
	customer := &model.Customer{}
	if err := dbmgr.First(customer, customerID).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

// SelectCustomerS find all customers with model.Customer{} conditions
func SelectCustomerS(customer model.Customer) ([]*model.Customer, error) {
	customers := make([]*model.Customer, 0, 4)
	if err := dbmgr.Where(&customer).Find(customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

// SelectCustomerM find all customers with map conditions
func SelectCustomerM(condition db.M) ([]*model.Customer, error) {
	customers := make([]*model.Customer, 0, 4)
	if err := dbmgr.Where(condition).Find(customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
