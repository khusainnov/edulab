package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile"))
	userId := r.Context().Value("userId")
	err := json.NewEncoder(w).Encode(&map[string]interface{}{
		"status": http.StatusOK,
		"userId": userId,
	})

	if err != nil {
		logrus.Errorf("Cannot encode data: %s", err.Error())
		json.NewEncoder(w).Encode(&map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
	}
}
