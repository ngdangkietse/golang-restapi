package errors

func IsSuccess(code int) bool {
	return code >= 0
}

func IsFailed(code int) bool {
	return !IsSuccess(code)
}
