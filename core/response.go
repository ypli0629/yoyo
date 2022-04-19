package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents processing result
type Response struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SucceedResponse(data interface{}) Response {
	return Response{
		Success: true,
		Code:    "0",
		Message: "",
		Data:    data,
	}
}

func FailedResponse(code string, message string) Response {
	return Response{
		Success: false,
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// OK returns processing result successfully
func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SucceedResponse(data))
}

// Fail returns error code and message
func Fail(c *gin.Context, code string, message string) {
	c.JSON(http.StatusOK, FailedResponse(code, message))
}
