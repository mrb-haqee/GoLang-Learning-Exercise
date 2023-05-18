package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type StudyData struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	StudyCode string `json:"study_code"`
}
