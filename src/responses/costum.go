package responses

import "snapp/validations"

type Response struct {
	Status          bool                          `json:"Status"`
	StatusCode      int                           `json:"StatusCode"`
	Result          any                           `json:"Result,omitempty"`
	Error           string                        `json:"Error,omitempty"`
	ValidationError *[]validations.ValidationError `json:"ValidationError,omitempty"`
}

func GenerateNormalResponse(status bool, statuscode int, result any) *Response {
	return &Response{
		Status: status,
		StatusCode: statuscode,
		Result: result,
	}
}

func GenerateResponseWithError(status bool, statuscode int, err error) *Response {
	return &Response{
		Status: status,
		StatusCode: statuscode,
		Error: err.Error(),
	}
}

func GenerateResponseWithValidationError(status bool, statuscode int, err error) *Response {
	return &Response{
		Status: status,
		StatusCode: statuscode,
		ValidationError: validations.GenerateValidationError(err),
	}
}