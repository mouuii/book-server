package context

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}
