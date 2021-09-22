package ec

// ErrCode ...
type ErrCode int16

// Error Codes
const (
	Success       ErrCode = 0
	Internal      ErrCode = 1
	Unauthorized  ErrCode = 2
	InvalidToken  ErrCode = 3
	InvalidFormat ErrCode = 4

	CustomerNameEmpty    ErrCode = 11
	CustomerEmailEmpty   ErrCode = 12
	CustomerPhoneEmpty   ErrCode = 13
	CustomerAddressEmpty ErrCode = 14
	CustomerDOBInvalid   ErrCode = 15
)
