package bl

import (
	"mcdserver/app/dao"
	"crypto/sha256"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"time"
	"errors"
	"mcdserver/app/models"
	"mcdserver/app"
	"unicode"
)

type UserBl struct {
	email string
	password string
	cryptedPassword string
}

func NewUserBl (email string, password string) UserBl {
	userBl := UserBl {email: email, password: password}
	userBl.cryptPassword()
	return userBl
}

func (userBl *UserBl) cryptPassword() {
	hasher := sha256.New()
	hasher.Write([]byte(userBl.password))
	userBl.cryptedPassword = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (userBl UserBl) getUserFromDB() (models.User, error) {
	userDao := new(dao.UserDao)

	user, err := userDao.GetUserByEmail(userBl.email)

	return user, err
}

func (userBl UserBl) insertUserInDB() error {
	userDao := new(dao.UserDao)

	err := userDao.InsertUser(userBl.email, userBl.cryptedPassword)

	return err
}

func (userBl UserBl) checkPassword(password string) bool {
	if password == userBl.cryptedPassword {
		return true
	}
	return false
}

func (userBl UserBl) verifyPassword() (sevenOrMore, number, lower, upper, special bool) {
	lettersCount := 0
	for _, char := range userBl.password {
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

func (userBl UserBl) GetToken() (string, error) {
	user, err := userBl.getUserFromDB()
	if err != nil {
		return "", err
	}

	passwordCorrect := userBl.checkPassword(user.GetPassword())
	if !passwordCorrect {
		return "", errors.New("password incorrect")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims {
		"admin": false,
		"email": userBl.email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := app.GetSecretKey()
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (userBl UserBl) Registration() error {
	sevenOrMore, number, lower, upper, _ := userBl.verifyPassword()
	if !(sevenOrMore && number && lower && upper) {
		app.Logs.Print(sevenOrMore)
		app.Logs.Print(number)
		app.Logs.Print(lower)
		app.Logs.Print(upper)
		return errors.New("password must contain number, lower, upper and length more or 7")
	}

	user, err := userBl.getUserFromDB()
	if err == nil {
		app.Logs.Print("User exist with email : " + user.GetEmail())
		return errors.New("user such email exist yet")
	}

	err = userBl.insertUserInDB()
	if err != nil {
		return err
	}

	return nil
}