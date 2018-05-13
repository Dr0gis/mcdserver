package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

func NewUser(email string, password string, name string, surname string) User {
	return User{Email: email, Password: password, Name: name, Surname: surname}
}

func (user User) GetEmail() string {
	return user.Email
}

func (user *User) SetEmail(email string) {
	user.Email = email
}

func (user User) GetPassword() string {
	return user.Password
}

func (user *User) SetPassword(password string) {
	user.Password = password
}

func (user User) GetName() string {
	return user.Name
}

func (user *User) SetName(name string) {
	user.Name = name
}

func (user User) GetSurname() string {
	return user.Surname
}

func (user *User) SetSurname(surname string) {
	user.Surname = surname
}