package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) GreetingPage(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("Executing template for greeting page")
	err := tml.ExecuteTemplate(w, "greeting_page.html", nil)
	if err != nil {
		logrus.Errorf("Cannot execute \"greeting_page\", due to error: %s", err.Error())
	}
}