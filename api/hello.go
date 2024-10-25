package api

// SayHelloReq 入参
type SayHelloReq struct {
	Name string `p:"name" validate:"required"`
}

// SayHelloRes 出参
type SayHelloRes struct {
	Say string `json:"say"`
}
