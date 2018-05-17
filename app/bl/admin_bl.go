package bl

import (
	"crypto/sha256"
	"encoding/base64"
	"mcdserver/app/models"
	"mcdserver/app/dao"
	"unicode"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"mcdserver/app"
)

type AdminBl struct {
	email string
	login string
	password string
	cryptedPassword string
}

func NewAdminBl(email string, login string, password string) AdminBl {
	adminBl := AdminBl {email: email, login: login, password: password}
	adminBl.cryptPassword()
	return adminBl
}

func NewAdminBlEmailOrLogin(emailOrLogin string, password string) AdminBl {
	var adminBl AdminBl

	fiveOrMore, at, dot := verifyEmail(emailOrLogin)
	if fiveOrMore && at && dot {
		adminBl = AdminBl {email: emailOrLogin, login: "", password: password}
	} else {
		adminBl = AdminBl {email: "", login: emailOrLogin, password: password}
	}

	adminBl.cryptPassword()
	return adminBl
}

func (adminBl *AdminBl) cryptPassword() {
	hasher := sha256.New()
	hasher.Write([]byte(adminBl.password))
	adminBl.cryptedPassword = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (adminBl AdminBl) getAdminFromDB() (models.Admin, error) {
	adminDao := new(dao.AdminDao)

	loginOrEmail := adminBl.email
	if loginOrEmail == "" {
		loginOrEmail = adminBl.login
	}

	admin, err := adminDao.GetAdminByEmailOrLogin(loginOrEmail)

	return admin, err
}

func (adminBl AdminBl) insertAdminInDB() error {
	adminDao := new(dao.AdminDao)

	err := adminDao.InsertAdmin(adminBl.email, adminBl.login, adminBl.cryptedPassword)

	return err
}

func (adminBl AdminBl) checkPassword(password string) bool {
	if password == adminBl.cryptedPassword {
		return true
	}
	return false
}

func verifyEmail(email string) (fiveOrMore, at, dot bool) {
	lettersCount := 0
	for _, char := range email {
		switch {
			case char == '@':
				at = true
			case char == '.':
				dot = true
			default:
			//return false, false, false, false
		}
		lettersCount++;
	}
	fiveOrMore = lettersCount >= 5
	return
}

func (adminBl AdminBl) verifyPassword() (sevenOrMore, number, lower, upper, special bool) {
	lettersCount := 0
	for _, char := range adminBl.password {
		switch {
		case unicode.IsNumber(char):
			number = true
		case unicode.IsUpper(char):
			upper = true
		case unicode.IsLower(char):
			lower = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		default:
			//return false, false, false, false
		}
		lettersCount++;
	}
	sevenOrMore = lettersCount >= 7
	return
}

func (adminBl AdminBl) GetToken() (string, error) {
	admin, err := adminBl.getAdminFromDB()
	if err != nil {
		return "", err
	}

	passwordCorrect := adminBl.checkPassword(admin.Password)
	if !passwordCorrect {
		return "", errors.New("password incorrect")
	}

	var emailOrLogin string
	if adminBl.email == "" {
		emailOrLogin = adminBl.login
	}
	if adminBl.login == "" {
		emailOrLogin = adminBl.email
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"admin": true,
		"email": emailOrLogin,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := app.GetSecretKey()
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (adminBl AdminBl) Registration() error {
	sevenOrMore, number, lower, upper, _ := adminBl.verifyPassword()
	if !(sevenOrMore && number && lower && upper) {
		app.Logs.Print(sevenOrMore)
		app.Logs.Print(number)
		app.Logs.Print(lower)
		app.Logs.Print(upper)
		return errors.New("password must contain number, lower, upper and length more or 7")
	}

	if adminBl.email == "" {
		return errors.New("email must'n empty")
	}

	if adminBl.login == "" {
		return errors.New("login must'n empty")
	}

	admin, err := adminBl.getAdminFromDB()
	if err == nil {
		app.Logs.Print("User exist with email : " + admin.Email)
		return errors.New("user such email exist yet")
	}

	err = adminBl.insertAdminInDB()
	if err != nil {
		return err
	}

	return nil
}

func (adminBl AdminBl) GetEmailAndLogin() (email string, login string, err error) {
	admin, err := adminBl.getAdminFromDB()
	if err != nil {
		return "", "", err
	}

	return admin.Email, admin.Login, nil
}