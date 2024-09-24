package utils

import "net/http"

func SendData(w http.ResponseWriter, status int, data interface{}) {
	SendJson(w, status, map[string]any{
		"status":  true,
		"message": "Success",
		"data":    data,
	})
}
