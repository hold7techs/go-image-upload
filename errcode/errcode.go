package errcode

import (
	"go-image-upload/internal/errors"
)

// 错误类型
//	- 系统错误(10)
//		- 参数错误: 102001
//	- 业务错误(50)
//		- Transform(50)
//		-

var (
	// ErrTransformImgType 10类型
	ErrTransformImgType    = errors.New(102010, "transform fail, image type can not be gif")
	ErrTransformMkdir      = errors.New(102011, "transform fail, mkdir got err")
	ErrTransformCreateFile = errors.New(102012, "transform fail, os creat")
)
