package handlers

import (
	"go-learning-project/web/utils"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, 200, "Program Is Running......................")
}
