package ec

import (
	"fmt"
)

// SvrErr server error
type SvrErr struct {
	err  error
	code ErrCode
}

// Error ...
func (se *SvrErr) Error() string {
	return se.err.Error()
}

// String ...
func (se *SvrErr) String() string {
	return fmt.Sprintf("server internal error. code[%d], %s", se.code, se.err.Error())
}

// Code ...
func (se *SvrErr) Code() ErrCode {
	return se.code
}

// NewError create a new server error
func NewError(code ErrCode, format string, args ...interface{}) *SvrErr {
	return &SvrErr{
		err:  fmt.Errorf(format, args...),
		code: code,
	}
}

// Error
var (
	ErrSuccess              = NewError(Success, "Success")
	ErrInternal             = NewError(Internal, "internal server error")
	ErrInvalidFormat        = NewError(InvalidFormat, "data format invalid")
	ErrInvalidParam         = NewError(InvalidParam, "parameter is invalid")
	ErrCustomerNameEmpty    = NewError(CustomerNameEmpty, "customer's name is empty")
	ErrCustomerEmailEmpty   = NewError(CustomerEmailEmpty, "customer's email is empty")
	ErrCustomerPhoneEmpty   = NewError(CustomerPhoneEmpty, "customer's phone is empty")
	ErrCustomerAddressEmpty = NewError(CustomerAddressEmpty, "customer's address is empty")
	ErrCustomerDOBInvalid   = NewError(CustomerDOBInvalid, "customer's DOB is invalid")
	ErrUnauthorized         = NewError(Unauthorized, "unauthorized user")
	ErrInvalidToken         = NewError(InvalidToken, "invalid token format")
)
