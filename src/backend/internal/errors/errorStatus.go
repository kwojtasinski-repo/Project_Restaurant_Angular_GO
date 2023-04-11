package errors

type ErrorStatus struct {
	Status  int
	Message string
}

func BadRequest(message string) *ErrorStatus {
	return &ErrorStatus{
		Status:  400,
		Message: message,
	}
}

func Forbidden() *ErrorStatus {
	return &ErrorStatus{
		Status: 403,
	}
}

func NotFound() *ErrorStatus {
	return &ErrorStatus{
		Status: 404,
	}
}

func NotFoundWithMessage(message string) *ErrorStatus {
	return &ErrorStatus{
		Status:  404,
		Message: message,
	}
}

func InternalError(message string) *ErrorStatus {
	return &ErrorStatus{
		Status:  500,
		Message: message,
	}
}
