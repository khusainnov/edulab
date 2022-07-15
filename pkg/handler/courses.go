package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) Courses(w http.ResponseWriter, r *http.Request) {
	logrus.Infoln("Executing template for courses page")
	err := tml.ExecuteTemplate(w, "courses_page.html", nil)
	if err != nil {
		logrus.Errorf("Cannot execute \"courses_page\", due to error: %s", err.Error())
	}
}