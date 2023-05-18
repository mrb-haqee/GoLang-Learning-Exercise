package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type UserLogin struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StudyData struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	StudyCode string `json:"study_code"`
}

type StudyDataDetail struct {
	Code  string          `json:"code"`
	Name  string          `json:"name"`
	Users []UserStudyData `json:"users"`
}

type UserStudyData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// -------- Weather response ----------

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Humidity int     `json:"Humidity"`
}

type MainWeather struct {
	ID       int       `json:"id"`
	Weather  []Weather `json:"weather"`
	Main     Main      `json:"main"`
	DateTime int64     `json:"dt"`
	Name     string    `json:"name"`
}

// ------------------------------------
