package api

import "mime/multipart"

// SayHelloReq 入参
type SayHelloReq struct {
	Name string                  `p:"name" v:"required"`
	File []*multipart.FileHeader `p:"file" v:"required,min=1" msg:"请至少上传一个文件"`
}

// SayHelloRes 出参
type SayHelloRes struct {
	Say string `json:"say"`
}
