package http

type SuccessCtx struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorCtx struct {
	Success bool                `json:"success"`
	Error   errorCodeAndMessage `json:"error"`
}

type errorCodeAndMessage struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
