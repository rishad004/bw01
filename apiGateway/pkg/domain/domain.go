package domain

type H map[string]any

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Response struct {
	Message    any `json:"message"`
	StatusCode int `json:"code"`
}

type Request struct {
	Method   string `json:"method"`
	WaitTime int    `json:"wait"`
}
