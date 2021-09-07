package ec

// ErrCode ...
type ErrCode int32

// Error Codes
const (
	Success              ErrCode = 0
	CustomerNameEmpty    ErrCode = 1
	CustomerEmailEmpty   ErrCode = 2
	CustomerPhoneEmpty   ErrCode = 3
	CustomerAddressEmpty ErrCode = 4
	CustomerDOBInvliad   ErrCode = 5
)
