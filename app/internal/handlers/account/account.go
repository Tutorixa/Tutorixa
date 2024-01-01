/*
Contains the base handlers for creating and authenticating accounts.
*/
package accounthandler

import (
	"app/internal/storage"
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"unicode/utf8"
)

type AccountHandler struct {
	store storage.Storage
}

func New(store storage.Storage) *AccountHandler {
	return &AccountHandler{
		store: store,
	}
}

func (handle *AccountHandler) CreateAccount(c echo.Context) error {
	//Get and check the email to see if the account exists
	email := c.FormValue("email")
	account, err := handle.store.GetAccount(email)
	if err != nil{
		fmt.Println("Account already exists")
		return err
	}
	//Check and hash the password
	password := c.FormValue("password")
	if utf8.RuneCountInString(password) < 8{
		fmt.Println("Password too short")
		return nil
	}
	if utf8.RuneCountInString(password) > 72{
		fmt.Println("Password too long")
		return nil
	}
	//Create a new account
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	handle.store.CreateAccount(c.FormValue("fname"), c.FormValue("lname"), email, hash)
	return nil
}

func (handle *AccountHandler) Verification(c echo.Context) error {
	// Pull the data from the request
	email := c.FormValue("email")
	account, err := handle.store.GetAccount(email)
	if err != nil {
		fmt.Println("No account found")
		return err
	}

	fmt.Println(account)
	fmt.Println(handle.store)
	
	// Check if the account password matches the hashed password
	
	password := c.FormValue("password")
	hash := handle.store.GetPassword(email)
	matched:= bcrypt.CompareHashAndPassword(hash, []byte(password))
	if matched != nil {
		err := c.Redirect(http.StatusOK, "<URL>")
		return err
	}

	htmlstr := "Incorrect Username or Password"
	tmpl, err := template.New("t").Parse(htmlstr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return tmpl.Execute(c.Response().Writer, nil)
}