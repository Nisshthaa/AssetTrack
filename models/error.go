package models

type ClientError struct {
	MessageToUser string `json:"messageToUser"`
	Err           string `json:"error"`
	StatusCode    int    `json:"statusCode"`
}

type ServiceResponse struct {
	Data       any
	StatusCode int
	Err        error
	Message    string
}
