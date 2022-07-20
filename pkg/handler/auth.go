package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/khusainnov/edulab/internal/entity/user"
	"github.com/sirupsen/logrus"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input user.User
	logrus.Infoln("SignUp page loading")

	if r.URL.Path != "/auth/sign-up" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	switch r.Method {
	case "GET":
		{
			logrus.Infoln("Executing auth page")
			err := tml.ExecuteTemplate(w, "signup_page.html", nil)
			if err != nil {
				logrus.Errorf("Cannot execute \"signup_page\" template, due to error: %s", err.Error())
			}
		}

	case "POST":
		{
			logrus.Infoln("SignUp method POST checking")
			if err := r.ParseForm(); err != nil {
				logrus.Errorf("Cannot parse form in signup_page, due to error: %s", err.Error())
				fmt.Fprintf(w, "%s", err.Error())
			}
			input = user.User{
				Name:     r.FormValue("fname"),
				Surname:  r.FormValue("fsurname"),
				Username: r.FormValue("fusername"),
				Email:    r.FormValue("femail"),
				Password: r.FormValue("fpassword"),
			}

			fmt.Println(input.Password)
			fmt.Printf("Username: %v\n", input.Username)

			id, err := h.services.CreateUser(input)
			if err != nil {
				logrus.Errorf("code: %d, hadnler/auth - SignUp: %s", http.StatusInternalServerError, err.Error())
			}
			fmt.Println(id)
			err = tml.ExecuteTemplate(w, "signup_thanks.html", &input)
			if err != nil {
				logrus.Errorf("Cannot execute \"signup_thanks\" template, due to error: %s", err.Error())
			}
			/*switch err {
			case err:
				if err != nil {
					logrus.Errorf("code: %d, hadnler/auth - SignUp: %s", http.StatusInternalServerError, err.Error())
				}
			default:
				fmt.Println(id)
				err = tml.ExecuteTemplate(w, "signup_thanks.html", &input)
				if err != nil {
					logrus.Errorf("Cannot execute \"signup_thanks\" template, due to error: %s", err.Error())
				}
			}*/
		}
	}
}

type SignInInput struct {
	Login    string `json:"username,password,login,omitempty" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	var input SignInInput

	logrus.Infoln("SignIn page loading")
	logrus.Infoln("Executing auth page")
	err := tml.ExecuteTemplate(w, "signin_page.html", nil)
	if err != nil {
		logrus.Errorf("Cannot execute \"signin_page\", due to error: %s", err.Error())
	}

	input = SignInInput{
		Login:    r.FormValue("flogin"),
		Password: r.FormValue("fpassword"),
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		logrus.Errorf("code: %d, hadnler/auth - SignUp: %s", http.StatusInternalServerError, err.Error())
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 12),
	}
	r.AddCookie(cookie)
	fmt.Println(token)
	http.Redirect(w, r, "/", http.StatusOK)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("Action Logout")
	http.Redirect(w, r, "/", http.StatusOK)
}
