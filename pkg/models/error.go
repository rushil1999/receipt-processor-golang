package models


type CustomError struct {
	Message string
	DebugMessage string
	HttpCode int
}

func (e CustomError) Error() string { // Implementing the error interface
	return e.Message
}

