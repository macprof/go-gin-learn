package utils

import "go-gin-workshop/types"

func BuildResponseSuccess(status int, message string, data any) types.Response {
	res := types.Response{
		Status:  status,
		Result:  true,
		Message: message,
		Data:    data,
	}
	return res
}

func BuildResponseFailed(status int, message string, err string, data any) types.Response {
	res := types.Response{
		Status:  status,
		Result:  false,
		Message: message,
		Error:   err,
		Data:    data,
	}
	return res
}
