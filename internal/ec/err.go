package ec

import "errors"

// Error
var (
	ErrSuccess              = errors.New("success")
	ErrCustomerNameEmpty    = errors.New("customer's name is empty")
	ErrCustomerEmailEmpty   = errors.New("customer's email is empty")
	ErrCustomerPhoneEmpty   = errors.New("customer's phone is empty")
	ErrCustomerAddressEmpty = errors.New("customer's address is empty")
	ErrCustomerDOBInvalid   = errors.New("customer's DOB is invalid")
)
