package api

import "mime/multipart"

// SayHelloReq 入参
type SayHelloReq struct {
	Name string                  `p:"name" validate:"required"`
	File []*multipart.FileHeader `p:"file" validate:"required,dive"`
}

// SayHelloRes 出参
type SayHelloRes struct {
	Say string `json:"say"`
}
