package response

import (
	"encoding/json"
	"net/http"
)

// StandardResponse 定義與 Java 類似的結構
type StandardResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// JSON 是一個封裝過的 Helper，簡化回傳流程
func JSON(w http.ResponseWriter, statusCode int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := StandardResponse[any]{
		Code:    statusCode,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(resp)
}
