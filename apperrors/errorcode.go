package apperrors

type ErrCode string

const (
	Unkown ErrCode = "U000"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed ErrCode = "S002"
	NAData ErrCode = "S003"
	NoTartegData ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"
)
