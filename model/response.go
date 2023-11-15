package model

import "encoding/json"

type WebResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type ErrorInputResponse struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

type ErrorInput struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
