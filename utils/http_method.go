package utils

import "net/http"

func HttpMethod(method string, res http.ResponseWriter, req *http.Request) {

	if req.Method != method {
		HandlerMessage := []byte(`
			{
				"success": false,
				"message": "Check your HTTP method: Invalid HTTP method executed",
			}
		`)

		ReturnJsonResponse(res, http.StatusMethodNotAllowed, HandlerMessage)
		return
	}
}
