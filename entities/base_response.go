package entities

type BaseResponse struct {
	Status  bool
	Message string
	Data    interface{}
}
