package common

type ResponseData struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
}
