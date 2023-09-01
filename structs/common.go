package structs

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}
