package command

type Response struct {
	Success bool `json:"success"`
}

func New() *Response {
	return &Response{
		Success: true,
	}
}
