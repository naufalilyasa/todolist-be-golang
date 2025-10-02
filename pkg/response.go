package pkg

import (
	"encoding/json"
	"net/http"
)

type Pagination struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	Total_pages int   `json:"total_pages"`
}

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

func JSON(w http.ResponseWriter, status int, payload Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func JSONSuccess(w http.ResponseWriter, status int, message string, data any, pagination *Pagination) {
	resp := Response{
		Status:     "success",
		StatusCode: status,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
	JSON(w, status, resp)
}

func JSONError(w http.ResponseWriter, status int, message string) {
	resp := Response{
		Status:     "error",
		StatusCode: status,
		Message:    message,
	}
	JSON(w, status, resp)
}
