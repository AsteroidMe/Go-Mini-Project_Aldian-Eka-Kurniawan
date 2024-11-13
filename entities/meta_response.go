package entities

type Meta struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}
