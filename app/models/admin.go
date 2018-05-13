package models

type Admin struct {
	Email    string `json:"email"`
	Login	 string `json:"login"`
	Password string `json:"password"`
}

func NewAdmin(email string, login string, password string) Admin {
	return Admin{Email: email, Login: login, Password: password}
}