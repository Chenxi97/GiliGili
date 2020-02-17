package main

type ApiBody struct {
	Url     string `json:"url"  binding:"required"`
	Method  string `json:"method"  binding:"required"`
	ReqBody string `json:"req_body"  binding:"required"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestNotRecognized   = Err{Error: "api not recognized, bad request", ErrorCode: "001"}
	ErrorRequestBodyParseFailed = Err{Error: "request body is not correct", ErrorCode: "002"}
	ErrorInternalFaults         = Err{Error: "internal server errror", ErrorCode: "003"}
)
