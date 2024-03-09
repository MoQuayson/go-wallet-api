package models

type APIResponse struct {
	Code         int64       `json:"code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	Errors       interface{} `json:"errors,omitempty"`
	CustomErrors interface{} `json:"cus_errors,omitempty"`
}

type ValidationError struct {
	Field   string
	Tag     string
	Message interface{}
}

// Only use this as a struct for channels
type DBResponse struct {
	Data  interface{}
	Error error
}
