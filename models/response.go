package models

import "time"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

type Meta struct {
	Total  int       `json:"total,omitempty"`
	Page   int       `json:"page,omitempty"`
	Limit  int       `json:"limit,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
}

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Received string `json:"received"`
	Echoed   string `json:"echoed"`
	At       int64  `json:"at"`
}

type PingResponse struct {
	Status   string `json:"status"`
	Pong     bool   `json:"pong"`
	Time     int64  `json:"time"`
	Uptime   string `json:"uptime"`
}

// Helper functions to create responses
func SuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
		Meta: &Meta{
			Timestamp: time.Now().Unix(),
		},
	}
}

func SuccessResponseWithMessage(data interface{}, message string) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta: &Meta{
			Timestamp: time.Now().Unix(),
		},
	}
}

func ErrorResponse(message string) Response {
	return Response{
		Success: false,
		Error: &ErrorInfo{
			Message: message,
		},
		Meta: &Meta{
			Timestamp: time.Now().Unix(),
		},
	}
}

func ErrorResponseWithCode(code, message string) Response {
	return Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
		},
		Meta: &Meta{
			Timestamp: time.Now().Unix(),
		},
	}
}
