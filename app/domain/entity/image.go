package entity

import (
	"image"
)

// 上传文件方式
const (
	UploadTypeIsURL     = 1 // 图片URL上传
	UploadTypeIsFile    = 2 // 图片文件上传
	UploadTypeIsRawData = 3 // 图片二进制流上传
)

// ImgType 图片类型
type ImgType int

const (
	ImageTypePNG  ImgType = 1
	ImageTypeJPEG ImgType = 2
	ImageTypeGIF  ImgType = 3
	ImageTypeWEBP ImgType = 4
)

var ImgFileExtMap = map[ImgType]string{
	ImageTypeGIF:  ".gif",
	ImageTypePNG:  ".png",
	ImageTypeJPEG: ".jpg",
	ImageTypeWEBP: ".webp",
}

// UploadConfig 上传配置
type UploadConfig struct {
	ImageRoot   string // 存储根位置
	MaxSize     int    // 上传最大文件大小限制
	JpegQuality int    // Jpeg默认质量，75
	PngQuality  int    // Png默认质量，75
}

// Watermark 水印图片
type Watermark struct {
	Word       string
	File       image.Image
	PosX, PosY int
}
