package utils

// Result common output
type Result struct {
	Data  interface{}
	Error error
}

// BaseWrapperModel data structure
type BaseWrapperModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Meta    interface{} `json:"meta,omitempty"`
}
