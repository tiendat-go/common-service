package model

type ResponseCode string

const (
	Success        ResponseCode = "SUCCESS"
	NotFound       ResponseCode = "NOT_FOUND"
	Unauthorized   ResponseCode = "UNAUTHORIZED"
	Forbidden      ResponseCode = "FORBIDDEN"
	BadRequest     ResponseCode = "BAD_REQUEST"
	TooManyRequest ResponseCode = "TOO_MANY_REQUEST"
	Unavailable    ResponseCode = "UNAVAILABLE"
	Timeout        ResponseCode = "TIMEOUT"
)

const (
	SuccessMsg        = "Success"
	NotFoundMsg       = "Not found"
	UnauthorizedMsg   = "Unauthorized"
	ForbiddenMsg      = "Forbidden"
	BadRequestMsg     = "Bad request"
	TooManyRequestMsg = "Too many requests"
	UnavailableMsg    = "Service unavailable"
	TimeoutMsg        = "Request timeout"
)

type Response struct {
	Code      ResponseCode `json:"code"`
	Data      any          `json:"data"`
	Success   bool         `json:"success"`
	Message   string       `json:"message"`
	Total     *int64       `json:"total,omitempty"`
	NextToken *string      `json:"nextToken,omitempty"`
}
