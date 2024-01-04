package jsonSuccessResponse

import (
	"encoding/json"
	projectTypes "main/src/ProjectTypes"
	"net/http"
)

func JsonSuccessResponse(w http.ResponseWriter) {
	response := projectTypes.ApiResponse{Status: "success"}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
