package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("SignUp page loading")
	logrus.Infoln("Executing auth page")
	err := tml.ExecuteTemplate(w, "signup_page.html", nil)
	if err != nil {
		logrus.Errorf("Cannot execute \"signup_page\" template, due to error: %s", err.Error())
	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("SignIn page loading")
	logrus.Infoln("Executing auth page")
	err := tml.ExecuteTemplate(w, "signin_page.html", nil)
	if err != nil {
		logrus.Errorf("Cannot execute \"signin_page\", due to error: %s", err.Error())
	}
}
