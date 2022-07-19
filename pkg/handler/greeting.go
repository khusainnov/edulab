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

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	/*logrus.Infoln("Executing template for About page")
	err := tml.ExecuteTemplate(w, "about_page.hmlt", nil)
	if err != nil {
		logrus.Errorf("Cannot execute \"about_page\", due to error: %s", err.Error())
	}*/
	w.Write([]byte("About page"))
}
