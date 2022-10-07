package controllers

import (
	"net/http"

	"main.go/utils"
)

func InitHandler(res http.ResponseWriter, req *http.Request) {

	utils.HttpMethod("GET", res, req)

	HandlerMessage := []byte(`
		{
			"success": true,
			"message": "The server is running properly"
		}
	`)

	utils.ReturnJsonResponse(res, http.StatusOK, HandlerMessage)
}
