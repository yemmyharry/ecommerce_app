package response

type Response struct {
	Messsage string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Error    error       `json:"error,omitempty"`
}
