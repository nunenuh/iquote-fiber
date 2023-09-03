package webutils

import "github.com/gofiber/fiber/v2"

type HTTPMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Meta struct {
	Path        string     `json:"path"`
	StatusCode  int        `json:"statusCode"`
	Status      string     `json:"status"`
	Message     string     `json:"message"`
	Timestamp   string     `json:"timestamp"`
	Error       *MetaError `json:"error,omitempty"`
	RequestID   string     `json:"requestId"`
	TimeElapsed string     `json:"timeElapsed"`
}

type MetaError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type HTTPResponse struct {
	Message    HTTPMessage `json:"message"`
	Meta       Meta        `json:"metadata"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// func NewSuccessResponse(data interface{}, pagination *Pagination, ctx ContextAdapter) *HTTPResponse {
// 	// Populate message and meta based on your logic
// 	message := HTTPMessage{
// 		Title: "Success",
// 		Body:  "Operation completed successfully",
// 	}
// 	meta := Meta{
// 		StatusCode: 200,
// 		//
// 	}
// 	return &HTTPResponse{
// 		Message:    message,
// 		Meta:       meta,
// 		Data:       data,
// 		Pagination: pagination,
// 	}
// }

// func NewErrorResponse(ctx ContextAdapter, errorMessage string, statusCode int) *HTTPResponse {
// 	// Populate message and meta based on your logic
// 	message := HTTPMessage{
// 		Title: "Error",
// 		Body:  errorMessage,
// 	}
// 	meta := Meta{
// 		StatusCode: statusCode,
// 		// ...
// 	}
// 	metaError := MetaError{
// 		Code:    statusCode,
// 		Message: errorMessage,
// 	}
// 	meta.Error = &metaError
// 	return &HTTPResponse{
// 		Message: message,
// 		Meta:    meta,
// 	}
// }

func NewSuccessResponseWithMessage(message string, data interface{}) *HTTPResponse {
	return &HTTPResponse{
		Message: HTTPMessage{
			Title: "Success",
			Body:  message,
		},
		Meta: Meta{
			StatusCode: 200,
			Message:    "OK",
		},
		Data: data,
	}
}

func NewSuccessResponse(data interface{}) *HTTPResponse {
	return &HTTPResponse{
		Message: HTTPMessage{
			Title: "Success",
			Body:  "Request was successful.",
		},
		Meta: Meta{
			StatusCode: 200,
			Message:    "OK",
		},
		Data: data,
	}
}

func NewSuccessResponseWithPagination(data interface{}, pagination *Pagination) *HTTPResponse {
	return &HTTPResponse{
		Message: HTTPMessage{
			Title: "Success",
			Body:  "Request was successful.",
		},
		Meta: Meta{
			StatusCode: 200,
			Message:    "OK",
		},
		Data:       data,
		Pagination: pagination,
	}
}

func NewErrorResponse(c *fiber.Ctx, statusCode int, message string, details string) error {
	resp := &HTTPResponse{
		Message: HTTPMessage{
			Title: "Error",
			Body:  message,
		},
		Meta: Meta{
			StatusCode: statusCode,
			Message:    message,
			Error: &MetaError{
				Code:    statusCode,
				Message: message,
				Detail:  details,
			},
		},
	}
	return c.Status(statusCode).JSON(resp)
}
