package query

type Response struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
}

func New(data map[string]interface{}) *Response {
	return &Response{
		Success: true,
		Data:    data,
	}
}
