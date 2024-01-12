package structures

import "strings"

type UserCreateRequest struct {
	FirstName string // не может быть пустым; не может содержать пробелы
	Age       int    // не может быть 0 или отрицательным; не может быть больше 150
}

func Validate(req UserCreateRequest) string {

	var errorMessage = "invalid request"

	if strings.Contains(req.FirstName, " ") || req.FirstName == "" {
		return errorMessage
	}

	if req.Age <= 0 || req.Age > 150 {
		return errorMessage
	}

	return ""
}
