package api

import (
	"mime/multipart"
)

// SayHelloReq 入参
type SayHelloReq struct {
	Name string                  `p:"name" v:"required#名称不能为空"`
	File []*multipart.FileHeader `p:"file" v:"required#文件不能为空"`
}

// SayHelloRes 出参
type SayHelloRes struct {
	Say string `json:"say"`
}
