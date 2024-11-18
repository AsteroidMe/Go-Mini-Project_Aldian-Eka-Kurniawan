package entities

type Meta struct {
	Status  bool
	Message string
}

type Response struct {
	Meta Meta
	Data interface{}
}
